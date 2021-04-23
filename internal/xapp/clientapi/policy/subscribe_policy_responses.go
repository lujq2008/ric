// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	clientmodel "nRIC/internal/xapp/clientmodel"
)

// SubscribePolicyReader is a Reader for the SubscribePolicy structure.
type SubscribePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscribePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSubscribePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSubscribePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSubscribePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscribePolicyCreated creates a SubscribePolicyCreated with default headers values
func NewSubscribePolicyCreated() *SubscribePolicyCreated {
	return &SubscribePolicyCreated{}
}

/*SubscribePolicyCreated handles this case with default header values.

Subscription successfully created
*/
type SubscribePolicyCreated struct {
	Payload *clientmodel.SubscriptionResponse
}

func (o *SubscribePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /subscriptions/policy][%d] subscribePolicyCreated  %+v", 201, o.Payload)
}

func (o *SubscribePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(clientmodel.SubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubscribePolicyBadRequest creates a SubscribePolicyBadRequest with default headers values
func NewSubscribePolicyBadRequest() *SubscribePolicyBadRequest {
	return &SubscribePolicyBadRequest{}
}

/*SubscribePolicyBadRequest handles this case with default header values.

Invalid input
*/
type SubscribePolicyBadRequest struct {
}

func (o *SubscribePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /subscriptions/policy][%d] subscribePolicyBadRequest ", 400)
}

func (o *SubscribePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubscribePolicyInternalServerError creates a SubscribePolicyInternalServerError with default headers values
func NewSubscribePolicyInternalServerError() *SubscribePolicyInternalServerError {
	return &SubscribePolicyInternalServerError{}
}

/*SubscribePolicyInternalServerError handles this case with default header values.

Internal error
*/
type SubscribePolicyInternalServerError struct {
}

func (o *SubscribePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /subscriptions/policy][%d] subscribePolicyInternalServerError ", 500)
}

func (o *SubscribePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
