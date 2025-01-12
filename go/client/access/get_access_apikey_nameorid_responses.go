// Code generated by go-swagger; DO NOT EDIT.

package access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/loopholelabs/scale/go/client/models"
)

// GetAccessApikeyNameoridReader is a Reader for the GetAccessApikeyNameorid structure.
type GetAccessApikeyNameoridReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessApikeyNameoridReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessApikeyNameoridOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccessApikeyNameoridUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccessApikeyNameoridNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccessApikeyNameoridInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccessApikeyNameoridOK creates a GetAccessApikeyNameoridOK with default headers values
func NewGetAccessApikeyNameoridOK() *GetAccessApikeyNameoridOK {
	return &GetAccessApikeyNameoridOK{}
}

/*
GetAccessApikeyNameoridOK describes a response with status code 200, with default header values.

OK
*/
type GetAccessApikeyNameoridOK struct {
	Payload *models.ModelsGetAPIKeyResponse
}

// IsSuccess returns true when this get access apikey nameorid o k response has a 2xx status code
func (o *GetAccessApikeyNameoridOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access apikey nameorid o k response has a 3xx status code
func (o *GetAccessApikeyNameoridOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey nameorid o k response has a 4xx status code
func (o *GetAccessApikeyNameoridOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access apikey nameorid o k response has a 5xx status code
func (o *GetAccessApikeyNameoridOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access apikey nameorid o k response a status code equal to that given
func (o *GetAccessApikeyNameoridOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access apikey nameorid o k response
func (o *GetAccessApikeyNameoridOK) Code() int {
	return 200
}

func (o *GetAccessApikeyNameoridOK) Error() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridOK  %+v", 200, o.Payload)
}

func (o *GetAccessApikeyNameoridOK) String() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridOK  %+v", 200, o.Payload)
}

func (o *GetAccessApikeyNameoridOK) GetPayload() *models.ModelsGetAPIKeyResponse {
	return o.Payload
}

func (o *GetAccessApikeyNameoridOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsGetAPIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessApikeyNameoridUnauthorized creates a GetAccessApikeyNameoridUnauthorized with default headers values
func NewGetAccessApikeyNameoridUnauthorized() *GetAccessApikeyNameoridUnauthorized {
	return &GetAccessApikeyNameoridUnauthorized{}
}

/*
GetAccessApikeyNameoridUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccessApikeyNameoridUnauthorized struct {
	Payload string
}

// IsSuccess returns true when this get access apikey nameorid unauthorized response has a 2xx status code
func (o *GetAccessApikeyNameoridUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access apikey nameorid unauthorized response has a 3xx status code
func (o *GetAccessApikeyNameoridUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey nameorid unauthorized response has a 4xx status code
func (o *GetAccessApikeyNameoridUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access apikey nameorid unauthorized response has a 5xx status code
func (o *GetAccessApikeyNameoridUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get access apikey nameorid unauthorized response a status code equal to that given
func (o *GetAccessApikeyNameoridUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get access apikey nameorid unauthorized response
func (o *GetAccessApikeyNameoridUnauthorized) Code() int {
	return 401
}

func (o *GetAccessApikeyNameoridUnauthorized) Error() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccessApikeyNameoridUnauthorized) String() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccessApikeyNameoridUnauthorized) GetPayload() string {
	return o.Payload
}

func (o *GetAccessApikeyNameoridUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessApikeyNameoridNotFound creates a GetAccessApikeyNameoridNotFound with default headers values
func NewGetAccessApikeyNameoridNotFound() *GetAccessApikeyNameoridNotFound {
	return &GetAccessApikeyNameoridNotFound{}
}

/*
GetAccessApikeyNameoridNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAccessApikeyNameoridNotFound struct {
	Payload string
}

// IsSuccess returns true when this get access apikey nameorid not found response has a 2xx status code
func (o *GetAccessApikeyNameoridNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access apikey nameorid not found response has a 3xx status code
func (o *GetAccessApikeyNameoridNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey nameorid not found response has a 4xx status code
func (o *GetAccessApikeyNameoridNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get access apikey nameorid not found response has a 5xx status code
func (o *GetAccessApikeyNameoridNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get access apikey nameorid not found response a status code equal to that given
func (o *GetAccessApikeyNameoridNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get access apikey nameorid not found response
func (o *GetAccessApikeyNameoridNotFound) Code() int {
	return 404
}

func (o *GetAccessApikeyNameoridNotFound) Error() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridNotFound  %+v", 404, o.Payload)
}

func (o *GetAccessApikeyNameoridNotFound) String() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridNotFound  %+v", 404, o.Payload)
}

func (o *GetAccessApikeyNameoridNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetAccessApikeyNameoridNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessApikeyNameoridInternalServerError creates a GetAccessApikeyNameoridInternalServerError with default headers values
func NewGetAccessApikeyNameoridInternalServerError() *GetAccessApikeyNameoridInternalServerError {
	return &GetAccessApikeyNameoridInternalServerError{}
}

/*
GetAccessApikeyNameoridInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccessApikeyNameoridInternalServerError struct {
	Payload string
}

// IsSuccess returns true when this get access apikey nameorid internal server error response has a 2xx status code
func (o *GetAccessApikeyNameoridInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get access apikey nameorid internal server error response has a 3xx status code
func (o *GetAccessApikeyNameoridInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access apikey nameorid internal server error response has a 4xx status code
func (o *GetAccessApikeyNameoridInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access apikey nameorid internal server error response has a 5xx status code
func (o *GetAccessApikeyNameoridInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get access apikey nameorid internal server error response a status code equal to that given
func (o *GetAccessApikeyNameoridInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get access apikey nameorid internal server error response
func (o *GetAccessApikeyNameoridInternalServerError) Code() int {
	return 500
}

func (o *GetAccessApikeyNameoridInternalServerError) Error() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccessApikeyNameoridInternalServerError) String() string {
	return fmt.Sprintf("[GET /access/apikey/{nameorid}][%d] getAccessApikeyNameoridInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccessApikeyNameoridInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetAccessApikeyNameoridInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
