package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	pbe2t "nRIC/api/v1/pb/nrice2t"
	"nRIC/pkg/nrice2t/route/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////

func MakeRouteTableInsertEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.RouteTableInsertRequest)
		fmt.Printf("RouteTableInsert")
		res, err := svc.RouteTableInsert(ctx, req)
		if err != nil {
			res := &pbe2t.RouteTableInsertResponse{}
			res.SubIdXapp = -1
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableReadEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.RouteTableReadRequest)
		fmt.Printf("RouteTableRead")
		res, err := svc.RouteTableRead(ctx, req)
		if err != nil {
			res := &pbe2t.RouteTableReadResponse{}
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableUpdateEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.RouteTableUpdateRequest)
		fmt.Printf("RouteTableUpdate")
		res, err := svc.RouteTableUpdate(ctx, req)
		if err != nil {
			res := &pbe2t.RouteTableUpdateResponse{}
			res.Updated = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableDeleteEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.RouteTableDeleteRequest)
		fmt.Printf("RouteTableDelete")
		res, err := svc.RouteTableDelete(ctx, req)
		if err != nil {
			res := &pbe2t.RouteTableDeleteResponse{}
			res.Deleted = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRouteTableReadAllEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbe2t.RouteTableReadAllRequest)
		fmt.Printf("RouteTableReadAll")
		res, err := svc.RouteTableReadAll(ctx, req)
		if err != nil {
			res := &pbe2t.RouteTableReadAllResponse{}
			return res, nil
		}
		return res, nil
	}
}

//////////////////////////////
