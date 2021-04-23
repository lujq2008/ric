
package xapp

import (
	"bytes"
	"crypto/md5"
	"fmt"
)


//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type MsgParams struct {
	Mtype      int
	Payload    []byte
	PayloadLen int
	Meid       *MsgMeid
	Xid        string
	SubId      int
	Src        string
	//Mbuf       *C.rmr_mbuf_t
	Whid       int
	Callid     int
	Timeout    int
	status     int
	Route      *Route
}

func (params *MsgParams) String() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "params(Src=%s Mtype=%d SubId=%d Xid=%s Meid=%s Paylens=%d/%d Paymd5=%x)", params.Src, params.Mtype, params.SubId, params.Xid, params.Meid, params.PayloadLen, len(params.Payload), md5.Sum(params.Payload))
	return b.String()
}

type Route struct {
	SubIdXapp	int64
	SubIdRan	int64
	Topic		string
}

type MsgMeid struct {
	PlmnID  string
	EnbID   string
	RanName string
}

func (meid *MsgMeid) String() string {
	str := "meid("
	pad := ""
	if len(meid.PlmnID) > 0 {
		str += pad + "PlmnID=" + meid.PlmnID
		pad = " "
	}
	if len(meid.EnbID) > 0 {
		str += pad + "EnbID=" + meid.EnbID
		pad = " "
	}
	if len(meid.RanName) > 0 {
		str += pad + "RanName=" + meid.RanName
		pad = " "
	}
	str += ")"
	return str
}

type MessageConsumerFunc func(*MsgParams) error

func (fn MessageConsumerFunc) Consume(params *MsgParams) error {
	return fn(params)
}

type MessageConsumer interface {
	Consume(params *MsgParams) error
}

type PortData struct {
	Name              string
	Port              int
	MaxSize           int
	ThreadType        int
	LowLatency        bool
	FastAck           bool
	Policies          []int
	MaxRetryOnFailure int
}

// @todo: read these from config or somewhere else
const (
	SERVICE_HTTP    = "SERVICE_%s_%s_HTTP_PORT"
)


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
