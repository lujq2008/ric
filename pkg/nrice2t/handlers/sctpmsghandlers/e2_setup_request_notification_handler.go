package sctpmsghandlers

// #cgo CFLAGS: -I../../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../../internal/asn1codec/   -lasn1objects
// #include "asn_application.h"
// #include "E2AP-PDU.h"
// #include "asn_codecs.h"
// #include "asn1codec_utils.h"
// #include "InitiatingMessage.h"
// #include "RANfunctions-List.h"
// #include "E2SM-TS-RANFunctionDefinition.h"
// #include "BuildRunName.h"
import "C"
import (
	"encoding/xml"
	"errors"
	"fmt"
	pbdb "nRIC/api/v1/pb/db"
	"nRIC/internal/configuration"
	"nRIC/internal/logger"
	"nRIC/internal/utils"
	"nRIC/internal/xapp"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"nRIC/pkg/nrice2t/models"
	"strconv"
	"strings"
	"unsafe"
)

var (
	emptyTagsToReplaceToSelfClosingTags = []string{"reject", "ignore", "transport-resource-unavailable", "om-intervention", "request-id-unknown",
		"v60s", "v20s", "v10s", "v5s", "v2s", "v1s"}
)

type E2SetupRequestNotificationHandler struct {
	logger                        *logger.Logger
	config                        *configuration.Configuration
	grpcSender                     *dbclient.MsgSender
}

func NewE2SetupRequestNotificationHandler(logger *logger.Logger, config *configuration.Configuration, grpcSender *dbclient.MsgSender) *E2SetupRequestNotificationHandler {
	return &E2SetupRequestNotificationHandler{
		logger:                        logger,
		config:                        config,
		grpcSender:                     grpcSender,
	}
}

func (h *E2SetupRequestNotificationHandler) Handle(Request *models.NotificationRequest) {
	Pdu := Request.Pdu
	h.logger.Infof("#E2SetupRequestNotificationHandler.Handle - pdu: %x", Pdu)

	var ranfunctions_slice []*C.RANfunction_Item_t
	var ranfGoslice []*pbdb.RANFunctionsTable
	var RanName []byte = make([]byte,256)
	var RanNameStr string

	initiatingMessage := *(**C.InitiatingMessage_t)(unsafe.Pointer(Request.GetChoice()))
	e2setupRequest := *(*C.E2setupRequest_t)(unsafe.Pointer(&initiatingMessage.value.choice[0]))
	if e2setupRequest.protocolIEs.list.count > 0 {
		count := int(e2setupRequest.protocolIEs.list.count)
		protocolIEs_slice := (*[1 << 30]*C.E2setupRequestIEs_t)(unsafe.Pointer(e2setupRequest.protocolIEs.list.array))[:count:count]
		for _,protocolIE := range protocolIEs_slice {
			switch protocolIE.value.present {
			case C.E2setupRequestIEs__value_PR_GlobalE2node_ID:
				fmt.Printf("%s\n","E2setupRequestIEs__value_PR_GlobalE2node_ID")
				size := C.buildRanName((*C.char)(unsafe.Pointer(&RanName[0])),protocolIE)
				RanName := RanName[:size]
				RanNameStr = string(RanName)
				xapp.Logger.Error("size = %d, RanName = %s\n",size,RanNameStr)

			case C.E2setupRequestIEs__value_PR_RANfunctions_List:
				fmt.Printf("%s\n","E2setupRequestIEs__value_PR_RANfunctions_List")
				ranfunctions_List := *(*C.RANfunctions_List_t)(unsafe.Pointer(&protocolIE.value.choice[0]))
				if ranfunctions_List.list.count > 0 {
					count := int(ranfunctions_List.list.count)
					//ranfunctions_slice := (*[1 << 30]*C.RANfunction_Item_t)(unsafe.Pointer(ranfunctions_List.list.array))[:count:count]
					ranfunctions_ItemIEs := (*[1 << 30]*C.RANfunction_ItemIEs_t)(unsafe.Pointer(ranfunctions_List.list.array))[:count:count]
					for _,ranfunctions_ItemIE := range ranfunctions_ItemIEs {
						if ranfunctions_ItemIE.value.present == C.RANfunction_ItemIEs__value_PR_RANfunction_Item {
							ranfunction := (*C.RANfunction_Item_t)(unsafe.Pointer(&ranfunctions_ItemIE.value.choice[0]))
							//ranfunctions_slice = append(ranfunctions_slice,ranfunction)
							ranFDbuf := C.GoBytes(unsafe.Pointer(ranfunction.ranFunctionDefinition.buf), C.int(ranfunction.ranFunctionDefinition.size))
							ranfGoslice = append(ranfGoslice,&pbdb.RANFunctionsTable{
								RanFunctionID:uint32(C.long(ranfunction.ranFunctionID)),
								RanFunctionRevision:uint32(C.long(ranfunction.ranFunctionRevision)),
								/*TBD : RanFunctionOID:ranfunction.ranFunctionOID.buf, */
								RanFunctionOID:"1.3.6.1.4.1.53148.1.2.255",
								RanFunctionDefinition:ranFDbuf,
							})
						}
						//return ranfunctions_slice
					}
				}
			default:
				fmt.Printf("%s\n","NOT E2setupRequestIEs__value_PR_RANfunctions_List")
			}
			xapp.Logger.Error("ranfGoslice len is %d\n",len(ranfGoslice))

			for _,r := range ranfGoslice {
				r.GlobalE2NodeIDStr = RanNameStr
				h.grpcSender.RANFunctionsTableInsert(&pbdb.RANFunctionsTableInsertRequest{RANFunctionsTable: r})
			}
		}
	}


	for _,ranfunction_Item := range ranfunctions_slice {
		ranfunction := ranfunction_Item.ranFunctionDefinition
		fmt.Printf("\nranfunction size: %d\n",ranfunction.size)
		fmt.Printf("\nranfunction buf : %x\n",C.GoBytes(unsafe.Pointer(ranfunction.buf), C.int(ranfunction.size)))
		_,err := h.ParseRANFunction(ranfunction)
		if err != nil {
			cause := models.Cause{RicRequest: &models.CauseRic{RequestIdUnknown: &struct{}{}}}
			h.handleUnsuccessfulResponse(Request, cause)
			return
		}

	}

	setupRequest := &models.E2SetupRequestMessage{}
	err := xml.Unmarshal(utils.NormalizeXml(Request.Xml_pduBuf), &setupRequest.E2APPDU)
	if err != nil {
		fmt.Printf("#E2SetupRequestNotificationHandler - Error unmarshalling E2 Setup Request payload: %x", Request.Xml_pduBuf)
	}

	h.handleSuccessfulResponse(Request, setupRequest)

	cause := models.Cause{Misc: &models.CauseMisc{OmIntervention: &struct{}{}}}
	h.handleUnsuccessfulResponse(Request, cause)

}


