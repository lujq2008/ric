

package control

import (
	"nRIC/internal/msgx"
	"nRIC/internal/xapp"
	"nRIC/pkg/nricsubs/e2ap"
	"strconv"
	"sync"
	"time"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type TransactionIf interface {
	String() string
	Release()
	SendEvent(interface{}, time.Duration) (bool, bool)
	WaitEvent(time.Duration) (interface{}, bool)
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------

type Transaction struct {
	mutex     sync.Mutex       //
	Seq       uint64           //transaction sequence
	tracker   *Tracker         //tracker instance
	Meid      *xapp.MsgMeid    //meid transaction related
	Mtype     int              //Encoded message type to be send
	Payload   *e2ap.PackedData //Encoded message to be send
	EventChan chan interface{}
	Endpoint msgx.SenderIf   //Topic ;send to xapp
}

func (t *Transaction) String() string {
	meidstr := "N/A"
	if t.Meid != nil {
		meidstr = t.Meid.String()
	}
	return "trans(" + strconv.FormatUint(uint64(t.Seq), 10) + "/" + meidstr + ")"
}

func (t *Transaction) SendEvent(event interface{}, waittime time.Duration) (bool, bool) {
	if waittime > 0 {
		select {
		case t.EventChan <- event:
			return true, false
		case <-time.After(waittime):
			return false, true
		}
		return false, false
	}
	t.EventChan <- event
	return true, false
}

func (t *Transaction) WaitEvent(waittime time.Duration) (interface{}, bool) {
	if waittime > 0 {
		select {
		case event := <-t.EventChan:
			return event, false
		case <-time.After(waittime):
			return nil, true
		}
	}
	event := <-t.EventChan
	return event, false
}

func (t *Transaction) GetMtype() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.Mtype
}

func (t *Transaction) GetMeid() *xapp.MsgMeid {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.Meid != nil {
		return t.Meid
	}
	return nil
}

func (t *Transaction) GetPayload() *e2ap.PackedData {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.Payload
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type TransactionSubs struct {
	Transaction //
}

func (t *TransactionSubs) String() string {
	return "transsubs(" + t.Transaction.String() + ")"
}

func (t *TransactionSubs) Release() {
	t.mutex.Lock()
	xapp.Logger.Debug("RELEASE %s", t.String())
	t.tracker = nil
	t.mutex.Unlock()
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type TransactionXappKey struct {
	Endpoint msgx.SenderIf
	Xid string // xapp xid in req
}


func (key *TransactionXappKey) String() string {
	return "transkey(" + key.Endpoint.String() + "/" + key.Xid + ")"
}
//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type TransactionXapp struct {
	Transaction
	XappKey *TransactionXappKey
	SubId   uint32
}

func (t *TransactionXapp) String() string {
	var transkey string = "transkey(N/A)"
	if t.XappKey != nil {
		transkey = t.XappKey.String()
	}
	return "transxapp(" + t.Transaction.String() + "/" + transkey + "/" + strconv.FormatUint(uint64(t.SubId), 10) + ")"
}

func (t *TransactionXapp) GetEndpoint() msgx.SenderIf {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.XappKey != nil {
		return t.XappKey.Endpoint
	}
	return nil
}



func (t *TransactionXapp) GetXid() string {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.XappKey != nil {
		return t.XappKey.Xid
	}
	return ""
}
func (t *TransactionXapp) GetSrc() string {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.XappKey != nil {
		return t.XappKey.Endpoint.String()
	}
	return ""
}

func (t *TransactionXapp) GetSubId() uint32 {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.SubId
}

func (t *TransactionXapp) Release() {
	t.mutex.Lock()
	xapp.Logger.Debug("RELEASE %s", t.String())
	tracker := t.tracker
	xappkey := t.XappKey
	t.tracker = nil
	t.mutex.Unlock()

	if tracker != nil && xappkey != nil {
		tracker.UnTrackTransaction(*xappkey)
	}
}
