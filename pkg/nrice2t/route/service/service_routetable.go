package service

import (
	"context"
	pbe2t "nRIC/api/v1/pb/nrice2t"
	dbx "nRIC/pkg/nrice2t/route/handleroute"
)

var Db *dbx.Route2Xapp

// 1 service  /////////////////////////////////////////////////////////////////////////////
//access to

//insert
func (w *msgService)RouteTableInsert(_ context.Context, r *pbe2t.RouteTableInsertRequest) (*pbe2t.RouteTableInsertResponse, error){
	return Db.Insert(r.RouteTable)
}
func (w *msgService)RouteTableRead(_ context.Context, r *pbe2t.RouteTableReadRequest) (*pbe2t.RouteTableReadResponse, error){
	return nil,nil
}
func (w *msgService)RouteTableUpdate(_ context.Context, r *pbe2t.RouteTableUpdateRequest) (*pbe2t.RouteTableUpdateResponse, error){
	return Db.Update(r.RouteTable)
}
func (w *msgService)RouteTableDelete(_ context.Context,r  *pbe2t.RouteTableDeleteRequest) (*pbe2t.RouteTableDeleteResponse, error){
	return nil,nil
}
func (w *msgService)RouteTableReadAll(_ context.Context,r *pbe2t.RouteTableReadAllRequest) (*pbe2t.RouteTableReadAllResponse, error){
	Routetalbe_list,err := Db.List()
	return &pbe2t.RouteTableReadAllResponse{RouteTables: Routetalbe_list},err
}

func init ()  {
	Db = dbx.NewRoute2Xapp()
}