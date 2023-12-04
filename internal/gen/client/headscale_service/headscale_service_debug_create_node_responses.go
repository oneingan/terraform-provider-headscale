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

// HeadscaleServiceDebugCreateNodeReader is a Reader for the HeadscaleServiceDebugCreateNode structure.
type HeadscaleServiceDebugCreateNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceDebugCreateNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceDebugCreateNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceDebugCreateNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceDebugCreateNodeOK creates a HeadscaleServiceDebugCreateNodeOK with default headers values
func NewHeadscaleServiceDebugCreateNodeOK() *HeadscaleServiceDebugCreateNodeOK {
	return &HeadscaleServiceDebugCreateNodeOK{}
}

/*
HeadscaleServiceDebugCreateNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceDebugCreateNodeOK struct {
	Payload *models.V1DebugCreateNodeResponse
}

// IsSuccess returns true when this headscale service debug create node o k response has a 2xx status code
func (o *HeadscaleServiceDebugCreateNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service debug create node o k response has a 3xx status code
func (o *HeadscaleServiceDebugCreateNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service debug create node o k response has a 4xx status code
func (o *HeadscaleServiceDebugCreateNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service debug create node o k response has a 5xx status code
func (o *HeadscaleServiceDebugCreateNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service debug create node o k response a status code equal to that given
func (o *HeadscaleServiceDebugCreateNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service debug create node o k response
func (o *HeadscaleServiceDebugCreateNodeOK) Code() int {
	return 200
}

func (o *HeadscaleServiceDebugCreateNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/debug/node][%d] headscaleServiceDebugCreateNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDebugCreateNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/debug/node][%d] headscaleServiceDebugCreateNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceDebugCreateNodeOK) GetPayload() *models.V1DebugCreateNodeResponse {
	return o.Payload
}

func (o *HeadscaleServiceDebugCreateNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1DebugCreateNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceDebugCreateNodeDefault creates a HeadscaleServiceDebugCreateNodeDefault with default headers values
func NewHeadscaleServiceDebugCreateNodeDefault(code int) *HeadscaleServiceDebugCreateNodeDefault {
	return &HeadscaleServiceDebugCreateNodeDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceDebugCreateNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceDebugCreateNodeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service debug create node default response has a 2xx status code
func (o *HeadscaleServiceDebugCreateNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service debug create node default response has a 3xx status code
func (o *HeadscaleServiceDebugCreateNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service debug create node default response has a 4xx status code
func (o *HeadscaleServiceDebugCreateNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service debug create node default response has a 5xx status code
func (o *HeadscaleServiceDebugCreateNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service debug create node default response a status code equal to that given
func (o *HeadscaleServiceDebugCreateNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service debug create node default response
func (o *HeadscaleServiceDebugCreateNodeDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceDebugCreateNodeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/debug/node][%d] HeadscaleService_DebugCreateNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDebugCreateNodeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/debug/node][%d] HeadscaleService_DebugCreateNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceDebugCreateNodeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceDebugCreateNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
