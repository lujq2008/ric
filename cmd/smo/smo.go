package main

import (
	"fmt"
	"github.com/spf13/viper"
	"nRIC/internal"
	"nRIC/internal/configuration"
	"nRIC/internal/msgx"
	"nRIC/internal/xapp"
	"nRIC/pkg/smo/control"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
)

func smo() {

	//msg client,to nirc mgmt
	MsgSendertoNricmgmt := msgx.NewMsgSender(internal.NricmgmtHost,internal.DefaultGRPCPort)


	c := control.NewControl(MsgSendertoNricmgmt)

	//msg server
	grpcAddr := net.JoinHostPort(internal.Smo, envString("GRPC_PORT", internal.DefaultGRPCPort))
	c.Run(grpcAddr)
}


func main() {
	config := configuration.ParseConfiguration("smo")

	xapp.Logger.SetLevel(int(config.Logging.LogLevel))   //golog.INFO is 3  ,DEBUG is 4 ,ERR is 1 ,WARN is 2
	xapp.Logger.Info("smo")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("smo")
	viper.AllowEmptyEnv(true)

	xapp.Logger.Info("#app.main - Configuration %s\n", config)


	go smo()

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