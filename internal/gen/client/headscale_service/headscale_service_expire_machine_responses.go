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

// HeadscaleServiceExpireMachineReader is a Reader for the HeadscaleServiceExpireMachine structure.
type HeadscaleServiceExpireMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceExpireMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceExpireMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceExpireMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceExpireMachineOK creates a HeadscaleServiceExpireMachineOK with default headers values
func NewHeadscaleServiceExpireMachineOK() *HeadscaleServiceExpireMachineOK {
	return &HeadscaleServiceExpireMachineOK{}
}

/*
HeadscaleServiceExpireMachineOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceExpireMachineOK struct {
	Payload *models.V1ExpireMachineResponse
}

// IsSuccess returns true when this headscale service expire machine o k response has a 2xx status code
func (o *HeadscaleServiceExpireMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service expire machine o k response has a 3xx status code
func (o *HeadscaleServiceExpireMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service expire machine o k response has a 4xx status code
func (o *HeadscaleServiceExpireMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service expire machine o k response has a 5xx status code
func (o *HeadscaleServiceExpireMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service expire machine o k response a status code equal to that given
func (o *HeadscaleServiceExpireMachineOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceExpireMachineOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/machine/{machineId}/expire][%d] headscaleServiceExpireMachineOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpireMachineOK) String() string {
	return fmt.Sprintf("[POST /api/v1/machine/{machineId}/expire][%d] headscaleServiceExpireMachineOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpireMachineOK) GetPayload() *models.V1ExpireMachineResponse {
	return o.Payload
}

func (o *HeadscaleServiceExpireMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ExpireMachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceExpireMachineDefault creates a HeadscaleServiceExpireMachineDefault with default headers values
func NewHeadscaleServiceExpireMachineDefault(code int) *HeadscaleServiceExpireMachineDefault {
	return &HeadscaleServiceExpireMachineDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceExpireMachineDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceExpireMachineDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service expire machine default response
func (o *HeadscaleServiceExpireMachineDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service expire machine default response has a 2xx status code
func (o *HeadscaleServiceExpireMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service expire machine default response has a 3xx status code
func (o *HeadscaleServiceExpireMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service expire machine default response has a 4xx status code
func (o *HeadscaleServiceExpireMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service expire machine default response has a 5xx status code
func (o *HeadscaleServiceExpireMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service expire machine default response a status code equal to that given
func (o *HeadscaleServiceExpireMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceExpireMachineDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/machine/{machineId}/expire][%d] HeadscaleService_ExpireMachine default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpireMachineDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/machine/{machineId}/expire][%d] HeadscaleService_ExpireMachine default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpireMachineDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceExpireMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
