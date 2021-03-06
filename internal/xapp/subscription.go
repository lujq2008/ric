
package xapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-openapi/loads"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
	//"errors"

	apiclient "nRIC/internal/xapp/clientapi"
	apicommon "nRIC/internal/xapp/clientapi/common"
	apipolicy "nRIC/internal/xapp/clientapi/policy"
	apireport "nRIC/internal/xapp/clientapi/report"
	apimodel "nRIC/internal/xapp/clientmodel"

	"nRIC/internal/xapp/models"
	"nRIC/internal/xapp/restapi"
	"nRIC/internal/xapp/restapi/operations"
	"nRIC/internal/xapp/restapi/operations/common"
	"nRIC/internal/xapp/restapi/operations/policy"
	"nRIC/internal/xapp/restapi/operations/query"
	"nRIC/internal/xapp/restapi/operations/report"
)

type SubscriptionHandler func(models.SubscriptionType, interface{}) (*models.SubscriptionResponse, error)
type SubscriptionQueryHandler func() (models.SubscriptionList, error)
type SubscriptionDeleteHandler func(string) error
type SubscriptionResponseCallback func(*apimodel.SubscriptionResponse)

type Subscriber struct {
	localAddr  string
	localPort  int
	remoteHost string
	remoteUrl  string
	remoteProt []string
	timeout    time.Duration
	clientUrl  string
	clientCB   SubscriptionResponseCallback
}

func NewSubscriber(host string, timo int) *Subscriber {
	if host == "" {
		pltnamespace := os.Getenv("PLT_NAMESPACE")
		if pltnamespace == "" {
			pltnamespace = "ricplt"
		}
		host = fmt.Sprintf("service-%s-submgr-http.%s:8088", pltnamespace, pltnamespace)
	}

	if timo == 0 {
		timo = 20
	}

	r := &Subscriber{
		remoteHost: host,
		remoteUrl:  "/ric/v1",
		remoteProt: []string{"http"},
		timeout:    time.Duration(timo) * time.Second,
		localAddr:  "0.0.0.0",
		localPort:  8088,
		clientUrl:  "/ric/v1/subscriptions/response",
	}
	Resource.InjectRoute(r.clientUrl, r.ResponseHandler, "POST")

	return r
}

func (r *Subscriber) ResponseHandler(w http.ResponseWriter, req *http.Request) {
	if req.Body != nil {
		var resp apimodel.SubscriptionResponse
		if err := json.NewDecoder(req.Body).Decode(&resp); err == nil {
			if r.clientCB != nil {
				r.clientCB(&resp)
			}
		}
		req.Body.Close()
	}
}

// Server interface: listen and receive subscription requests
func (r *Subscriber) Listen(createSubscription SubscriptionHandler, getSubscription SubscriptionQueryHandler, delSubscription SubscriptionDeleteHandler) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	api := operations.NewXappFrameworkAPI(swaggerSpec)

	// Subscription: Query
	api.QueryGetAllSubscriptionsHandler = query.GetAllSubscriptionsHandlerFunc(
		func(p query.GetAllSubscriptionsParams) middleware.Responder {
			if resp, err := getSubscription(); err == nil {
				return query.NewGetAllSubscriptionsOK().WithPayload(resp)
			}
			return query.NewGetAllSubscriptionsInternalServerError()
		})

	// SubscriptionType: Report
	api.ReportSubscribeReportHandler = report.SubscribeReportHandlerFunc(
		func(p report.SubscribeReportParams) middleware.Responder {
			if resp, err := createSubscription(models.SubscriptionTypeReport, p.ReportParams); err == nil {
				return report.NewSubscribeReportCreated().WithPayload(resp)
			}
			return report.NewSubscribeReportInternalServerError()
		})

	// SubscriptionType: Policy
	api.PolicySubscribePolicyHandler = policy.SubscribePolicyHandlerFunc(
		func(p policy.SubscribePolicyParams) middleware.Responder {
			if resp, err := createSubscription(models.SubscriptionTypePolicy, p.PolicyParams); err == nil {
				return policy.NewSubscribePolicyCreated().WithPayload(resp)
			}
			return policy.NewSubscribePolicyInternalServerError()
		})

	// SubscriptionType: Delete
	api.CommonUnsubscribeHandler = common.UnsubscribeHandlerFunc(
		func(p common.UnsubscribeParams) middleware.Responder {
			if err := delSubscription(p.SubscriptionID); err == nil {
				return common.NewUnsubscribeNoContent()
			}
			return common.NewUnsubscribeInternalServerError()
		})



	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Host = r.localAddr
	server.Port = r.localPort

	Logger.Info("Serving subscriptions on %s:%d\n", server.Host, server.Port)
	if err := server.Serve(); err != nil {
		return err
	}
	return nil
}

