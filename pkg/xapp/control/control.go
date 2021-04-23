/*
#cgo CFLAGS: -I../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
#cgo LDFLAGS: -L../../../internal/asn1codec/   -lasn1objects
#include <RIC_message_types.h>
*/
package control

import "C"
import (
	"context"
	"errors"
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"nRIC/internal"
	"nRIC/internal/kafka-go"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/endpoint"
	"nRIC/internal/msgx/service"
	"nRIC/internal/msgx/transport"
	"nRIC/internal/utils"
	"nRIC/internal/xapp"
	"nRIC/internal/xapp/golog"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Control struct {
	ranList               	[]string             //nodeB list
	eventCreateExpired    	int32                //maximum time for the RIC Subscription Request event creation procedure in the E2 Node
	eventDeleteExpired    	int32                //maximum time for the RIC Subscription Request event deletion procedure in the E2 Node
	rcChan                	chan *xapp.MsgParams //channel for receiving  message
	Msg2SubmgrClient        *msgx.MsgSender
	Msg2CflmgrClient        *msgx.MsgSender
	Msg2MgmtClient          *msgx.MsgSender
	eventCreateExpiredMap 	map[string]bool      //map for recording the RIC Subscription Request event creation procedure is expired or not
	eventDeleteExpiredMap 	map[string]bool      //map for recording the RIC Subscription Request event deletion procedure is expired or not
	eventCreateExpiredMu 	*sync.Mutex          //mutex for eventCreateExpiredMap
	eventDeleteExpiredMu  	*sync.Mutex          //mutex for eventDeleteExpiredMap
	Topic			  		string
	RegMsgServer      		*grpc.Server
	XappID					uint16
	Registered	            bool
}

func init() {
	file := "/app/cmd/xapp/xapp_ts.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	xapp.Logger.SetLevel(int(golog.ERR))
}

func NewControl(MsgClient2Submgr *msgx.MsgSender,MsgClient2Cflmgr *msgx.MsgSender) Control {
	//str := os.Getenv("ranList")
	str := "RAN1"
	return Control{strings.Split(str, ","),
		5, 5,
		make(chan *xapp.MsgParams),
		MsgClient2Submgr,
		MsgClient2Cflmgr,
		nil,
		make(map[string]bool),
		make(map[string]bool),
		&sync.Mutex{},
		&sync.Mutex{},
		"",
		nil,
		0,
		false}
}

func (c *Control)  IsRegistered() bool {
	return c.Registered
}

func (c *Control)  SetMsg2SubmgrClient (client *msgx.MsgSender) {
	c.Msg2SubmgrClient = client
}

func (c *Control)  SetMsg2CflmgrClient (client *msgx.MsgSender) {
	c.Msg2CflmgrClient = client
}

func ReadyCB(i interface{}) {
	//c := i.(*Control)

	//c.startTimerSubReq()
	//go c.controlLoop()
}


func (c *Control)  CreateAndStartMsgServer (grpcAddr string) {

	svc := service.NewMsgService(c)
	ep  := endpoint.NewMsgServiceEndpoint(svc)
	s   := transport.NewMsgServer(ep)


	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		xapp.Logger.Info("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	//xapp.Logger.Info("transport", "gRPC", "addr", grpcAddr)
	// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
	// the here demonstrated zipkin tracing middleware.
	c.RegMsgServer = grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	msgx.RegisterMsgServiceServer(c.RegMsgServer, s)
	c.RegMsgServer.Serve(grpcListener)
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		//GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1, //10e3, // 10KB   //cnbu
		MaxBytes: 10e6, // 10MB
		PartitionWatchInterval: 1,  //cnbu
	})
}

