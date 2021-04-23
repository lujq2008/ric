package service

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
	db "nRIC/pkg/dbagent/driver/mariadb"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
//access to mariadb
//insert
func (w *msgService)MOITableInsert(_ context.Context, r *pbdb.MOITableInsertRequest) (*pbdb.MOITableInsertResponse, error){
	//write to db
	return db.InsertMOITable(r.MoiTable)
}
func (w *msgService)MOITableRead(_ context.Context, r *pbdb.MOITableReadRequest) (*pbdb.MOITableReadResponse, error){
	route ,err := db.GetMOITable(r.XappID)
	return &pbdb.MOITableReadResponse{MoiTable: route},err
}
func (w *msgService)MOITableUpdate(_ context.Context, r *pbdb.MOITableUpdateRequest) (*pbdb.MOITableUpdateResponse, error){
	//write to db
	return db.UpdateMOITable(r.MoiTable)
}
func (w *msgService)MOITableDelete(_ context.Context,r  *pbdb.MOITableDeleteRequest) (*pbdb.MOITableDeleteResponse, error){
	return db.MOITableDelete(r.XappID)
}
func (w *msgService)MOITableReadAll(_ context.Context,r *pbdb.MOITableReadAllRequest) (*pbdb.MOITableReadAllResponse, error){
	MOItalbe_list,err := db.ListMOITable()
	return &pbdb.MOITableReadAllResponse{MoiTables: MOItalbe_list},err
}