// Server interface: send notification to client
func (r *Subscriber) Notify(resp *models.SubscriptionResponse, clientEndpoint string) (err error) {
	respData, err := json.Marshal(resp)
	if err != nil {
		Logger.Error("json.Marshal failed: %v", err)
		return err
	}

	ep, _, _ := net.SplitHostPort(clientEndpoint)
	_, port, _ := net.SplitHostPort(fmt.Sprintf(":%d", GetPortData("http").Port))
	clientUrl := fmt.Sprintf("http://%s:%s%s", ep, port, r.clientUrl)

	retries := viper.GetInt("subscription.retryCount")
	if retries == 0 {
		retries = 10
	}

	delay := viper.GetInt("subscription.retryDelay")
	if delay == 0 {
		delay = 5
	}

	for i := 0; i < retries; i++ {
		r, err := http.Post(clientUrl, "application/json", bytes.NewBuffer(respData))
		if err == nil && (r != nil && r.StatusCode == http.StatusOK) {
			break
		}

		if err != nil {
			Logger.Error("%v", err)
		}
		if r != nil && r.StatusCode != http.StatusOK {
			Logger.Error("clientUrl=%s statusCode=%d", clientUrl, r.StatusCode)
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}

	return err
}

// Subscription interface for xApp: Response callback
func (r *Subscriber) SetResponseCB(c SubscriptionResponseCallback) {
	r.clientCB = c
}

// Subscription interface for xApp: REPORT
func (r *Subscriber) SubscribeReport(p *apimodel.ReportParams) (*apimodel.SubscriptionResponse, error) {
	params := apireport.NewSubscribeReportParamsWithTimeout(r.timeout).WithReportParams(p)
	result, err := r.CreateTransport().Report.SubscribeReport(params)
	if err != nil {
		return &apimodel.SubscriptionResponse{}, err
	}

	return result.Payload, err
}

// Subscription interface for xApp: POLICY
func (r *Subscriber) SubscribePolicy(p *apimodel.PolicyParams) (*apimodel.SubscriptionResponse, error) {
	params := apipolicy.NewSubscribePolicyParamsWithTimeout(r.timeout).WithPolicyParams(p)
	result, err := r.CreateTransport().Policy.SubscribePolicy(params)
	if err != nil {
		return &apimodel.SubscriptionResponse{}, err
	}

	return result.Payload, err
}

// Subscription interface for xApp: DELETE
func (r *Subscriber) UnSubscribe(subId string) error {
	params := apicommon.NewUnsubscribeParamsWithTimeout(r.timeout).WithSubscriptionID(subId)
	_, err := r.CreateTransport().Common.Unsubscribe(params)

	return err
}

// Subscription interface for xApp: QUERY
func (r *Subscriber) QuerySubscriptions() (models.SubscriptionList, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/%s/subscriptions", r.remoteHost, r.remoteUrl))
	if err != nil {
		return models.SubscriptionList{}, err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.SubscriptionList{}, err
	}

	subscriptions := models.SubscriptionList{}
	err = json.Unmarshal([]byte(string(contents)), &subscriptions)
	if err != nil {
		return models.SubscriptionList{}, err
	}

	return subscriptions, nil
}

func (r *Subscriber) CreateTransport() *apiclient.RICSubscription {
	return apiclient.New(httptransport.New(r.remoteHost, r.remoteUrl, r.remoteProt), strfmt.Default)
}