func (h *E2SetupRequestNotificationHandler) ParseRANFunction(ranfunction C.OCTET_STRING_t) (*C.E2SM_TS_RANFunctionDefinition_t,error) {
	ts_ranfunction := C.new_ranfunction(C.size_t(1))
	ranbuf := C.GoBytes(unsafe.Pointer(ranfunction.buf), C.int(ranfunction.size))

	errBuf := make([]C.char, 4096)
	ranfunctionBuf := make([]C.char, 4096)
	err := errors.New("uppack ranfunction failed!")

	if C.per_unpack_ranfunction(ts_ranfunction, ranfunction.size, unsafe.Pointer(&ranbuf[0]) , C.size_t(len(errBuf)), (*C.char)(unsafe.Pointer(&errBuf[0]))) {
		C.asn1_ranfunction_printer(ts_ranfunction, C.size_t(len(ranfunctionBuf)), &ranfunctionBuf[0])
		pduAsString := C.GoString(&ranfunctionBuf[0])
		//m.logger.Infof("CUBU3:  %s",pduAsString)
		fmt.Printf("\nE2SM_TS_RANfunction_Description:\n%s\n", pduAsString)

	} else {
		h.logger.Errorf(("CUBU : FAILED1 of C.per_unpack_ranfunction"))
		h.logger.Errorf("errBuf: %v", C.GoString(&errBuf[0]))
		return nil,err
	}
	return ts_ranfunction,nil
}


func (h *E2SetupRequestNotificationHandler) handleSuccessfulResponse(req *models.NotificationRequest, setupRequest *models.E2SetupRequestMessage) {

	plmnId := buildPlmnId(h.config.GlobalRicId.Mcc, h.config.GlobalRicId.Mnc)

	ricNearRtId, err := convertTo20BitString(h.config.GlobalRicId.RicId)
	if err != nil {
		return
	}
	successResponse := models.NewE2SetupSuccessResponseMessage(plmnId, ricNearRtId, setupRequest)
	h.logger.Infof("#E2SetupRequestNotificationHandler.handleSuccessfulResponse - E2_SETUP_RESPONSE has been built successfully %+v", successResponse)

	responsePayload, err := xml.Marshal(&successResponse.E2APPDU)
	if err != nil {
		h.logger.Warnf("#E2SetupRequestNotificationHandler.handleSuccessfulResponse -  Error marshalling RIC_E2_SETUP_RESP. Payload: %s",  responsePayload)
	}

	responsePayload = utils.ReplaceEmptyTagsWithSelfClosing(responsePayload,emptyTagsToReplaceToSelfClosingTags)

	h.logger.Infof("#E2SetupRequestNotificationHandler.handleSuccessfulResponse - payload: %s", responsePayload)

	h.sctpSendPackage(responsePayload,req)

}


