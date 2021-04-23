// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "nRIC/internal/xapp/models"
)

// SubscribePolicyCreatedCode is the HTTP code returned for type SubscribePolicyCreated
const SubscribePolicyCreatedCode int = 201

/*SubscribePolicyCreated Subscription successfully created

swagger:response subscribePolicyCreated
*/
type SubscribePolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SubscriptionResponse `json:"body,omitempty"`
}

// NewSubscribePolicyCreated creates SubscribePolicyCreated with default headers values
func NewSubscribePolicyCreated() *SubscribePolicyCreated {

	return &SubscribePolicyCreated{}
}

// WithPayload adds the payload to the subscribe policy created response
func (o *SubscribePolicyCreated) WithPayload(payload *models.SubscriptionResponse) *SubscribePolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subscribe policy created response
func (o *SubscribePolicyCreated) SetPayload(payload *models.SubscriptionResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubscribePolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubscribePolicyBadRequestCode is the HTTP code returned for type SubscribePolicyBadRequest
const SubscribePolicyBadRequestCode int = 400

/*SubscribePolicyBadRequest Invalid input

swagger:response subscribePolicyBadRequest
*/
type SubscribePolicyBadRequest struct {
}

// NewSubscribePolicyBadRequest creates SubscribePolicyBadRequest with default headers values
func NewSubscribePolicyBadRequest() *SubscribePolicyBadRequest {

	return &SubscribePolicyBadRequest{}
}

// WriteResponse to the client
func (o *SubscribePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SubscribePolicyInternalServerErrorCode is the HTTP code returned for type SubscribePolicyInternalServerError
const SubscribePolicyInternalServerErrorCode int = 500

/*SubscribePolicyInternalServerError Internal error

swagger:response subscribePolicyInternalServerError
*/
type SubscribePolicyInternalServerError struct {
}

// NewSubscribePolicyInternalServerError creates SubscribePolicyInternalServerError with default headers values
func NewSubscribePolicyInternalServerError() *SubscribePolicyInternalServerError {

	return &SubscribePolicyInternalServerError{}
}

// WriteResponse to the client
func (o *SubscribePolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
