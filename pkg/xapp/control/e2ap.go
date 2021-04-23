
package control
// #cgo CFLAGS: -I../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../internal/asn1codec/   -lasn1objects
/*
#include <stdlib.h>
#include <wrapper_e2ap.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type E2ap struct {
}

/* RICsubscriptionRequest */

func (c *E2ap) GetSubscriptionRequestSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_request_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Request Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

func (c *E2ap) SetSubscriptionRequestSequenceNumber(payload []byte, newSubscriptionid uint16) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	size := C.e2ap_set_ric_subscription_request_sequence_number(cptr, C.size_t(len(payload)), C.long(newSubscriptionid))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Request Sequence Number due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2ap) SetSubscriptionRequestPayload(payload []byte, ricRequestorID uint16, ricRequestSequenceNumber uint16, ranFunctionID uint16,
	eventTriggerDefinition []byte, eventTriggerDefinitionSize int, actionCount int, actionIds []int64, actionTypes []int64, actionDefinitions []ActionDefinition, subsequentActions []SubsequentAction) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	eventTrigger := unsafe.Pointer(&eventTriggerDefinition[0])
	actIds := unsafe.Pointer(&actionIds[0])
	actTypes := unsafe.Pointer(&actionTypes[0])

	count := len(actionDefinitions)
	actDefs := (*C.RICactionDefinition)(C.calloc(C.size_t(len(actionDefinitions)), C.sizeof_RICactionDefinition))
	for index := 0; index < count; index++ {
		ptr := *(*C.RICactionDefinition)(unsafe.Pointer((uintptr)(unsafe.Pointer(actDefs)) + (uintptr)(C.sizeof_RICactionDefinition*C.int(index))))
		ptr.size = C.int(actionDefinitions[index].Size)
		if ptr.size != 0 {
			ptr.actionDefinition = (*C.uint8_t)(C.CBytes(actionDefinitions[index].Buf))
		}
	}
	defer C.free(unsafe.Pointer(actDefs))

	count = len(subsequentActions)
	subActs := (*C.RICSubsequentAction)(C.calloc(C.size_t(len(subsequentActions)), C.sizeof_RICSubsequentAction))
	for index := 0; index < count; index++ {
		ptr := *(*C.RICSubsequentAction)(unsafe.Pointer((uintptr)(unsafe.Pointer(subActs)) + (uintptr)(C.sizeof_RICSubsequentAction*C.int(index))))
		ptr.isValid = C.int(subsequentActions[index].IsValid)
		ptr.subsequentActionType = C.long(subsequentActions[index].SubsequentActionType)
		ptr.timeToWait = C.long(subsequentActions[index].TimeToWait)
	}
	defer C.free(unsafe.Pointer(subActs))

	size := C.e2ap_encode_ric_subscription_request_message(cptr, C.size_t(len(payload)), C.long(ricRequestorID), C.long(ricRequestSequenceNumber),
		C.long(ranFunctionID), eventTrigger, C.size_t(eventTriggerDefinitionSize), C.int(actionCount), (*C.long)(actIds), (*C.long)(actTypes), actDefs, subActs)
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Request Payload due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

/* RICsubscriptionResponse */

func (c *E2ap) GetSubscriptionResponseSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_response_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Response Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

func (c *E2ap) SetSubscriptionResponseSequenceNumber(payload []byte, newSubscriptionid uint16) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	size := C.e2ap_set_ric_subscription_response_sequence_number(cptr, C.size_t(len(payload)), C.long(newSubscriptionid))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Response Sequence Number due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2ap) GetSubscriptionResponseMessage(payload []byte) (decodedMsg *DecodedSubscriptionResponseMessage, err error) {
	cptr := unsafe.Pointer(&payload[0])
	decodedMsg = &DecodedSubscriptionResponseMessage{}
	decodedCMsg := C.e2ap_decode_ric_subscription_response_message(cptr, C.size_t(len(payload)))
	defer C.free(unsafe.Pointer(decodedCMsg))

	if decodedCMsg == nil {
		return decodedMsg, errors.New("e2ap wrapper is unable to decode subscription response message due to wrong or invalid payload")
	}

	decodedMsg.RequestID = int32(decodedCMsg.requestorID)
	decodedMsg.RequestSequenceNumber = int32(decodedCMsg.requestSequenceNumber)
	decodedMsg.FuncID = int32(decodedCMsg.ranfunctionID)

	admittedCount := int(decodedCMsg.ricActionAdmittedList.count)
	for index := 0; index < admittedCount; index++ {
		decodedMsg.ActionAdmittedList.ActionID[index] = int32(decodedCMsg.ricActionAdmittedList.ricActionID[index])
	}
	decodedMsg.ActionAdmittedList.Count = admittedCount

	notAdmittedCount := int(decodedCMsg.ricActionNotAdmittedList.count)
	for index := 0; index < notAdmittedCount; index++ {
		decodedMsg.ActionNotAdmittedList.ActionID[index] = int32(decodedCMsg.ricActionNotAdmittedList.ricActionID[index])
		decodedMsg.ActionNotAdmittedList.Cause[index].CauseType = int32(decodedCMsg.ricActionNotAdmittedList.ricCause[index].ricCauseType)
		decodedMsg.ActionNotAdmittedList.Cause[index].CauseID = int32(decodedCMsg.ricActionNotAdmittedList.ricCause[index].ricCauseID)
	}
	decodedMsg.ActionNotAdmittedList.Count = notAdmittedCount

	return
}

