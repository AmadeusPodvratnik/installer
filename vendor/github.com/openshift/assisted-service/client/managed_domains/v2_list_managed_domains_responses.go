// Code generated by go-swagger; DO NOT EDIT.

package managed_domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2ListManagedDomainsReader is a Reader for the V2ListManagedDomains structure.
type V2ListManagedDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2ListManagedDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2ListManagedDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewV2ListManagedDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2ListManagedDomainsOK creates a V2ListManagedDomainsOK with default headers values
func NewV2ListManagedDomainsOK() *V2ListManagedDomainsOK {
	return &V2ListManagedDomainsOK{}
}

/* V2ListManagedDomainsOK describes a response with status code 200, with default header values.

Success.
*/
type V2ListManagedDomainsOK struct {
	Payload models.ListManagedDomains
}

func (o *V2ListManagedDomainsOK) Error() string {
	return fmt.Sprintf("[GET /v2/domains][%d] v2ListManagedDomainsOK  %+v", 200, o.Payload)
}
func (o *V2ListManagedDomainsOK) GetPayload() models.ListManagedDomains {
	return o.Payload
}

func (o *V2ListManagedDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2ListManagedDomainsInternalServerError creates a V2ListManagedDomainsInternalServerError with default headers values
func NewV2ListManagedDomainsInternalServerError() *V2ListManagedDomainsInternalServerError {
	return &V2ListManagedDomainsInternalServerError{}
}

/* V2ListManagedDomainsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2ListManagedDomainsInternalServerError struct {
	Payload *models.Error
}

func (o *V2ListManagedDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/domains][%d] v2ListManagedDomainsInternalServerError  %+v", 500, o.Payload)
}
func (o *V2ListManagedDomainsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2ListManagedDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
