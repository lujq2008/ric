package transport

import (
	"context"
	pbe2t "nRIC/api/v1/pb/nrice2t"
)

//  3 transport  /////////////////////////////////////////////////////////////////////////////
func (g *msgServer)RouteTableInsert(ctx context.Context, r *pbe2t.RouteTableInsertRequest) (*pbe2t.RouteTableInsertResponse, error){
	_, rep, err := g.routeTableCreate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.RouteTableInsertResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableRead(ctx context.Context, r *pbe2t.RouteTableReadRequest) (*pbe2t.RouteTableReadResponse, error){
	_, rep, err := g.routeTableRead.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.RouteTableReadResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableUpdate(ctx context.Context, r *pbe2t.RouteTableUpdateRequest) (*pbe2t.RouteTableUpdateResponse, error){
	_, rep, err := g.routeTableUpdate.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.RouteTableUpdateResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableDelete(ctx context.Context, r *pbe2t.RouteTableDeleteRequest) (*pbe2t.RouteTableDeleteResponse, error){
	_, rep, err := g.routeTableDelete.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.RouteTableDeleteResponse)
	return Resp, nil
}
func (g *msgServer)RouteTableReadAll(ctx context.Context, r *pbe2t.RouteTableReadAllRequest) (*pbe2t.RouteTableReadAllResponse, error){
	_, rep, err := g.routeTableReadAll.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	Resp := rep.(*pbe2t.RouteTableReadAllResponse)
	return Resp, nil
}