package models

// #cgo CFLAGS: -I../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../internal/asn1codec/   -lasn1objects
// #include "asn_application.h"
// #include "E2AP-PDU.h"
// #include "asn_codecs.h"
// #include "asn1codec_utils.h"
// #include "InitiatingMessage.h"
// #include "RANfunctions-List.h"
// #include "E2SM-TS-RANFunctionDefinition.h"
import "C"
import (
	"net"
	"unsafe"
)

type NotificationRequest struct {
	Pdu	  unsafe.Pointer
	conn  net.Conn
	Xml_pduBuf []byte
}

func NewNotificationRequest(Pdu	unsafe.Pointer,conn net.Conn,Xml_pduBuf []byte) *NotificationRequest {
	return &NotificationRequest{
		Pdu: Pdu,
		conn:conn,
		Xml_pduBuf: Xml_pduBuf,
	}
}

func (r NotificationRequest) GetConn() net.Conn {
	return r.conn
}

func (r NotificationRequest) GetChoice() *byte {
	return &(((*C.E2AP_PDU_t)(r.Pdu)).choice[0])
}