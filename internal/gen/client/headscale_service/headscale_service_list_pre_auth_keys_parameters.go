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

// NewHeadscaleServiceListPreAuthKeysParams creates a new HeadscaleServiceListPreAuthKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceListPreAuthKeysParams() *HeadscaleServiceListPreAuthKeysParams {
	return &HeadscaleServiceListPreAuthKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceListPreAuthKeysParamsWithTimeout creates a new HeadscaleServiceListPreAuthKeysParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceListPreAuthKeysParamsWithTimeout(timeout time.Duration) *HeadscaleServiceListPreAuthKeysParams {
	return &HeadscaleServiceListPreAuthKeysParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceListPreAuthKeysParamsWithContext creates a new HeadscaleServiceListPreAuthKeysParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceListPreAuthKeysParamsWithContext(ctx context.Context) *HeadscaleServiceListPreAuthKeysParams {
	return &HeadscaleServiceListPreAuthKeysParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceListPreAuthKeysParamsWithHTTPClient creates a new HeadscaleServiceListPreAuthKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceListPreAuthKeysParamsWithHTTPClient(client *http.Client) *HeadscaleServiceListPreAuthKeysParams {
	return &HeadscaleServiceListPreAuthKeysParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceListPreAuthKeysParams contains all the parameters to send to the API endpoint

	for the headscale service list pre auth keys operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceListPreAuthKeysParams struct {

	// User.
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service list pre auth keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListPreAuthKeysParams) WithDefaults() *HeadscaleServiceListPreAuthKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service list pre auth keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceListPreAuthKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) WithTimeout(timeout time.Duration) *HeadscaleServiceListPreAuthKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) WithContext(ctx context.Context) *HeadscaleServiceListPreAuthKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) WithHTTPClient(client *http.Client) *HeadscaleServiceListPreAuthKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) WithUser(user *string) *HeadscaleServiceListPreAuthKeysParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the headscale service list pre auth keys params
func (o *HeadscaleServiceListPreAuthKeysParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceListPreAuthKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
