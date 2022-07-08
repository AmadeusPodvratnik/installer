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

// NewV2UpdateClusterInstallConfigParams creates a new V2UpdateClusterInstallConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2UpdateClusterInstallConfigParams() *V2UpdateClusterInstallConfigParams {
	return &V2UpdateClusterInstallConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2UpdateClusterInstallConfigParamsWithTimeout creates a new V2UpdateClusterInstallConfigParams object
// with the ability to set a timeout on a request.
func NewV2UpdateClusterInstallConfigParamsWithTimeout(timeout time.Duration) *V2UpdateClusterInstallConfigParams {
	return &V2UpdateClusterInstallConfigParams{
		timeout: timeout,
	}
}

// NewV2UpdateClusterInstallConfigParamsWithContext creates a new V2UpdateClusterInstallConfigParams object
// with the ability to set a context for a request.
func NewV2UpdateClusterInstallConfigParamsWithContext(ctx context.Context) *V2UpdateClusterInstallConfigParams {
	return &V2UpdateClusterInstallConfigParams{
		Context: ctx,
	}
}

// NewV2UpdateClusterInstallConfigParamsWithHTTPClient creates a new V2UpdateClusterInstallConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2UpdateClusterInstallConfigParamsWithHTTPClient(client *http.Client) *V2UpdateClusterInstallConfigParams {
	return &V2UpdateClusterInstallConfigParams{
		HTTPClient: client,
	}
}

/* V2UpdateClusterInstallConfigParams contains all the parameters to send to the API endpoint
   for the v2 update cluster install config operation.

   Typically these are written to a http.Request.
*/
type V2UpdateClusterInstallConfigParams struct {

	/* ClusterID.

	   The cluster whose install config is being updated.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* InstallConfigParams.

	   Install config overrides.
	*/
	InstallConfigParams string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 update cluster install config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2UpdateClusterInstallConfigParams) WithDefaults() *V2UpdateClusterInstallConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 update cluster install config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2UpdateClusterInstallConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) WithTimeout(timeout time.Duration) *V2UpdateClusterInstallConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) WithContext(ctx context.Context) *V2UpdateClusterInstallConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) WithHTTPClient(client *http.Client) *V2UpdateClusterInstallConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) WithClusterID(clusterID strfmt.UUID) *V2UpdateClusterInstallConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithInstallConfigParams adds the installConfigParams to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) WithInstallConfigParams(installConfigParams string) *V2UpdateClusterInstallConfigParams {
	o.SetInstallConfigParams(installConfigParams)
	return o
}

// SetInstallConfigParams adds the installConfigParams to the v2 update cluster install config params
func (o *V2UpdateClusterInstallConfigParams) SetInstallConfigParams(installConfigParams string) {
	o.InstallConfigParams = installConfigParams
}

// WriteToRequest writes these params to a swagger request
func (o *V2UpdateClusterInstallConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.InstallConfigParams); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
