package transport

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////
func (g *msgServer)MOITableInsert(ctx context.Context, r *pbdb.MOITableInsertRequest) (*pbdb.MOITableInsertResponse, error){
	_, rep, err := g.moiTableCreate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.MOITableInsertResponse)
	return Resp, nil
}
func (g *msgServer)MOITableRead(ctx context.Context, r *pbdb.MOITableReadRequest) (*pbdb.MOITableReadResponse, error){
	_, rep, err := g.moiTableRead.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.MOITableReadResponse)
	return Resp, nil
}
func (g *msgServer)MOITableUpdate(ctx context.Context, r *pbdb.MOITableUpdateRequest) (*pbdb.MOITableUpdateResponse, error){
	_, rep, err := g.moiTableUpdate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.MOITableUpdateResponse)
	return Resp, nil
}
func (g *msgServer)MOITableDelete(ctx context.Context, r *pbdb.MOITableDeleteRequest) (*pbdb.MOITableDeleteResponse, error){
	_, rep, err := g.moiTableDelete.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.MOITableDeleteResponse)
	return Resp, nil
}
func (g *msgServer)MOITableReadAll(ctx context.Context, r *pbdb.MOITableReadAllRequest) (*pbdb.MOITableReadAllResponse, error){
	_, rep, err := g.moiTableReadAll.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	//空表
	if rep == nil {
		return nil ,nil
	}
	Resp := rep.(*pbdb.MOITableReadAllResponse)
	return Resp, nil
}