
package control

import (
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/endpoint"
	"nRIC/internal/msgx/service"
	"nRIC/internal/msgx/transport"
	"nRIC/internal/xapp/golog"
	"nRIC/internal/xapp/models"
	"nRIC/internal/xapp"
	"github.com/spf13/viper"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	e2tclient "nRIC/pkg/nrice2t/route"
	"nRIC/pkg/nricsubs/e2ap"
	"net"
	"os"
	"time"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

func idstring(err error, entries ...fmt.Stringer) string {
	var retval string = ""
	var filler string = ""
	for _, entry := range entries {
		retval += filler + entry.String()
		filler = " "
	}
	if err != nil {
		retval += filler + "err(" + err.Error() + ")"
		filler = " "

	}
	return retval
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

var e2tSubReqTimeout time.Duration
var e2tSubDelReqTime time.Duration
var e2tRecvMsgTimeout time.Duration
var e2tMaxSubReqTryCount uint64    // Initial try + retry
var e2tMaxSubDelReqTryCount uint64 // Initial try + retry

type Control struct {
	//*xapp.MsgClientToXapp
	MsgClientToE2T *e2tclient.MsgSender //*msgx.MsgSender
	e2ap     *E2ap
	registry *Registry
	tracker  *Tracker
	//subscriber *xapp.Subscriber
	CntRecvMsg uint64
	AccessDbagent *dbclient.MsgSender
}

type MsgMeid struct {
	PlmnID  string
	EnbID   string
	RanName string
}

func init() {
	xapp.Logger.SetLevel(int(golog.ERR))
	xapp.Logger.Info("SUBMGR")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("submgr")
	viper.AllowEmptyEnv(true)
}

func NewControl(/*MsgClientToXapp msgx.SenderIf,*/ MsgClientToE2T *e2tclient.MsgSender /* *msgx.MsgSender*/,AcessDbAgent *dbclient.MsgSender) *Control {
	// viper.GetDuration returns nanoseconds
	e2tSubReqTimeout = viper.GetDuration("controls.e2tSubReqTimeout_ms") * 1000000
	if e2tSubReqTimeout == 0 {
		e2tSubReqTimeout = 2000 * 1000000
	}
	xapp.Logger.Info("e2tSubReqTimeout %v", e2tSubReqTimeout)
	e2tSubDelReqTime = viper.GetDuration("controls.e2tSubDelReqTime_ms") * 1000000
	if e2tSubDelReqTime == 0 {
		e2tSubDelReqTime = 2000 * 1000000
	}
	xapp.Logger.Info("e2tSubDelReqTime %v", e2tSubDelReqTime)
	e2tRecvMsgTimeout = viper.GetDuration("controls.e2tRecvMsgTimeout_ms") * 1000000
	if e2tRecvMsgTimeout == 0 {
		e2tRecvMsgTimeout = 2000 * 1000000
	}
	xapp.Logger.Info("e2tRecvMsgTimeout %v", e2tRecvMsgTimeout)
	e2tMaxSubReqTryCount = viper.GetUint64("controls.e2tMaxSubReqTryCount")
	if e2tMaxSubReqTryCount == 0 {
		e2tMaxSubReqTryCount = 1
	}
	xapp.Logger.Info("e2tMaxSubReqTryCount %v", e2tMaxSubReqTryCount)
	e2tMaxSubDelReqTryCount = viper.GetUint64("controls.e2tMaxSubDelReqTryCount")
	if e2tMaxSubDelReqTryCount == 0 {
		e2tMaxSubDelReqTryCount = 1
	}
	xapp.Logger.Info("e2tMaxSubDelReqTryCount %v", e2tMaxSubDelReqTryCount)

	registry := new(Registry)
	registry.Initialize()
	registry.AcessDbAgent = AcessDbAgent

	tracker := new(Tracker)
	tracker.Init()

	//subscriber := xapp.NewSubscriber(viper.GetString("subscription.host"), viper.GetInt("subscription.timeout"))

	c := &Control{e2ap: new(E2ap),
		registry: registry,
		tracker:  tracker,
		MsgClientToE2T:	MsgClientToE2T,
		AccessDbagent: AcessDbAgent,
	}
	return c
}

func (c *Control) ReadyCB(data interface{}) {
	if c.MsgClientToE2T == nil {
	}
}


func (c *Control)  CreateAndRunMsgServer (grpcAddr string) {

	svc := service.NewMsgService(c)
	ep  := endpoint.NewMsgServiceEndpoint(svc)
	s   := transport.NewMsgServer(ep)


	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		xapp.Logger.Info("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	xapp.Logger.Info("transport", "gRPC", "addr", grpcAddr)
	// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
	// the here demonstrated zipkin tracing middleware.
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	msgx.RegisterMsgServiceServer(baseServer, s)
	baseServer.Serve(grpcListener)
}

func (c *Control) Run(grpcAddr string) {
	c.CreateAndRunMsgServer(grpcAddr)
}

//-------------------------------------------------------------------
//
//-------------------------------------------------------------------
func (c *Control) SubscriptionHandler(stype models.SubscriptionType, params interface{}) (*models.SubscriptionResponse, error) {

	return &models.SubscriptionResponse{}, fmt.Errorf("Subscription rest interface not implemented")
}


func (c *Control) SubscriptionDeleteHandler(string) error {
	return fmt.Errorf("Subscription rest interface not implemented")
}

func (c *Control) QueryHandler() (models.SubscriptionList, error) {
	return c.registry.QueryHandler()
}

//-------------------------------------------------------------------
//
//-------------------------------------------------------------------

func (c *Control) msgSendToE2T(desc string, subs *Subscription, trans *TransactionSubs) (err error) {
	params := &xapp.MsgParams{}
	params.Mtype = trans.GetMtype()
	params.SubId = int(subs.GetReqId().InstanceId)
	params.Xid = ""
	params.Meid = subs.GetMeid()
	params.PayloadLen = len(trans.Payload.Buf)
	params.Payload = trans.Payload.Buf
	//params.Mbuf = nil
	xapp.Logger.Info("MSG to E2T: %s %s %s", desc, trans.String(), params.String())
	return c.MsgClientToE2T.SendMsg(params)
}

func (c *Control) msgSendToXapp(desc string, subs *Subscription, trans *TransactionXapp) (err error) {

	params := &xapp.MsgParams{}
	params.Mtype = trans.GetMtype()
	params.SubId = int(subs.GetReqId().InstanceId)
	params.Xid = trans.GetXid()
	params.Meid = trans.GetMeid()
	params.PayloadLen = len(trans.Payload.Buf)
	params.Payload = trans.Payload.Buf
	//params.Mbuf = nil
	xapp.Logger.Info("MSG to XAPP: %s %s %s", desc, trans.String(), params.String())
	return trans.Endpoint.SendMsg(params)
}

func (c *Control) Consume(msg *xapp.MsgParams) (err error) {
	if c.MsgClientToE2T == nil {
		err = fmt.Errorf("Msg object nil can handle %s", msg.String())
		xapp.Logger.Error("%s", err.Error())
		return
	}
	c.CntRecvMsg++

	switch msg.Mtype {

	case xapp.RIC_SUB_REQ:
		go c.handleXAPPSubscriptionRequest(msg)

	case xapp.RIC_SUB_RESP:
		go c.handleE2TSubscriptionResponse(msg)

	case xapp.RIC_SUB_FAILURE:
		go c.handleE2TSubscriptionFailure(msg)
	case xapp.RIC_SUB_DEL_REQ:
		go c.handleXAPPSubscriptionDeleteRequest(msg)

	case xapp.RIC_SUB_DEL_RESP:
		go c.handleE2TSubscriptionDeleteResponse(msg)
	case xapp.RIC_SUB_DEL_FAILURE:
		go c.handleE2TSubscriptionDeleteFailure(msg)


	default:
		xapp.Logger.Info("Unknown Message Type '%d', discarding", msg.Mtype)
	}
	return
}

//-------------------------------------------------------------------
// handle from XAPP Subscription Request
//------------------------------------------------------------------
func (c *Control) handleXAPPSubscriptionRequest(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from XAPP: %s", params.String())

	subReqMsg, err := c.e2ap.UnpackSubscriptionRequest(params.Payload)
	if err != nil {
		xapp.Logger.Error("XAPP-SubReq: %s", idstring(err, params))
		return
	}
	Endpoint := msgx.NewKafkaMsgSender(params.Src)
	trans := c.tracker.NewXappTransaction(Endpoint, params.Xid, subReqMsg.RequestId.InstanceId, params.Meid)
	if trans == nil {
		xapp.Logger.Error("XAPP-SubReq: %s", idstring(fmt.Errorf("transaction not created"), params))
		return
	}

	defer trans.Release()

	err = c.tracker.Track(trans)
	if err != nil {
		xapp.Logger.Error("XAPP-SubReq: %s", idstring(err, trans))
		return
	}

	//TODO handle subscription toward e2term inside AssignToSubscription / hide handleSubscriptionCreate in it?
	subs, err := c.registry.AssignToSubscription(trans, subReqMsg)
	if err != nil {
		xapp.Logger.Error("XAPP-SubReq: %s", idstring(err, trans))
		return
	}

	//
	// Wake subs request
	//
	go c.handleSubscriptionCreate(subs, trans)
	event, _ := trans.WaitEvent(0) //blocked wait as timeout is handled in subs side

	err = nil
	if event != nil {
		switch themsg := event.(type) {
		case *e2ap.E2APSubscriptionResponse:
			trans.Mtype, trans.Payload, err = c.e2ap.PackSubscriptionResponse(themsg)
			if err == nil {
				trans.Release()
				c.msgSendToXapp("", subs, trans)
				return
			}
		case *e2ap.E2APSubscriptionFailure:
			trans.Mtype, trans.Payload, err = c.e2ap.PackSubscriptionFailure(themsg)
			if err == nil {
				c.msgSendToXapp("", subs, trans)
			}
		default:
			break
		}
	}
	xapp.Logger.Info("XAPP-SubReq: failed %s", idstring(err, trans, subs))

}



//-------------------------------------------------------------------
// handle from XAPP Subscription Delete Request
//------------------------------------------------------------------
func (c *Control) handleXAPPSubscriptionDeleteRequest(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from XAPP: %s", params.String())

	subDelReqMsg, err := c.e2ap.UnpackSubscriptionDeleteRequest(params.Payload)
	if err != nil {
		xapp.Logger.Error("XAPP-SubDelReq %s", idstring(err, params))
		return
	}
	Endpoint := msgx.NewKafkaMsgSender(params.Src)
	trans := c.tracker.NewXappTransaction(Endpoint, params.Xid, subDelReqMsg.RequestId.InstanceId, params.Meid)
	if trans == nil {
		xapp.Logger.Error("XAPP-SubDelReq: %s", idstring(fmt.Errorf("transaction not created"), params))
		return
	}
	defer trans.Release()

	err = c.tracker.Track(trans)
	if err != nil {
		xapp.Logger.Error("XAPP-SubReq: %s", idstring(err, trans))
		return
	}

	subs, err := c.registry.GetSubscriptionFirstMatch([]uint32{trans.GetSubId()})
	if err != nil {
		xapp.Logger.Error("XAPP-SubDelReq: %s", idstring(err, trans))
		return
	}

	//
	// Wake subs delete
	//
	go c.handleSubscriptionDelete(subs, trans)
	trans.WaitEvent(0) //blocked wait as timeout is handled in subs side

	xapp.Logger.Debug("XAPP-SubDelReq: Handling event %s ", idstring(nil, trans, subs))

	// Whatever is received send ok delete response
	subDelRespMsg := &e2ap.E2APSubscriptionDeleteResponse{}
	subDelRespMsg.RequestId = subs.GetReqId().RequestId
	subDelRespMsg.FunctionId = subs.SubReqMsg.FunctionId
	trans.Mtype, trans.Payload, err = c.e2ap.PackSubscriptionDeleteResponse(subDelRespMsg)
	if err == nil {
		c.msgSendToXapp("", subs, trans)
	}

}

//-------------------------------------------------------------------
// SUBS CREATE Handling
//-------------------------------------------------------------------
func (c *Control) handleSubscriptionCreate(subs *Subscription, parentTrans *TransactionXapp) {

	trans := c.tracker.NewSubsTransaction(subs)
	subs.WaitTransactionTurn(trans)
	defer subs.ReleaseTransactionTurn(trans)
	defer trans.Release()

	xapp.Logger.Debug("SUBS-SubReq: Handling %s ", idstring(nil, trans, subs, parentTrans))

	subRfMsg, valid := subs.GetCachedResponse()
	if subRfMsg == nil && valid == true {

		//
		// In case of failure
		// - make internal delete
		// - in case duplicate cause, retry (currently max 1 retry)
		//
		maxRetries := uint64(1)
		doRetry := true
		for retries := uint64(0); retries <= maxRetries && doRetry; retries++ {
			doRetry = false

			event := c.sendE2TSubscriptionRequest(subs, trans, parentTrans)
			switch themsg := event.(type) {
			case *e2ap.E2APSubscriptionResponse:
				subRfMsg, valid = subs.SetCachedResponse(event, true)
			case *e2ap.E2APSubscriptionFailure:
				subRfMsg, valid = subs.SetCachedResponse(event, false)
				doRetry = true
				for _, item := range themsg.ActionNotAdmittedList.Items {
					if item.Cause.Content != e2ap.E2AP_CauseContent_Ric || (item.Cause.Value != e2ap.E2AP_CauseValue_Ric_duplicate_action && item.Cause.Value != e2ap.E2AP_CauseValue_Ric_duplicate_event) {
						doRetry = false
						break
					}
				}
				xapp.Logger.Info("SUBS-SubReq: internal delete and possible retry due event(%s) retry(%t,%d/%d) %s", typeofSubsMessage(event), doRetry, retries, maxRetries, idstring(nil, trans, subs, parentTrans))
				c.sendE2TSubscriptionDeleteRequest(subs, trans, parentTrans)
			default:
				xapp.Logger.Info("SUBS-SubReq: internal delete due event(%s) %s", typeofSubsMessage(event), idstring(nil, trans, subs, parentTrans))
				subRfMsg, valid = subs.SetCachedResponse(nil, false)
				c.sendE2TSubscriptionDeleteRequest(subs, trans, parentTrans)
			}
		}

		xapp.Logger.Debug("SUBS-SubReq: Handling (e2t response %s) %s", typeofSubsMessage(subRfMsg), idstring(nil, trans, subs, parentTrans))
	} else {
		xapp.Logger.Debug("SUBS-SubReq: Handling (cached response %s) %s", typeofSubsMessage(subRfMsg), idstring(nil, trans, subs, parentTrans))
	}

	//Now RemoveFromSubscription in here to avoid race conditions (mostly concerns delete)
	if valid == false {
		c.registry.RemoveFromSubscription(subs, parentTrans, 5*time.Second)
	}
	parentTrans.SendEvent(subRfMsg, 0)
}

//-------------------------------------------------------------------
// SUBS DELETE Handling
//-------------------------------------------------------------------

func (c *Control) handleSubscriptionDelete(subs *Subscription, parentTrans *TransactionXapp) {

	trans := c.tracker.NewSubsTransaction(subs)
	subs.WaitTransactionTurn(trans)
	defer subs.ReleaseTransactionTurn(trans)
	defer trans.Release()

	xapp.Logger.Debug("SUBS-SubDelReq: Handling %s", idstring(nil, trans, subs, parentTrans))

	subs.mutex.Lock()
	if subs.valid && subs.EpList.HasEndpoint(parentTrans.GetEndpoint()) && subs.EpList.Size() == 1 {
		subs.valid = false
		subs.mutex.Unlock()
		c.sendE2TSubscriptionDeleteRequest(subs, trans, parentTrans)
	} else {
		subs.mutex.Unlock()
	}
	//Now RemoveFromSubscription in here to avoid race conditions (mostly concerns delete)
	//  If parallel deletes ongoing both might pass earlier sendE2TSubscriptionDeleteRequest(...) if
	//  RemoveFromSubscription locates in caller side (now in handleXAPPSubscriptionDeleteRequest(...))
	c.registry.RemoveFromSubscription(subs, parentTrans, 5*time.Second)
	parentTrans.SendEvent(nil, 0)
}

//-------------------------------------------------------------------
// send to E2T Subscription Request
//-------------------------------------------------------------------
func (c *Control) sendE2TSubscriptionRequest(subs *Subscription, trans *TransactionSubs, parentTrans *TransactionXapp) interface{} {
	var err error
	var event interface{} = nil
	var timedOut bool = false

	subReqMsg := subs.SubReqMsg
	subReqMsg.RequestId = subs.GetReqId().RequestId
	trans.Mtype, trans.Payload, err = c.e2ap.PackSubscriptionRequest(subReqMsg)
	if err != nil {
		xapp.Logger.Error("SUBS-SubReq: %s", idstring(err, trans, subs, parentTrans))
		return event
	}

	for retries := uint64(0); retries < e2tMaxSubReqTryCount; retries++ {
		desc := fmt.Sprintf("(retry %d)", retries)
		c.msgSendToE2T(desc, subs, trans)
		event, timedOut = trans.WaitEvent(e2tSubReqTimeout)
		if timedOut {
			continue
		}
		break
	}
	xapp.Logger.Debug("SUBS-SubReq: Response handling event(%s) %s", typeofSubsMessage(event), idstring(nil, trans, subs, parentTrans))
	return event
}

//-------------------------------------------------------------------
// send to E2T Subscription Delete Request
//-------------------------------------------------------------------

func (c *Control) sendE2TSubscriptionDeleteRequest(subs *Subscription, trans *TransactionSubs, parentTrans *TransactionXapp) interface{} {
	var err error
	var event interface{}
	var timedOut bool

	subDelReqMsg := &e2ap.E2APSubscriptionDeleteRequest{}
	subDelReqMsg.RequestId = subs.GetReqId().RequestId
	subDelReqMsg.FunctionId = subs.SubReqMsg.FunctionId
	trans.Mtype, trans.Payload, err = c.e2ap.PackSubscriptionDeleteRequest(subDelReqMsg)
	if err != nil {
		xapp.Logger.Error("SUBS-SubDelReq: %s", idstring(err, trans, subs, parentTrans))
		return event
	}

	for retries := uint64(0); retries < e2tMaxSubDelReqTryCount; retries++ {
		desc := fmt.Sprintf("(retry %d)", retries)
		c.msgSendToE2T(desc, subs, trans)
		event, timedOut = trans.WaitEvent(e2tSubDelReqTime)
		if timedOut {
			continue
		}
		break
	}
	xapp.Logger.Debug("SUBS-SubDelReq: Response handling event(%s) %s", typeofSubsMessage(event), idstring(nil, trans, subs, parentTrans))
	return event
}

//-------------------------------------------------------------------
// handle from E2T Subscription Reponse
//-------------------------------------------------------------------
func (c *Control) handleE2TSubscriptionResponse(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from E2T: %s", params.String())
	subRespMsg, err := c.e2ap.UnpackSubscriptionResponse(params.Payload)
	if err != nil {
		xapp.Logger.Error("MSG-SubResp %s", idstring(err, params))
		return
	}
	subs, err := c.registry.GetSubscriptionFirstMatch([]uint32{subRespMsg.RequestId.InstanceId})
	if err != nil {
		xapp.Logger.Error("MSG-SubResp: %s", idstring(err, params))
		return
	}
	trans := subs.GetTransaction()
	if trans == nil {
		err = fmt.Errorf("Ongoing transaction not found")
		xapp.Logger.Error("MSG-SubResp: %s", idstring(err, params, subs))
		return
	}
	sendOk, timedOut := trans.SendEvent(subRespMsg, e2tRecvMsgTimeout)
	if sendOk == false {
		err = fmt.Errorf("Passing event to transaction failed: sendOk(%t) timedOut(%t)", sendOk, timedOut)
		xapp.Logger.Error("MSG-SubResp: %s", idstring(err, trans, subs))
	}
	return
}

//-------------------------------------------------------------------
// handle from E2T Subscription Failure
//-------------------------------------------------------------------
func (c *Control) handleE2TSubscriptionFailure(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from E2T: %s", params.String())
	subFailMsg, err := c.e2ap.UnpackSubscriptionFailure(params.Payload)
	if err != nil {
		xapp.Logger.Error("MSG-SubFail %s", idstring(err, params))
		return
	}
	subs, err := c.registry.GetSubscriptionFirstMatch([]uint32{subFailMsg.RequestId.InstanceId})
	if err != nil {
		xapp.Logger.Error("MSG-SubFail: %s", idstring(err, params))
		return
	}
	trans := subs.GetTransaction()
	if trans == nil {
		err = fmt.Errorf("Ongoing transaction not found")
		xapp.Logger.Error("MSG-SubFail: %s", idstring(err, params, subs))
		return
	}
	sendOk, timedOut := trans.SendEvent(subFailMsg, e2tRecvMsgTimeout)
	if sendOk == false {
		err = fmt.Errorf("Passing event to transaction failed: sendOk(%t) timedOut(%t)", sendOk, timedOut)
		xapp.Logger.Error("MSG-SubFail: %s", idstring(err, trans, subs))
	}
	return
}

//-------------------------------------------------------------------
// handle from E2T Subscription Delete Response
//-------------------------------------------------------------------
func (c *Control) handleE2TSubscriptionDeleteResponse(params *xapp.MsgParams) (err error) {
	xapp.Logger.Info("MSG from E2T: %s", params.String())
	subDelRespMsg, err := c.e2ap.UnpackSubscriptionDeleteResponse(params.Payload)
	if err != nil {
		xapp.Logger.Error("MSG-SubDelResp: %s", idstring(err, params))
		return
	}
	subs, err := c.registry.GetSubscriptionFirstMatch([]uint32{subDelRespMsg.RequestId.InstanceId})
	if err != nil {
		xapp.Logger.Error("MSG-SubDelResp: %s", idstring(err, params))
		return
	}
	trans := subs.GetTransaction()
	if trans == nil {
		err = fmt.Errorf("Ongoing transaction not found")
		xapp.Logger.Error("MSG-SubDelResp: %s", idstring(err, params, subs))
		return
	}
	sendOk, timedOut := trans.SendEvent(subDelRespMsg, e2tRecvMsgTimeout)
	if sendOk == false {
		err = fmt.Errorf("Passing event to transaction failed: sendOk(%t) timedOut(%t)", sendOk, timedOut)
		xapp.Logger.Error("MSG-SubDelResp: %s", idstring(err, trans, subs))
	}
	return
}

//-------------------------------------------------------------------
// handle from E2T Subscription Delete Failure
//-------------------------------------------------------------------
func (c *Control) handleE2TSubscriptionDeleteFailure(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from E2T: %s", params.String())
	subDelFailMsg, err := c.e2ap.UnpackSubscriptionDeleteFailure(params.Payload)
	if err != nil {
		xapp.Logger.Error("MSG-SubDelFail: %s", idstring(err, params))
		return
	}
	subs, err := c.registry.GetSubscriptionFirstMatch([]uint32{subDelFailMsg.RequestId.InstanceId})
	if err != nil {
		xapp.Logger.Error("MSG-SubDelFail: %s", idstring(err, params))
		return
	}
	trans := subs.GetTransaction()
	if trans == nil {
		err = fmt.Errorf("Ongoing transaction not found")
		xapp.Logger.Error("MSG-SubDelFail: %s", idstring(err, params, subs))
		return
	}
	sendOk, timedOut := trans.SendEvent(subDelFailMsg, e2tRecvMsgTimeout)
	if sendOk == false {
		err = fmt.Errorf("Passing event to transaction failed: sendOk(%t) timedOut(%t)", sendOk, timedOut)
		xapp.Logger.Error("MSG-SubDelFail: %s", idstring(err, trans, subs))
	}
	return
}

//-------------------------------------------------------------------
//
//-------------------------------------------------------------------
func typeofSubsMessage(v interface{}) string {
	if v == nil {
		return "NIL"
	}
	switch v.(type) {
	case *e2ap.E2APSubscriptionRequest:
		return "SubReq"
	case *e2ap.E2APSubscriptionResponse:
		return "SubResp"
	case *e2ap.E2APSubscriptionFailure:
		return "SubFail"
	case *e2ap.E2APSubscriptionDeleteRequest:
		return "SubDelReq"
	case *e2ap.E2APSubscriptionDeleteResponse:
		return "SubDelResp"
	case *e2ap.E2APSubscriptionDeleteFailure:
		return "SubDelFail"
	default:
		return "Unknown"
	}
}