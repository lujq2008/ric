package service

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/internal/xapp"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
type MsgService interface {
	RpcHandleMsg(context.Context, *pbdb.GrpcMsg) (*pbdb.GrpcReply, error)
	RouteTableInsert(context.Context, *pbdb.RouteTableInsertRequest) (*pbdb.RouteTableInsertResponse, error)
	RouteTableRead(context.Context, *pbdb.RouteTableReadRequest) (*pbdb.RouteTableReadResponse, error)
	RouteTableUpdate(context.Context, *pbdb.RouteTableUpdateRequest) (*pbdb.RouteTableUpdateResponse, error)
	RouteTableDelete(context.Context, *pbdb.RouteTableDeleteRequest) (*pbdb.RouteTableDeleteResponse, error)
	RouteTableReadAll(context.Context, *pbdb.RouteTableReadAllRequest) (*pbdb.RouteTableReadAllResponse, error)

	RANFunctionsTableInsert(context.Context, *pbdb.RANFunctionsTableInsertRequest) (*pbdb.RANFunctionsTableInsertResponse, error)
	RANFunctionsTableRead(context.Context, *pbdb.RANFunctionsTableReadRequest) (*pbdb.RANFunctionsTableReadResponse, error)
	RANFunctionsTableUpdate(context.Context, *pbdb.RANFunctionsTableUpdateRequest) (*pbdb.RANFunctionsTableUpdateResponse, error)
	RANFunctionsTableDelete(context.Context, *pbdb.RANFunctionsTableDeleteRequest) (*pbdb.RANFunctionsTableDeleteResponse, error)
	RANFunctionsTableReadAll(context.Context, *pbdb.RANFunctionsTableReadAllRequest) (*pbdb.RANFunctionsTableReadAllResponse, error)

	MOITableInsert(context.Context, *pbdb.MOITableInsertRequest) (*pbdb.MOITableInsertResponse, error)
	MOITableRead(context.Context, *pbdb.MOITableReadRequest) (*pbdb.MOITableReadResponse, error)
	MOITableUpdate(context.Context, *pbdb.MOITableUpdateRequest) (*pbdb.MOITableUpdateResponse, error)
	MOITableDelete(context.Context, *pbdb.MOITableDeleteRequest) (*pbdb.MOITableDeleteResponse, error)
	MOITableReadAll(context.Context, *pbdb.MOITableReadAllRequest) (*pbdb.MOITableReadAllResponse, error)
}


type msgService struct{
	c xapp.MessageConsumer
}

func NewmsgService(c xapp.MessageConsumer) MsgService {
	return &msgService{
		c:	c,
	}
}

func (w *msgService) RpcHandleMsg(_ context.Context, msg *pbdb.GrpcMsg) (*pbdb.GrpcReply, error) {
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

	resp := &pbdb.GrpcReply{}
	err := w.c.Consume(params)
	resp.Code = 0
	return resp, err
}
