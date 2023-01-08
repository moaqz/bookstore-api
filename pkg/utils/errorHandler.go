package utils

import (
	"database/sql"
	"errors"
	"net/http"
)

type ApiError struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Errors  *map[string]string `json:"errors,omitempty"`
}

func ParseErrors(err error) (int, *ApiError) {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNotFound, &ApiError{
			Code:    http.StatusNotFound,
			Message: "Not Found",
		}
	case IsValidationError(err):
		return http.StatusBadRequest, &ApiError{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Errors:  ValidatorErrors(err),
		}
	default:
		return http.StatusInternalServerError, &ApiError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
}
