
package control

import "C"
import (
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/endpoint"
	"nRIC/internal/msgx/service"
	"nRIC/internal/msgx/transport"
	"nRIC/internal/xapp/golog"
	"nRIC/internal/xapp"
	"github.com/spf13/viper"
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
	MsgClientToNricmgmt *msgx.MsgSender
	//subscriber *xapp.Subscriber
	CntRecvMsg uint64
	//AccessDbagent *dbclient.MsgSender
}

type MsgMeid struct {
	PlmnID  string
	EnbID   string
	RanName string
}

func init() {
	xapp.Logger.SetLevel(int(golog.ERR))
	xapp.Logger.Info("SMO")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("smo")
	viper.AllowEmptyEnv(true)
}

func NewControl( MsgClientToNricmgmt *msgx.MsgSender ) *Control {
	c := &Control{
		//MsgClientToXapp: MsgClientToXapp,
		MsgClientToNricmgmt:	MsgClientToNricmgmt,
	}
	return c
}

func (c *Control) ReadyCB(data interface{}) {
	if c.MsgClientToNricmgmt == nil {
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


func (c *Control) Consume(msg *xapp.MsgParams) (err error) {
	xapp.Logger.Debug("Received message type: %s", xapp.RicMessageTypeToName[msg.Mtype])

	if c.MsgClientToNricmgmt == nil {
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
	case xapp.RIC_O1_REGISTER:
		go c.handleXappRegister(msg)


	default:
		xapp.Logger.Info("Unknown Message Type '%d', discarding", msg.Mtype)
	}
	return
}

//-------------------------------------------------------------------
// handle from nRIC MGMT Register Request
//------------------------------------------------------------------
func (c *Control) handleXappRegister(params *xapp.MsgParams) {
	xapp.Logger.Info("MSG from nRIC Mgmt: %s", params.String())

	M := &msgx.XappRegMsg{}
	proto.Unmarshal(params.Payload,M)

	c.InitXapp(M)

	c.EnableXapp(M)
}


//-------------------------------------------------------------------
// Init xapp
//------------------------------------------------------------------
func (c *Control) InitXapp(M *msgx.XappRegMsg) {
	params := &xapp.MsgParams{
		Meid: &xapp.MsgMeid{RanName: ""},
	}
	IinitM 							:=&msgx.SMOInitMsg{}
	IinitM.Header 					= &msgx.RICMsgHeader{}
	IinitM.Header.XappRequestID		= &msgx.XAPPRequestID{}
	IinitM.Header.XappRequestID.XappInstanceID	= M.Header.XappRequestID.XappInstanceID
	IinitM.Header.XappRequestID.XappID 			= M.Header.XappRequestID.XappID
	IinitM.Header.MsgType       = xapp.RIC_O1_INIT
	IinitM.Header.MsgVer        = 1
	params.Mtype   		= xapp.RIC_O1_INIT
	IM, err	:= proto.Marshal(IinitM)
	if err != nil {
		xapp.Logger.Error(err.Error())
		return
	}
	params.Payload		= IM
	params.PayloadLen	= len(IM)
	xapp.Logger.Info("Send Msg to nRIC Mgmt: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	err = c.MsgClientToNricmgmt.SendMsg(params)
	if err != nil {
		xapp.Logger.Error(" Message Send to nRIC Mgmt failed :  '%s' \n", params.String())
	}
}


//-------------------------------------------------------------------
// Enable xapp
//------------------------------------------------------------------
func (c *Control) EnableXapp(M *msgx.XappRegMsg) {
	params := &xapp.MsgParams{
		Meid: &xapp.MsgMeid{RanName: ""},
	}
	MEnable := &msgx.SMOEnableMsg{Header:
		&msgx.RICMsgHeader{
			MsgType: xapp.RIC_O1_ENABLE,
			XappRequestID: &msgx.XAPPRequestID{
				XappInstanceID: M.Header.XappRequestID.XappInstanceID,
				XappID: M.Header.XappRequestID.XappID,
			},
		},
	}
	protoM,err := proto.Marshal(MEnable)
	if err != nil {
		xapp.Logger.Error(" Marshal Enable Message failed :  '%s' \n", err.Error())
	}

	params.Mtype 		= xapp.RIC_O1_ENABLE
	params.Payload		= protoM
	params.PayloadLen   = len(protoM)

	xapp.Logger.Info("Send Enable Msg to nRIC Mgmt: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	err = c.MsgClientToNricmgmt.SendMsg(params)
	if err != nil {
		xapp.Logger.Error(" Enable Message Send to nRIC Mgmt failed :  '%s' \n", params.String())
	}
}


//-------------------------------------------------------------------
// Disable xapp
//------------------------------------------------------------------
func (c *Control) DisableXapp(params *xapp.MsgParams) {
	M := &msgx.SMOEnableMsg{Header: &msgx.RICMsgHeader{MsgType: xapp.RIC_O1_DISABLE}}
	protoM,err := proto.Marshal(M)
	if err != nil {
		xapp.Logger.Error(" Marshal DISABLE Message failed :  '%s' \n", err.Error())
	}

	params.Mtype 		= xapp.RIC_O1_DISABLE
	params.Payload		= protoM
	params.PayloadLen   = len(protoM)

	xapp.Logger.Info("Send DISABLE Msg to nRIC Mgmt: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	err = c.MsgClientToNricmgmt.SendMsg(params)
	if err != nil {
		xapp.Logger.Error(" DISABLE Message Send to nRIC Mgmt failed :  '%s' \n", params.String())
	}
}
