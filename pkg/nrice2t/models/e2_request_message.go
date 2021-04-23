package models

import (
	"fmt"
	"nRIC/internal/logger"
)

type E2RequestMessage struct {
	transactionId string
	ranIp         string
	ranPort       uint16
	ranName       string
	payload       []byte
}

func (e2RequestMessage E2RequestMessage) RanName() string {
	return e2RequestMessage.ranName
}

func (e2RequestMessage E2RequestMessage) TransactionId() string {
	return e2RequestMessage.transactionId
}

func NewE2RequestMessage(transactionId string, ranIp string, ranPort uint16, ranName string, payload []byte) *E2RequestMessage {
	return &E2RequestMessage{transactionId: transactionId, ranIp: ranIp, ranPort: ranPort, ranName: ranName, payload: payload}
}

// TODO: this shouldn't receive logger
func (e2RequestMessage E2RequestMessage) GetMessageAsBytes(logger *logger.Logger) []byte {
	messageStringWithoutPayload := fmt.Sprintf("%s|%d|%s|%d|", e2RequestMessage.ranIp, e2RequestMessage.ranPort, e2RequestMessage.ranName, len(e2RequestMessage.payload))
	logger.Debugf("#e2_request_message.GetMessageAsBytes - messageStringWithoutPayload: %s", messageStringWithoutPayload)
	messageBytesWithoutPayload := []byte(messageStringWithoutPayload)
	return append(messageBytesWithoutPayload, e2RequestMessage.payload...)
}
