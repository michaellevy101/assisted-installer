// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// V2PostStepReplyNoContentCode is the HTTP code returned for type V2PostStepReplyNoContent
const V2PostStepReplyNoContentCode int = 204

/*V2PostStepReplyNoContent Success.

swagger:response v2PostStepReplyNoContent
*/
type V2PostStepReplyNoContent struct {
}

// NewV2PostStepReplyNoContent creates V2PostStepReplyNoContent with default headers values
func NewV2PostStepReplyNoContent() *V2PostStepReplyNoContent {

	return &V2PostStepReplyNoContent{}
}

// WriteResponse to the client
func (o *V2PostStepReplyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// V2PostStepReplyBadRequestCode is the HTTP code returned for type V2PostStepReplyBadRequest
const V2PostStepReplyBadRequestCode int = 400

/*V2PostStepReplyBadRequest Error.

swagger:response v2PostStepReplyBadRequest
*/
type V2PostStepReplyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyBadRequest creates V2PostStepReplyBadRequest with default headers values
func NewV2PostStepReplyBadRequest() *V2PostStepReplyBadRequest {

	return &V2PostStepReplyBadRequest{}
}

// WithPayload adds the payload to the v2 post step reply bad request response
func (o *V2PostStepReplyBadRequest) WithPayload(payload *models.Error) *V2PostStepReplyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply bad request response
func (o *V2PostStepReplyBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyUnauthorizedCode is the HTTP code returned for type V2PostStepReplyUnauthorized
const V2PostStepReplyUnauthorizedCode int = 401

/*V2PostStepReplyUnauthorized Unauthorized.

swagger:response v2PostStepReplyUnauthorized
*/
type V2PostStepReplyUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2PostStepReplyUnauthorized creates V2PostStepReplyUnauthorized with default headers values
func NewV2PostStepReplyUnauthorized() *V2PostStepReplyUnauthorized {

	return &V2PostStepReplyUnauthorized{}
}

// WithPayload adds the payload to the v2 post step reply unauthorized response
func (o *V2PostStepReplyUnauthorized) WithPayload(payload *models.InfraError) *V2PostStepReplyUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply unauthorized response
func (o *V2PostStepReplyUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyForbiddenCode is the HTTP code returned for type V2PostStepReplyForbidden
const V2PostStepReplyForbiddenCode int = 403

/*V2PostStepReplyForbidden Forbidden.

swagger:response v2PostStepReplyForbidden
*/
type V2PostStepReplyForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewV2PostStepReplyForbidden creates V2PostStepReplyForbidden with default headers values
func NewV2PostStepReplyForbidden() *V2PostStepReplyForbidden {

	return &V2PostStepReplyForbidden{}
}

// WithPayload adds the payload to the v2 post step reply forbidden response
func (o *V2PostStepReplyForbidden) WithPayload(payload *models.InfraError) *V2PostStepReplyForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply forbidden response
func (o *V2PostStepReplyForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyNotFoundCode is the HTTP code returned for type V2PostStepReplyNotFound
const V2PostStepReplyNotFoundCode int = 404

/*V2PostStepReplyNotFound Error.

swagger:response v2PostStepReplyNotFound
*/
type V2PostStepReplyNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyNotFound creates V2PostStepReplyNotFound with default headers values
func NewV2PostStepReplyNotFound() *V2PostStepReplyNotFound {

	return &V2PostStepReplyNotFound{}
}

// WithPayload adds the payload to the v2 post step reply not found response
func (o *V2PostStepReplyNotFound) WithPayload(payload *models.Error) *V2PostStepReplyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply not found response
func (o *V2PostStepReplyNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyMethodNotAllowedCode is the HTTP code returned for type V2PostStepReplyMethodNotAllowed
const V2PostStepReplyMethodNotAllowedCode int = 405

/*V2PostStepReplyMethodNotAllowed Method Not Allowed.

swagger:response v2PostStepReplyMethodNotAllowed
*/
type V2PostStepReplyMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyMethodNotAllowed creates V2PostStepReplyMethodNotAllowed with default headers values
func NewV2PostStepReplyMethodNotAllowed() *V2PostStepReplyMethodNotAllowed {

	return &V2PostStepReplyMethodNotAllowed{}
}

// WithPayload adds the payload to the v2 post step reply method not allowed response
func (o *V2PostStepReplyMethodNotAllowed) WithPayload(payload *models.Error) *V2PostStepReplyMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply method not allowed response
func (o *V2PostStepReplyMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyInternalServerErrorCode is the HTTP code returned for type V2PostStepReplyInternalServerError
const V2PostStepReplyInternalServerErrorCode int = 500

/*V2PostStepReplyInternalServerError Error.

swagger:response v2PostStepReplyInternalServerError
*/
type V2PostStepReplyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyInternalServerError creates V2PostStepReplyInternalServerError with default headers values
func NewV2PostStepReplyInternalServerError() *V2PostStepReplyInternalServerError {

	return &V2PostStepReplyInternalServerError{}
}

// WithPayload adds the payload to the v2 post step reply internal server error response
func (o *V2PostStepReplyInternalServerError) WithPayload(payload *models.Error) *V2PostStepReplyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply internal server error response
func (o *V2PostStepReplyInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyNotImplementedCode is the HTTP code returned for type V2PostStepReplyNotImplemented
const V2PostStepReplyNotImplementedCode int = 501

/*V2PostStepReplyNotImplemented Not implemented.

swagger:response v2PostStepReplyNotImplemented
*/
type V2PostStepReplyNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyNotImplemented creates V2PostStepReplyNotImplemented with default headers values
func NewV2PostStepReplyNotImplemented() *V2PostStepReplyNotImplemented {

	return &V2PostStepReplyNotImplemented{}
}

// WithPayload adds the payload to the v2 post step reply not implemented response
func (o *V2PostStepReplyNotImplemented) WithPayload(payload *models.Error) *V2PostStepReplyNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply not implemented response
func (o *V2PostStepReplyNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// V2PostStepReplyServiceUnavailableCode is the HTTP code returned for type V2PostStepReplyServiceUnavailable
const V2PostStepReplyServiceUnavailableCode int = 503

/*V2PostStepReplyServiceUnavailable Unavailable.

swagger:response v2PostStepReplyServiceUnavailable
*/
type V2PostStepReplyServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewV2PostStepReplyServiceUnavailable creates V2PostStepReplyServiceUnavailable with default headers values
func NewV2PostStepReplyServiceUnavailable() *V2PostStepReplyServiceUnavailable {

	return &V2PostStepReplyServiceUnavailable{}
}

// WithPayload adds the payload to the v2 post step reply service unavailable response
func (o *V2PostStepReplyServiceUnavailable) WithPayload(payload *models.Error) *V2PostStepReplyServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the v2 post step reply service unavailable response
func (o *V2PostStepReplyServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *V2PostStepReplyServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
