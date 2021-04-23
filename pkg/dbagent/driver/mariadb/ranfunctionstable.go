package mariadb

import (
	"fmt"
	"log"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/internal/xapp"
)


func CreateRANFunctionsTable() {
	stmt, err := db.Prepare(createRANFunctionsTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

var createRANFunctionsTable = `
CREATE TABLE IF NOT EXISTS ranfunctions (
     GlobalE2NodeIDStr        	VARCHAR(256),
	 RanFunctionID            	INT,   
	 RanFunctionRevision		INT,
	 RanFunctionOID				VARCHAR(256),
	 RanFunctionDefinition	    BLOB(65535),
	 primary key (GlobalE2NodeIDStr,RanFunctionID,RanFunctionRevision,RanFunctionOID)
);
`

var RANFunctionsTableInsert = `
INSERT INTO
    ranfunctions(
        GlobalE2NodeIDStr,
        RanFunctionID,
        RanFunctionRevision,
		RanFunctionOID,
		RanFunctionDefinition
    )
VALUES(?, ?, ?, ?, ?)
`
var RANFunctionsTableQuery = `
SELECT
        GlobalE2NodeIDStr,
        RanFunctionID,
        RanFunctionRevision,
        RanFunctionOID,
        RanFunctionDefinition
FROM
    ranfunctions
`

var dropRANFunctionsTable = `
DROP TABLE ranfunctions;
`


func DropRANFunctionsTable() {
	stmt, err := db.Prepare(dropRANFunctionsTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	CreateRANFunctionsTable()
}


func InsertRANFunctionsTable(ranfunctions *pbdb.RANFunctionsTable) (*pbdb.RANFunctionsTableInsertResponse, error){
	fmt.Printf("Insert GlobalE2NodeIDStr = %s,RanFunctionID = %d,RanFunctionRevision = %d, RanFunctionOID=%s\n",
		ranfunctions.GlobalE2NodeIDStr,ranfunctions.RanFunctionID,ranfunctions.RanFunctionRevision,ranfunctions.RanFunctionOID)

	m,err := GetRANFunctionsTable(ranfunctions.GlobalE2NodeIDStr,ranfunctions.RanFunctionID,ranfunctions.RanFunctionRevision,ranfunctions.RanFunctionOID)
	if m != nil {
		//panic("insert conflict")
		r,err := UpdateRANFunctionsTable(ranfunctions)
		return  &pbdb.RANFunctionsTableInsertResponse{Api:r.Api,ResultCode: -1},err
	}

	stmt, err := db.Prepare(RANFunctionsTableInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ranfunctions.GlobalE2NodeIDStr,
		ranfunctions.RanFunctionID,
		ranfunctions.RanFunctionRevision,
		ranfunctions.RanFunctionOID,
		ranfunctions.RanFunctionDefinition,
	)
	if err != nil {
		panic(err)
		return nil, err
	}
	return &pbdb.RANFunctionsTableInsertResponse{ResultCode: 0},nil
}


func UpdateRANFunctionsTable(ranfunctions *pbdb.RANFunctionsTable) (*pbdb.RANFunctionsTableUpdateResponse, error){
	fmt.Printf("Update1 GlobalE2NodeIDStr = %s,RanFunctionID = %d,RanFunctionRevision = %d, RanFunctionOID=%s\n",
		ranfunctions.GlobalE2NodeIDStr,ranfunctions.RanFunctionID,ranfunctions.RanFunctionRevision,ranfunctions.RanFunctionOID)

	m,err := GetRANFunctionsTable(ranfunctions.GlobalE2NodeIDStr,ranfunctions.RanFunctionID,ranfunctions.RanFunctionRevision,ranfunctions.RanFunctionOID)
	if m == nil {
		//panic("insert conflict")
		xapp.Logger.Error("Can not update:GlobalE2NodeIDStr = %s,RanFunctionID = %d,RanFunctionRevision = %d, RanFunctionOID=%s, RanFunctionDefinition=%v\n",
			ranfunctions.GlobalE2NodeIDStr,ranfunctions.RanFunctionID,ranfunctions.RanFunctionRevision,ranfunctions.RanFunctionOID,ranfunctions.RanFunctionDefinition)
		return nil,err
	}

	stmt, err := db.Prepare("UPDATE ranfunctions SET RanFunctionDefinition=? WHERE GlobalE2NodeIDStr = ? AND RanFunctionID = ? AND RanFunctionRevision = ? AND RanFunctionOID = ?")
	if err != nil {
		//panic(err)
		fmt.Printf("Update2:%s",err.Error())
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ranfunctions.RanFunctionDefinition,
		&m.GlobalE2NodeIDStr,
		&m.RanFunctionID,
		&m.RanFunctionRevision,
		&m.RanFunctionOID,
	)
	if err != nil {
		//panic(err)
		return nil, err
	}
	return &pbdb.RANFunctionsTableUpdateResponse{Updated: 0},nil
}


func ListRANFunctionsTable() ([]*pbdb.RANFunctionsTable,error) {
	rows, err := db.Query(RANFunctionsTableQuery)
	if err != nil {
		return nil ,err
	}
	defer rows.Close()

	list := []*pbdb.RANFunctionsTable{}
	for rows.Next() {
		m := new(pbdb.RANFunctionsTable)
		if err := rows.Scan(
			&m.GlobalE2NodeIDStr,
			&m.RanFunctionID,
			&m.RanFunctionRevision,
			&m.RanFunctionOID,
			&m.RanFunctionDefinition,
		); err != nil {
			log.Println(err)
			//panic(err)
			return nil ,err
		}
		list = append(list, m)
	}

	return list,err
}

func GetRANFunctionsTable(GlobalE2NodeIDStr string,RanFunctionID uint32,RanFunctionRevision uint32,RanFunctionOID string) (*pbdb.RANFunctionsTable,error)  {
	rows, err := db.Query(RANFunctionsTableQuery + " WHERE GlobalE2NodeIDStr=? AND RanFunctionID=? AND RanFunctionRevision=? AND RanFunctionOID=?",
		GlobalE2NodeIDStr,RanFunctionID,RanFunctionRevision,RanFunctionOID)
	//rows, err := db.Query(RANFunctionsTableQuery + " WHERE GlobalE2NodeIDStr='?'",GlobalE2NodeIDStr)
	if err != nil {
		//panic(err)
		fmt.Printf("Get1:%s",err.Error())
		return nil,err
	}
	defer rows.Close()

	m := new(pbdb.RANFunctionsTable)
	if rows.Next() {
		if err := rows.Scan(
			&m.GlobalE2NodeIDStr,
			&m.RanFunctionID,
			&m.RanFunctionRevision,
			&m.RanFunctionOID,
			&m.RanFunctionDefinition,
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


