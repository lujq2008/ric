package msgx

import (
	"context"
	"fmt"
	"nRIC/internal"
	"nRIC/internal/kafka-go"
	"nRIC/internal/xapp"
	"time"
)

type KafkaMsgSender struct {
	Client *kafka.Writer
}
func NewKafkaMsgSender(topic string) *KafkaMsgSender {
	kafkaURL := internal.KafkaURL //os.Getenv("kafkaURL")
	return &KafkaMsgSender{
		&kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}}
}


func(s *KafkaMsgSender)String() string {
	return s.Client.Topic
}

func(s *KafkaMsgSender)SendMsg(req *xapp.MsgParams) error {
	msg := kafka.Message{}
	msg.Topic 			 = req.Src
	msg.Headers = make([]kafka.Header,4)
	msg.Headers[0].Value = []byte(fmt.Sprint(req.Mtype))
	msg.Headers[1].Value = []byte(fmt.Sprint(req.PayloadLen))
	msg.Headers[2].Value = []byte(fmt.Sprint(req.SubId))
	msg.Headers[3].Value = []byte(fmt.Sprint(req.Meid.RanName))
	msg.Value 			 = req.Payload

	var i int
	var err error
	for i=0;i<5;i++ {
		err = s.Client.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println("%d time(s): error while calling gRPC: %v\n", i+1,err)
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
	return err
}