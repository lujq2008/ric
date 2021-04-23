

package e2ap

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APSubscriptionRequest struct {
	RequestId
	FunctionId
	EventTriggerDefinition
	ActionSetups []ActionToBeSetupItem
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APSubscriptionResponse struct {
	RequestId
	FunctionId
	ActionAdmittedList
	ActionNotAdmittedList
}

//-----------------------------------------------------------------------------
//
//-----------------------------------------------------------------------------
type E2APSubscriptionFailure struct {
	RequestId
	FunctionId
	ActionNotAdmittedList
	CriticalityDiagnostics
}
