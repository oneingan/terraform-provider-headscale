// Code generated by go-swagger; DO NOT EDIT.

package headscale_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHeadscaleServiceSetTagsParams creates a new HeadscaleServiceSetTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceSetTagsParams() *HeadscaleServiceSetTagsParams {
	return &HeadscaleServiceSetTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceSetTagsParamsWithTimeout creates a new HeadscaleServiceSetTagsParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceSetTagsParamsWithTimeout(timeout time.Duration) *HeadscaleServiceSetTagsParams {
	return &HeadscaleServiceSetTagsParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceSetTagsParamsWithContext creates a new HeadscaleServiceSetTagsParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceSetTagsParamsWithContext(ctx context.Context) *HeadscaleServiceSetTagsParams {
	return &HeadscaleServiceSetTagsParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceSetTagsParamsWithHTTPClient creates a new HeadscaleServiceSetTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceSetTagsParamsWithHTTPClient(client *http.Client) *HeadscaleServiceSetTagsParams {
	return &HeadscaleServiceSetTagsParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceSetTagsParams contains all the parameters to send to the API endpoint

	for the headscale service set tags operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceSetTagsParams struct {

	// Body.
	Body HeadscaleServiceSetTagsBody

	// MachineID.
	//
	// Format: uint64
	MachineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service set tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceSetTagsParams) WithDefaults() *HeadscaleServiceSetTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service set tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceSetTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) WithTimeout(timeout time.Duration) *HeadscaleServiceSetTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) WithContext(ctx context.Context) *HeadscaleServiceSetTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) WithHTTPClient(client *http.Client) *HeadscaleServiceSetTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) WithBody(body HeadscaleServiceSetTagsBody) *HeadscaleServiceSetTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) SetBody(body HeadscaleServiceSetTagsBody) {
	o.Body = body
}

// WithMachineID adds the machineID to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) WithMachineID(machineID string) *HeadscaleServiceSetTagsParams {
	o.SetMachineID(machineID)
	return o
}

// SetMachineID adds the machineId to the headscale service set tags params
func (o *HeadscaleServiceSetTagsParams) SetMachineID(machineID string) {
	o.MachineID = machineID
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceSetTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param machineId
	if err := r.SetPathParam("machineId", o.MachineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
