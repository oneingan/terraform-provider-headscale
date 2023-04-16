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

// HeadscaleServiceEnableRouteReader is a Reader for the HeadscaleServiceEnableRoute structure.
type HeadscaleServiceEnableRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceEnableRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceEnableRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceEnableRouteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceEnableRouteOK creates a HeadscaleServiceEnableRouteOK with default headers values
func NewHeadscaleServiceEnableRouteOK() *HeadscaleServiceEnableRouteOK {
	return &HeadscaleServiceEnableRouteOK{}
}

/*
HeadscaleServiceEnableRouteOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceEnableRouteOK struct {
	Payload models.V1EnableRouteResponse
}

// IsSuccess returns true when this headscale service enable route o k response has a 2xx status code
func (o *HeadscaleServiceEnableRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service enable route o k response has a 3xx status code
func (o *HeadscaleServiceEnableRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service enable route o k response has a 4xx status code
func (o *HeadscaleServiceEnableRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service enable route o k response has a 5xx status code
func (o *HeadscaleServiceEnableRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service enable route o k response a status code equal to that given
func (o *HeadscaleServiceEnableRouteOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceEnableRouteOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/enable][%d] headscaleServiceEnableRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceEnableRouteOK) String() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/enable][%d] headscaleServiceEnableRouteOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceEnableRouteOK) GetPayload() models.V1EnableRouteResponse {
	return o.Payload
}

func (o *HeadscaleServiceEnableRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceEnableRouteDefault creates a HeadscaleServiceEnableRouteDefault with default headers values
func NewHeadscaleServiceEnableRouteDefault(code int) *HeadscaleServiceEnableRouteDefault {
	return &HeadscaleServiceEnableRouteDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceEnableRouteDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceEnableRouteDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service enable route default response
func (o *HeadscaleServiceEnableRouteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service enable route default response has a 2xx status code
func (o *HeadscaleServiceEnableRouteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service enable route default response has a 3xx status code
func (o *HeadscaleServiceEnableRouteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service enable route default response has a 4xx status code
func (o *HeadscaleServiceEnableRouteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service enable route default response has a 5xx status code
func (o *HeadscaleServiceEnableRouteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service enable route default response a status code equal to that given
func (o *HeadscaleServiceEnableRouteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceEnableRouteDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/enable][%d] HeadscaleService_EnableRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceEnableRouteDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/routes/{routeId}/enable][%d] HeadscaleService_EnableRoute default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceEnableRouteDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceEnableRouteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
