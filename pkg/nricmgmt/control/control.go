
package control

import (
	"context"
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"nRIC/api/v1/pb/db"
	"nRIC/internal"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/endpoint"
	"nRIC/internal/msgx/service"
	"nRIC/internal/msgx/transport"
	"nRIC/internal/xapp"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"net"
	"os"
	"sort"
	"strconv"
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
	MsgSendertoSMO *msgx.MsgSender //*msgx.MsgSender
	CntRecvMsg uint64
	AccessDbagent *dbclient.MsgSender
	Endpoint  map[uint32]*msgx.KafkaMsgSender   //key: XappID ,value : kafka writer
}

type MsgMeid struct {
	PlmnID  string
	EnbID   string
	RanName string
}

func init() {
}

func NewControl(MsgSendertoSMO *msgx.MsgSender ,AcessDbAgent *dbclient.MsgSender) *Control {

	endpoint := make(map[uint32]*msgx.KafkaMsgSender)

	c := &Control{
		//MsgClientToXapp: MsgClientToXapp,
		MsgSendertoSMO:	MsgSendertoSMO,
		//subscriber: subscriber,
		AccessDbagent: AcessDbAgent,
		Endpoint: endpoint,
	}
	return c
}

func (c *Control) ReadyCB(data interface{}) {
	if c.MsgSendertoSMO == nil {
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

	if c.MsgSendertoSMO == nil {
		err = fmt.Errorf("Msg object nil can handle %s", msg.String())
		xapp.Logger.Error("%s", err.Error())
		return
	}
	c.CntRecvMsg++

	switch msg.Mtype {

	case xapp.RIC_O1_REGISTER:
		go c.handleXappRegisterRequest(msg)

	case xapp.RIC_O1_INIT:
		go c.handleSMOInit(msg)

	case xapp.RIC_O1_ENABLE:
		go c.handleSMOEnableOrDisable(msg)

	case xapp.RIC_O1_DISABLE:
		go c.handleSMOEnableOrDisable(msg)

	default:
		xapp.Logger.Info("Unknown Message Type '%d', discarding", msg.Mtype)
	}
	return
}

//分配xappID
//XappID 合法值 （1...65535），0 为 非法值
func (c *Control) allocXappID(RegMsg msgx.XappRegMsg) (uint32,error, bool){
	isRegistered := false
	resp,err := c.AccessDbagent.Client.MOITableReadAll(context.Background(),&db.MOITableReadAllRequest{Api: "1",})
	if err != nil {
		xapp.Logger.Error(err.Error())
		return 0 ,err,isRegistered
	}
	//第一个注册的xapp，直接分配 xappID = 1
	if len(resp.MoiTables) == 0 {
		return 1,nil,isRegistered
	}

	for _,r := range resp.MoiTables{
		//重复注册，走正常返回流程
		if r.XappName == RegMsg.XappName {
			isRegistered = true
			return r.XappID,nil,isRegistered
		}
	}
	//新注册 ，先按从小到大排序(只有1个表项也可以排序，不会返回错误)，然后找到最小的可用xappID
	ps := resp.MoiTables
	sort.SliceStable(ps, func(i, j int) bool {
		return ps[i].XappID < ps[j].XappID
	})

	var  i uint32
	//正常的分配为: ps[0].XappID = 1 ,ps[1].XappID = 2 ,ps[2].XappID = 3 ,
	//如果出现ps[1].XappID = 3 ,说明原来的XappID = 2 表项已经释放，可以被再次分配使用
	for  i = 1; i <= 65535;i++ {
		if i != ps[i-1].XappID {
			// i 值未被使用，可以被分配
			return i,nil,isRegistered
		}
	}
	return 0 , fmt.Errorf("分配失败"),isRegistered
}

//新增该xApp的管理对象实例（xApp MOI）表项
func (c *Control) addXappMOI(XappID uint32,RegMsg msgx.XappRegMsg) error {
	r := &db.MOITableInsertRequest{}
	r.Api = "1"
	m := &db.MOITable{}
	m.XappID 		= XappID
	m.XappName		= RegMsg.XappName
	m.XappVer		= RegMsg.XappVer
	m.Functions	    = RegMsg.XappFunctions
	m.RunningStatus = "inactive"
	m.IsReady		= "false"
	m.Topic			= "Xapp_"+strconv.Itoa(int(XappID))+"_topic"
	r.MoiTable 		= m
	_, err := c.AccessDbagent.Client.MOITableInsert(context.Background(),r)
	return err
}

//通知网管该xApp在nRT RIC平台上的部署
func (c *Control) Register2SMO (RegMsg *msgx.XappRegMsg,params *xapp.MsgParams){
	RICO1RegMsg,err := proto.Marshal(RegMsg)
	if err != nil {
		xapp.Logger.Error("Marshal RICO1RegMsg failed! %s",err.Error())
		return
	}
	params.Payload 		= RICO1RegMsg
	params.PayloadLen 	= len(RICO1RegMsg)
	//确保消息发送成功，否则每隔 5 秒再次发送
	for {
		err := c.MsgSendertoSMO.SendMsg(params)
		if err == nil {
			break
		}
		xapp.Logger.Error("Register2SMO:",err.Error())
		time.Sleep( 5 * time.Second)
	}
}
//-------------------------------------------------------------------
// handle from XAPP Register Request
//------------------------------------------------------------------
func (c *Control) handleXappRegisterRequest(params *xapp.MsgParams) {
	xapp.Logger.Info("Register MSG from XAPP: %s", params.String())
	var RegMsg msgx.XappRegMsg
	err := proto.Unmarshal(params.Payload,&RegMsg)
	if err != nil {
		xapp.Logger.Error("Unmarshal XappRegMsg failed! %s",err.Error())
		//（解析不到xapp的ip，无法返回响应消息）xapp 接收不到注册成功响应消息，会继续发起注册
		return
	}
	xapp.Logger.Info("XappName = %s,XappRequestID = %d,Token = %s /n",
		RegMsg.XappName,RegMsg.Header.XappRequestID,RegMsg.Header.Token)

	//第一个消息，xapp还没获取到topic，需要通过grpc来返回注册响应消息
	Client2Xapp := msgx.NewMsgSender(RegMsg.XappIpaddr,RegMsg.XappPort)

	//分配xappID ; 并判断是否重复注册
	XappID,err,isRegistered := c.allocXappID(RegMsg)
	if err != nil {
		xapp.Logger.Error("Alloc XappID failed! %s",err.Error())

		return
	}

	//非重复注册，新增该xApp的管理对象实例（xApp MOI）表项
	if !isRegistered {
		err = c.addXappMOI(XappID ,RegMsg)
		if err != nil {
			xapp.Logger.Error("Add Xapp MOI failed! %s",err.Error())
			return
		}
	}

	//通知网管该xApp在nRT RIC平台上的部署
	RegMsg.Header.XappRequestID.XappID = XappID
	go c.Register2SMO(&RegMsg,params)

	//response: 携带xApp所需服务（如数据库、冲突解决功能）的信息（服务名称、版本、详细信息等）
	Topic	 := "Xapp_"+strconv.Itoa(int(XappID))+"_topic"

	//除第一个RegisterResp消息外，第二个及以后的消息返回xapp，都通过xapp专有的kafka消息通道返回
	Endpoint := msgx.NewKafkaMsgSender(Topic)
	c.Endpoint[XappID] = Endpoint

	//
	var RicServices  map [string]*msgx.RICService
	RicServices = make(map[string]*msgx.RICService)
	RicSubsmgr := msgx.RICService{Name:"nricsubs",ServiceVer: 1,IpAddr: internal.SubmgrHost,Port: internal.DefaultGRPCPort}
	RicServices["nricsubs"] = &RicSubsmgr

	RicCflmgr  := msgx.RICService{Name:"nriccflm",ServiceVer: 1,IpAddr: internal.NriccflmHost,Port: internal.DefaultGRPCPort}
	RicServices["nriccflm"] = &RicCflmgr

	RicDbagent  := msgx.RICService{Name:"nricdbagent",ServiceVer: 1,IpAddr: internal.DbagentHost,Port: internal.DefaultGRPCPort}
	RicServices["nricdbagent"] = &RicDbagent

	XappRegResp := msgx.XappRegResp{
		Header: &msgx.RICMsgHeader{
			MsgType: xapp.RIC_O1_REGISTER_RESP,
			MsgVer: 1,
			XappRequestID: &msgx.XAPPRequestID{
				XappID: XappID,   //返回分配的XappID
				XappInstanceID: RegMsg.Header.XappRequestID.XappInstanceID,
			},
		},
		RicServices: RicServices,
		Topic:Topic,
		KafkaURL: internal.KafkaURL,
	}


	pbXappRegResp,err := proto.Marshal(&XappRegResp)
	if err != nil {
		xapp.Logger.Error("Marshal XappRegResp failed! %s",err.Error())
		// 释放MOI表项，释放XappID
		c.AccessDbagent.Client.MOITableDelete(context.Background(),&db.MOITableDeleteRequest{XappID: XappID,Api: "1"})
		return
	}
	params.Mtype = xapp.RIC_O1_REGISTER_RESP
	params.Payload 		= pbXappRegResp
	params.PayloadLen 	= len(pbXappRegResp)


	err = Client2Xapp.SendMsg(params)
	if err != nil {
		xapp.Logger.Error("Send  RIC_O1_REGISTER_RESP to Xapp failed! %s",err.Error())
		// 释放MOI表项，释放XappID
		c.AccessDbagent.Client.MOITableDelete(context.Background(),&db.MOITableDeleteRequest{XappID: XappID,Api: "1"})
		return
	}
}

//-------------------------------------------------------------------
// handle from SMO Init Request
//------------------------------------------------------------------
func (c *Control) SendRegisterFailureResp(Client2Xapp *msgx.MsgSender,Cause string,params *xapp.MsgParams) {
	xapp.Logger.Info("Send RegisterFailureResp Msg to Xapp: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	XappRegResp := msgx.XappRegResp{
		Header: &msgx.RICMsgHeader{
			MsgType: xapp.RIC_O1_REGISTER_FAILURE,
			MsgVer: 1,
		},
		Cause: Cause,
	}
	pbXappRegResp,err := proto.Marshal(&XappRegResp)
	if err != nil {
		xapp.Logger.Error("Marshal XappRegResp failed! %s",err.Error())
		// 释放MOI表项，释放XappID
		c.AccessDbagent.Client.MOITableDelete(context.Background(),&db.MOITableDeleteRequest{XappID: XappID,Api: "1"})
		return
	}
	params.Mtype = xapp.RIC_O1_REGISTER_RESP
	params.Payload 		= pbXappRegResp
	params.PayloadLen 	= len(pbXappRegResp)


	err = Client2Xapp.SendMsg(params)
	if err != nil {
		xapp.Logger.Error("Send  RIC_O1_REGISTER_RESP to Xapp failed! %s",err.Error())
		// 释放MOI表项，释放XappID
		c.AccessDbagent.Client.MOITableDelete(context.Background(),&db.MOITableDeleteRequest{XappID: XappID,Api: "1"})
		return
	}

}


//-------------------------------------------------------------------
// handle from SMO Init Request
//------------------------------------------------------------------
func (c *Control) handleSMOInit(params *xapp.MsgParams) {
	xapp.Logger.Info("Recv Msg From SMO: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	M := &msgx.SMOInitMsg{}
	err := proto.Unmarshal(params.Payload,M)
	if err != nil {
		xapp.Logger.Error(err.Error())
		return
	}
	//send Init msg to Xapp
	if e, ok := c.Endpoint[M.Header.XappRequestID.XappID]; ok {
		err = e.SendMsg(params)
		if err != nil {
			xapp.Logger.Error(err.Error())
			return
		}
	}else{
		xapp.Logger.Error("Endpoint is nil ,M.Header.XappRequestID.XappID = %d",M.Header.XappRequestID.XappID)
		return
	}
}


//-------------------------------------------------------------------
// handle from SMO Enable or Disable Xapp Request
//------------------------------------------------------------------
func (c *Control) handleSMOEnableOrDisable(params *xapp.MsgParams) {
	xapp.Logger.Info("Recv Msg From SMO: %s\n",xapp.RicMessageTypeToName[params.Mtype])
	M := &msgx.SMOEnableMsg{}
	err := proto.Unmarshal(params.Payload,M)
	if err != nil {
		xapp.Logger.Error(err.Error())
		return
	}

	//send Enable or Disable msg to Xapp
	if e, ok := c.Endpoint[M.Header.XappRequestID.XappID]; ok {
		err = e.SendMsg(params)
		if err != nil {
			xapp.Logger.Error(err.Error())
			return
		}
	}else{
		xapp.Logger.Error("Endpoint is nil ,M.Header.XappRequestID.XappID = %d",M.Header.XappRequestID.XappID)
		return
	}
}
