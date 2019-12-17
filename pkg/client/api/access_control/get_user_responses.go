// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

A user representation
*/
type GetUserOK struct {
	Payload *GetUserOKBody
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /api/users/{userId}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) GetPayload() *GetUserOKBody {
	return o.Payload
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDefault creates a GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	return &GetUserDefault{
		_statusCode: code,
	}
}

/*GetUserDefault handles this case with default header values.

An error occurred
*/
type GetUserDefault struct {
	_statusCode int

	Payload *GetUserDefaultBody
}

// Code gets the status code for the get user default response
func (o *GetUserDefault) Code() int {
	return o._statusCode
}

func (o *GetUserDefault) Error() string {
	return fmt.Sprintf("[GET /api/users/{userId}][%d] getUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) GetPayload() *GetUserDefaultBody {
	return o.Payload
}

func (o *GetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetUserDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserDefaultBody Error response
swagger:model GetUserDefaultBody
*/
type GetUserDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get user default body
func (o *GetUserDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUser default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetUserDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetUserOKBody get user o k body
swagger:model GetUserOKBody
*/
type GetUserOKBody struct {
	models.UserEntity

	// user
	User *models.User `json:"user,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetUserOKBody) UnmarshalJSON(raw []byte) error {
	// GetUserOKBodyAO0
	var getUserOKBodyAO0 models.UserEntity
	if err := swag.ReadJSON(raw, &getUserOKBodyAO0); err != nil {
		return err
	}
	o.UserEntity = getUserOKBodyAO0

	// GetUserOKBodyAO1
	var dataGetUserOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetUserOKBodyAO1); err != nil {
		return err
	}

	o.User = dataGetUserOKBodyAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetUserOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getUserOKBodyAO0, err := swag.WriteJSON(o.UserEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getUserOKBodyAO0)

	var dataGetUserOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}

	dataGetUserOKBodyAO1.User = o.User

	jsonDataGetUserOKBodyAO1, errGetUserOKBodyAO1 := swag.WriteJSON(dataGetUserOKBodyAO1)
	if errGetUserOKBodyAO1 != nil {
		return nil, errGetUserOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetUserOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get user o k body
func (o *GetUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.UserEntity
	if err := o.UserEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUserOKBody) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getUserOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}