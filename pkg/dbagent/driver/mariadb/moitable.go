package mariadb

import (
	"fmt"
	"log"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/internal/xapp"
)


func CreateMOITable() {
	stmt, err := db.Prepare(createMOITable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

var createMOITable = `
CREATE TABLE IF NOT EXISTS moi (
     XappID        				INT,
	 XappVer            		FLOAT,   
	 XappName		 			VARCHAR(256),
	 Functions					VARCHAR(256),
	 RunningStatus	    		VARCHAR(256),
	 IsReady					VARCHAR(256),
	 Topic						VARCHAR(256),
	 primary key (XappID)
);
`

var MOITableInsert = `
INSERT INTO
    moi(
 	 XappID,
	 XappVer,   
	 XappName,
	 Functions,
	 RunningStatus,
	 IsReady,
	 Topic
    )
VALUES(?, ?, ?, ?, ?, ?, ?)
`
var MOITableQuery = `
SELECT
 	 XappID,
	 XappVer,   
	 XappName,
	 Functions,
	 RunningStatus,
	 IsReady,
	 Topic
FROM
    moi
`

var dropMOITable = `
DROP TABLE moi;
`


func DropMOITable() {
	stmt, err := db.Prepare(dropMOITable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	CreateMOITable()
}


func InsertMOITable(moi *pbdb.MOITable) (*pbdb.MOITableInsertResponse, error){
	fmt.Printf("Insert XappID = %d,XappVer = %f,XappName = %s, RunningStatus=%s\n",
		moi.XappID,moi.XappVer,moi.XappName,moi.RunningStatus)

	m,err := GetMOITable(moi.XappID)
	if m != nil {
		//panic("insert conflict")
		r,err := UpdateMOITable(moi)
		//XappID 合法值 （1...），0 为 非法值
		return  &pbdb.MOITableInsertResponse{Api:r.Api,XappID: 0},err
	}

	stmt, err := db.Prepare(MOITableInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		moi.XappID,
		moi.XappVer,
		moi.XappName,
		moi.Functions,
		moi.RunningStatus,
		moi.IsReady,
		moi.Topic,
	)
	if err != nil {
		panic(err)
		return nil, err
	}
	return &pbdb.MOITableInsertResponse{XappID: moi.XappID},nil
}


func UpdateMOITable(moi *pbdb.MOITable) (*pbdb.MOITableUpdateResponse, error){
	fmt.Printf("Update XappID = %d,XappVer = %f,XappName = %s, RunningStatus=%s\n",
		moi.XappID,moi.XappVer,moi.XappName,moi.RunningStatus)

	m,err := GetMOITable(moi.XappID)
	if m == nil {
		//panic("insert conflict")
		xapp.Logger.Error("Insert XappID = %d,XappVer = %f,XappName = %s, RunningStatus=%s\n",
			moi.XappID,moi.XappVer,moi.XappName,moi.RunningStatus)
		return nil,err
	}

	stmt, err := db.Prepare("UPDATE moi SET XappVer=?,XappName=?,Functions=?,RunningStatus=?,IsReady=?,Topic=? WHERE XappID = ?")
	if err != nil {
		//panic(err)
		fmt.Printf("Update2:%s",err.Error())
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		moi.XappVer,
		moi.XappName,
		moi.Functions,
		moi.RunningStatus,
		moi.IsReady,
		moi.Topic,
		moi.XappID,
	)
	if err != nil {
		//panic(err)
		return nil, err
	}
	return &pbdb.MOITableUpdateResponse{Updated: 0},nil
}


func ListMOITable() ([]*pbdb.MOITable,error) {
	rows, err := db.Query(MOITableQuery)
	if err != nil {
		return nil ,err
	}
	//空表
	if rows == nil {
		return nil , nil
	}
	defer rows.Close()

	list := []*pbdb.MOITable{}
	for rows.Next() {
		m := new(pbdb.MOITable)
		if err := rows.Scan(
			&m.XappID,
			&m.XappVer,
			&m.XappName,
			&m.Functions,
			&m.RunningStatus,
			&m.IsReady,
			&m.Topic,
		); err != nil {
			log.Println(err)
			//panic(err)
			return nil ,err
		}
		list = append(list, m)
	}

	return list,err
}

func GetMOITable(XappID uint32) (*pbdb.MOITable,error)  {
	rows, err := db.Query(MOITableQuery + " WHERE XappID = ?", int(XappID))
	//rows, err := db.Query(MOITableQuery + " WHERE GlobalE2NodeIDStr='?'",GlobalE2NodeIDStr)
	if err != nil {
		//panic(err)
		fmt.Printf("GetMOITable:%s",err.Error())
		return nil,err
	}
	defer rows.Close()

	m := new(pbdb.MOITable)
	if rows.Next() {
		if err := rows.Scan(
			&m.XappID,
			&m.XappVer,
			&m.XappName,
			&m.Functions,
			&m.RunningStatus,
			&m.IsReady,
			&m.Topic,
		); err != nil {
			//panic(err)
			fmt.Printf("Get2:%s",err.Error())
			return nil,err
		}
	}else {
		return  nil,nil
	}

	return m,err
}


func MOITableDelete(XappID uint32) (*pbdb.MOITableDeleteResponse, error) {
	stmt, err := db.Prepare("delete from moi where XappID=?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(XappID)
	if err != nil {
		panic(err)
	}

	// affected rows
	a, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("RowsAffected:%d\n",a)     // 1
	resp := &pbdb.MOITableDeleteResponse{
		Deleted: int32(a),
	}
	return resp,err
}
