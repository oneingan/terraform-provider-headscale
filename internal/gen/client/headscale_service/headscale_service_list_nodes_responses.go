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

// HeadscaleServiceListNodesReader is a Reader for the HeadscaleServiceListNodes structure.
type HeadscaleServiceListNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceListNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceListNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceListNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceListNodesOK creates a HeadscaleServiceListNodesOK with default headers values
func NewHeadscaleServiceListNodesOK() *HeadscaleServiceListNodesOK {
	return &HeadscaleServiceListNodesOK{}
}

/*
HeadscaleServiceListNodesOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceListNodesOK struct {
	Payload *models.V1ListNodesResponse
}

// IsSuccess returns true when this headscale service list nodes o k response has a 2xx status code
func (o *HeadscaleServiceListNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service list nodes o k response has a 3xx status code
func (o *HeadscaleServiceListNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service list nodes o k response has a 4xx status code
func (o *HeadscaleServiceListNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service list nodes o k response has a 5xx status code
func (o *HeadscaleServiceListNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service list nodes o k response a status code equal to that given
func (o *HeadscaleServiceListNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service list nodes o k response
func (o *HeadscaleServiceListNodesOK) Code() int {
	return 200
}

func (o *HeadscaleServiceListNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/node][%d] headscaleServiceListNodesOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListNodesOK) String() string {
	return fmt.Sprintf("[GET /api/v1/node][%d] headscaleServiceListNodesOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListNodesOK) GetPayload() *models.V1ListNodesResponse {
	return o.Payload
}

func (o *HeadscaleServiceListNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListNodesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceListNodesDefault creates a HeadscaleServiceListNodesDefault with default headers values
func NewHeadscaleServiceListNodesDefault(code int) *HeadscaleServiceListNodesDefault {
	return &HeadscaleServiceListNodesDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceListNodesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceListNodesDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service list nodes default response has a 2xx status code
func (o *HeadscaleServiceListNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service list nodes default response has a 3xx status code
func (o *HeadscaleServiceListNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service list nodes default response has a 4xx status code
func (o *HeadscaleServiceListNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service list nodes default response has a 5xx status code
func (o *HeadscaleServiceListNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service list nodes default response a status code equal to that given
func (o *HeadscaleServiceListNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service list nodes default response
func (o *HeadscaleServiceListNodesDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceListNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/node][%d] HeadscaleService_ListNodes default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListNodesDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/node][%d] HeadscaleService_ListNodes default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListNodesDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceListNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
