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

// HeadscaleServiceRegisterNodeReader is a Reader for the HeadscaleServiceRegisterNode structure.
type HeadscaleServiceRegisterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceRegisterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceRegisterNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceRegisterNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceRegisterNodeOK creates a HeadscaleServiceRegisterNodeOK with default headers values
func NewHeadscaleServiceRegisterNodeOK() *HeadscaleServiceRegisterNodeOK {
	return &HeadscaleServiceRegisterNodeOK{}
}

/*
HeadscaleServiceRegisterNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceRegisterNodeOK struct {
	Payload *models.V1RegisterNodeResponse
}

// IsSuccess returns true when this headscale service register node o k response has a 2xx status code
func (o *HeadscaleServiceRegisterNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service register node o k response has a 3xx status code
func (o *HeadscaleServiceRegisterNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service register node o k response has a 4xx status code
func (o *HeadscaleServiceRegisterNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service register node o k response has a 5xx status code
func (o *HeadscaleServiceRegisterNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service register node o k response a status code equal to that given
func (o *HeadscaleServiceRegisterNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service register node o k response
func (o *HeadscaleServiceRegisterNodeOK) Code() int {
	return 200
}

func (o *HeadscaleServiceRegisterNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/node/register][%d] headscaleServiceRegisterNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceRegisterNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/node/register][%d] headscaleServiceRegisterNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceRegisterNodeOK) GetPayload() *models.V1RegisterNodeResponse {
	return o.Payload
}

func (o *HeadscaleServiceRegisterNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RegisterNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceRegisterNodeDefault creates a HeadscaleServiceRegisterNodeDefault with default headers values
func NewHeadscaleServiceRegisterNodeDefault(code int) *HeadscaleServiceRegisterNodeDefault {
	return &HeadscaleServiceRegisterNodeDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceRegisterNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceRegisterNodeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service register node default response has a 2xx status code
func (o *HeadscaleServiceRegisterNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service register node default response has a 3xx status code
func (o *HeadscaleServiceRegisterNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service register node default response has a 4xx status code
func (o *HeadscaleServiceRegisterNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service register node default response has a 5xx status code
func (o *HeadscaleServiceRegisterNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service register node default response a status code equal to that given
func (o *HeadscaleServiceRegisterNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service register node default response
func (o *HeadscaleServiceRegisterNodeDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceRegisterNodeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/node/register][%d] HeadscaleService_RegisterNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceRegisterNodeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/node/register][%d] HeadscaleService_RegisterNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceRegisterNodeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceRegisterNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