func ListTopic(conn *kafka.Conn)  {
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

//func  ReadMessage()  {
func (c *Control)  KafkaCreateAndStartMsgServer (kafkaURL string,topic string) {
	// to create topics when auto.create.topics.enable='true'
	var err error
	c.Topic = topic //os.Getenv("topic")
	if c.Topic == "" {
		c.Topic,err = os.Hostname()
		if err != nil {
			panic(err.Error())
		}
	}
	//fmt.Printf("KafkaCreateAndStartMsgServer topic Name:%s\n",c.Topic)
	//create topics
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaURL, c.Topic, 0)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	ListTopic(conn)

	//topic := os.Getenv("topic")
	groupID := "nric-group-id" //os.Getenv("groupID")

	reader := getKafkaReader(kafkaURL, c.Topic, groupID)

	defer reader.Close()

	fmt.Println("\nstart consuming ... !!\n")
	fmt.Printf("message at topic:%s\n",c.Topic)
	reader.SetOffsetAt(context.Background(),time.Now())

	for {

		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			xapp.Logger.Error(err.Error())
			return
		}
		/*
		fmt.Printf("message at topic:%v partition:%v offset:%v	%v = %v,\n %s = %s \n %s = %s \n",
			m.Topic, m.Partition, m.Offset, m.Key, m.Value,
			m.Headers[0].Key,m.Headers[0].Value,m.Headers[1].Key,m.Headers[1].Value)

		 */

		mtype,_ := strconv.Atoi(string(m.Headers[0].Value))
		payloadLen,_ := strconv.Atoi(string(m.Headers[1].Value))
		subId,_ := strconv.Atoi(string(m.Headers[2].Value))
		params := &xapp.MsgParams{
			Mtype: mtype,
			PayloadLen: payloadLen,
			SubId: subId,
			Meid:  &xapp.MsgMeid{RanName: string(m.Headers[3].Value)},
			Payload: m.Value,

		}
		/*
		fmt.Printf("message at topic:%v partition:%v offset:%v	%v = %v,\n %s = %v \n %s = %v \n",
			m.Topic, m.Partition, m.Offset, m.Key, m.Value,
			m.Headers[0].Key,mtype,m.Headers[1].Key,params.PayloadLen)

		 */

		//params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}

		c.Consume(params)
	}
}

func (c *Control) RunKafka(Addr string ,topic string)  {
	for {
		c.KafkaCreateAndStartMsgServer(Addr,topic)
		time.Sleep(2 * time.Second)
	}
}


func (c *Control) Run(grpcAddr string) {

	if len(c.ranList) > 0 {
		//xapp.SetReadyCB(ReadyCB, c)


		//go c.RunKafka(Addr,topic)

		//c.startTimerSubReq()
		go c.CreateAndStartMsgServer(grpcAddr)

		go c.controlLoop()

	} else {
		xapp.Logger.Error("gNodeB not set for subscription")
		log.Printf("gNodeB not set for subscription")
	}

}

func (c *Control) startTimerSubReq() {
	timerSR := time.NewTimer(5 * time.Second)
	count := 0

	go func(t *time.Timer) {
		defer timerSR.Stop()
		for {
			<-t.C
			count++
			xapp.Logger.Debug("send RIC_SUB_REQ to gNodeB with cnt=%d", count)
			log.Printf("send RIC_SUB_REQ to gNodeB with cnt=%d", count)
			err := c.sendRicSubRequest(1001, 1001, 0)
			if err != nil && count < MAX_SUBSCRIPTION_ATTEMPTS {
				t.Reset(5 * time.Second)
			} else {
				break
			}
		}
	}(timerSR)
}

func (c *Control) Consume(rp *xapp.MsgParams) (err error) {
	c.rcChan <- rp
	return
}


func (c *Control) controlLoop() {
	for {
		msg := <-c.rcChan
		xapp.Logger.Debug("Received message type: %s", xapp.RicMessageTypeToName[msg.Mtype])
		log.Printf("Received message type: %d", msg.Mtype)
		switch msg.Mtype {
		case xapp.RIC_O1_REGISTER_RESP:
			c.handleRicRegResp(msg)
		case xapp.RIC_O1_INIT:
			c.handleSMOInitMsg(msg)
		case xapp.RIC_O1_ENABLE:
			c.handleSMOEnable(msg)
		case xapp.RIC_O1_DISABLE:
			c.handleSMODisable(msg)

		case xapp.RIC_SUB_RESP:
			c.handleSubscriptionResponse(msg)
		case xapp.RIC_SUB_FAILURE:
			c.handleSubscriptionFailure(msg)
		case xapp.RIC_SUB_DEL_RESP:
			c.handleSubscriptionDeleteResponse(msg)
		case xapp.RIC_SUB_DEL_FAILURE:
			c.handleSubscriptionDeleteFailure(msg)
		case xapp.RIC_INDICATION:
			c.handleIndication(msg)
		default:
			err := errors.New("Message Type " + strconv.Itoa(msg.Mtype) + " is discarded")
			xapp.Logger.Error("Unknown message type: %v", err)
			log.Printf("Unknown message type: %v", err)
		}
	}
}

func IsContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//保存注册响应消息带回的RIC平台信息
var (
	NricSubmgrHost 		string
	NricSubmgrPort		string

	NricCflmgrHost 		string
	NricCflmgrPort		string

	NricDbagentHost		string
	NricDbagentPort		string
)

