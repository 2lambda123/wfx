// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package workflows

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
	"github.com/siemens/wfx/generated/model"
)

// NewPostWorkflowsParams creates a new PostWorkflowsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWorkflowsParams() *PostWorkflowsParams {
	return &PostWorkflowsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkflowsParamsWithTimeout creates a new PostWorkflowsParams object
// with the ability to set a timeout on a request.
func NewPostWorkflowsParamsWithTimeout(timeout time.Duration) *PostWorkflowsParams {
	return &PostWorkflowsParams{
		timeout: timeout,
	}
}

// NewPostWorkflowsParamsWithContext creates a new PostWorkflowsParams object
// with the ability to set a context for a request.
func NewPostWorkflowsParamsWithContext(ctx context.Context) *PostWorkflowsParams {
	return &PostWorkflowsParams{
		Context: ctx,
	}
}

// NewPostWorkflowsParamsWithHTTPClient creates a new PostWorkflowsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWorkflowsParamsWithHTTPClient(client *http.Client) *PostWorkflowsParams {
	return &PostWorkflowsParams{
		HTTPClient: client,
	}
}

/*
PostWorkflowsParams contains all the parameters to send to the API endpoint

	for the post workflows operation.

	Typically these are written to a http.Request.
*/
type PostWorkflowsParams struct {

	/* Workflow.

	   Workflow object to be added
	*/
	Workflow *model.Workflow

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkflowsParams) WithDefaults() *PostWorkflowsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post workflows params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWorkflowsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post workflows params
func (o *PostWorkflowsParams) WithTimeout(timeout time.Duration) *PostWorkflowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workflows params
func (o *PostWorkflowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workflows params
func (o *PostWorkflowsParams) WithContext(ctx context.Context) *PostWorkflowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workflows params
func (o *PostWorkflowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workflows params
func (o *PostWorkflowsParams) WithHTTPClient(client *http.Client) *PostWorkflowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workflows params
func (o *PostWorkflowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflow adds the workflow to the post workflows params
func (o *PostWorkflowsParams) WithWorkflow(workflow *model.Workflow) *PostWorkflowsParams {
	o.SetWorkflow(workflow)
	return o
}

// SetWorkflow adds the workflow to the post workflows params
func (o *PostWorkflowsParams) SetWorkflow(workflow *model.Workflow) {
	o.Workflow = workflow
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Workflow != nil {
		if err := r.SetBodyParam(o.Workflow); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