/* RICsubscriptionFailure */

func (c *E2ap) GetSubscriptionFailureSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_failure_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Failure Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

/* RICsubscriptionDeleteRequest */

func (c *E2ap) GetSubscriptionDeleteRequestSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_delete_request_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Delete Request Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

func (c *E2ap) SetSubscriptionDeleteRequestSequenceNumber(payload []byte, newSubscriptionid uint16) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	size := C.e2ap_set_ric_subscription_delete_request_sequence_number(cptr, C.size_t(len(payload)), C.long(newSubscriptionid))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Delete Request Sequence Number due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2ap) SetSubscriptionDeleteRequestPayload(payload []byte, ricRequestorID uint16, ricRequestSequenceNumber uint16, ranFunctionID uint16) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	size := C.e2ap_encode_ric_subscription_delete_request_message(cptr, C.size_t(len(payload)), C.long(ricRequestorID), C.long(ricRequestSequenceNumber), C.long(ranFunctionID))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Delete Request Payload due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

/* RICsubscriptionDeleteResponse */

func (c *E2ap) GetSubscriptionDeleteResponseSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_delete_response_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Delete Response Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

func (c *E2ap) SetSubscriptionDeleteResponseSequenceNumber(payload []byte, newSubscriptionid uint16) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	size := C.e2ap_set_ric_subscription_delete_response_sequence_number(cptr, C.size_t(len(payload)), C.long(newSubscriptionid))
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Subscription Delete Response Sequence Number due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

/* RICsubscriptionDeleteFailure */

func (c *E2ap) GetSubscriptionDeleteFailureSequenceNumber(payload []byte) (subId uint16, err error) {
	cptr := unsafe.Pointer(&payload[0])
	cret := C.e2ap_get_ric_subscription_delete_failure_sequence_number(cptr, C.size_t(len(payload)))
	if cret < 0 {
		return 0, errors.New("e2ap wrapper is unable to get Subscirption Failure Sequence Number due to wrong or invalid payload")
	}
	subId = uint16(cret)
	return
}

/* RICindication */

func (c *E2ap) GetIndicationMessage(payload []byte) (decodedMsg *DecodedIndicationMessage, err error) {
	cptr := unsafe.Pointer(&payload[0])
	decodedMsg = &DecodedIndicationMessage{}
	decodedCMsg := C.e2ap_decode_ric_indication_message(cptr, C.size_t(len(payload)))
	if decodedCMsg == nil {
		return decodedMsg, errors.New("e2ap wrapper is unable to decode indication message due to wrong or invalid payload")
	}
	defer C.e2ap_free_decoded_ric_indication_message(decodedCMsg)

	decodedMsg.RequestID = int32(decodedCMsg.requestorID)
	decodedMsg.RequestSequenceNumber = int32(decodedCMsg.requestSequenceNumber)
	decodedMsg.FuncID = int32(decodedCMsg.ranfunctionID)
	decodedMsg.ActionID = int32(decodedCMsg.actionID)
	decodedMsg.IndSN = int32(decodedCMsg.indicationSN)
	decodedMsg.IndType = int32(decodedCMsg.indicationType)
	indhdr := unsafe.Pointer(decodedCMsg.indicationHeader)
	decodedMsg.IndHeader = C.GoBytes(indhdr, C.int(decodedCMsg.indicationHeaderSize))
	decodedMsg.IndHeaderLength = int32(decodedCMsg.indicationHeaderSize)
	indmsg := unsafe.Pointer(decodedCMsg.indicationMessage)
	decodedMsg.IndMessage = C.GoBytes(indmsg, C.int(decodedCMsg.indicationMessageSize))
	decodedMsg.IndMessageLength = int32(decodedCMsg.indicationMessageSize)
	callproc := unsafe.Pointer(decodedCMsg.callProcessID)
	decodedMsg.CallProcessID = C.GoBytes(callproc, C.int(decodedCMsg.callProcessIDSize))
	decodedMsg.CallProcessIDLength = int32(decodedCMsg.callProcessIDSize)
	return
}

//eventTriggerDefinition []byte, eventTriggerDefinitionSize int, actionCount int, actionIds []int64, actionTypes []int64,
//	actionDefinitions []ActionDefinition, subsequentActions []SubsequentAction
func (c *E2ap) SetControlRequestPayload(payload []byte, ricRequestorID uint16, ricRequestSequenceNumber uint16, ranFunctionID uint16,
	ControlHeader []byte,ControlHeaderLen C.long, ControlMessage []byte,ControlMessageLen C.long) (newPayload []byte, err error) {
	cptr := unsafe.Pointer(&payload[0])
	controlHeader := unsafe.Pointer(&ControlHeader[0])
	controlMessage := unsafe.Pointer(&ControlMessage[0])

	size := C.e2ap_encode_ric_control_request_message(cptr, C.size_t(len(payload)), C.long(ricRequestorID), C.long(ricRequestSequenceNumber),C.long(ranFunctionID),
		controlHeader, ControlHeaderLen, controlMessage,ControlMessageLen)
	if size < 0 {
		return make([]byte, 0), errors.New("e2ap wrapper is unable to set Control Request Payload due to wrong or invalid payload")
	}
	newPayload = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}