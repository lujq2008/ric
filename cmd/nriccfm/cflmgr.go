package main

import (
	"fmt"
	"nRIC/internal"
	"nRIC/internal/configuration"
	"nRIC/internal/logger"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"nRIC/pkg/nriccfm/control"
	e2tclient "nRIC/pkg/nrice2t/route"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
)

func cflmgr() {

	//msg client,to xapp
	//MsgSendertoXapp := msgx.NewMsgSender(internal.XappHost,internal.DefaultGRPCPort)

	//db client,to dbagent
	AccessDbAgent := dbclient.NewMsgSender(internal.DbagentHost,internal.DefaultGRPCPort)

	//msgSendertoXapp := msgx.SenderIf(MsgSendertoXapp)
	//MsgSendertoE2T := msgx.NewMsgSender(internal.Nrice2tHost,internal.DefaultGRPCPort)

	MsgSendertoE2T := e2tclient.NewMsgSender(internal.Nrice2tHost,internal.DefaultGRPCPort)
	c := control.NewControl(MsgSendertoE2T,AccessDbAgent)

	//msg server
	grpcAddr := net.JoinHostPort(internal.NriccflmHost, envString("GRPC_PORT", internal.DefaultGRPCPort))
	c.Run(grpcAddr)
}


func main() {
	config := configuration.ParseConfiguration("nriccfm")
	logLevel, _ := logger.LogLevelTokenToLevel("info")
	logger, err := logger.InitLogger(logLevel)
	if err != nil {
		fmt.Printf("#app.main - failed to initialize logger, error: %s\n", err)
		os.Exit(1)
	}
	logger.Infof("#app.main - Configuration %s\n", config)


	go cflmgr()

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
	logger.Infof("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}