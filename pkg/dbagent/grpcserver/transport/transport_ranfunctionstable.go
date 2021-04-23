package transport

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////
func (g *msgServer)RANFunctionsTableInsert(ctx context.Context, r *pbdb.RANFunctionsTableInsertRequest) (*pbdb.RANFunctionsTableInsertResponse, error){
	_, rep, err := g.ranFunctionsTableCreate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RANFunctionsTableInsertResponse)
	return Resp, nil
}
func (g *msgServer)RANFunctionsTableRead(ctx context.Context, r *pbdb.RANFunctionsTableReadRequest) (*pbdb.RANFunctionsTableReadResponse, error){
	_, rep, err := g.ranFunctionsTableRead.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RANFunctionsTableReadResponse)
	return Resp, nil
}
func (g *msgServer)RANFunctionsTableUpdate(ctx context.Context, r *pbdb.RANFunctionsTableUpdateRequest) (*pbdb.RANFunctionsTableUpdateResponse, error){
	_, rep, err := g.ranFunctionsTableUpdate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RANFunctionsTableUpdateResponse)
	return Resp, nil
}
func (g *msgServer)RANFunctionsTableDelete(ctx context.Context, r *pbdb.RANFunctionsTableDeleteRequest) (*pbdb.RANFunctionsTableDeleteResponse, error){
	_, rep, err := g.ranFunctionsTableDelete.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RANFunctionsTableDeleteResponse)
	return Resp, nil
}
func (g *msgServer)RANFunctionsTableReadAll(ctx context.Context, r *pbdb.RANFunctionsTableReadAllRequest) (*pbdb.RANFunctionsTableReadAllResponse, error){
	_, rep, err := g.ranFunctionsTableReadAll.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RANFunctionsTableReadAllResponse)
	return Resp, nil
}