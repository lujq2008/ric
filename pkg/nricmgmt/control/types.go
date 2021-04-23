

package control

import (
	"nRIC/pkg/nricsubs/e2ap"
)

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type RequestId struct {
	e2ap.RequestId
}

func (rid *RequestId) String() string {
	return "reqid(" + rid.RequestId.String() + ")"
}
