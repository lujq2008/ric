package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"nRIC/pkg/nrice2t/route/service"
	pbe2t "nRIC/api/v1/pb/nrice2t"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////
type MsgServiceEndpoint struct {
	GrpcMsgServiceEndpoint           	endpoint.Endpoint
	RouteTableInsertEndpoint            endpoint.Endpoint
	RouteTableReadEndpoint              endpoint.Endpoint
	RouteTableUpdateEndpoint       	    endpoint.Endpoint
	RouteTableDeleteEndpoint            endpoint.Endpoint
	RouteTableReadAllEndpoint           endpoint.Endpoint
}

func NewMsgServiceEndpoint(svc service.MsgService) MsgServiceEndpoint {
	return MsgServiceEndpoint{
		GrpcMsgServiceEndpoint:           MakeGrpcMsgServiceEndpoint(svc),
		RouteTableInsertEndpoint:         MakeRouteTableInsertEndpoint(svc),
		RouteTableReadEndpoint:           MakeRouteTableReadEndpoint(svc),
		RouteTableUpdateEndpoint:         MakeRouteTableUpdateEndpoint(svc),
		RouteTableDeleteEndpoint:         MakeRouteTableDeleteEndpoint(svc),
		RouteTableReadAllEndpoint:        MakeRouteTableReadAllEndpoint(svc),
	}
}

func MakeGrpcMsgServiceEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.GrpcMsg)
		//fmt.Printf("RpcHandleMsg")
		res, err := svc.RpcHandleMsg(ctx, req)
		if err != nil {
			resp := &pbe2t.GrpcReply{}
			resp.Code = 1
			return res, nil
		}
		return res, nil
	}
}

