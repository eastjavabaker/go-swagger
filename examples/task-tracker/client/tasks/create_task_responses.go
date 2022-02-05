// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

// CreateTaskReader is a Reader for the CreateTask structure.
type CreateTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateTaskCreated creates a CreateTaskCreated with default headers values
func NewCreateTaskCreated() *CreateTaskCreated {
	return &CreateTaskCreated{}
}

/* CreateTaskCreated describes a response with status code 201, with default header values.

Task created
*/
type CreateTaskCreated struct {

	/* URL to the newly added Task

	   Format: uri
	*/
	Location strfmt.URI
}

// IsSuccess returns true when this create task created response returns a 2xx status code
func (o *CreateTaskCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create task created response returns a 3xx status code
func (o *CreateTaskCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create task created response returns a 4xx status code
func (o *CreateTaskCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create task created response returns a 5xx status code
func (o *CreateTaskCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create task created response returns a 4xx status code
func (o *CreateTaskCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTaskCreated) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] createTaskCreated ", 201)
}

func (o *CreateTaskCreated) String() string {
	return fmt.Sprintf("[POST /tasks][%d] createTaskCreated ", 201)
}

func (o *CreateTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		vallocation, err := formats.Parse("uri", hdrLocation)
		if err != nil {
			return errors.InvalidType("Location", "header", "strfmt.URI", hdrLocation)
		}
		o.Location = *(vallocation.(*strfmt.URI))
	}

	return nil
}

// NewCreateTaskDefault creates a CreateTaskDefault with default headers values
func NewCreateTaskDefault(code int) *CreateTaskDefault {
	return &CreateTaskDefault{
		_statusCode: code,
	}
}

/* CreateTaskDefault describes a response with status code -1, with default header values.

Error response
*/
type CreateTaskDefault struct {
	_statusCode int
	XErrorCode  string

	Payload *models.Error
}

// Code gets the status code for the create task default response
func (o *CreateTaskDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create task default response returns a 2xx status code
func (o *CreateTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create task default response returns a 3xx status code
func (o *CreateTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create task default response returns a 4xx status code
func (o *CreateTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create task default response returns a 5xx status code
func (o *CreateTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create task default response returns a 4xx status code
func (o *CreateTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateTaskDefault) Error() string {
	return fmt.Sprintf("[POST /tasks][%d] createTask default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTaskDefault) String() string {
	return fmt.Sprintf("[POST /tasks][%d] createTask default  %+v", o._statusCode, o.Payload)
}

func (o *CreateTaskDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Error-Code
	hdrXErrorCode := response.GetHeader("X-Error-Code")

	if hdrXErrorCode != "" {
		o.XErrorCode = hdrXErrorCode
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
