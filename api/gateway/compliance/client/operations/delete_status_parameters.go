// Code generated by go-swagger; DO NOT EDIT.

package operations

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"

	models "github.com/panther-labs/panther/api/gateway/compliance/models"
)

// NewDeleteStatusParams creates a new DeleteStatusParams object
// with the default values initialized.
func NewDeleteStatusParams() *DeleteStatusParams {
	var ()
	return &DeleteStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStatusParamsWithTimeout creates a new DeleteStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStatusParamsWithTimeout(timeout time.Duration) *DeleteStatusParams {
	var ()
	return &DeleteStatusParams{

		timeout: timeout,
	}
}

// NewDeleteStatusParamsWithContext creates a new DeleteStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStatusParamsWithContext(ctx context.Context) *DeleteStatusParams {
	var ()
	return &DeleteStatusParams{

		Context: ctx,
	}
}

// NewDeleteStatusParamsWithHTTPClient creates a new DeleteStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStatusParamsWithHTTPClient(client *http.Client) *DeleteStatusParams {
	var ()
	return &DeleteStatusParams{
		HTTPClient: client,
	}
}

/*DeleteStatusParams contains all the parameters to send to the API endpoint
for the delete status operation typically these are written to a http.Request
*/
type DeleteStatusParams struct {

	/*Body*/
	Body *models.DeleteStatusBatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete status params
func (o *DeleteStatusParams) WithTimeout(timeout time.Duration) *DeleteStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete status params
func (o *DeleteStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete status params
func (o *DeleteStatusParams) WithContext(ctx context.Context) *DeleteStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete status params
func (o *DeleteStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete status params
func (o *DeleteStatusParams) WithHTTPClient(client *http.Client) *DeleteStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete status params
func (o *DeleteStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete status params
func (o *DeleteStatusParams) WithBody(body *models.DeleteStatusBatch) *DeleteStatusParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete status params
func (o *DeleteStatusParams) SetBody(body *models.DeleteStatusBatch) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
