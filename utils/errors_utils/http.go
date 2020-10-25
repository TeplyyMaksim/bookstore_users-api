package errors_utils

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Status 		int			`json:"status"`
	Message 	string		`json:"message"`
	Error		string		`json:"error"`
}

func NewError (message string) error {
	return errors.New(message)
}

func NewBadRequestError(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *HttpError {
	return &HttpError{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "internal_server_error",
	}
}