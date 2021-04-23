package dbclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"nRIC/internal/xapp"
	"time"
	pbdb "nRIC/api/v1/pb/db"
)

type MsgSender struct {
	Client pbdb.DbServiceClient
}

func NewMsgSender(hostip string,port string) *MsgSender {
	cc, err := grpc.Dial(hostip+":"+port, grpc.WithInsecure())
	if err != nil {
		xapp.Logger.Info("could not connect: %v", err.Error())
	}
	//defer cc.Close()

	c := pbdb.NewDbServiceClient(cc)
	fmt.Printf("Created client: %f", c)
	return &MsgSender{Client: c}
}

func(s *MsgSender)String() string{
	return "GrpcMsgSender"
}

func(s *MsgSender)SendMsg(req *xapp.MsgParams) error {
	msg := &pbdb.GrpcMsg{}
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


func (s *MsgSender) RouteTableInsert(r *pbdb.RouteTableInsertRequest)()  {
	_,err := s.Client.RouteTableInsert(context.Background(),r)
	if err != nil {
		xapp.Logger.Error(err.Error())
	}
}


func (s *MsgSender) RouteTableUpdate(r *pbdb.RouteTableUpdateRequest)()  {
	_, err := s.Client.RouteTableUpdate(context.Background(),r)
	if err != nil {
		xapp.Logger.Error(err.Error())
	}
}


func (s *MsgSender) RANFunctionsTableInsert(r *pbdb.RANFunctionsTableInsertRequest)()  {
	_,err := s.Client.RANFunctionsTableInsert(context.Background(),r)
	if err != nil {
		xapp.Logger.Error(err.Error())
	}
}

func (s *MsgSender) RANFunctionsTableUpdate(r *pbdb.RANFunctionsTableUpdateRequest)()  {
	_,err := s.Client.RANFunctionsTableUpdate(context.Background(),r)
	if err != nil {
		xapp.Logger.Error(err.Error())
	}
}