func (c *Control) handleRicRegResp(params *xapp.MsgParams) (err error) {
	M := msgx.XappRegResp{}
	err = proto.Unmarshal(params.Payload,&M)
	xapp.Logger.Info(M.String())
	if err != nil {
		xapp.Logger.Error("Unmarshal XappRegResp msg failed!\n")
		return err
	}
	c.Registered   = true  //注册成功

	NricSubmgrHost = M.RicServices["nricsubs"].IpAddr
	NricSubmgrPort = M.RicServices["nricsubs"].Port
	NricCflmgrHost = M.RicServices["nriccflm"].IpAddr
	NricCflmgrPort = M.RicServices["nriccflm"].Port
	NricDbagentHost = M.RicServices["nricdbagent"].IpAddr
	NricDbagentPort = M.RicServices["nricdbagent"].Port
	c.Topic	 		= M.Topic
	c.XappID		= uint16(M.Header.XappRequestID.XappID)

	//msg client
	MsgSender2Subsmgr := msgx.NewMsgSender(utils.GetValue(NricSubmgrHost,internal.SubmgrHost),internal.DefaultGRPCPort)
	MsgSender2Cflmgr := msgx.NewMsgSender(utils.GetValue(NricCflmgrHost,internal.NriccflmHost),internal.DefaultGRPCPort)
	c.SetMsg2SubmgrClient(MsgSender2Subsmgr)
	c.SetMsg2CflmgrClient(MsgSender2Cflmgr)

	go c.RunKafka(M.KafkaURL,M.Topic)

	c.startTimerSubReq()

	//注册完成，释放 Server 资源
	c.RegMsgServer.GracefulStop()
	c.RegMsgServer = nil

	return nil
}


func (c *Control) handleSMOInitMsg(params *xapp.MsgParams) (err error) {
	M := msgx.SMOInitMsg{}
	err = proto.Unmarshal(params.Payload,&M)
	if err != nil {
		xapp.Logger.Error("Unmarshal SMOInitMsg failed!\n")
		return err
	}
	xapp.Logger.Info("XappID = %d,XappInstanceID = %d,InitConfig = %v\n",
		M.Header.XappRequestID.XappID,M.Header.XappRequestID.XappInstanceID,M.InitConfig)
	return nil
}

// Enable xapp by SMO through nRIC
func (c *Control) handleSMOEnable(params *xapp.MsgParams) (err error) {
	M := msgx.SMOEnableMsg{}
	err = proto.Unmarshal(params.Payload,&M)
	if err != nil {
		xapp.Logger.Error("Unmarshal SMOEnableMsg failed!\n")
		return err
	}
	xapp.Logger.Info("XappID = %d,XappInstanceID = %d,MsgType = %s\n",
		M.Header.XappRequestID.XappID,M.Header.XappRequestID.XappInstanceID, xapp.RicMessageTypeToName[int(M.Header.MsgType)])

	// TBD: send heartbeat msg to nric mgmt
	//c.Msg2MgmtClient.SendMsg()

	return nil
}


// Disable xapp by SMO through nRIC
func (c *Control) handleSMODisable(params *xapp.MsgParams) (err error) {
	M := msgx.SMOEnableMsg{}
	err = proto.Unmarshal(params.Payload,&M)
	if err != nil {
		xapp.Logger.Error("Unmarshal SMODisableMsg failed!\n")
		return err
	}
	xapp.Logger.Info("XappID = %d,XappInstanceID = %d,MsgType = %s\n",
		M.Header.XappRequestID.XappID,M.Header.XappRequestID.XappInstanceID, xapp.RicMessageTypeToName[int(M.Header.MsgType)])

	// TBD: stop heartbeat msg
	return nil
}


