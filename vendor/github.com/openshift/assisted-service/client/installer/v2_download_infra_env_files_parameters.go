// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewV2DownloadInfraEnvFilesParams creates a new V2DownloadInfraEnvFilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2DownloadInfraEnvFilesParams() *V2DownloadInfraEnvFilesParams {
	return &V2DownloadInfraEnvFilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2DownloadInfraEnvFilesParamsWithTimeout creates a new V2DownloadInfraEnvFilesParams object
// with the ability to set a timeout on a request.
func NewV2DownloadInfraEnvFilesParamsWithTimeout(timeout time.Duration) *V2DownloadInfraEnvFilesParams {
	return &V2DownloadInfraEnvFilesParams{
		timeout: timeout,
	}
}

// NewV2DownloadInfraEnvFilesParamsWithContext creates a new V2DownloadInfraEnvFilesParams object
// with the ability to set a context for a request.
func NewV2DownloadInfraEnvFilesParamsWithContext(ctx context.Context) *V2DownloadInfraEnvFilesParams {
	return &V2DownloadInfraEnvFilesParams{
		Context: ctx,
	}
}

// NewV2DownloadInfraEnvFilesParamsWithHTTPClient creates a new V2DownloadInfraEnvFilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2DownloadInfraEnvFilesParamsWithHTTPClient(client *http.Client) *V2DownloadInfraEnvFilesParams {
	return &V2DownloadInfraEnvFilesParams{
		HTTPClient: client,
	}
}

/* V2DownloadInfraEnvFilesParams contains all the parameters to send to the API endpoint
   for the v2 download infra env files operation.

   Typically these are written to a http.Request.
*/
type V2DownloadInfraEnvFilesParams struct {

	/* FileName.

	   The file to be downloaded.
	*/
	FileName string

	/* InfraEnvID.

	   The infra-env whose file should be downloaded.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 download infra env files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2DownloadInfraEnvFilesParams) WithDefaults() *V2DownloadInfraEnvFilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 download infra env files params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2DownloadInfraEnvFilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) WithTimeout(timeout time.Duration) *V2DownloadInfraEnvFilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) WithContext(ctx context.Context) *V2DownloadInfraEnvFilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) WithHTTPClient(client *http.Client) *V2DownloadInfraEnvFilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileName adds the fileName to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) WithFileName(fileName string) *V2DownloadInfraEnvFilesParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithInfraEnvID adds the infraEnvID to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2DownloadInfraEnvFilesParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 download infra env files params
func (o *V2DownloadInfraEnvFilesParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *V2DownloadInfraEnvFilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {

		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
