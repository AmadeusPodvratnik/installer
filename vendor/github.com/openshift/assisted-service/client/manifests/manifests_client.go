// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the manifests client
type API interface {
	/*
	   CreateClusterManifest Creates a manifest for customizing cluster installation.*/
	CreateClusterManifest(ctx context.Context, params *CreateClusterManifestParams) (*CreateClusterManifestCreated, error)
	/*
	   DeleteClusterManifest Deletes a manifest from the cluster.*/
	DeleteClusterManifest(ctx context.Context, params *DeleteClusterManifestParams) (*DeleteClusterManifestOK, error)
	/*
	   DownloadClusterManifest Downloads cluster manifest.*/
	DownloadClusterManifest(ctx context.Context, params *DownloadClusterManifestParams, writer io.Writer) (*DownloadClusterManifestOK, error)
	/*
	   ListClusterManifests Lists manifests for customizing cluster installation.*/
	ListClusterManifests(ctx context.Context, params *ListClusterManifestsParams) (*ListClusterManifestsOK, error)
	/*
	   V2CreateClusterManifest Creates a manifest for customizing cluster installation.*/
	V2CreateClusterManifest(ctx context.Context, params *V2CreateClusterManifestParams) (*V2CreateClusterManifestCreated, error)
	/*
	   V2DeleteClusterManifest Deletes a manifest from the cluster.*/
	V2DeleteClusterManifest(ctx context.Context, params *V2DeleteClusterManifestParams) (*V2DeleteClusterManifestOK, error)
	/*
	   V2ListClusterManifests Lists manifests for customizing cluster installation.*/
	V2ListClusterManifests(ctx context.Context, params *V2ListClusterManifestsParams) (*V2ListClusterManifestsOK, error)
	/*
	   V2DownloadClusterManifest Downloads cluster manifest.*/
	V2DownloadClusterManifest(ctx context.Context, params *V2DownloadClusterManifestParams, writer io.Writer) (*V2DownloadClusterManifestOK, error)
}

// New creates a new manifests API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for manifests API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateClusterManifest Creates a manifest for customizing cluster installation.
*/
func (a *Client) CreateClusterManifest(ctx context.Context, params *CreateClusterManifestParams) (*CreateClusterManifestCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateClusterManifest",
		Method:             "POST",
		PathPattern:        "/v1/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateClusterManifestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterManifestCreated), nil

}

/*
DeleteClusterManifest Deletes a manifest from the cluster.
*/
func (a *Client) DeleteClusterManifest(ctx context.Context, params *DeleteClusterManifestParams) (*DeleteClusterManifestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterManifest",
		Method:             "DELETE",
		PathPattern:        "/v1/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteClusterManifestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterManifestOK), nil

}

/*
DownloadClusterManifest Downloads cluster manifest.
*/
func (a *Client) DownloadClusterManifest(ctx context.Context, params *DownloadClusterManifestParams, writer io.Writer) (*DownloadClusterManifestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DownloadClusterManifest",
		Method:             "GET",
		PathPattern:        "/v1/clusters/{cluster_id}/manifests/files",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DownloadClusterManifestReader{formats: a.formats, writer: writer},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadClusterManifestOK), nil

}

/*
ListClusterManifests Lists manifests for customizing cluster installation.
*/
func (a *Client) ListClusterManifests(ctx context.Context, params *ListClusterManifestsParams) (*ListClusterManifestsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClusterManifests",
		Method:             "GET",
		PathPattern:        "/v1/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListClusterManifestsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClusterManifestsOK), nil

}

/*
V2CreateClusterManifest Creates a manifest for customizing cluster installation.
*/
func (a *Client) V2CreateClusterManifest(ctx context.Context, params *V2CreateClusterManifestParams) (*V2CreateClusterManifestCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2CreateClusterManifest",
		Method:             "POST",
		PathPattern:        "/v2/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2CreateClusterManifestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2CreateClusterManifestCreated), nil

}

/*
V2DeleteClusterManifest Deletes a manifest from the cluster.
*/
func (a *Client) V2DeleteClusterManifest(ctx context.Context, params *V2DeleteClusterManifestParams) (*V2DeleteClusterManifestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2DeleteClusterManifest",
		Method:             "DELETE",
		PathPattern:        "/v2/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2DeleteClusterManifestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2DeleteClusterManifestOK), nil

}

/*
V2ListClusterManifests Lists manifests for customizing cluster installation.
*/
func (a *Client) V2ListClusterManifests(ctx context.Context, params *V2ListClusterManifestsParams) (*V2ListClusterManifestsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "V2ListClusterManifests",
		Method:             "GET",
		PathPattern:        "/v2/clusters/{cluster_id}/manifests",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2ListClusterManifestsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2ListClusterManifestsOK), nil

}

/*
V2DownloadClusterManifest Downloads cluster manifest.
*/
func (a *Client) V2DownloadClusterManifest(ctx context.Context, params *V2DownloadClusterManifestParams, writer io.Writer) (*V2DownloadClusterManifestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "v2DownloadClusterManifest",
		Method:             "GET",
		PathPattern:        "/v2/clusters/{cluster_id}/manifests/files",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &V2DownloadClusterManifestReader{formats: a.formats, writer: writer},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*V2DownloadClusterManifestOK), nil

}
