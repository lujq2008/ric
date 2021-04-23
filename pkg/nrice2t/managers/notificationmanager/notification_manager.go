package notificationmanager

//-lasncodec -le2sim

// #cgo CFLAGS: -I../../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../../internal/asn1codec/   -lasn1objects
// #include "asn_application.h"
// #include "E2AP-PDU.h"
// #include "asn_codecs.h"
// #include "asn1codec_utils.h"
// #include "InitiatingMessage.h"
// #include "SuccessfulOutcome.h"
// #include "UnsuccessfulOutcome.h"
// #include "RANfunctions-List.h"
// #include "E2SM-TS-RANFunctionDefinition.h"
// #include <stdlib.h>
// #include <stddef.h>
import "C"
import (
	"fmt"
	"nRIC/internal/logger"
	"nRIC/internal/xapp"
	"nRIC/pkg/nrice2t/models"
	"nRIC/pkg/nrice2t/providers/sctpmsghandlerprovider"
	"net"
	"unsafe"
)

type NotificationManager struct {
	logger                      *logger.Logger
	notificationHandlerProvider *sctpmsghandlerprovider.NotificationHandlerProvider
	c xapp.MessageConsumer
}

func NewNotificationManager(logger *logger.Logger, notificationHandlerProvider *sctpmsghandlerprovider.NotificationHandlerProvider,c xapp.MessageConsumer) *NotificationManager {
	return &NotificationManager{
		logger:                      logger,
		notificationHandlerProvider: notificationHandlerProvider,
		c: 							 c,
	}
}

func (m NotificationManager) HandleMessage(mbuf []byte,conn net.Conn) error {

	//m.logger.Infof("CUBU : NotificationManager  HandleMessage")
	//m.logger.Infof("mbuf len: %v", len(mbuf))
	//fmt.Printf("\nmbuf====:%x\n",mbuf)

	errBuf := make([]C.char, 4096)
	//pduBuf := make([]C.char, 4096)

	Pdu := C.new_pdu(C.size_t(1)) //TODO: change signature  (*C.uchar)

	if C.per_unpack_pdu(Pdu, (C.size_t)(len(mbuf)), unsafe.Pointer(&mbuf[0]), C.size_t(len(errBuf)), (*C.char)(unsafe.Pointer(&errBuf[0]))) {
		//C.asn1_pdu_printer(Pdu, C.size_t(len(pduBuf)), &pduBuf[0])
		//pduAsString := C.GoString(&pduBuf[0])
		//m.logger.Infof("CUBU3:  %s",pduAsString)
		//fmt.Printf("\n%s\n",pduAsString)

	} else {
		m.logger.Errorf(("CUBU : FAILED3 of C.per_unpack_pdu"))
		m.logger.Errorf("errBuf: %v",C.GoString(&errBuf[0]))
	}
	//initiatingMessage := *(**C.InitiatingMessage_t)(unsafe.Pointer(&pdu.choice[0]))
	//fmt.Printf("initiatingMessage.value.present = %d\n",initiatingMessage.value.present)
	/*TBD:
		如果是E2SETUP请求消息：
			需要提取其中的GlobalE2node_ID，并和conn对象映射起来。
			后续来自xapp，需要发往基站侧的消息，都需要携带GlobalE2node_ID，以其为index查找映射表获取conn对象，才能将消息发往正确的基站。

			提取RANparameter_Item信息，记录到以GlobalE2node_ID为index的表中。这些信息代表了基站支持的参数能力集。xapp只能订阅这些参数集范围内的功能。

			最后，发出响应消息给基站侧。
	*/

	//xml encode
	// er = asn_encode_to_buffer(nullptr, ATS_BASIC_XER, &asn_DEF_E2AP_PDU, pdu, buffer, buffer_size);
	var xml_pduBuf_size *C.ulong
	xml_pduBuf_size = (*C.ulong)(C.calloc(1,8))
	*xml_pduBuf_size = 8192  //512*1024
	//xml_pduBuf := make([]C.uchar, *xml_pduBuf_size)
	xml_pduBuf := make([]byte, *xml_pduBuf_size)


	if C.pack_pdu_aux(Pdu, (*C.size_t)(xml_pduBuf_size), (*C.uchar)(unsafe.Pointer(&xml_pduBuf[0])), C.size_t(len(errBuf)), (*C.char)(unsafe.Pointer(&errBuf[0])),C.ATS_BASIC_XER) {
		//C.asn1_pdu_xer_printer((*C.E2AP_PDU_t)(unsafe.Pointer(&xml_pduBuf[0])), C.size_t(*xml_pduBuf_size), &pduBuf[0])
		pduAsString := C.GoString((*C.char)(unsafe.Pointer(&xml_pduBuf[0])))
		//m.logger.Infof("CUBU3:  %s",pduAsString)
		xapp.Logger.Debug("\nxml_pdu: %s\n",pduAsString)

	} else {
		m.logger.Errorf(("CUBU : FAILED4 of C.pack_pdu_aux"))
		m.logger.Errorf("errBuf: %v",C.GoString(&errBuf[0]))
	}

	m.DispatchMessage(Pdu,conn,xml_pduBuf,mbuf)

	return nil
}


