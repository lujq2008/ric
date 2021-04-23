package service

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/internal/xapp"
	db "nRIC/pkg/dbagent/driver/mariadb"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
//access to mariadb
//insert
func (w *msgService)RouteTableInsert(_ context.Context, r *pbdb.RouteTableInsertRequest) (*pbdb.RouteTableInsertResponse, error){
	//publish route to E2T
	params := &xapp.MsgParams{}
	params.Mtype = xapp.RIC_ROUTE_INSERT
	rt := &xapp.Route{}
	rt.SubIdXapp = r.RouteTable.SubIdXapp
	rt.SubIdRan = r.RouteTable.SubIdRan
	rt.Topic = r.RouteTable.Topic
	params.Route = rt

	w.c.Consume(params)

	//write to db
	return db.Insert(r.RouteTable)
}
func (w *msgService)RouteTableRead(_ context.Context, r *pbdb.RouteTableReadRequest) (*pbdb.RouteTableReadResponse, error){
	route ,err := db.Get(r.SubIdXapp)
	return &pbdb.RouteTableReadResponse{RouteTable: route},err
}
func (w *msgService)RouteTableUpdate(_ context.Context, r *pbdb.RouteTableUpdateRequest) (*pbdb.RouteTableUpdateResponse, error){
	//publish route to E2T
	params := &xapp.MsgParams{}
	params.Mtype = xapp.RIC_ROUTE_UPDATE
	rt := &xapp.Route{}
	rt.SubIdXapp = r.RouteTable.SubIdXapp
	rt.SubIdRan = r.RouteTable.SubIdRan
	rt.Topic = r.RouteTable.Topic
	params.Route = rt

	w.c.Consume(params)
	//write to db

	return db.Update(r.RouteTable)
}
func (w *msgService)RouteTableDelete(_ context.Context,r  *pbdb.RouteTableDeleteRequest) (*pbdb.RouteTableDeleteResponse, error){
	return nil,nil
}
func (w *msgService)RouteTableReadAll(_ context.Context,r *pbdb.RouteTableReadAllRequest) (*pbdb.RouteTableReadAllResponse, error){
	Routetalbe_list,err := db.List()
	return &pbdb.RouteTableReadAllResponse{RouteTables: Routetalbe_list},err
}
