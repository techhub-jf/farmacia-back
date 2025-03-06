package response

import (
	"errors"
	"net/http"
)

type Response struct { //nolint:errname
	Status      int
	Payload     any
	Headers     map[string]string
	InternalErr error
	LogAttrs    map[string]any
	OmitLogs    bool
}

type Error struct {
	Status  int32  `json:"status"`
	Message string `json:"message,omitempty"`
}

func (r *Response) Error() string {
	return r.InternalErr.Error()
}

func (r *Response) WithHeaders(header map[string]string) *Response {
	r.Headers = header

	return r
}

func (r *Response) WithLogAttrs(attrs map[string]any) *Response {
	r.LogAttrs = attrs

	return r
}

func (r *Response) WithOmittedLogs() *Response {
	r.OmitLogs = true

	return r
}

// Success

func OK(payload any) *Response {
	return &Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

func Created(payload any) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

func Accepted(payload any) *Response {
	return &Response{
		Status:  http.StatusAccepted,
		Payload: payload,
	}
}

func NoContent() *Response {
	return &Response{
		Status: http.StatusNoContent,
	}
}

// Failure

func BadRequest(err error, message string) *Response {
	return &Response{
		Status: http.StatusBadRequest,
		Payload: Error{
			Status:  http.StatusBadRequest,
			Message: message,
		},
		InternalErr: err,
	}
}

func Conflict(err error, message string) *Response {
	return &Response{
		Status: http.StatusConflict,
		Payload: Error{
			Status:  http.StatusConflict,
			Message: message,
		},
		InternalErr: err,
	}
}

func Unauthorized(message string) *Response {
	if message == "" {
		message = "user is not authorized to perform this operation"
	}

	return &Response{
		Status: http.StatusUnauthorized,
		Payload: Error{
			Status:  http.StatusUnauthorized,
			Message: message,
		},
		InternalErr: errors.New("unauthorized"),
	}
}

func NotFound(err error, message string) *Response {
	return &Response{
		Status: http.StatusNotFound,
		Payload: Error{
			Status:  http.StatusNotFound,
			Message: message,
		},
		InternalErr: err,
	}
}

func MethodNotAllowed() Response {
	return Response{
		Status: http.StatusMethodNotAllowed,
		Payload: Error{
			Status:  http.StatusMethodNotAllowed,
			Message: "the http method used is not supported by this resource",
		},
		InternalErr: errors.New("method not allowed"),
	}
}

func InternalServerError(err error) *Response {
	return &Response{
		Status: http.StatusInternalServerError,
		Payload: Error{
			Status:  http.StatusInternalServerError,
			Message: "an unexpected error has occurred",
		},
		InternalErr: err,
	}
}
