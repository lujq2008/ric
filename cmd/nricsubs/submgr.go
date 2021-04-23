package main

import (
	"fmt"
	"nRIC/internal"
	"nRIC/internal/configuration"
	"nRIC/internal/xapp"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	e2tclient "nRIC/pkg/nrice2t/route"
	"nRIC/pkg/nricsubs/control"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
)

func submgr() {

	//db client,to dbagent
	AccessDbAgent := dbclient.NewMsgSender(internal.DbagentHost,internal.DefaultGRPCPort)

	MsgSendertoE2T := e2tclient.NewMsgSender(internal.Nrice2tHost,internal.DefaultGRPCPort)
	c := control.NewControl(MsgSendertoE2T,AccessDbAgent)

	//msg server
	grpcAddr := net.JoinHostPort(internal.SubmgrHost, envString("GRPC_PORT", internal.DefaultGRPCPort))
	c.Run(grpcAddr)
}


func main() {
	config := configuration.ParseConfiguration("nricsubs")
	xapp.Logger.Info("#app.main - Configuration %s\n", config)


	go submgr()

	var g group.Group


	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	xapp.Logger.Info("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}