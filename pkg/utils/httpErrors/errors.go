package httpErrors

import "errors"

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrNotFound            = errors.New("Not Found")
	ErrBadQueryParams      = errors.New("Invalid query params")
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrValidation          = errors.New("Validation Failed")
	ErrConflict            = errors.New("A Conflict ocurred")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

type RestError struct {
	ErrStatus int         `json:"status"`
	ErrError  string      `json:"error"`
	ErrCauses interface{} `json:"causes,omitempty"`
}

type EmptyBody struct{}