func (m NotificationManager) DispatchMessage(Pdu *C.E2AP_PDU_t,conn net.Conn ,xml_pduBuf []byte,mbuf []byte) ([]*C.RANfunction_Item_t) {

	if Pdu.present == C.E2AP_PDU_PR_initiatingMessage {
		initiatingMessage := *(**C.InitiatingMessage_t)(unsafe.Pointer(&Pdu.choice[0]))
		switch(initiatingMessage.value.present){
		case C.InitiatingMessage__value_PR_E2setupRequest:
			//fmt.Printf("%s\n","InitiatingMessage__value_PR_E2setupRequest")
			handler,_ := m.notificationHandlerProvider.GetNotificationHandler(C.InitiatingMessage__value_PR_E2setupRequest)
			NotificationRequest := models.NewNotificationRequest((unsafe.Pointer)(Pdu),conn,xml_pduBuf)
			handler.Handle(NotificationRequest)

		case C.InitiatingMessage__value_PR_RICindication:
			var RicInstanceID int
			//fmt.Printf("%s\n","InitiatingMessage__value_PR_RICindication")
			RICindication := *(*C.RICindication_t)(unsafe.Pointer(&initiatingMessage.value.choice[0]))
			count := RICindication.protocolIEs.list.count
			RICindication_slice := (*[1 << 30]*C.RICindication_IEs_t)(unsafe.Pointer(RICindication.protocolIEs.list.array))[:count:count]

			for _,RICindication_IE := range RICindication_slice{
				switch (RICindication_IE.value.present){
				case C.RICindication_IEs__value_PR_RICrequestID:
					RICrequestID := *(*C.RICrequestID_t)(unsafe.Pointer(&RICindication_IE.value.choice[0]))
					//RicRequestorID := RICrequestID.ricRequestorID
					RicInstanceID = int(RICrequestID.ricInstanceID)
				}

			}

			//send mbuf to submgr
			//fmt.Printf("\nConsume Msg: %x \n",mbuf)
			params := &xapp.MsgParams{}
			params.Mtype 		= int(xapp.RIC_INDICATION)
			params.Payload 		= mbuf
			params.PayloadLen 	= len(mbuf)
			params.Meid		    = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}  //tbd: Globle gnoteb id
			params.SubId		= RicInstanceID
			err := m.c.Consume(params)
			//fmt.Printf("\nConsume Msg: %x \n",mbuf)
			if err != nil {
				xapp.Logger.Error("Error: %s", err)
			}
			return  nil


		default:
			fmt.Printf("%s\n","Not E2setupRequest")
		}
	} else if (Pdu.present == C.E2AP_PDU_PR_successfulOutcome) {
		successfulOutcome := *(**C.SuccessfulOutcome_t)(unsafe.Pointer(&Pdu.choice[0]))
		switch (successfulOutcome.value.present) {
		case C.SuccessfulOutcome__value_PR_RICsubscriptionResponse:
			//fmt.Printf("%s\n", "SuccessfulOutcome__value_PR_RICsubscriptionResponse")

			//send mbuf to submgr
			//fmt.Printf("\nConsume Msg: %x \n",mbuf)
			params := &xapp.MsgParams{}
			params.Mtype 		= int(xapp.RIC_SUB_RESP)
			params.Payload 		= mbuf
			params.PayloadLen 	= len(mbuf)
			params.Meid		    = &xapp.MsgMeid{PlmnID: "373437", EnbID: "10110101110001100111011110001", RanName: "gnb_734_733_b5c67788"}  //tbd: Globle gnoteb id
			err := m.c.Consume(params)
			//fmt.Printf("\nConsume Msg: %x \n",mbuf)
			if err != nil {
				xapp.Logger.Error("Error: %s", err)
			}
			return  nil

		default:
			fmt.Printf("%s\n", "Not SuccessfulOutcome__value_PR_RICsubscriptionResponse")
		}
	} else {
		fmt.Printf("%s\n","NOT initiatingMessage,NOT successfulOutcome")
	}


	return nil
}

