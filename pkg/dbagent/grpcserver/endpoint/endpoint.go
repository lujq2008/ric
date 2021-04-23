package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/grpcserver/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////
type MsgServiceEndpoint struct {
	GrpcMsgServiceEndpoint           	endpoint.Endpoint
	RouteTableInsertEndpoint            endpoint.Endpoint
	RouteTableReadEndpoint              endpoint.Endpoint
	RouteTableUpdateEndpoint       	    endpoint.Endpoint
	RouteTableDeleteEndpoint            endpoint.Endpoint
	RouteTableReadAllEndpoint           endpoint.Endpoint

	RANFunctionsTableInsertEndpoint     endpoint.Endpoint
	RANFunctionsTableReadEndpoint       endpoint.Endpoint
	RANFunctionsTableUpdateEndpoint     endpoint.Endpoint
	RANFunctionsTableDeleteEndpoint     endpoint.Endpoint
	RANFunctionsTableReadAllEndpoint    endpoint.Endpoint

	MOITableInsertEndpoint            	endpoint.Endpoint
	MOITableReadEndpoint              	endpoint.Endpoint
	MOITableUpdateEndpoint       	    endpoint.Endpoint
	MOITableDeleteEndpoint            	endpoint.Endpoint
	MOITableReadAllEndpoint           	endpoint.Endpoint

}

func NewMsgServiceEndpoint(svc service.MsgService) MsgServiceEndpoint {
	return MsgServiceEndpoint{
		GrpcMsgServiceEndpoint:           MakeGrpcMsgServiceEndpoint(svc),
		RouteTableInsertEndpoint:         MakeRouteTableInsertEndpoint(svc),
		RouteTableReadEndpoint:           MakeRouteTableReadEndpoint(svc),
		RouteTableUpdateEndpoint:         MakeRouteTableUpdateEndpoint(svc),
		RouteTableDeleteEndpoint:         MakeRouteTableDeleteEndpoint(svc),
		RouteTableReadAllEndpoint:        MakeRouteTableReadAllEndpoint(svc),

		RANFunctionsTableInsertEndpoint:         MakeRANFunctionsTableInsertEndpoint(svc),
		RANFunctionsTableReadEndpoint:           MakeRANFunctionsTableReadEndpoint(svc),
		RANFunctionsTableUpdateEndpoint:         MakeRANFunctionsTableUpdateEndpoint(svc),
		RANFunctionsTableDeleteEndpoint:         MakeRANFunctionsTableDeleteEndpoint(svc),
		RANFunctionsTableReadAllEndpoint:        MakeRANFunctionsTableReadAllEndpoint(svc),

		MOITableInsertEndpoint:         MakeMOITableInsertEndpoint(svc),
		MOITableReadEndpoint:           MakeMOITableReadEndpoint(svc),
		MOITableUpdateEndpoint:         MakeMOITableUpdateEndpoint(svc),
		MOITableDeleteEndpoint:         MakeMOITableDeleteEndpoint(svc),
		MOITableReadAllEndpoint:        MakeMOITableReadAllEndpoint(svc),

	}
}

func MakeGrpcMsgServiceEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.GrpcMsg)
		//fmt.Printf("RpcHandleMsg")
		res, err := svc.RpcHandleMsg(ctx, req)
		if err != nil {
			resp := &pbdb.GrpcReply{}
			resp.Code = 1
			return res, nil
		}
		return res, nil
	}
}

