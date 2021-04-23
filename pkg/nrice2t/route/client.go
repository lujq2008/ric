package e2tclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"nRIC/internal/xapp"
	"time"
	pbe2t "nRIC/api/v1/pb/nrice2t"
)

type MsgSender struct {
	Client pbe2t.E2TServiceClient
}

func NewMsgSender(hostip string,port string) *MsgSender {
	cc, err := grpc.Dial(hostip+":"+port, grpc.WithInsecure())
	if err != nil {
		xapp.Logger.Info("could not connect: %v", err.Error())
	}
	//defer cc.Close()

	c := pbe2t.NewE2TServiceClient(cc)
	fmt.Printf("Created client: %f", c)
	return &MsgSender{Client: c}
}

func(s *MsgSender)String() string{
	return "GrpcMsgSender"
}

func(s *MsgSender)SendMsg(req *xapp.MsgParams) error {
	msg := &pbe2t.GrpcMsg{}
	msg.Topic 		= req.Src
	msg.SubId 		= int32(req.SubId)
	msg.Mtype 		= int32(req.Mtype)
	msg.PayloadLen 	= int32(req.PayloadLen)
	msg.Payload 	= req.Payload
	msg.Meid 		= req.Meid.RanName

	var i int
	var err error
	for i=0;i<5;i++ {
		_, err = s.Client.RpcHandleMsg(context.Background(), msg)
		if err != nil {
			fmt.Println("%d time(s): error while calling gRPC: %v\n", i+1,err)
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	return err
}


func (s *MsgSender) RouteTableInsert(r *pbe2t.RouteTableInsertRequest)()  {
	s.Client.RouteTableInsert(context.Background(),r)
}


func (s *MsgSender) RouteTableUpdate(r *pbe2t.RouteTableUpdateRequest)()  {
	s.Client.RouteTableUpdate(context.Background(),r)
}
