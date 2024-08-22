package model

import "net/http"

var (
	ErrUnhealthy      = NewError(http.StatusInternalServerError, "something went wrong")
	ErrRefreshExpired = NewError(http.StatusUnauthorized, "refresh")
	ErrUnauthorized   = NewError(http.StatusUnauthorized, "user unauthorized")
	ErrInvalidBody    = NewError(http.StatusBadRequest, "request invalid body")
	ErrNotFoundUser   = NewError(http.StatusNotFound, "The email address entered is not recognized. Please double-check your entry and try again.")
	ErrInvalidRole    = NewError(http.StatusBadRequest, "user invalid role")
	ErrUsenameExist   = NewError(http.StatusBadRequest, "username exist")
)

const (
	NotFound = "not found"
)

type Error interface {
	error
	Status() int
}

// StatusError represents HTTP error.
type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Message
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

//nolint:ireturn
func NewError(code int, message string) Error {
	return StatusError{
		Code:    code,
		Message: message,
	}
}
