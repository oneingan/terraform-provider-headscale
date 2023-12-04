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

// HeadscaleServiceRenameNodeReader is a Reader for the HeadscaleServiceRenameNode structure.
type HeadscaleServiceRenameNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceRenameNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceRenameNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceRenameNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceRenameNodeOK creates a HeadscaleServiceRenameNodeOK with default headers values
func NewHeadscaleServiceRenameNodeOK() *HeadscaleServiceRenameNodeOK {
	return &HeadscaleServiceRenameNodeOK{}
}

/*
HeadscaleServiceRenameNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceRenameNodeOK struct {
	Payload *models.V1RenameNodeResponse
}

// IsSuccess returns true when this headscale service rename node o k response has a 2xx status code
func (o *HeadscaleServiceRenameNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service rename node o k response has a 3xx status code
func (o *HeadscaleServiceRenameNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service rename node o k response has a 4xx status code
func (o *HeadscaleServiceRenameNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service rename node o k response has a 5xx status code
func (o *HeadscaleServiceRenameNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service rename node o k response a status code equal to that given
func (o *HeadscaleServiceRenameNodeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service rename node o k response
func (o *HeadscaleServiceRenameNodeOK) Code() int {
	return 200
}

func (o *HeadscaleServiceRenameNodeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/rename/{newName}][%d] headscaleServiceRenameNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceRenameNodeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/rename/{newName}][%d] headscaleServiceRenameNodeOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceRenameNodeOK) GetPayload() *models.V1RenameNodeResponse {
	return o.Payload
}

func (o *HeadscaleServiceRenameNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RenameNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceRenameNodeDefault creates a HeadscaleServiceRenameNodeDefault with default headers values
func NewHeadscaleServiceRenameNodeDefault(code int) *HeadscaleServiceRenameNodeDefault {
	return &HeadscaleServiceRenameNodeDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceRenameNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceRenameNodeDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service rename node default response has a 2xx status code
func (o *HeadscaleServiceRenameNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service rename node default response has a 3xx status code
func (o *HeadscaleServiceRenameNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service rename node default response has a 4xx status code
func (o *HeadscaleServiceRenameNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service rename node default response has a 5xx status code
func (o *HeadscaleServiceRenameNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service rename node default response a status code equal to that given
func (o *HeadscaleServiceRenameNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service rename node default response
func (o *HeadscaleServiceRenameNodeDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceRenameNodeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/rename/{newName}][%d] HeadscaleService_RenameNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceRenameNodeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/node/{nodeId}/rename/{newName}][%d] HeadscaleService_RenameNode default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceRenameNodeDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceRenameNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
