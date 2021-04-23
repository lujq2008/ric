package service

import (
	"context"
	"nRIC/internal/msgx"
	"nRIC/internal/xapp"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
type MsgService interface {
	RpcHandleMsg(context.Context, *msgx.GrpcMsg) (*msgx.GrpcReply, error)
}


type msgService struct{
	c xapp.MessageConsumer
}

func NewMsgService(c xapp.MessageConsumer) MsgService {
	return &msgService{
		c:	c,
	}
}

func (w *msgService) RpcHandleMsg(_ context.Context, msg *msgx.GrpcMsg) (*msgx.GrpcReply, error) {
	// query the database using the filters and return the list of documents
	// return error if the filter (key) is invalid and also return error if no item found
	//response := "cnbu define response222:" + input
	//fmt.Printf("\nRpcHandleMsg Server: %x \n",msg.Payload)
	params := &xapp.MsgParams{}
	params.Mtype 		= int(msg.Mtype)
	params.Payload 		= msg.Payload
	params.PayloadLen 	= int(msg.PayloadLen)
	params.Src 			= msg.Topic
	params.SubId		= int(msg.SubId)
	params.Meid = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}
	params.Meid.RanName = msg.Meid

	resp := &msgx.GrpcReply{}
	err := w.c.Consume(params)
	resp.Code = 0
	return resp, err
}