func (c *Control) handleIndication(params *xapp.MsgParams) (err error) {
	var e2ap *E2ap
	var e2sm *E2sm
	log.Printf("RIC Indication message from {%s} received", params.Meid.RanName)

	indicationMsg, err := e2ap.GetIndicationMessage(params.Payload)
	if err != nil {
		xapp.Logger.Error("Failed to decode RIC Indication message: %v", err)
		log.Printf("Failed to decode RIC Indication message: %v", err)
		return
	}

	log.Printf("RIC Indication message from {%s} received", params.Meid.RanName)
	log.Printf("RequestID: %d", indicationMsg.RequestID)
	log.Printf("RequestSequenceNumber: %d", indicationMsg.RequestSequenceNumber)
	log.Printf("FunctionID: %d", indicationMsg.FuncID)
	log.Printf("ActionID: %d", indicationMsg.ActionID)
	log.Printf("IndicationSN: %d", indicationMsg.IndSN)
	log.Printf("IndicationType: %d", indicationMsg.IndType)
	log.Printf("IndicationHeader: %x", indicationMsg.IndHeader)
	log.Printf("IndicationMessage: %x", indicationMsg.IndMessage)
	log.Printf("CallProcessID: %x", indicationMsg.CallProcessID)

	 _, err = e2sm.GetIndicationHeader(indicationMsg.IndHeader)
	if err != nil {
		xapp.Logger.Error("Failed to decode RIC Indication Header: %v", err)
		log.Printf("Failed to decode RIC Indication Header: %v", err)
		return
	}

	indMsg, err := e2sm.GetIndicationMessage(indicationMsg.IndMessage)
	if err != nil {
		xapp.Logger.Error("Failed to decode RIC Indication Message: %v", err)
		log.Printf("Failed to decode RIC Indication Message: %v", err)
		return
	}


	log.Printf("-----------RIC Indication Message-----------")
	log.Printf("IndMsgType: %d", indMsg.IndMsgType)
	GnoteB.Lock.Lock()
	defer GnoteB.Lock.Unlock()

	if indMsg.IndMsgType == 1 {
		IndMsg := indMsg.IndMsg.(*IndicationMessageFormat1)
		for i:=0; i < IndMsg.ContainerCount;i++ {
			v := IndMsg.o_Cu_CpCellMeasurement[i]
			Cellid, err := GetCellid(v.NRCellid)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				//是一个新的cell，从 0 开始分配id
				Cellid = GnoteB.CellNum
				GnoteB.CellNum++
				SelectUE := make(map[int] []int)
				for i:=0;i<127;i++{
					SelectUE[i] = []int{}
				}
				GnoteB.Cells[Cellid] = &Cell{UEs: make(map[int] *UE ),SelectUE: SelectUE}
				GnoteB.Cells[Cellid].Cellid = Cellid
				GnoteB.Cells[Cellid].NRCellID = v.NRCellid
				GnoteB.Cells[Cellid].NRCGI = v.NRCGI
				GnoteB.Cells[Cellid].NumberOfRrcConnections = v.NumberOfRrcConnections
				GnoteB.Cells[Cellid].NumberOfSupportedRrcConnections = v.NumberOfSupportedRrcConnections
				GnoteB.Cells[Cellid].Load = v.NumberOfRrcConnections * 100 / v.NumberOfSupportedRrcConnections
			} else {
				//是一个老的cell，仅需更新必要的信息
				GnoteB.Cells[Cellid].NumberOfRrcConnections = v.NumberOfRrcConnections
				GnoteB.Cells[Cellid].NumberOfSupportedRrcConnections = v.NumberOfSupportedRrcConnections
				GnoteB.Cells[Cellid].Load = v.NumberOfRrcConnections * 100 / v.NumberOfSupportedRrcConnections
			}
			if (GnoteB.Cells[Cellid].Load < 0 || GnoteB.Cells[Cellid].Load > 100) {
				xapp.Logger.Error("v.NumberOfRrcConnections = %d,v.NumberOfSupportedRrcConnections = %d, Load = %d\n",
					v.NumberOfRrcConnections,v.NumberOfSupportedRrcConnections, GnoteB.Cells[Cellid].Load)
			}
			//xapp.Logger.Debug("cell info : Cellid = %d, Load = %d\n",Cellid, GnoteB.Cells[Cellid].Load)
		}
	} else if (indMsg.IndMsgType == 4 ) {
		IndMsg := indMsg.IndMsg.(*IndicationMessageFormat4)
		for i := 0; i < IndMsg.ContainerCount; i++ {
			v := IndMsg.o_Cu_CpUeMeasurement[i]
			Cellid, err := GetCellid(v.NRCellid)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
				//是一个新的cell，从 0 开始分配id
				Cellid = GnoteB.CellNum
				GnoteB.CellNum++
				SelectUE := make(map[int] []int)
				for i:=0;i<127;i++{
					SelectUE[i] = []int{}
				}
				GnoteB.Cells[Cellid] = &Cell{UEs: make(map[int] *UE ),SelectUE: SelectUE}
				GnoteB.Cells[Cellid].Cellid = Cellid
				GnoteB.Cells[Cellid].NRCellID = v.NRCellid
				GnoteB.Cells[Cellid].NRCGI = v.NRCGI
			}

			//更新UE信息
			c := GnoteB.Cells[Cellid]
			//c.Lock.Lock()
			rSRP := v.MResult.measurementResultServingCell.basedSSB.rSRP
			c.UEs[v.UEId] = &UE{
				UEId: v.UEId,   //c_RNTI
				SrcCellID: Cellid,
				MResult: MeasurementResult{
										measurementResultServingCell: MeasurementResultServingCell{
												basedSSB: BasedSSB{
														rSRP: rSRP,
													},
										},
				},
			}
			if !IsContain(c.SelectUE[rSRP],v.UEId){
				c.SelectUE[rSRP] = append(c.SelectUE[rSRP],v.UEId)
			}
			//c.Lock.Unlock()

			//fmt.Printf("cell info : %v\n", c)
		}
	}
	return nil
}


