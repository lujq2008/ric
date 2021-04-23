package control

import (
	"fmt"
	"nRIC/internal/xapp"
	"nRIC/internal/xapp/models"
	dbclient "nRIC/pkg/dbagent/grpcserver"
	"nRIC/pkg/nricsubs/e2ap"
	"sync"
	"time"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

type Registry struct {
	mutex       sync.Mutex
	register    map[uint32]*Subscription
	subIds      []uint32
	AcessDbAgent *dbclient.MsgSender
}

func (r *Registry) Initialize() {
	r.register = make(map[uint32]*Subscription)
	var i uint32
	for i = 0; i < 65535; i++ {
		r.subIds = append(r.subIds, i+1)
	}
}

func (r *Registry) QueryHandler() (models.SubscriptionList, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	resp := models.SubscriptionList{}
	return resp, nil
}

func (r *Registry) allocateSubs(trans *TransactionXapp, subReqMsg *e2ap.E2APSubscriptionRequest) (*Subscription, error) {
	if len(r.subIds) > 0 {
		subId := r.subIds[0]
		r.subIds = r.subIds[1:]
		if _, ok := r.register[subId]; ok == true {
			r.subIds = append(r.subIds, subId)
			return nil, fmt.Errorf("Registry: Failed to reserve subscription exists")
		}
		subs := &Subscription{
			registry:  r,
			Meid:      trans.Meid,
			SubReqMsg: subReqMsg,
			valid:     true,
		}
		subs.ReqId.Id = 123
		subs.ReqId.InstanceId = subId

		if subs.EpList.AddEndpoint(trans.GetEndpoint()) == false {
			r.subIds = append(r.subIds, subs.ReqId.InstanceId)
			return nil, fmt.Errorf("Registry: Endpoint existing already in subscription")
		}

		return subs, nil
	}
	return nil, fmt.Errorf("Registry: Failed to reserve subscription no free ids")
}

func (r *Registry) findExistingSubs(trans *TransactionXapp, subReqMsg *e2ap.E2APSubscriptionRequest) (*Subscription, bool) {

	for _, subs := range r.register {
		if subs.IsMergeable(trans, subReqMsg) {

			//
			// check if there has been race conditions
			//
			subs.mutex.Lock()
			//subs has been set to invalid
			if subs.valid == false {
				subs.mutex.Unlock()
				continue
			}
			// If size is zero, entry is to be deleted
			if subs.EpList.Size() == 0 {
				subs.mutex.Unlock()
				continue
			}
			// Try to add to endpointlist. Adding fails if endpoint is already in the list
			if subs.EpList.AddEndpoint(trans.GetEndpoint()) == false {
				subs.mutex.Unlock()
				xapp.Logger.Debug("Registry: Subs with requesting endpoint found. %s for %s", subs.String(), trans.String())
				return subs, true
			}
			subs.mutex.Unlock()

			xapp.Logger.Debug("Registry: Mergeable subs found. %s for %s", subs.String(), trans.String())
			return subs, false
		}
	}
	return nil, false
}


func (r *Registry) CheckActionTypes(subReqMsg *e2ap.E2APSubscriptionRequest) (uint64, error) {
	var reportFound bool = false
	var policyFound bool = false
	var insertFound bool = false

	for _, acts := range subReqMsg.ActionSetups {
		if acts.ActionType == e2ap.E2AP_ActionTypeReport {
			reportFound = true
		}
		if acts.ActionType == e2ap.E2AP_ActionTypePolicy {
			policyFound = true
		}
		if acts.ActionType == e2ap.E2AP_ActionTypeInsert {
			insertFound = true
		}
	}
	if reportFound == true && policyFound == true || reportFound == true && insertFound == true || policyFound == true && insertFound == true {
		return e2ap.E2AP_ActionTypeInvalid, fmt.Errorf("Different action types (Report, Policy or Insert) in same RICactions-ToBeSetup-List")
	}
	if reportFound == true {
		return e2ap.E2AP_ActionTypeReport, nil
	}
	if policyFound == true {
		return e2ap.E2AP_ActionTypePolicy, nil
	}
	if insertFound == true {
		return e2ap.E2AP_ActionTypeInsert, nil
	}
	return e2ap.E2AP_ActionTypeInvalid, fmt.Errorf("Invalid action type in RICactions-ToBeSetup-List")
}

// TODO: Works with concurrent calls, but check if can be improved
func (r *Registry) RemoveFromSubscription(subs *Subscription, trans *TransactionXapp, waitRouteClean time.Duration) error {


	return nil
}

func (r *Registry) GetSubscription(subId uint32) *Subscription {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if _, ok := r.register[subId]; ok {
		return r.register[subId]
	}
	return nil
}

func (r *Registry) GetSubscriptionFirstMatch(subIds []uint32) (*Subscription, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, subId := range subIds {
		if _, ok := r.register[subId]; ok {
			return r.register[subId], nil
		}
	}
	return nil, fmt.Errorf("No valid subscription found with subIds %v", subIds)
}
