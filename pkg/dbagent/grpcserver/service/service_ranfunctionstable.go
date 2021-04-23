package service

import (
	"context"
	pbdb "nRIC/api/v1/pb/db"
	db "nRIC/pkg/dbagent/driver/mariadb"
)

// 1 service  /////////////////////////////////////////////////////////////////////////////
//access to mariadb
//insert
func (w *msgService)RANFunctionsTableInsert(_ context.Context, r *pbdb.RANFunctionsTableInsertRequest) (*pbdb.RANFunctionsTableInsertResponse, error){
	//write to db
	return db.InsertRANFunctionsTable(r.RANFunctionsTable)
}
func (w *msgService)RANFunctionsTableRead(_ context.Context, r *pbdb.RANFunctionsTableReadRequest) (*pbdb.RANFunctionsTableReadResponse, error){
	route ,err := db.GetRANFunctionsTable(r.GlobalE2NodeIDStr,r.RanFunctionID,r.RanFunctionRevision,r.RanFunctionOID)
	return &pbdb.RANFunctionsTableReadResponse{RANFunctionsTable: route},err
}
func (w *msgService)RANFunctionsTableUpdate(_ context.Context, r *pbdb.RANFunctionsTableUpdateRequest) (*pbdb.RANFunctionsTableUpdateResponse, error){
	//write to db
	return db.UpdateRANFunctionsTable(r.RANFunctionsTable)
}
func (w *msgService)RANFunctionsTableDelete(_ context.Context,r  *pbdb.RANFunctionsTableDeleteRequest) (*pbdb.RANFunctionsTableDeleteResponse, error){
	return nil,nil
}
func (w *msgService)RANFunctionsTableReadAll(_ context.Context,r *pbdb.RANFunctionsTableReadAllRequest) (*pbdb.RANFunctionsTableReadAllResponse, error){
	Routetalbe_list,err := db.ListRANFunctionsTable()
	return &pbdb.RANFunctionsTableReadAllResponse{RANFunctionsTables: Routetalbe_list},err
}
