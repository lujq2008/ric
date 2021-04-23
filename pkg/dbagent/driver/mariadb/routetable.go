package mariadb

import (
	"fmt"
	"log"
	"nRIC/internal/xapp"
	pbdb "nRIC/api/v1/pb/db"
)


func CreateRouteTable() {
	stmt, err := db.Prepare(createRouteTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

var createRouteTable = `
CREATE TABLE IF NOT EXISTS route (
     SubIdXapp        INT PRIMARY KEY,
	 Topic            VARCHAR(256),   
	 SubIdRan		  INT	    
);
`

func Drop() {
	stmt, err := db.Prepare(dropRouteTable)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	CreateRouteTable()
}

var dropRouteTable = `
DROP TABLE route;
`

func Insert(route *pbdb.RouteTable) (*pbdb.RouteTableInsertResponse, error){
	fmt.Printf("Insert SubIdXapp = %d,Topic = %s,SubIdRan = %d\n",route.SubIdXapp,route.Topic,route.SubIdRan)
	m,err := Get(route.SubIdXapp)
	if m != nil {
		//panic("insert conflict")
		r,err := Update(route)
		return  &pbdb.RouteTableInsertResponse{Api:r.Api,SubIdXapp: route.SubIdXapp},err
	}

	stmt, err := db.Prepare(routeInsert)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		route.SubIdXapp,
		route.Topic,
		route.SubIdRan,
	)
	if err != nil {
		panic(err)
		return nil, err
	}
	return &pbdb.RouteTableInsertResponse{SubIdXapp: route.SubIdXapp},nil
}


func Update(route *pbdb.RouteTable) (*pbdb.RouteTableUpdateResponse, error){
	fmt.Printf("Update1 SubIdXapp = %d,Topic = %s,SubIdRan = %d\n",route.SubIdXapp,route.Topic,route.SubIdRan)
	m,err := Get(route.SubIdXapp)
	if m == nil {
		//panic("insert conflict")
		xapp.Logger.Error("Can not update:SubIdXapp = %d,Topic = %s,SubIdRan = %d\n",route.SubIdXapp,route.Topic,route.SubIdRan)
		return nil,err
	}

	stmt, err := db.Prepare("UPDATE route SET SubIdRan=?, Topic=? WHERE SubIdXapp=?")
	if err != nil {
		//panic(err)
		fmt.Printf("Update2:%s",err.Error())
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		route.SubIdRan,
		route.Topic,
		route.SubIdXapp,
	)
	if err != nil {
		//panic(err)
		return nil, err
	}
	return &pbdb.RouteTableUpdateResponse{Updated: route.SubIdXapp},nil
}


func List() ([]*pbdb.RouteTable,error) {
	rows, err := db.Query(routeQuery)
	if err != nil {
		return nil ,err
	}
	defer rows.Close()

	list := []*pbdb.RouteTable{}
	for rows.Next() {
		m := new(pbdb.RouteTable)
		if err := rows.Scan(
			&m.SubIdXapp,
			&m.Topic,
			&m.SubIdRan,
		); err != nil {
			log.Println(err)
			//panic(err)
			return nil ,err
		}
		list = append(list, m)
	}

	return list,err
}

func Get(SubIdXapp int64) (*pbdb.RouteTable,error)  {
	rows, err := db.Query(routeQuery+" WHERE SubIdXapp = ?", SubIdXapp)
	if err != nil {
		//panic(err)
		fmt.Printf("Get1:%s",err.Error())
		return nil,err
	}
	defer rows.Close()

	m := new(pbdb.RouteTable)
	if rows.Next() {
		if err := rows.Scan(
			&m.SubIdXapp,
			&m.Topic,
			&m.SubIdRan,
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

var routeInsert = `
INSERT INTO
    route(
        SubIdXapp,
        Topic,
        SubIdRan
    )
VALUES(?, ?, ?)
`
var routeQuery = `
SELECT
        SubIdXapp,
        Topic,
        SubIdRan
FROM
    route
`

