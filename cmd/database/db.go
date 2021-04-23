package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"nRIC/internal"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/pkg/dbagent/control"
	"nRIC/pkg/dbagent/driver/mariadb"
	e2tclient "nRIC/pkg/nrice2t/route"
	"net"
	"os"
	_ "github.com/go-sql-driver/mysql"
)
func Setup () {
	//MsgSendertoE2T := msgx.NewMsgSender(internal.Nrice2tHost,internal.DefaultGRPCPort)
	//publish route to E2T
	MsgSendertoE2T := e2tclient.NewMsgSender(internal.Nrice2tHost,internal.DefaultGRPCPort)
	c := control.NewControl(MsgSendertoE2T)

	//msg server
	grpcAddr := net.JoinHostPort(internal.DbagentHost, envString("GRPC_PORT", internal.DefaultGRPCPort))
	c.Run(grpcAddr)

}
func main() {
											//id:password@tcp(your-amazonaws-uri.com:3306)/dbname
	//db, err := sql.Open("mysql", "root:D5uxIiry0n@tcp(opulent-tuatara-mariadb-galera:3306)/mysql")
	db := mariadb.MustConnectDB()
	defer db.Close()

	mariadb.CreateRouteTable()
	mariadb.CreateRANFunctionsTable()
	mariadb.Insert(&pbdb.RouteTable{SubIdXapp: 101,Topic: "xapp1_topic",SubIdRan: 1})
	mariadb.InsertRANFunctionsTable(&pbdb.RANFunctionsTable{GlobalE2NodeIDStr:"gnodeb-00001-00001",RanFunctionID:1,RanFunctionRevision:1,
		RanFunctionOID:"1.3.6.1.4.1.53148.1.2.255",
		RanFunctionDefinition:[]byte{'1','1','2','2','3','3','4','4','5','5','6','6','7','7','8','8','9','9','0','0'}})
	mariadb.CreateMOITable()
	mariadb.InsertMOITable(&pbdb.MOITable{XappID:1, Topic: "init",IsReady: "NA",XappName: "xxx",XappVer: 0.1,RunningStatus: "Inactive",Functions: "1,2"})

	routes,_ := mariadb.List()
	for _,r := range routes {
		fmt.Printf("SubIdXapp = %d, dest = %s , SubIdRan = %d ",r.SubIdXapp,r.Topic,r.SubIdRan)
	}
	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)

	selDB, _ := db.Query("SELECT * FROM wsrep_cluster ")
	var cluster_uuid string
	var (
		view_id int
		view_seqno int
		protocol_version int
		capabilities int
	)
	for selDB.Next() {
		selDB.Scan(&cluster_uuid,&view_id,&view_seqno,&protocol_version,&capabilities)
		fmt.Println(cluster_uuid,view_id,view_seqno,protocol_version,capabilities)
	}

	Setup()


}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func init() {
	var logger   log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	//internal.DbagentHost = internal.GetIPandSetenv("DbagentHost")

}