func (c *Control) handleSubscriptionResponse(params *xapp.MsgParams) (err error) {
	xapp.Logger.Debug("The SubId in RIC_SUB_RESP is %d", params.SubId)
	log.Printf("The SubId in RIC_SUB_RESP is %d", params.SubId)

	ranName := "gnb_734_733_b5c67788" //params.Meid.RanName
	c.eventCreateExpiredMu.Lock()
	_, ok := c.eventCreateExpiredMap[ranName]
	if !ok {
		c.eventCreateExpiredMu.Unlock()
		xapp.Logger.Debug("RIC_SUB_REQ has been deleted!")
		log.Printf("RIC_SUB_REQ has been deleted!")
		return nil
	} else {
		c.eventCreateExpiredMap[ranName] = true
		c.eventCreateExpiredMu.Unlock()
	}

	var cep *E2ap
	subscriptionResp, err := cep.GetSubscriptionResponseMessage(params.Payload)
	if err != nil {
		xapp.Logger.Error("Failed to decode RIC Subscription Response message: %v", err)
		log.Printf("Failed to decode RIC Subscription Response message: %v", err)
		return
	}

	//log.Printf("RIC Subscription Response message from {%s} received", params.Meid.RanName)
	log.Printf("SubscriptionID: %d", params.SubId)
	log.Printf("RequestID: %d", subscriptionResp.RequestID)
	log.Printf("RequestSequenceNumber: %d", subscriptionResp.RequestSequenceNumber)
	log.Printf("FunctionID: %d", subscriptionResp.FuncID)

	log.Printf("ActionAdmittedList:")
	for index := 0; index < subscriptionResp.ActionAdmittedList.Count; index++ {
		log.Printf("[%d]ActionID: %d", index, subscriptionResp.ActionAdmittedList.ActionID[index])
	}

	log.Printf("ActionNotAdmittedList:")
	for index := 0; index < subscriptionResp.ActionNotAdmittedList.Count; index++ {
		log.Printf("[%d]ActionID: %d", index, subscriptionResp.ActionNotAdmittedList.ActionID[index])
		log.Printf("[%d]CauseType: %d    CauseID: %d", index, subscriptionResp.ActionNotAdmittedList.Cause[index].CauseType, subscriptionResp.ActionNotAdmittedList.Cause[index].CauseID)
	}

	return nil
}

func (c *Control) handleSubscriptionFailure(params *xapp.MsgParams) (err error) {
	xapp.Logger.Debug("The SubId in RIC_SUB_FAILURE is %d", params.SubId)
	log.Printf("The SubId in RIC_SUB_FAILURE is %d", params.SubId)

	ranName := params.Meid.RanName
	c.eventCreateExpiredMu.Lock()
	_, ok := c.eventCreateExpiredMap[ranName]
	if !ok {
		c.eventCreateExpiredMu.Unlock()
		xapp.Logger.Debug("RIC_SUB_REQ has been deleted!")
		log.Printf("RIC_SUB_REQ has been deleted!")
		return nil
	} else {
		c.eventCreateExpiredMap[ranName] = true
		c.eventCreateExpiredMu.Unlock()
	}

	return nil
}

func (c *Control) handleSubscriptionDeleteResponse(params *xapp.MsgParams) (err error) {
	xapp.Logger.Debug("The SubId in RIC_SUB_DEL_RESP is %d", params.SubId)
	log.Printf("The SubId in RIC_SUB_DEL_RESP is %d", params.SubId)

	ranName := params.Meid.RanName
	c.eventDeleteExpiredMu.Lock()
	_, ok := c.eventDeleteExpiredMap[ranName]
	if !ok {
		c.eventDeleteExpiredMu.Unlock()
		xapp.Logger.Debug("RIC_SUB_DEL_REQ has been deleted!")
		log.Printf("RIC_SUB_DEL_REQ has been deleted!")
		return nil
	} else {
		c.eventDeleteExpiredMap[ranName] = true
		c.eventDeleteExpiredMu.Unlock()
	}

	return nil
}

