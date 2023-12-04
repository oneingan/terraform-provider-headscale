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

// HeadscaleServiceListAPIKeysReader is a Reader for the HeadscaleServiceListAPIKeys structure.
type HeadscaleServiceListAPIKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadscaleServiceListAPIKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadscaleServiceListAPIKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHeadscaleServiceListAPIKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHeadscaleServiceListAPIKeysOK creates a HeadscaleServiceListAPIKeysOK with default headers values
func NewHeadscaleServiceListAPIKeysOK() *HeadscaleServiceListAPIKeysOK {
	return &HeadscaleServiceListAPIKeysOK{}
}

/*
HeadscaleServiceListAPIKeysOK describes a response with status code 200, with default header values.

A successful response.
*/
type HeadscaleServiceListAPIKeysOK struct {
	Payload *models.V1ListAPIKeysResponse
}

// IsSuccess returns true when this headscale service list Api keys o k response has a 2xx status code
func (o *HeadscaleServiceListAPIKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this headscale service list Api keys o k response has a 3xx status code
func (o *HeadscaleServiceListAPIKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this headscale service list Api keys o k response has a 4xx status code
func (o *HeadscaleServiceListAPIKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this headscale service list Api keys o k response has a 5xx status code
func (o *HeadscaleServiceListAPIKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this headscale service list Api keys o k response a status code equal to that given
func (o *HeadscaleServiceListAPIKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the headscale service list Api keys o k response
func (o *HeadscaleServiceListAPIKeysOK) Code() int {
	return 200
}

func (o *HeadscaleServiceListAPIKeysOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikey][%d] headscaleServiceListApiKeysOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListAPIKeysOK) String() string {
	return fmt.Sprintf("[GET /api/v1/apikey][%d] headscaleServiceListApiKeysOK  %+v", 200, o.Payload)
}

func (o *HeadscaleServiceListAPIKeysOK) GetPayload() *models.V1ListAPIKeysResponse {
	return o.Payload
}

func (o *HeadscaleServiceListAPIKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListAPIKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHeadscaleServiceListAPIKeysDefault creates a HeadscaleServiceListAPIKeysDefault with default headers values
func NewHeadscaleServiceListAPIKeysDefault(code int) *HeadscaleServiceListAPIKeysDefault {
	return &HeadscaleServiceListAPIKeysDefault{
		_statusCode: code,
	}
}

/*
HeadscaleServiceListAPIKeysDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HeadscaleServiceListAPIKeysDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this headscale service list Api keys default response has a 2xx status code
func (o *HeadscaleServiceListAPIKeysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this headscale service list Api keys default response has a 3xx status code
func (o *HeadscaleServiceListAPIKeysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this headscale service list Api keys default response has a 4xx status code
func (o *HeadscaleServiceListAPIKeysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this headscale service list Api keys default response has a 5xx status code
func (o *HeadscaleServiceListAPIKeysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this headscale service list Api keys default response a status code equal to that given
func (o *HeadscaleServiceListAPIKeysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the headscale service list Api keys default response
func (o *HeadscaleServiceListAPIKeysDefault) Code() int {
	return o._statusCode
}

func (o *HeadscaleServiceListAPIKeysDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikey][%d] HeadscaleService_ListApiKeys default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListAPIKeysDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/apikey][%d] HeadscaleService_ListApiKeys default  %+v", o._statusCode, o.Payload)
}

func (o *HeadscaleServiceListAPIKeysDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *HeadscaleServiceListAPIKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
