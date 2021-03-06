package kafka

import (
	"bufio"
)

type listGroupsRequestV1 struct {
}

func (t listGroupsRequestV1) size() int32 {
	return 0
}

func (t listGroupsRequestV1) writeTo(wb *writeBuffer) {
}

type listGroupsResponseGroupV1 struct {
	// GroupID holds the unique group identifier
	GroupID      string
	ProtocolType string
}

func (t listGroupsResponseGroupV1) size() int32 {
	return sizeofString(t.GroupID) + sizeofString(t.ProtocolType)
}

func (t listGroupsResponseGroupV1) writeTo(wb *writeBuffer) {
	wb.writeString(t.GroupID)
	wb.writeString(t.ProtocolType)
}

func (t *listGroupsResponseGroupV1) readFrom(r *bufio.Reader, size int) (remain int, err error) {
	if remain, err = readString(r, size, &t.GroupID); err != nil {
		return
	}
	if remain, err = readString(r, remain, &t.ProtocolType); err != nil {
		return
	}
	return
}

type listGroupsResponseV1 struct {
	// ThrottleTimeMS holds the duration in milliseconds for which the request
	// was throttled due to quota violation (Zero if the request did not violate
	// any quota)
	ThrottleTimeMS int32

	// ErrorCode holds response error code
	ErrorCode int16
	Groups    []listGroupsResponseGroupV1
}

func (t listGroupsResponseV1) size() int32 {
	return sizeofInt32(t.ThrottleTimeMS) +
		sizeofInt16(t.ErrorCode) +
		sizeofArray(len(t.Groups), func(i int) int32 { return t.Groups[i].size() })
}

func (t listGroupsResponseV1) writeTo(wb *writeBuffer) {
	wb.writeInt32(t.ThrottleTimeMS)
	wb.writeInt16(t.ErrorCode)
	wb.writeArray(len(t.Groups), func(i int) { t.Groups[i].writeTo(wb) })
}

func (t *listGroupsResponseV1) readFrom(r *bufio.Reader, size int) (remain int, err error) {
	if remain, err = readInt32(r, size, &t.ThrottleTimeMS); err != nil {
		return
	}
	if remain, err = readInt16(r, remain, &t.ErrorCode); err != nil {
		return
	}

	fn := func(withReader *bufio.Reader, withSize int) (fnRemain int, fnErr error) {
		var item listGroupsResponseGroupV1
		if fnRemain, fnErr = (&item).readFrom(withReader, withSize); err != nil {
			return
		}
		t.Groups = append(t.Groups, item)
		return
	}
	if remain, err = readArrayWith(r, remain, fn); err != nil {
		return
	}

	return
}
