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

// HeadscaleServiceExpireAPIKeyReader is a Reader for the HeadscaleServiceExpireAPIKey structure.
type HeadscaleServiceExpireAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceExpireAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceExpireAPIKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceExpireAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceExpireAPIKeyOK creates a HeadscaleServiceExpireAPIKeyOK with default headers values
func NewHeadscaleServiceExpireAPIKeyOK() *HeadscaleServiceExpireAPIKeyOK {
	return &HeadscaleServiceExpireAPIKeyOK{}
}

/*
HeadscaleServiceExpireAPIKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceExpireAPIKeyOK struct {
	Payload models.V1ExpireAPIKeyResponse
}

// IsSuccess returns true when this headscale service expire Api key o k response has a 2xx status code
func (o *HeadscaleServiceExpireAPIKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service expire Api key o k response has a 3xx status code
func (o *HeadscaleServiceExpireAPIKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service expire Api key o k response has a 4xx status code
func (o *HeadscaleServiceExpireAPIKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service expire Api key o k response has a 5xx status code
func (o *HeadscaleServiceExpireAPIKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service expire Api key o k response a status code equal to that given
func (o *HeadscaleServiceExpireAPIKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *HeadscaleServiceExpireAPIKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/apikey/expire][%d] headscaleServiceExpireApiKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpireAPIKeyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/apikey/expire][%d] headscaleServiceExpireApiKeyOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceExpireAPIKeyOK) GetPayload() models.V1ExpireAPIKeyResponse {
	return o.Payload
}

func (o *HeadscaleServiceExpireAPIKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceExpireAPIKeyDefault creates a HeadscaleServiceExpireAPIKeyDefault with default headers values
func NewHeadscaleServiceExpireAPIKeyDefault(code int) *HeadscaleServiceExpireAPIKeyDefault {
	return &HeadscaleServiceExpireAPIKeyDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceExpireAPIKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceExpireAPIKeyDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the headscale service expire Api key default response
func (o *HeadscaleServiceExpireAPIKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this headscale service expire Api key default response has a 2xx status code
func (o *HeadscaleServiceExpireAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service expire Api key default response has a 3xx status code
func (o *HeadscaleServiceExpireAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service expire Api key default response has a 4xx status code
func (o *HeadscaleServiceExpireAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service expire Api key default response has a 5xx status code
func (o *HeadscaleServiceExpireAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service expire Api key default response a status code equal to that given
func (o *HeadscaleServiceExpireAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HeadscaleServiceExpireAPIKeyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/apikey/expire][%d] HeadscaleService_ExpireApiKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpireAPIKeyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/apikey/expire][%d] HeadscaleService_ExpireApiKey default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceExpireAPIKeyDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceExpireAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
