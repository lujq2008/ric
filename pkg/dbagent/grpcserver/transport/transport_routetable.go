package transport

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////
func (g *msgServer)RouteTableInsert(ctx context.Context, r *pbdb.RouteTableInsertRequest) (*pbdb.RouteTableInsertResponse, error){
	_, rep, err := g.routeTableCreate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RouteTableInsertResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableRead(ctx context.Context, r *pbdb.RouteTableReadRequest) (*pbdb.RouteTableReadResponse, error){
	_, rep, err := g.routeTableRead.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RouteTableReadResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableUpdate(ctx context.Context, r *pbdb.RouteTableUpdateRequest) (*pbdb.RouteTableUpdateResponse, error){
	_, rep, err := g.routeTableUpdate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RouteTableUpdateResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableDelete(ctx context.Context, r *pbdb.RouteTableDeleteRequest) (*pbdb.RouteTableDeleteResponse, error){
	_, rep, err := g.routeTableDelete.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RouteTableDeleteResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableReadAll(ctx context.Context, r *pbdb.RouteTableReadAllRequest) (*pbdb.RouteTableReadAllResponse, error){
	_, rep, err := g.routeTableReadAll.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbdb.RouteTableReadAllResponse)
	return Resp, nil
}