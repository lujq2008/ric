

package sctpmsghandlerprovider

// #cgo CFLAGS: -I../../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../../internal/asn1codec/   -lasn1objects
// #include "asn_application.h"
// #include "E2AP-PDU.h"
// #include "asn_codecs.h"
// #include "asn1codec_utils.h"
// #include "InitiatingMessage.h"
// #include "RANfunctions-List.h"
// #include "E2SM-TS-RANFunctionDefinition.h"
import "C"
import (
	"fmt"
	"nRIC/internal/configuration"
	"nRIC/internal/logger"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"nRIC/pkg/nrice2t/handlers/sctpmsghandlers"
)

type NotificationHandlerProvider struct {
	notificationHandlers map[int]sctpmsghandlers.NotificationHandler
}

func NewNotificationHandlerProvider() *NotificationHandlerProvider {
	return &NotificationHandlerProvider{
		notificationHandlers: map[int]sctpmsghandlers.NotificationHandler{},
	}
}

// TODO: check whether it has been initialized
func (provider NotificationHandlerProvider) GetNotificationHandler(messageType int) (sctpmsghandlers.NotificationHandler, error) {
	handler, ok := provider.notificationHandlers[messageType]

	if !ok {
		return nil, fmt.Errorf("notification handler not found for message %d", messageType)
	}

	return handler, nil
}

func (provider *NotificationHandlerProvider) Register(msgType int, handler sctpmsghandlers.NotificationHandler) {
	provider.notificationHandlers[msgType] = handler
}



func (provider *NotificationHandlerProvider) Init(logger *logger.Logger, config *configuration.Configuration, grpcSender *dbclient.MsgSender) {

	e2SetupRequestNotificationHandler := sctpmsghandlers.NewE2SetupRequestNotificationHandler(logger, config, grpcSender)

	provider.Register(C.InitiatingMessage__value_PR_E2setupRequest, e2SetupRequestNotificationHandler)

}
