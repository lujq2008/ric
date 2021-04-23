package mariadb

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func MustConnectDB() *sql.DB {
	if err := ConnectDatabase(); err != nil {
		panic(err)
	}
	return db
}

var (
	username string
	password string
	host     string
	port     string
	database string
)

func Config() {
	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")

	host = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	if host == "" {
		host = "localhost"
	}

	port = os.Getenv("MYSQL_PORT_3306_TCP_PORT")
	if port == "" {
		port = "3306"
	}

	database = os.Getenv("MYSQL_INSTANCE_NAME")
}

func ConnectDatabase() (err error) {
	//uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
	uri :=  "root:D5uxIiry0n@tcp(opulent-tuatara-mariadb-galera:3306)/mysql"
	if db, err = sql.Open("mysql", uri); err != nil {
		return err
	}
	err = db.Ping()
	return err
}

type RouteTable struct {
	SubId int64
	Dest  string
}
func InitDB() {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()

	CreateRouteTable()
	//Insert(&msgx.RouteTable{SubId: 2, Dest: })
}

