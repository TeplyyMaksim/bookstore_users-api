package utils

import "net/http"

type HttpError struct {
	Status 		int			`json:"status"`
	Message 	string		`json:"message"`
	Error		string		`json:"error"`
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