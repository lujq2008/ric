package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/grpcserver/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////

func MakeRouteTableInsertEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RouteTableInsertRequest)
		fmt.Printf("RouteTableInsert")
		res, err := svc.RouteTableInsert(ctx, req)
		if err != nil {
			res := &pbdb.RouteTableInsertResponse{}
			res.SubIdXapp = -1
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableReadEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RouteTableReadRequest)
		fmt.Printf("RouteTableRead")
		res, err := svc.RouteTableRead(ctx, req)
		if err != nil {
			res := &pbdb.RouteTableReadResponse{}
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableUpdateEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RouteTableUpdateRequest)
		fmt.Printf("RouteTableUpdate")
		res, err := svc.RouteTableUpdate(ctx, req)
		if err != nil {
			res := &pbdb.RouteTableUpdateResponse{}
			res.Updated = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableDeleteEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RouteTableDeleteRequest)
		fmt.Printf("RouteTableDelete")
		res, err := svc.RouteTableDelete(ctx, req)
		if err != nil {
			res := &pbdb.RouteTableDeleteResponse{}
			res.Deleted = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableReadAllEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RouteTableReadAllRequest)
		fmt.Printf("RouteTableReadAll")
		res, err := svc.RouteTableReadAll(ctx, req)
		if err != nil {
			res := &pbdb.RouteTableReadAllResponse{}
			return res, nil
		}
		return res, nil
	}
}

//////////////////////////////