func (c *Control) handleSubscriptionDeleteFailure(params *xapp.MsgParams) (err error) {
	xapp.Logger.Debug("The SubId in RIC_SUB_DEL_FAILURE is %d", params.SubId)
	log.Printf("The SubId in RIC_SUB_DEL_FAILURE is %d", params.SubId)

	ranName := params.Meid.RanName
	c.eventDeleteExpiredMu.Lock()
	_, ok := c.eventDeleteExpiredMap[ranName]
	if !ok {
		c.eventDeleteExpiredMu.Unlock()
		xapp.Logger.Debug("RIC_SUB_DEL_REQ has been deleted!")
		log.Printf("RIC_SUB_DEL_REQ has been deleted!")
		return nil
	} else {
		c.eventDeleteExpiredMap[ranName] = true
		c.eventDeleteExpiredMu.Unlock()
	}

	return nil
}

func (c *Control) setEventCreateExpiredTimer(ranName string) {
	c.eventCreateExpiredMu.Lock()
	c.eventCreateExpiredMap[ranName] = false
	c.eventCreateExpiredMu.Unlock()

	timer := time.NewTimer(time.Duration(c.eventCreateExpired) * time.Second)
	go func(t *time.Timer) {
		defer t.Stop()
		xapp.Logger.Debug("RIC_SUB_REQ[%s]: Waiting for RIC_SUB_RESP...", ranName)
		log.Printf("RIC_SUB_REQ[%s]: Waiting for RIC_SUB_RESP...", ranName)
		for {
			select {
			case <-t.C:
				c.eventCreateExpiredMu.Lock()
				isResponsed := c.eventCreateExpiredMap[ranName]
				delete(c.eventCreateExpiredMap, ranName)
				c.eventCreateExpiredMu.Unlock()
				if !isResponsed {
					xapp.Logger.Debug("RIC_SUB_REQ[%s]: RIC Event Create Timer experied!", ranName)
					log.Printf("RIC_SUB_REQ[%s]: RIC Event Create Timer experied!", ranName)
					// c.sendRicSubDelRequest(subID, requestSN, funcID)
					return
				}
			default:
				c.eventCreateExpiredMu.Lock()
				flag := c.eventCreateExpiredMap[ranName]
				if flag {
					delete(c.eventCreateExpiredMap, ranName)
					c.eventCreateExpiredMu.Unlock()
					xapp.Logger.Debug("RIC_SUB_REQ[%s]: RIC Event Create Timer canceled!", ranName)
					log.Printf("RIC_SUB_REQ[%s]: RIC Event Create Timer canceled!", ranName)
					return
				} else {
					c.eventCreateExpiredMu.Unlock()
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(timer)
}

func (c *Control) setEventDeleteExpiredTimer(ranName string) {
	c.eventDeleteExpiredMu.Lock()
	c.eventDeleteExpiredMap[ranName] = false
	c.eventDeleteExpiredMu.Unlock()

	timer := time.NewTimer(time.Duration(c.eventDeleteExpired) * time.Second)
	go func(t *time.Timer) {
		defer t.Stop()
		xapp.Logger.Debug("RIC_SUB_DEL_REQ[%s]: Waiting for RIC_SUB_DEL_RESP...", ranName)
		log.Printf("RIC_SUB_DEL_REQ[%s]: Waiting for RIC_SUB_DEL_RESP...", ranName)
		for {
			select {
			case <-t.C:
				c.eventDeleteExpiredMu.Lock()
				isResponsed := c.eventDeleteExpiredMap[ranName]
				delete(c.eventDeleteExpiredMap, ranName)
				c.eventDeleteExpiredMu.Unlock()
				if !isResponsed {
					xapp.Logger.Debug("RIC_SUB_DEL_REQ[%s]: RIC Event Delete Timer experied!", ranName)
					log.Printf("RIC_SUB_DEL_REQ[%s]: RIC Event Delete Timer experied!", ranName)
					return
				}
			default:
				c.eventDeleteExpiredMu.Lock()
				flag := c.eventDeleteExpiredMap[ranName]
				if flag {
					delete(c.eventDeleteExpiredMap, ranName)
					c.eventDeleteExpiredMu.Unlock()
					xapp.Logger.Debug("RIC_SUB_DEL_REQ[%s]: RIC Event Delete Timer canceled!", ranName)
					log.Printf("RIC_SUB_DEL_REQ[%s]: RIC Event Delete Timer canceled!", ranName)
					return
				} else {
					c.eventDeleteExpiredMu.Unlock()
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}(timer)
}

func (c *Control) sendRicSubRequest(subID int, requestSN int, funcID int) (err error) {
	var e2ap *E2ap
	var e2sm *E2sm

	var eventTriggerCount int = 1
	var periods []int64 = []int64{13}
	var eventTriggerDefinition []byte = make([]byte, 8)
	_, err = e2sm.SetEventTriggerDefinition(eventTriggerDefinition, eventTriggerCount, periods)
	if err != nil {
		xapp.Logger.Error("Failed to send RIC_SUB_REQ: %v", err)
		log.Printf("Failed to send RIC_SUB_REQ: %v", err)
		return err
	}
	log.Printf("Set EventTriggerDefinition: %x", eventTriggerDefinition)

	var actionCount int = 1
	var ricStyleType []int64 = []int64{0}
	var actionIds []int64 = []int64{0}
	var actionTypes []int64 = []int64{0}
	var actionDefinitions []ActionDefinition = make([]ActionDefinition, actionCount)
	var subsequentActions []SubsequentAction = []SubsequentAction{SubsequentAction{0, 0, 0}}

	for index := 0; index < actionCount; index++ {
		if ricStyleType[index] == 0 {
			actionDefinitions[index].Buf = nil
			actionDefinitions[index].Size = 0
		} else {
			actionDefinitions[index].Buf = make([]byte, 8)
			_, err = e2sm.SetActionDefinition(actionDefinitions[index].Buf, ricStyleType[index])
			if err != nil {
				xapp.Logger.Error("Failed to send RIC_SUB_REQ: %v", err)
				log.Printf("Failed to send RIC_SUB_REQ: %v", err)
				return err
			}
			actionDefinitions[index].Size = len(actionDefinitions[index].Buf)

			log.Printf("Set ActionDefinition[%d]: %x", index, actionDefinitions[index].Buf)
		}
	}

	for index := 0; index < 1; index++ { //len(c.ranList)
		params := &xapp.MsgParams{}
		params.Mtype = 12010
		params.SubId = subID
		params.Src = c.Topic

		//xapp.Logger.Debug("Send RIC_SUB_REQ to {%s}", c.ranList[index])
		//log.Printf("Send RIC_SUB_REQ to {%s}", c.ranList[index])

		params.Payload = make([]byte, 1024)
		params.Payload, err = e2ap.SetSubscriptionRequestPayload(params.Payload, 1001, uint16(requestSN), uint16(funcID), eventTriggerDefinition, len(eventTriggerDefinition), actionCount, actionIds, actionTypes, actionDefinitions, subsequentActions)
		if err != nil {
			xapp.Logger.Error("Failed to send RIC_SUB_REQ: %v", err)
			log.Printf("Failed to send RIC_SUB_REQ: %v", err)
			return err
		}

		log.Printf("Set Payload: %x", params.Payload)

		//params.Meid = &xapp.MsgMeid{RanName: c.ranList[index]}
		params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}
		xapp.Logger.Debug("The Msg message to be sent is %d with SubId=%d", params.Mtype, params.SubId)
		log.Printf("The Msg message to be sent is %d with SubId=%d", params.Mtype, params.SubId)

		err = c.Msg2SubmgrClient.SendMsg(params)
		if err != nil {
			xapp.Logger.Error("Failed to send RIC_SUB_REQ: %v", err)
			log.Printf("Failed to send RIC_SUB_REQ: %v", err)
			return err
		}

		c.setEventCreateExpiredTimer(params.Meid.RanName)
		//c.ranList = append(c.ranList[:index], c.ranList[index+1:]...)
		//index--
	}

	return nil
}

func (c *Control) sendRicSubDelRequest(subID int, requestSN int, funcID int) (err error) {
	params := &xapp.MsgParams{}
	params.Mtype = 12020
	params.SubId = subID
	var e2ap *E2ap

	params.Payload = make([]byte, 1024)
	params.Payload, err = e2ap.SetSubscriptionDeleteRequestPayload(params.Payload, 100, uint16(requestSN), uint16(funcID))
	if err != nil {
		xapp.Logger.Error("Failed to send RIC_SUB_DEL_REQ: %v", err)
		return err
	}

	log.Printf("Set Payload: %x", params.Payload)

	if funcID == 0 {
		//params.Meid = &xapp.MsgMeid{PlmnID: "::", EnbID: "::", RanName: "0"}
		params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}
	} else {
		//params.Meid = &xapp.MsgMeid{PlmnID: "::", EnbID: "::", RanName: "3"}
		params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}
	}

	xapp.Logger.Debug("The Msg message to be sent is %d with SubId=%d", params.Mtype, params.SubId)
	log.Printf("The Msg message to be sent is %d with SubId=%d", params.Mtype, params.SubId)


	c.setEventDeleteExpiredTimer(params.Meid.RanName)

	return nil
}

//actionType :ReleaseUE  ,  HandOverUE

func (c *Control) SendRicControlRequest(ricRequestorID uint16, requestSN uint16, funcID uint16,ReleaseUEs map[int]*UE, HandOverUEs map[int]*UE) (err error) {
	fmt.Printf("==================SendRicControlRequest ===================\n")

	InstructRelease += len(ReleaseUEs)
	InstructHandOver += len(HandOverUEs)
	fmt.Printf("Release:%d, HandOver:%d \n", InstructRelease, InstructHandOver)

	var e2ap *E2ap
	var e2sm *E2sm
	var ControlHeader []byte = make([]byte, 1024)
	var ControlMessageFormat6 []byte = make([]byte, 4096)

	ControlHeaderNew,err  := e2sm.SetControlHeader(ControlHeader,13)
	if err != nil {
		xapp.Logger.Error("Failed to SetControlHeader: %v", err)
		log.Printf("Failed to SetControlHeader: %v", err)
		return err
	}

	ControlMessageFormat6New, err := e2sm.SetControlMessageFormat6(ControlMessageFormat6, ReleaseUEs, HandOverUEs)
	if err != nil {
		xapp.Logger.Error("Failed to SetControlMessageFormat6: %v", err)
		log.Printf("Failed to SetControlMessageFormat6: %v", err)
		return err
	}
	log.Printf("Set ControlMessageFormat6new: %x", ControlMessageFormat6New)
	log.Printf("Set ControlMessageFormat6   : %x", ControlMessageFormat6)


	for index := 0; index < 1; index++ { //len(c.ranList)
		params := &xapp.MsgParams{}
		params.Mtype = xapp.RIC_CONTROL_REQ
		params.SubId = int(ricRequestorID)
		params.Src = c.Topic

		//xapp.Logger.Debug("Send RIC_SUB_REQ to {%s}", c.ranList[index])
		//log.Printf("Send RIC_SUB_REQ to {%s}", c.ranList[index])
		//fmt.Printf("Len of ControlHeaderNew = %d\n",len(ControlHeaderNew))
		//fmt.Printf("Len of ControlMessageFormat6New = %d\n",len(ControlMessageFormat6New))

		params.Payload = make([]byte, 8192)
		params.Payload, err = e2ap.SetControlRequestPayload(params.Payload, c.XappID, uint16(requestSN), uint16(funcID),
			ControlHeader,C.long(len(ControlHeaderNew)),ControlMessageFormat6,C.long(len(ControlMessageFormat6New)))
		if err != nil {
			xapp.Logger.Error("Failed to send RIC_CONTROL_REQ: %v", err)
			log.Printf("Failed to send RIC_CONTROL_REQ: %v", err)
			return err
		}


		log.Printf("Set Payload: %x\n", params.Payload)
		//fmt.Printf("Len of params.Payload = %d\n",len(params.Payload))
		params.PayloadLen = len(params.Payload)

		//params.Meid = &xapp.MsgMeid{RanName: c.ranList[index]}
		params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}
		xapp.Logger.Debug("The Msg message to be sent is %d with SubId=%d", params.Mtype, params.SubId)
		log.Printf("The Msg message to be sent is %s with SubId=%d", xapp.RicMessageTypeToName[params.Mtype], params.SubId)

		err = c.Msg2CflmgrClient.SendMsg(params)
		if err != nil {
			xapp.Logger.Error("Failed to send RIC_CONTROL_REQ: %v", err)
			log.Printf("Failed to send RIC_CONTROL_REQ: %v", err)
			return err
		}

		//c.setEventCreateExpiredTimer(params.Meid.RanName)

		//c.ranList = append(c.ranList[:index], c.ranList[index+1:]...)
		//index--
	}

	return nil
}