func buildPlmnId(mmc string, mnc string) string {
	var b strings.Builder

	b.WriteByte(mmc[1])
	b.WriteByte(mmc[0])
	if len(mnc) == 2 {
		b.WriteString("F")
	} else {
		b.WriteByte(mnc[2])
	}
	b.WriteByte(mmc[2])
	b.WriteByte(mnc[1])
	b.WriteByte(mnc[0])

	return b.String()
}


func convertTo20BitString(ricNearRtId string) (string, error) {
	r, err := strconv.ParseUint(ricNearRtId, 16, 32)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%020b", r)[:20], nil
}


func (h *E2SetupRequestNotificationHandler) handleUnsuccessfulResponse(req *models.NotificationRequest, cause models.Cause) {
	failureResponse := models.NewE2SetupFailureResponseMessage(models.TimeToWaitEnum.V60s, cause)
	h.logger.Debugf("#E2SetupRequestNotificationHandler.handleUnsuccessfulResponse - E2_SETUP_RESPONSE has been built successfully %+v", failureResponse)

	responsePayload, err := xml.Marshal(&failureResponse.E2APPDU)
	if err != nil {
		h.logger.Warnf("#E2SetupRequestNotificationHandler.handleUnsuccessfulResponse - Error marshalling RIC_E2_SETUP_RESP. Payload: %s",  responsePayload)
	}

	responsePayload = utils.ReplaceEmptyTagsWithSelfClosing(responsePayload,emptyTagsToReplaceToSelfClosingTags)

	h.logger.Infof("#E2SetupRequestNotificationHandler.handleUnsuccessfulResponse - payload: %s", responsePayload)

	h.sctpSendPackage(responsePayload,req)
}

func (h *E2SetupRequestNotificationHandler)sctpSendPackage(payload []byte,req *models.NotificationRequest){

	//c unpack xml
	// unpack_pdu_aux(E2AP_PDU_t *pdu, size_t packed_buf_size, const void* packed_buf,size_t err_buf_size, char* err_buf,enum asn_transfer_syntax syntax);

	mbuf := payload
	errBuf := make([]C.char, 4096)
	pduBuf := make([]C.char, 4096)

	Pdu := C.new_pdu(C.size_t(1)) //TODO: change signature  (*C.uchar)

	if C.unpack_pdu_aux(Pdu, (C.size_t)(len(mbuf)), unsafe.Pointer(&mbuf[0]), C.size_t(len(errBuf)), (*C.char)(unsafe.Pointer(&errBuf[0])),C.ATS_BASIC_XER) {
		C.asn1_pdu_printer(Pdu, C.size_t(len(pduBuf)), &pduBuf[0])
		//pduAsString := C.GoString(&pduBuf[0])
		//fmt.Printf("\n XML C.unpack_pdu_aux: %s\n",pduAsString)

	} else {
		h.logger.Errorf(("CUBU : FAILED3 of xml C.unpack_pdu_aux"))
		h.logger.Errorf("errBuf: %v",C.GoString(&errBuf[0]))
	}

	// c pack per
	// per_pack_pdu(E2AP_PDU_t *pdu, size_t *packed_buf_size, unsigned char* packed_buf,size_t err_buf_size, char* err_buf)
	var sendbuf_size *C.ulong
	sendbuf_size  = (*C.ulong)(C.calloc(1,8))
	*sendbuf_size = 4096
	sendbuf := make([]byte, 4096)

	if C.per_pack_pdu(Pdu, (*C.size_t)(sendbuf_size), (*C.uchar)(unsafe.Pointer(&sendbuf[0])), C.size_t(len(errBuf)), (*C.char)(unsafe.Pointer(&errBuf[0]))) {
		//pduAsString := C.GoString((*C.char)(unsafe.Pointer(&sendbuf[0])))
		//m.logger.Infof("CUBU3:  %s",pduAsString)
		fmt.Printf("\nPER sendbuf : %x \n",sendbuf[:(*sendbuf_size)])

	} else {
		h.logger.Errorf(("CUBU : FAILED of PER C.pack_pdu_aux"))
		h.logger.Errorf("errBuf: %v",C.GoString(&errBuf[0]))
	}

	prefix := make([]byte,32)
	var data []byte =  sendbuf[:(*sendbuf_size)]
	send := append(prefix,data...)
	n, err := req.GetConn().Write(send)
	fmt.Printf("\n number = %d,err =%s\n",n,err)
	//fmt.Printf("\nPER sendbuf : %x \n",send)

	//h.grpcSender.Callmgmtservice()
}
