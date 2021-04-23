
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
	"nRIC/internal/xapp"
	"github.com/spf13/viper"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	e2tclient "nRIC/pkg/nrice2t/route"
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
	xapp.Logger.Info("CONFLICTMGR")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cflmgr")
	viper.AllowEmptyEnv(true)
}

func NewControl(/*MsgClientToXapp msgx.SenderIf,*/ MsgClientToE2T *e2tclient.MsgSender /* *msgx.MsgSender*/,AcessDbAgent *dbclient.MsgSender) *Control {

	registry := new(Registry)
	registry.Initialize()
	registry.AcessDbAgent = AcessDbAgent

	tracker := new(Tracker)
	tracker.Init()

	//subscriber := xapp.NewSubscriber(viper.GetString("subscription.host"), viper.GetInt("subscription.timeout"))

	c := &Control{e2ap: new(E2ap),
		registry: registry,
		tracker:  tracker,
		//MsgClientToXapp: MsgClientToXapp,
		MsgClientToE2T:	MsgClientToE2T,
		//subscriber: subscriber,
		AccessDbagent: AcessDbAgent,
	}
	//go xapp.Subscription.Listen(c.SubscriptionHandler, c.QueryHandler, c.SubscriptionDeleteHandler)

	//go c.subscriber.Listen(c.SubscriptionHandler, c.QueryHandler)
	return c
}

func (c *Control) ReadyCB(data interface{}) {
	if c.MsgClientToE2T == nil {
		//c.MsgClientToXapp = xapp.Msg
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
	//xapp.SetReadyCB(c.ReadyCB, nil)
	//xapp.Run(c,grpcAddr)
	c.CreateAndRunMsgServer(grpcAddr)
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
	return c.MsgClientToE2T.SendMsg(params) //c.SendWithRetry(params, false, 5)
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
	return trans.Endpoint.SendMsg(params) //c.SendWithRetry(params, false, 5)
}

func (c *Control) Consume(msg *xapp.MsgParams) (err error) {
	if c.MsgClientToE2T == nil {
		err = fmt.Errorf("Msg object nil can handle %s", msg.String())
		xapp.Logger.Error("%s", err.Error())
		return
	}
	c.CntRecvMsg++

	//	defer c.MsgClientTo.Free(msg.Mbuf)

	// xapp-frame might use direct access to c buffer and
	// when msg.Mbuf is freed, someone might take it into use
	// and payload data might be invalid inside message handle function
	//
	// subscriptions won't load system a lot so there is no
	// real performance hit by cloning buffer into new go byte slice
	/*
		cPay := append(msg.Payload[:0:0], msg.Payload...)
		msg.Payload = cPay
		msg.PayloadLen = len(cPay)
	*/

	switch msg.Mtype {
	case xapp.RIC_CONTROL_REQ:
		go c.handleXAPPControlRequest(msg)

	default:
		xapp.Logger.Info("Unknown Message Type '%d', discarding", msg.Mtype)
	}
	return
}

//-------------------------------------------------------------------
// handle from XAPP Subscription Request
//------------------------------------------------------------------
func (c *Control) handleXAPPControlRequest(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from XAPP: %s", params.String())
	err := c.MsgClientToE2T.SendMsg(params)
	if err != nil {
		//xapp.Logger.Error(" Message Type '%d', Send to E2T failed!\n", params.Mtype)
		xapp.Logger.Error(" Message Send to E2T failed :  '%s' \n", params.String())
	}
}

