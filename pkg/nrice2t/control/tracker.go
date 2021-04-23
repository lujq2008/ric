

package control

import (
	"fmt"
	"nRIC/internal/msgx"
	"nRIC/internal/xapp"
	"sync"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type Tracker struct {
	mutex                sync.Mutex
	transactionXappTable map[TransactionXappKey]*TransactionXapp
	transSeq             uint64
}

func (t *Tracker) Init() {
	t.transactionXappTable = make(map[TransactionXappKey]*TransactionXapp)
}

func (t *Tracker) initTransaction(transBase *Transaction) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	transBase.EventChan = make(chan interface{})
	transBase.tracker = t
	transBase.Seq = t.transSeq
	t.transSeq++
}

func (t *Tracker) NewSubsTransaction(subs *Subscription) *TransactionSubs {
	trans := &TransactionSubs{}
	trans.Meid = subs.GetMeid()
	t.initTransaction(&trans.Transaction)
	xapp.Logger.Debug("CREATE %s", trans.String())
	return trans
}

func (t *Tracker) NewXappTransaction(
	endpoint *msgx.MsgSender,
	xid string,
	subid uint32,
	meid *xapp.MsgMeid) *TransactionXapp {

	trans := &TransactionXapp{}
	trans.XappKey = &TransactionXappKey{*endpoint, xid}
	trans.Meid = meid
	trans.SubId = subid
	t.initTransaction(&trans.Transaction)
	xapp.Logger.Debug("CREATE %s", trans.String())
	return trans
}

func (t *Tracker) Track(trans *TransactionXapp) error {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	theKey := *trans.XappKey

	if othtrans, ok := t.transactionXappTable[theKey]; ok {
		err := fmt.Errorf("Tracker: %s is ongoing, not tracking %s", othtrans, trans)
		return err
	}

	trans.tracker = t
	t.transactionXappTable[theKey] = trans
	xapp.Logger.Debug("Tracker: Append %s", trans.String())
	//xapp.Logger.Debug("Tracker: transtable=%v", t.transactionXappTable)
	return nil
}

func (t *Tracker) UnTrackTransaction(xappKey TransactionXappKey) (*TransactionXapp, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if trans, ok2 := t.transactionXappTable[xappKey]; ok2 {
		xapp.Logger.Debug("Tracker: Remove %s", trans.String())
		delete(t.transactionXappTable, xappKey)
		//xapp.Logger.Debug("Tracker: transtable=%v", t.transactionXappTable)
		return trans, nil
	}
	return nil, fmt.Errorf("Tracker: No record %s", xappKey)
}
