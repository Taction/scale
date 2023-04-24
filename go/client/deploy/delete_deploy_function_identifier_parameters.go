// Code generated by go-swagger; DO NOT EDIT.

package deploy

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

// NewDeleteDeployFunctionIdentifierParams creates a new DeleteDeployFunctionIdentifierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeployFunctionIdentifierParams() *DeleteDeployFunctionIdentifierParams {
	return &DeleteDeployFunctionIdentifierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeployFunctionIdentifierParamsWithTimeout creates a new DeleteDeployFunctionIdentifierParams object
// with the ability to set a timeout on a request.
func NewDeleteDeployFunctionIdentifierParamsWithTimeout(timeout time.Duration) *DeleteDeployFunctionIdentifierParams {
	return &DeleteDeployFunctionIdentifierParams{
		timeout: timeout,
	}
}

// NewDeleteDeployFunctionIdentifierParamsWithContext creates a new DeleteDeployFunctionIdentifierParams object
// with the ability to set a context for a request.
func NewDeleteDeployFunctionIdentifierParamsWithContext(ctx context.Context) *DeleteDeployFunctionIdentifierParams {
	return &DeleteDeployFunctionIdentifierParams{
		Context: ctx,
	}
}

// NewDeleteDeployFunctionIdentifierParamsWithHTTPClient creates a new DeleteDeployFunctionIdentifierParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeployFunctionIdentifierParamsWithHTTPClient(client *http.Client) *DeleteDeployFunctionIdentifierParams {
	return &DeleteDeployFunctionIdentifierParams{
		HTTPClient: client,
	}
}

/*
DeleteDeployFunctionIdentifierParams contains all the parameters to send to the API endpoint

	for the delete deploy function identifier operation.

	Typically these are written to a http.Request.
*/
type DeleteDeployFunctionIdentifierParams struct {

	/* Identifier.

	   identifier
	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete deploy function identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeployFunctionIdentifierParams) WithDefaults() *DeleteDeployFunctionIdentifierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete deploy function identifier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeployFunctionIdentifierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) WithTimeout(timeout time.Duration) *DeleteDeployFunctionIdentifierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) WithContext(ctx context.Context) *DeleteDeployFunctionIdentifierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) WithHTTPClient(client *http.Client) *DeleteDeployFunctionIdentifierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) WithIdentifier(identifier string) *DeleteDeployFunctionIdentifierParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the delete deploy function identifier params
func (o *DeleteDeployFunctionIdentifierParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeployFunctionIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}