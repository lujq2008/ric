package msgx

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"nRIC/internal/xapp"
	"time"
)

type MsgSender struct {
	Client MsgServiceClient
}

func NewMsgSender(hostip string,port string) *MsgSender {
	cc, err := grpc.Dial(hostip+":"+port, grpc.WithInsecure())
	if err != nil {
		xapp.Logger.Info("could not connect: %v", err.Error())
	}
	//defer cc.Close()

	c := NewMsgServiceClient(cc)
	fmt.Printf("Created client: %f", c)
	return &MsgSender{Client: c}
}

func(s *MsgSender)String() string{
	return "GrpcMsgSender"
}

func(s *MsgSender)SendMsg(req *xapp.MsgParams) error {
	msg := &GrpcMsg{}
	msg.Topic 		= req.Src
	msg.SubId 		= int32(req.SubId)
	msg.Mtype 		= int32(req.Mtype)
	msg.PayloadLen 	= int32(req.PayloadLen)
	msg.Payload 	= req.Payload
	msg.Meid 		= req.Meid.RanName

	var i int
	var err error
	for i=0;i<5;i++ {
		xapp.Logger.Debug("msg: %s",msg.String())
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

