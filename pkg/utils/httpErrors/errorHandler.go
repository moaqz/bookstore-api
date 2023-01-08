package httpErrors

import (
	"database/sql"
	"errors"
	"net/http"

	v "github.com/techwithmat/bookery-api/pkg/utils/validation"
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
	case v.IsValidationError(err):
		return http.StatusBadRequest, &ApiError{
			Code:    http.StatusBadRequest,
			Message: "Bad Request",
			Errors:  v.ValidatorErrors(err),
		}
	default:
		return http.StatusInternalServerError, &ApiError{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}
}
