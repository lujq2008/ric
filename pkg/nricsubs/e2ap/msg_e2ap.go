

package e2ap

import (
	"nRIC/internal/conv"
	"strconv"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
const (
	E2AP_InitiatingMessage   uint64 = 1
	E2AP_SuccessfulOutcome   uint64 = 2
	E2AP_UnsuccessfulOutcome uint64 = 3
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
// E2AP messages
// Initiating message
const (
	E2AP_RICSubscriptionRequest       uint64 = 1
	E2AP_RICSubscriptionDeleteRequest uint64 = 2

)

// E2AP messages
// Successful outcome
const (
	E2AP_RICSubscriptionResponse       uint64 = 1
	E2AP_RICSubscriptionDeleteResponse uint64 = 2

)

// E2AP messages
// Unsuccessful outcome
const (
	E2AP_RICSubscriptionFailure       uint64 = 1
	E2AP_RICSubscriptionDeleteFailure uint64 = 2

)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type PackedData struct {
	Buf []byte
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type MessageInfo struct {
	MsgType uint64
	MsgId   uint64
}

func (msgInfo *MessageInfo) String() string {
	return "msginfo(" + strconv.FormatUint((uint64)(msgInfo.MsgType), 10) + string(":") + strconv.FormatUint((uint64)(msgInfo.MsgId), 10) + ")"
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type RequestId struct {
	Id         uint32
	InstanceId uint32
}

func (rid *RequestId) String() string {
	return strconv.FormatUint((uint64)(rid.Id), 10) + string(":") + strconv.FormatUint((uint64)(rid.InstanceId), 10)
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type FunctionId uint16

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

const (
	E2AP_ENBIDMacroPBits20    uint8 = 20
	E2AP_ENBIDHomeBits28      uint8 = 28
	E2AP_ENBIDShortMacroits18 uint8 = 18
	E2AP_ENBIDlongMacroBits21 uint8 = 21
)

type NodeId struct {
	Bits uint8
	Id   uint32
}

func (nid *NodeId) String() string {
	return strconv.FormatUint((uint64)(nid.Id), 10)
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type GlobalNodeId struct {
	Present      bool
	PlmnIdentity conv.PlmnIdentityTbcd
	NodeId       NodeId
}

func (gnid *GlobalNodeId) String() string {
	return gnid.PlmnIdentity.String() + string(":") + gnid.NodeId.String()
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type InterfaceId struct {
	GlobalEnbId GlobalNodeId
	GlobalGnbId GlobalNodeId
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

const (
	E2AP_InterfaceDirectionIncoming uint32 = 0
	E2AP_InterfaceDirectionOutgoing uint32 = 1
)

type EventTriggerDefinition struct {
	Data OctetString
}

/*
//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type CallProcessId struct {
  CallProcessIDVal uint32
}
*/

type ActionDefinitionChoice struct {
	Data OctetString
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type BitString struct {
	UnusedBits uint8
	Length     uint64
	Data       []uint8
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type OctetString struct {
	Length uint64
	Data   []uint8
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
const (
	E2AP_SubSeqActionTypeContinue uint64 = 0
	E2AP_SubSeqActionTypeWait     uint64 = 1
)

const (
	E2AP_TimeToWaitZero   uint64 = 0
	E2AP_TimeToWaitW1ms   uint64 = 1
	E2AP_TimeToWaitW2ms   uint64 = 2
	E2AP_TimeToWaitW5ms   uint64 = 3
	E2AP_TimeToWaitW10ms  uint64 = 4
	E2AP_TimeToWaitW20ms  uint64 = 4
	E2AP_TimeToWaitW30ms  uint64 = 5
	E2AP_TimeToWaitW40ms  uint64 = 6
	E2AP_TimeToWaitW50ms  uint64 = 7
	E2AP_TimeToWaitW100ms uint64 = 8
	E2AP_TimeToWaitW200ms uint64 = 9
	E2AP_TimeToWaitW500ms uint64 = 10
	E2AP_TimeToWaitW1s    uint64 = 11
	E2AP_TimeToWaitW2s    uint64 = 12
	E2AP_TimeToWaitW5s    uint64 = 13
	E2AP_TimeToWaitW10s   uint64 = 14
	E2AP_TimeToWaitW20s   uint64 = 15
	E2AP_TimeToWaitW60    uint64 = 16
)

type SubsequentAction struct {
	Present    bool
	Type       uint64
	TimetoWait uint64
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

const (
	E2AP_ActionTypeReport  uint64 = 0
	E2AP_ActionTypeInsert  uint64 = 1
	E2AP_ActionTypePolicy  uint64 = 2
	E2AP_ActionTypeInvalid uint64 = 99 // For RIC internal usage only
)

type ActionToBeSetupItem struct {
	ActionId                   uint64
	ActionType                 uint64
	RicActionDefinitionPresent bool
	ActionDefinitionChoice
	SubsequentAction
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

const (
	E2AP_CauseContent_RadioNetwork uint8 = 1
	E2AP_CauseContent_Transport    uint8 = 2
	E2AP_CauseContent_Protocol     uint8 = 3
	E2AP_CauseContent_Misc         uint8 = 4
	E2AP_CauseContent_Ric          uint8 = 5
)

//const (
//	E2AP_CauseValue_RadioNetwork_ uint8 = 0
//)
//const (
//	E2AP_CauseValue_Transport_ uint8 = 0
//)
//const (
//	E2AP_CauseValue_Protocol_ uint8 = 0
//)
//const (
//	E2AP_CauseValue_Misc_ uint8 = 0
//)

const (
	E2AP_CauseValue_Ric_function_id_Invalid                            uint8 = 0
	E2AP_CauseValue_Ric_action_not_supported                           uint8 = 1
	E2AP_CauseValue_Ric_excessive_actions                              uint8 = 2
	E2AP_CauseValue_Ric_duplicate_action                               uint8 = 3
	E2AP_CauseValue_Ric_duplicate_event                                uint8 = 4
	E2AP_CauseValue_Ric_function_resource_limit                        uint8 = 5
	E2AP_CauseValue_Ric_request_id_unknown                             uint8 = 6
	E2AP_CauseValue_Ric_inconsistent_action_subsequent_action_sequence uint8 = 7
	E2AP_CauseValue_Ric_control_message_invalid                        uint8 = 8
	E2AP_CauseValue_Ric_call_process_id_invalid                        uint8 = 9
	E2AP_CauseValue_Ric_function_not_required                          uint8 = 10
	E2AP_CauseValue_Ric_excessive_functions                            uint8 = 11
	E2AP_CauseValue_Ric_ric_resource_limi                              uint8 = 12
)

type Cause struct {
	Content uint8
	Value   uint8
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type ActionAdmittedItem struct {
	ActionId uint64
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type ActionAdmittedList struct {
	Items []ActionAdmittedItem //16
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type ActionNotAdmittedItem struct {
	ActionId uint64
	Cause    Cause
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type ActionNotAdmittedList struct {
	Items []ActionNotAdmittedItem //16
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
const (
	E2AP_CriticalityReject uint8 = 0
	E2AP_CriticalityIgnore uint8 = 1
	E2AP_CriticalityNotify uint8 = 2
)

type CriticalityDiagnosticsIEListItem struct {
	IeCriticality uint8 //Crit
	IeID          uint32
	TypeOfError   uint8
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type CriticalityDiagnosticsIEList struct {
	Items []CriticalityDiagnosticsIEListItem //256
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type CriticalityDiagnostics struct {
	Present         bool
	ProcCodePresent bool
	ProcCode        uint64
	TrigMsgPresent  bool
	TrigMsg         uint64
	ProcCritPresent bool
	ProcCrit        uint8 //Crit
	CriticalityDiagnosticsIEList
}
