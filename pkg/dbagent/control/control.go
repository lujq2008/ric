package control

import (
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	pbdb "nRIC/api/v1/pb/db"
	pbe2t "nRIC/api/v1/pb/nrice2t"
	"nRIC/internal/xapp"
	"github.com/spf13/viper"
	e2tclient "nRIC/pkg/nrice2t/route"
	"nRIC/pkg/dbagent/grpcserver/endpoint"
	"nRIC/pkg/dbagent/grpcserver/service"
	"nRIC/pkg/dbagent/grpcserver/transport"
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
	MsgClientToE2T *e2tclient.MsgSender
	//subscriber *xapp.Subscriber
	CntRecvMsg uint64
}

type MsgMeid struct {
	PlmnID  string
	EnbID   string
	RanName string
}

func init() {
	xapp.Logger.Info("SUBMGR")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("submgr")
	viper.AllowEmptyEnv(true)
}

func NewControl(/*MsgClientToXapp msgx.SenderIf,*/ MsgClientToE2T *e2tclient.MsgSender) *Control {
	//subscriber := xapp.NewSubscriber(viper.GetString("subscription.host"), viper.GetInt("subscription.timeout"))

	c := &Control{
		MsgClientToE2T:	MsgClientToE2T,
		//subscriber: subscriber,
	}

	return c
}

func (c *Control) ReadyCB(data interface{}) {
	if c.MsgClientToE2T == nil {
		//c.MsgClientToXapp = xapp.Msg
	}
}


func (c *Control)  CreateAndRunDbServer (grpcAddr string) {

	svc := service.NewmsgService(c)
	ep  := endpoint.NewMsgServiceEndpoint(svc)
	s   := transport.NewmsgServer(ep)


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
	pbdb.RegisterDbServiceServer(baseServer, s)
	baseServer.Serve(grpcListener)
}

func (c *Control) Run(grpcAddr string) {
	//xapp.SetReadyCB(c.ReadyCB, nil)
	//xapp.Run(c,grpcAddr)
	c.CreateAndRunDbServer(grpcAddr)
}


func (c *Control) Consume(msg *xapp.MsgParams) (err error) {
	if c.MsgClientToE2T == nil {
		err = fmt.Errorf("Msg object nil can handle %s", msg.String())
		xapp.Logger.Error("%s", err.Error())
		return
	}
	c.CntRecvMsg++

	switch msg.Mtype {
	case xapp.RIC_ROUTE_INSERT: //publish route to E2T
		RouteTable := &pbe2t.RouteTable{SubIdXapp: msg.Route.SubIdXapp,Topic: msg.Route.Topic,SubIdRan: msg.Route.SubIdRan}
		go c.MsgClientToE2T.RouteTableInsert(&pbe2t.RouteTableInsertRequest{RouteTable: RouteTable})

	case xapp.RIC_ROUTE_UPDATE: //publish route to E2T
		RouteTable := &pbe2t.RouteTable{SubIdXapp: msg.Route.SubIdXapp,Topic: msg.Route.Topic,SubIdRan: msg.Route.SubIdRan}
		go c.MsgClientToE2T.RouteTableUpdate(&pbe2t.RouteTableUpdateRequest{RouteTable: RouteTable})

	default:
		xapp.Logger.Info("Unknown Message Type '%d', discarding", msg.Mtype)
	}
	return
}
