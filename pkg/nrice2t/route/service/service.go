package service

import (
	"context"
	pbe2t "nRIC/api/v1/pb/nrice2t"
	"nRIC/internal/xapp"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
type MsgService interface {
	RpcHandleMsg(context.Context, *pbe2t.GrpcMsg) (*pbe2t.GrpcReply, error)
	RouteTableInsert(context.Context, *pbe2t.RouteTableInsertRequest) (*pbe2t.RouteTableInsertResponse, error)
	RouteTableRead(context.Context, *pbe2t.RouteTableReadRequest) (*pbe2t.RouteTableReadResponse, error)
	RouteTableUpdate(context.Context, *pbe2t.RouteTableUpdateRequest) (*pbe2t.RouteTableUpdateResponse, error)
	RouteTableDelete(context.Context, *pbe2t.RouteTableDeleteRequest) (*pbe2t.RouteTableDeleteResponse, error)
	RouteTableReadAll(context.Context, *pbe2t.RouteTableReadAllRequest) (*pbe2t.RouteTableReadAllResponse, error)
}


type msgService struct{
	c xapp.MessageConsumer
}

func NewmsgService(c xapp.MessageConsumer) MsgService {
	return &msgService{
		c:	c,
	}
}

func (w *msgService) RpcHandleMsg(_ context.Context, msg *pbe2t.GrpcMsg) (*pbe2t.GrpcReply, error) {
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

	resp := &pbe2t.GrpcReply{}
	err := w.c.Consume(params)
	resp.Code = 0
	return resp, err
}
