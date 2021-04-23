package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"nRIC/internal/msgx"
	"nRIC/internal/msgx/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////
type MsgServiceEndpoint struct {
	GrpcMsgServiceEndpoint           	endpoint.Endpoint
}

func NewMsgServiceEndpoint(svc service.MsgService) MsgServiceEndpoint {
	return MsgServiceEndpoint{
		GrpcMsgServiceEndpoint:           MakeGrpcMsgServiceEndpoint(svc),
	}
}

func MakeGrpcMsgServiceEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*msgx.GrpcMsg)
		//fmt.Printf("RpcHandleMsg")
		res, err := svc.RpcHandleMsg(ctx, req)
		if err != nil {
			resp := &msgx.GrpcReply{}
			resp.Code = 1
			return res, nil
		}
		return res, nil
	}
}

