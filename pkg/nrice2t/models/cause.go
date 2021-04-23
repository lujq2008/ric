

package models

import "encoding/xml"

type Cause struct {
	XMLName    xml.Name         `xml:"Cause"`
	Text       string           `xml:",chardata"`
	RicRequest *CauseRic        `xml:"ricRequest"`
	RicService *CauseRicService `xml:"ricService"`
	Transport  *CauseTransport  `xml:"transport"`
	Protocol   *CauseProtocol   `xml:"protocol"`
	Misc       *CauseMisc       `xml:"misc"`
}

type CauseTransport struct {
	Text                         string    `xml:",chardata"`
	TransportResourceUnavailable *struct{} `xml:"transport-resource-unavailable"`
	Unspecified                  *struct{} `xml:"unspecified"`
}

type CauseMisc struct {
	Text                      string    `xml:",chardata"`
	ControlProcessingOverload *struct{} `xml:"control-processing-overload"`
	HardwareFailure           *struct{} `xml:"hardware-failure"`
	OmIntervention            *struct{} `xml:"om-intervention"`
	Unspecified               *struct{} `xml:"unspecified"`
}

type CauseProtocol struct {
	Text                                         string    `xml:",chardata"`
	TransferSyntaxError                          *struct{} `xml:"transfer-syntax-error"`
	AbstractSyntaxErrorReject                    *struct{} `xml:"abstract-syntax-error-reject"`
	AbstractSyntaxErrorIgnoreAndNotify           *struct{} `xml:"abstract-syntax-error-ignore-and-notify"`
	MessageNotCompatibleWithReceiverState        *struct{} `xml:"message-not-compatible-with-receiver-state"`
	SemanticError                                *struct{} `xml:"semantic-error"`
	AbstractSyntaxErrorFalselyConstructedMessage *struct{} `xml:"abstract-syntax-error-falsely-constructed-message"`
	Unspecified                                  *struct{} `xml:"unspecified"`
}

type CauseRicService struct {
	Text                string    `xml:",chardata"`
	FunctionNotRequired *struct{} `xml:"function-not-required"`
	ExcessiveFunctions  *struct{} `xml:"excessive-functions"`
	RicResourceLimit    *struct{} `xml:"ric-resource-limit"`
}

type CauseRic struct {
	Text                                       string    `xml:",chardata"`
	RanFunctionIdInvalid                       *struct{} `xml:"ran-function-id-Invalid"`
	ActionNotSupported                         *struct{} `xml:"action-not-supported"`
	ExcessiveActions                           *struct{} `xml:"excessive-actions"`
	DuplicateAction                            *struct{} `xml:"duplicate-action"`
	DuplicateEvent                             *struct{} `xml:"duplicate-event"`
	FunctionResourceLimit                      *struct{} `xml:"function-resource-limit"`
	RequestIdUnknown                           *struct{} `xml:"request-id-unknown"`
	InconsistentActionSubsequentActionSequence *struct{} `xml:"inconsistent-action-subsequent-action-sequence"`
	ControlMessageInvalid                      *struct{} `xml:"control-message-invalid"`
	CallProcessIdInvalid                       *struct{} `xml:"call-process-id-invalid"`
	Unspecified                                *struct{} `xml:"unspecified"`
}
