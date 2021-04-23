package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/oklog/oklog/pkg/group"
	"github.com/spf13/viper"
	"nRIC/internal"
	"nRIC/internal/configuration"
	"nRIC/internal/msgx"
	"nRIC/internal/utils"
	"nRIC/internal/xapp"
	"nRIC/pkg/xapp/control"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Register(c control.Control){
	//msg client ：nric mgmt ip/port 是知名的服务地址，预配置
	nricmgmthost := viper.GetString("nricmgmthost")
	xapp.Logger.Error("nricmgmthost = %s\n",nricmgmthost)
	XappIpaddr,_ := utils.GetHostIp()
	xapp.Logger.Info("XappIpaddr = %s\n",XappIpaddr)

	c.Msg2MgmtClient = msgx.NewMsgSender(nricmgmthost,internal.DefaultGRPCPort)

	XappRegMsg := msgx.XappRegMsg{
		Header:&msgx.XappMsgHeader{
			MsgType:xapp.RIC_O1_REGISTER,
			MsgVer: 0.1,
			Token: "0011223333221100",
			XappRequestID: &msgx.XAPPRequestID{
				XappID: 0,
				XappInstanceID: 1,
			},
		},
			XappVer:0.1,
			XappName:"Xapp_TS",
			XappFunctions: "1,2",
			XappIpaddr:XappIpaddr,
			XappPort:internal.DefaultGRPCPort,
	}
	xappmsgreg,err := proto.Marshal(&XappRegMsg)
	if err != nil {
		xapp.Logger.Error("Failed to encode xappmsgreg:", err)
	}

	xapp.Logger.Info("Send RIC_O1_REGISTER msg to Nricmgmt !")
	xapp.NewRouter()

	for {
		time.Sleep(5 * time.Second)
		if !xapp.IsHealthProbeReady() {
			xapp.Logger.Info("xApp is not ready yet, waiting ...")
			continue
		}
		xapp.Logger.Info("xApp is now up and ready, continue with registration ...")
		break
	}

	for {
		err := c.Msg2MgmtClient.SendMsg(&xapp.MsgParams{
			Mtype: xapp.RIC_O1_REGISTER,
			Payload: xappmsgreg,
			PayloadLen: len(xappmsgreg),
			Meid: &xapp.MsgMeid{RanName: ""},
		})
		//注册消息发送失败
		if err != nil {
			xapp.Logger.Warn("xApp registration failed, continue with registration ...")
			time.Sleep(5 * time.Second)
			continue
		}

		time.Sleep(5 * time.Second)
		//收到了注册成功响应消息，注册成功
		if c.Registered {
			xapp.Logger.Info("xApp registration done, proceeding with startup ...")
			break
		}
		//没收到注册响应消息，或者收到的是注册失败的响应消息，继续发起注册
	}
}

func xapp_ts() {

	c := control.NewControl(nil,nil)

    d := &control.Display{&c}
	go d.DisplayHttpThread()
	go d.Workloop()
	//msg server addr
	//grpcXappAddr := net.JoinHostPort(internal.XappHost, envString("GRPC_PORT", internal.DefaultGRPCPort))

	// get kafka reader using environment variables. "nric-kafka.default.svc.cluster.local:9092"
	//kafkaURL := os.Getenv("kafkaURL")
	//msg server

	grpcAddr := net.JoinHostPort("0.0.0.0", envString("GRPC_PORT", internal.DefaultGRPCPort))
	c.Run(grpcAddr)

	Register(c)
}

func main() {

	config := configuration.ParseConfiguration("xapp")
	xapp.Logger.SetLevel(int(config.Logging.LogLevel))   //golog.INFO is 3  ,DEBUG is 4 ,ERR is 1 ,WARN is 2
	xapp.Logger.Info("xapp")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("xapp")
	viper.AllowEmptyEnv(true)

	xapp.Logger.Info("#app.main - Configuration %s\n", config)
	var g group.Group

	go xapp_ts()

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

