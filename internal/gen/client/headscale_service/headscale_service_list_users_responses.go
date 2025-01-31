// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/awlsring/terraform-provider-headscale/internal/gen/models"
)

// HeadscaleServiceListUsersReader is a Reader for the HeadscaleServiceListUsers structure.
type HeadscaleServiceListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceListUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceListUsersOK creates a HeadscaleServiceListUsersOK with default headers values
func NewHeadscaleServiceListUsersOK() *HeadscaleServiceListUsersOK {
	return &HeadscaleServiceListUsersOK{}
}

/*
HeadscaleServiceListUsersOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceListUsersOK struct {
	Payload *models.V1ListUsersResponse
}

// IsSuccess returns true when this headscale service list users o k response has a 2xx status code
func (o *HeadscaleServiceListUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service list users o k response has a 3xx status code
func (o *HeadscaleServiceListUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service list users o k response has a 4xx status code
func (o *HeadscaleServiceListUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service list users o k response has a 5xx status code
func (o *HeadscaleServiceListUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service list users o k response a status code equal to that given
func (o *HeadscaleServiceListUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service list users o k response
func (o *HeadscaleServiceListUsersOK) Code() int {
	return 200
}

func (o *HeadscaleServiceListUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user][%d] headscaleServiceListUsersOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListUsersOK) String() string {
	return fmt.Sprintf("[GET /api/v1/user][%d] headscaleServiceListUsersOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListUsersOK) GetPayload() *models.V1ListUsersResponse {
	return o.Payload
}

func (o *HeadscaleServiceListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListUsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceListUsersDefault creates a HeadscaleServiceListUsersDefault with default headers values
func NewHeadscaleServiceListUsersDefault(code int) *HeadscaleServiceListUsersDefault {
	return &HeadscaleServiceListUsersDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceListUsersDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceListUsersDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service list users default response has a 2xx status code
func (o *HeadscaleServiceListUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service list users default response has a 3xx status code
func (o *HeadscaleServiceListUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service list users default response has a 4xx status code
func (o *HeadscaleServiceListUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service list users default response has a 5xx status code
func (o *HeadscaleServiceListUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service list users default response a status code equal to that given
func (o *HeadscaleServiceListUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service list users default response
func (o *HeadscaleServiceListUsersDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceListUsersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/user][%d] HeadscaleService_ListUsers default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListUsersDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/user][%d] HeadscaleService_ListUsers default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListUsersDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceListUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
