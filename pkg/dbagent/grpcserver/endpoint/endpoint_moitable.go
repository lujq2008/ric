package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/grpcserver/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////

func MakeMOITableInsertEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.MOITableInsertRequest)
		//fmt.Printf("MOITableInsert")
		res, err := svc.MOITableInsert(ctx, req)
		if err != nil {
			res := &pbdb.MOITableInsertResponse{}
			res.XappID = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeMOITableReadEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.MOITableReadRequest)
		fmt.Printf("MOITableRead")
		res, err := svc.MOITableRead(ctx, req)
		if err != nil {
			res := &pbdb.MOITableReadResponse{}
			return res, nil
		}
		return res, nil
	}
}

func MakeMOITableUpdateEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.MOITableUpdateRequest)
		fmt.Printf("MOITableUpdate")
		res, err := svc.MOITableUpdate(ctx, req)
		if err != nil {
			res := &pbdb.MOITableUpdateResponse{}
			res.Updated = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeMOITableDeleteEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.MOITableDeleteRequest)
		fmt.Printf("MOITableDelete")
		res, err := svc.MOITableDelete(ctx, req)
		if err != nil {
			res := &pbdb.MOITableDeleteResponse{}
			res.Deleted = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeMOITableReadAllEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.MOITableReadAllRequest)
		//fmt.Printf("MOITableReadAll")
		res, err := svc.MOITableReadAll(ctx, req)
		if err != nil {
			res := &pbdb.MOITableReadAllResponse{}
			return res, nil
		}
		return res, nil
	}
}

//////////////////////////////
