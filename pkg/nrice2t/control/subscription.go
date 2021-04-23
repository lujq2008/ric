

package control

import (
	"nRIC/internal/msgx"
	"nRIC/pkg/nricsubs/e2ap"
	"nRIC/internal/xapp"
	"sync"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type Subscription struct {
	mutex     sync.Mutex                    // Lock
	valid     bool                          // valid
	registry  *Registry                     // Registry
	ReqId     RequestId                     // ReqId (Requestor Id + Seq Nro a.k.a subsid)
	Meid      *xapp.MsgMeid                 // Meid/ RanName
	EpList	  msgx.EndpointList
	TransLock sync.Mutex                    // Lock transactions, only one executed per time for subs
	TheTrans  TransactionIf                 // Ongoing transaction
	SubReqMsg *e2ap.E2APSubscriptionRequest // Subscription information
	SubRFMsg  interface{}                   // Subscription information
}

func (s *Subscription) String() string {
	meidstr := "N/A"
	if s.Meid != nil {
		meidstr = s.Meid.String()
	}
	return "subs(" + s.ReqId.String() + "/" + meidstr + "/"  + ")"
}

func (s *Subscription) GetCachedResponse() (interface{}, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.SubRFMsg, s.valid
}

func (s *Subscription) SetCachedResponse(subRFMsg interface{}, valid bool) (interface{}, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.SubRFMsg = subRFMsg
	s.valid = valid
	return s.SubRFMsg, s.valid
}

func (s *Subscription) GetReqId() *RequestId {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return &s.ReqId
}

func (s *Subscription) GetMeid() *xapp.MsgMeid {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.Meid != nil {
		return s.Meid
	}
	return nil
}

func (s *Subscription) GetTransaction() TransactionIf {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.TheTrans
}

func (s *Subscription) WaitTransactionTurn(trans TransactionIf) {
	s.TransLock.Lock()
	s.mutex.Lock()
	s.TheTrans = trans
	s.mutex.Unlock()
}

func (s *Subscription) ReleaseTransactionTurn(trans TransactionIf) {
	s.mutex.Lock()
	if trans != nil && trans == s.TheTrans {
		s.TheTrans = nil
	}
	s.mutex.Unlock()
	s.TransLock.Unlock()
}

func (s *Subscription) IsMergeable(trans *TransactionXapp, subReqMsg *e2ap.E2APSubscriptionRequest) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.valid == false {
		return false
	}

	if s.SubReqMsg == nil {
		return false
	}

	if s.Meid.RanName != trans.Meid.RanName {
		return false
	}

	// EventTrigger check
	if s.SubReqMsg.EventTriggerDefinition.Data.Length != subReqMsg.EventTriggerDefinition.Data.Length {
		return false
	}
	for i := uint64(0); i < s.SubReqMsg.EventTriggerDefinition.Data.Length; i++ {
		if s.SubReqMsg.EventTriggerDefinition.Data.Data[i] != subReqMsg.EventTriggerDefinition.Data.Data[i] {
			return false
		}
	}

	// Actions check
	if len(s.SubReqMsg.ActionSetups) != len(subReqMsg.ActionSetups) {
		return false
	}

	for _, acts := range s.SubReqMsg.ActionSetups {
		for _, actt := range subReqMsg.ActionSetups {
			if acts.ActionId != actt.ActionId {
				return false
			}
			if acts.ActionType != actt.ActionType {
				return false
			}

			if acts.ActionType != e2ap.E2AP_ActionTypeReport {
				return false
			}

			if acts.RicActionDefinitionPresent != actt.RicActionDefinitionPresent {
				return false
			}

			if acts.ActionDefinitionChoice.Data.Length != actt.ActionDefinitionChoice.Data.Length {
				return false
			}
			for i := uint64(0); i < acts.ActionDefinitionChoice.Data.Length; i++ {
				if acts.ActionDefinitionChoice.Data.Data[i] != actt.ActionDefinitionChoice.Data.Data[i] {
					return false
				}
			}
			//reflect.DeepEqual(acts.ActionDefinitionChoice, actt.ActionDefinitionChoice)

			if acts.SubsequentAction.Present != actt.SubsequentAction.Present ||
				acts.SubsequentAction.Type != actt.SubsequentAction.Type ||
				acts.SubsequentAction.TimetoWait != actt.SubsequentAction.TimetoWait {
				return false
			}
		}
	}

	return true
}
