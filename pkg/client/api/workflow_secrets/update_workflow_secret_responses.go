// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// UpdateWorkflowSecretReader is a Reader for the UpdateWorkflowSecret structure.
type UpdateWorkflowSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWorkflowSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWorkflowSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateWorkflowSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateWorkflowSecretOK creates a UpdateWorkflowSecretOK with default headers values
func NewUpdateWorkflowSecretOK() *UpdateWorkflowSecretOK {
	return &UpdateWorkflowSecretOK{}
}

/*UpdateWorkflowSecretOK handles this case with default header values.

Secret successfully updated
*/
type UpdateWorkflowSecretOK struct {
	Payload *UpdateWorkflowSecretOKBody
}

func (o *UpdateWorkflowSecretOK) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflowName}/secrets/{workflowSecretKey}][%d] updateWorkflowSecretOK  %+v", 200, o.Payload)
}

func (o *UpdateWorkflowSecretOK) GetPayload() *UpdateWorkflowSecretOKBody {
	return o.Payload
}

func (o *UpdateWorkflowSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateWorkflowSecretOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateWorkflowSecretDefault creates a UpdateWorkflowSecretDefault with default headers values
func NewUpdateWorkflowSecretDefault(code int) *UpdateWorkflowSecretDefault {
	return &UpdateWorkflowSecretDefault{
		_statusCode: code,
	}
}

/*UpdateWorkflowSecretDefault handles this case with default header values.

An error occurred
*/
type UpdateWorkflowSecretDefault struct {
	_statusCode int

	Payload *UpdateWorkflowSecretDefaultBody
}

// Code gets the status code for the update workflow secret default response
func (o *UpdateWorkflowSecretDefault) Code() int {
	return o._statusCode
}

func (o *UpdateWorkflowSecretDefault) Error() string {
	return fmt.Sprintf("[PUT /api/workflows/{workflowName}/secrets/{workflowSecretKey}][%d] updateWorkflowSecret default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateWorkflowSecretDefault) GetPayload() *UpdateWorkflowSecretDefaultBody {
	return o.Payload
}

func (o *UpdateWorkflowSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateWorkflowSecretDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateWorkflowSecretBody The request type for updating a workflow secret
swagger:model UpdateWorkflowSecretBody
*/
type UpdateWorkflowSecretBody struct {

	// value
	// Required: true
	Value models.BinaryString `json:"value"`
}

// Validate validates this update workflow secret body
func (o *UpdateWorkflowSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowSecretBody) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowSecretBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowSecretBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateWorkflowSecretDefaultBody Error response
swagger:model UpdateWorkflowSecretDefaultBody
*/
type UpdateWorkflowSecretDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this update workflow secret default body
func (o *UpdateWorkflowSecretDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowSecretDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateWorkflowSecret default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowSecretDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowSecretDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowSecretDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateWorkflowSecretOKBody update workflow secret o k body
swagger:model UpdateWorkflowSecretOKBody
*/
type UpdateWorkflowSecretOKBody struct {
	models.Entity

	// secret
	Secret *models.WorkflowSecretSummary `json:"secret,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *UpdateWorkflowSecretOKBody) UnmarshalJSON(raw []byte) error {
	// UpdateWorkflowSecretOKBodyAO0
	var updateWorkflowSecretOKBodyAO0 models.Entity
	if err := swag.ReadJSON(raw, &updateWorkflowSecretOKBodyAO0); err != nil {
		return err
	}
	o.Entity = updateWorkflowSecretOKBodyAO0

	// UpdateWorkflowSecretOKBodyAO1
	var dataUpdateWorkflowSecretOKBodyAO1 struct {
		Secret *models.WorkflowSecretSummary `json:"secret,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataUpdateWorkflowSecretOKBodyAO1); err != nil {
		return err
	}

	o.Secret = dataUpdateWorkflowSecretOKBodyAO1.Secret

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o UpdateWorkflowSecretOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	updateWorkflowSecretOKBodyAO0, err := swag.WriteJSON(o.Entity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, updateWorkflowSecretOKBodyAO0)

	var dataUpdateWorkflowSecretOKBodyAO1 struct {
		Secret *models.WorkflowSecretSummary `json:"secret,omitempty"`
	}

	dataUpdateWorkflowSecretOKBodyAO1.Secret = o.Secret

	jsonDataUpdateWorkflowSecretOKBodyAO1, errUpdateWorkflowSecretOKBodyAO1 := swag.WriteJSON(dataUpdateWorkflowSecretOKBodyAO1)
	if errUpdateWorkflowSecretOKBodyAO1 != nil {
		return nil, errUpdateWorkflowSecretOKBodyAO1
	}
	_parts = append(_parts, jsonDataUpdateWorkflowSecretOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update workflow secret o k body
func (o *UpdateWorkflowSecretOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Entity
	if err := o.Entity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateWorkflowSecretOKBody) validateSecret(formats strfmt.Registry) error {

	if swag.IsZero(o.Secret) { // not required
		return nil
	}

	if o.Secret != nil {
		if err := o.Secret.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateWorkflowSecretOK" + "." + "secret")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateWorkflowSecretOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateWorkflowSecretOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateWorkflowSecretOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}