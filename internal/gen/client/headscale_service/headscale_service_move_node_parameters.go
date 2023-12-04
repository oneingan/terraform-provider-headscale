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

// NewHeadscaleServiceMoveNodeParams creates a new HeadscaleServiceMoveNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadscaleServiceMoveNodeParams() *HeadscaleServiceMoveNodeParams {
	return &HeadscaleServiceMoveNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadscaleServiceMoveNodeParamsWithTimeout creates a new HeadscaleServiceMoveNodeParams object
// with the ability to set a timeout on a request.
func NewHeadscaleServiceMoveNodeParamsWithTimeout(timeout time.Duration) *HeadscaleServiceMoveNodeParams {
	return &HeadscaleServiceMoveNodeParams{
		timeout: timeout,
	}
}

// NewHeadscaleServiceMoveNodeParamsWithContext creates a new HeadscaleServiceMoveNodeParams object
// with the ability to set a context for a request.
func NewHeadscaleServiceMoveNodeParamsWithContext(ctx context.Context) *HeadscaleServiceMoveNodeParams {
	return &HeadscaleServiceMoveNodeParams{
		Context: ctx,
	}
}

// NewHeadscaleServiceMoveNodeParamsWithHTTPClient creates a new HeadscaleServiceMoveNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadscaleServiceMoveNodeParamsWithHTTPClient(client *http.Client) *HeadscaleServiceMoveNodeParams {
	return &HeadscaleServiceMoveNodeParams{
		HTTPClient: client,
	}
}

/*
HeadscaleServiceMoveNodeParams contains all the parameters to send to the API endpoint

	for the headscale service move node operation.

	Typically these are written to a http.Request.
*/
type HeadscaleServiceMoveNodeParams struct {

	// NodeID.
	//
	// Format: uint64
	NodeID string

	// User.
	User *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the headscale service move node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceMoveNodeParams) WithDefaults() *HeadscaleServiceMoveNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the headscale service move node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadscaleServiceMoveNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) WithTimeout(timeout time.Duration) *HeadscaleServiceMoveNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) WithContext(ctx context.Context) *HeadscaleServiceMoveNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) WithHTTPClient(client *http.Client) *HeadscaleServiceMoveNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeID adds the nodeID to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) WithNodeID(nodeID string) *HeadscaleServiceMoveNodeParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WithUser adds the user to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) WithUser(user *string) *HeadscaleServiceMoveNodeParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the headscale service move node params
func (o *HeadscaleServiceMoveNodeParams) SetUser(user *string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *HeadscaleServiceMoveNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param nodeId
	if err := r.SetPathParam("nodeId", o.NodeID); err != nil {
		return err
	}

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
