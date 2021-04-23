
package models

import (
	"encoding/xml"
	"strings"
)

type Gnb struct {
	Text        string `xml:",chardata"`
	GlobalGNBID struct {
		Text   string `xml:",chardata"`
		PlmnID string `xml:"plmn-id"`
		GnbID  struct {
			Text  string `xml:",chardata"`
			GnbID string `xml:"gnb-ID"`
		} `xml:"gnb-id"`
	} `xml:"global-gNB-ID"`
}

type EnGnb struct {
	Text        string `xml:",chardata"`
	GlobalGNBID struct {
		Text   string `xml:",chardata"`
		PlmnID string `xml:"pLMN-Identity"`
		GnbID  struct {
			Text  string `xml:",chardata"`
			GnbID string `xml:"gNB-ID"`
		} `xml:"gNB-ID"`
	} `xml:"global-gNB-ID"`
}

type NgEnbId struct {
	Text            string `xml:",chardata"`
	EnbIdMacro      string `xml:"enb-ID-macro"`
	EnbIdShortMacro string `xml:"enb-ID-shortmacro"`
	EnbIdLongMacro  string `xml:"enb-ID-longmacro"`
}

type NgEnb struct {
	Text          string `xml:",chardata"`
	GlobalNgENBID struct {
		Text   string  `xml:",chardata"`
		PlmnID string  `xml:"plmn-id"`
		EnbID  NgEnbId `xml:"enb-id"`
	} `xml:"global-ng-eNB-ID"`
}

type EnbId struct {
	Text            string `xml:",chardata"`
	MacroEnbId      string `xml:"macro-eNB-ID"`
	HomeEnbId       string `xml:"home-eNB-ID"`
	ShortMacroEnbId string `xml:"short-Macro-eNB-ID"`
	LongMacroEnbId  string `xml:"long-Macro-eNB-ID"`
}

type Enb struct {
	Text        string `xml:",chardata"`
	GlobalENBID struct {
		Text   string `xml:",chardata"`
		PlmnID string `xml:"pLMN-Identity"`
		EnbID  EnbId  `xml:"eNB-ID"`
	} `xml:"global-eNB-ID"`
}

type GlobalE2NodeId struct {
	Text  string `xml:",chardata"`
	GNB   Gnb    `xml:"gNB"`
	EnGNB EnGnb  `xml:"en-gNB"`
	NgENB NgEnb  `xml:"ng-eNB"`
	ENB   Enb    `xml:"eNB"`
}

type E2SetupRequest struct {
	Text        string `xml:",chardata"`
	ProtocolIEs struct {
		Text              string `xml:",chardata"`
		E2setupRequestIEs []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Criticality struct {
				Text   string `xml:",chardata"`
				Reject string `xml:"reject"`
			} `xml:"criticality"`
			Value struct {
				Text             string           `xml:",chardata"`
				GlobalE2nodeID   GlobalE2NodeId   `xml:"GlobalE2node-ID"`
				RANfunctionsList RANfunctionsList `xml:"RANfunctions-List"`
			} `xml:"value"`
		} `xml:"E2setupRequestIEs"`
	} `xml:"protocolIEs"`
}

type E2SetupRequestMessage struct {
	XMLName xml.Name `xml:"E2SetupRequestMessage"`
	Text    string   `xml:",chardata"`
	E2APPDU struct {
		Text              string `xml:",chardata"`
		InitiatingMessage struct {
			Text          string `xml:",chardata"`
			ProcedureCode string `xml:"procedureCode"`
			Criticality   struct {
				Text   string `xml:",chardata"`
				Reject string `xml:"reject"`
			} `xml:"criticality"`
			Value struct {
				Text           string         `xml:",chardata"`
				E2setupRequest E2SetupRequest `xml:"E2setupRequest"`
			} `xml:"value"`
		} `xml:"initiatingMessage"`
	} `xml:"E2AP-PDU"`
}

type RanFunctionItem struct {
	Text                  string `xml:",chardata"`
	RanFunctionID         uint32 `xml:"ranFunctionID"`
	RanFunctionDefinition string `xml:"ranFunctionDefinition"`
	RanFunctionRevision   uint32 `xml:"ranFunctionRevision"`
}

type RANfunctionsList struct {
	Text                      string `xml:",chardata"`
	ProtocolIESingleContainer []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id"`
		Criticality struct {
			Text   string `xml:",chardata"`
			Reject string `xml:"reject"`
		} `xml:"criticality"`
		Value struct {
			Text            string          `xml:",chardata"`
			RANfunctionItem RanFunctionItem `xml:"RANfunction-Item"`
		} `xml:"value"`
	} `xml:"ProtocolIE-SingleContainer"`
}

func (m *E2SetupRequestMessage) getGlobalE2NodeId() GlobalE2NodeId {
	return m.E2APPDU.InitiatingMessage.Value.E2setupRequest.ProtocolIEs.E2setupRequestIEs[0].Value.GlobalE2nodeID
}

func (m *E2SetupRequestMessage) GetPlmnId() string {
	globalE2NodeId := m.getGlobalE2NodeId()
	if id := globalE2NodeId.GNB.GlobalGNBID.PlmnID; id != "" {
		return m.trimSpaces(id)
	}
	if id := globalE2NodeId.EnGNB.GlobalGNBID.PlmnID; id != "" {
		return m.trimSpaces(id)
	}
	if id := globalE2NodeId.ENB.GlobalENBID.PlmnID; id != "" {
		return m.trimSpaces(id)
	}
	if id := globalE2NodeId.NgENB.GlobalNgENBID.PlmnID; id != "" {
		return m.trimSpaces(id)
	}
	return ""
}

func (m *E2SetupRequestMessage) getInnerEnbId(enbId EnbId) string {

	if id := enbId.HomeEnbId; id != "" {
		return id
	}

	if id := enbId.LongMacroEnbId; id != "" {
		return id
	}

	if id := enbId.MacroEnbId; id != "" {
		return id
	}

	if id := enbId.ShortMacroEnbId; id != "" {
		return id
	}

	return ""
}

func (m *E2SetupRequestMessage) getInnerNgEnbId(enbId NgEnbId) string {
	if id := enbId.EnbIdLongMacro; id != "" {
		return id
	}

	if id := enbId.EnbIdMacro; id != "" {
		return id
	}

	if id := enbId.EnbIdShortMacro; id != "" {
		return id
	}

	return ""
}

func (m *E2SetupRequestMessage) GetNbId() string {
	globalE2NodeId := m.getGlobalE2NodeId()

	if id := globalE2NodeId.GNB.GlobalGNBID.GnbID.GnbID; id != "" {
		return m.trimSpaces(id)
	}

	if id := globalE2NodeId.EnGNB.GlobalGNBID.GnbID.GnbID; id != "" {
		return m.trimSpaces(id)
	}

	if id := m.getInnerEnbId(globalE2NodeId.ENB.GlobalENBID.EnbID); id != "" {
		return m.trimSpaces(id)
	}

	if id := m.getInnerNgEnbId(globalE2NodeId.NgENB.GlobalNgENBID.EnbID); id != "" {
		return m.trimSpaces(id)
	}

	return ""
}

func (m *E2SetupRequestMessage) trimSpaces(str string) string {
	return strings.NewReplacer(" ", "", "\n", "").Replace(str)
}
