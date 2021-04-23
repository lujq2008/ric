package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/grpcserver/service"
)

// 2 endpoint  /////////////////////////////////////////////////////////////////////////////

func MakeRANFunctionsTableInsertEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RANFunctionsTableInsertRequest)
		fmt.Printf("RANFunctionsTableInsert")
		res, err := svc.RANFunctionsTableInsert(ctx, req)
		if err != nil {
			res := &pbdb.RANFunctionsTableInsertResponse{}
			res.ResultCode = -1
			return res, nil
		}
		return res, nil
	}
}

func MakeRANFunctionsTableReadEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RANFunctionsTableReadRequest)
		fmt.Printf("RANFunctionsTableRead")
		res, err := svc.RANFunctionsTableRead(ctx, req)
		if err != nil {
			res := &pbdb.RANFunctionsTableReadResponse{}
			return res, nil
		}
		return res, nil
	}
}

func MakeRANFunctionsTableUpdateEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RANFunctionsTableUpdateRequest)
		fmt.Printf("RANFunctionsTableUpdate")
		res, err := svc.RANFunctionsTableUpdate(ctx, req)
		if err != nil {
			res := &pbdb.RANFunctionsTableUpdateResponse{}
			res.Updated = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRANFunctionsTableDeleteEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RANFunctionsTableDeleteRequest)
		fmt.Printf("RANFunctionsTableDelete")
		res, err := svc.RANFunctionsTableDelete(ctx, req)
		if err != nil {
			res := &pbdb.RANFunctionsTableDeleteResponse{}
			res.Deleted = 0
			return res, nil
		}
		return res, nil
	}
}

func MakeRANFunctionsTableReadAllEndpoint(svc service.MsgService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pbdb.RANFunctionsTableReadAllRequest)
		fmt.Printf("RANFunctionsTableReadAll")
		res, err := svc.RANFunctionsTableReadAll(ctx, req)
		if err != nil {
			res := &pbdb.RANFunctionsTableReadAllResponse{}
			return res, nil
		}
		return res, nil
	}
}

//////////////////////////////
