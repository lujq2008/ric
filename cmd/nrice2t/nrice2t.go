package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
	"nRIC/internal"
	"nRIC/internal/kafka-go"
	"nRIC/internal/msgx"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"nRIC/pkg/nrice2t/control"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)



func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}


func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		BatchTimeout: 0,
		BatchSize:1, //cnbu
	}
}



func main() {
	var (
		logger   log.Logger
	)
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var g group.Group
	{
		//go nrice2t.Initsctp(&logger)
		MsgSendertoSubmgr := msgx.NewMsgSender(internal.SubmgrHost,internal.DefaultGRPCPort)
		MsgSendertoDbagent := dbclient.NewMsgSender(internal.DbagentHost,internal.DefaultGRPCPort)

		c := control.NewControl(MsgSendertoSubmgr,MsgSendertoDbagent)
		c.Sctpreceiver = c.SetupSctp(MsgSendertoDbagent)

		//msg and route server
		grpcAddr := net.JoinHostPort(internal.Nrice2tHost, envString("GRPC_PORT", internal.DefaultGRPCPort))
		c.Run(grpcAddr)

	}

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
	logger.Log("exit", g.Run())

/*
	//kafka
	{
		// get kafka writer using environment variables.
		kafkaURL := "nric-kafka-0.nric-kafka-headless.default.svc.cluster.local:9092" //os.Getenv("kafkaURL")
		topic := os.Getenv("topic")
		writer := newKafkaWriter(kafkaURL, topic)
		defer writer.Close()
		fmt.Println("start producing ... !!")
		for i := 0; i < 100 ; i++ {
			msg := kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(fmt.Sprint(uuid.New())),
			}
			err := writer.WriteMessages(context.Background(), msg)
			if err != nil {
				fmt.Println(err)
			}
			//time.Sleep(1 * time.Second)
			fmt.Println("send msg: %s = %s\n",string(msg.Key),string(msg.Value))
		}

	}
*/

}



func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e

}

func init(){
	//internal.Nrice2tHost = internal.GetIPandSetenv("Nrice2tHost")
}