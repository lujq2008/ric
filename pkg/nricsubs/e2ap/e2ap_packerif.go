

package e2ap

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionRequestIf interface {
	Pack(*E2APSubscriptionRequest) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionRequest)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionResponseIf interface {
	Pack(*E2APSubscriptionResponse) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionResponse)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionFailureIf interface {
	Pack(*E2APSubscriptionFailure) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionFailure)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionDeleteRequestIf interface {
	Pack(*E2APSubscriptionDeleteRequest) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionDeleteRequest)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionDeleteResponseIf interface {
	Pack(*E2APSubscriptionDeleteResponse) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionDeleteResponse)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APMsgPackerSubscriptionDeleteFailureIf interface {
	Pack(*E2APSubscriptionDeleteFailure) (error, *PackedData)
	UnPack(msg *PackedData) (error, *E2APSubscriptionDeleteFailure)
	String() string
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APPackerIf interface {
	NewPackerSubscriptionRequest() E2APMsgPackerSubscriptionRequestIf
	NewPackerSubscriptionResponse() E2APMsgPackerSubscriptionResponseIf
	NewPackerSubscriptionFailure() E2APMsgPackerSubscriptionFailureIf
	NewPackerSubscriptionDeleteRequest() E2APMsgPackerSubscriptionDeleteRequestIf
	NewPackerSubscriptionDeleteResponse() E2APMsgPackerSubscriptionDeleteResponseIf
	NewPackerSubscriptionDeleteFailure() E2APMsgPackerSubscriptionDeleteFailureIf
	//UnPack(*PackedData) (error, interface{})
	//Pack(interface{}, *PackedData) (error, *PackedData)
}
