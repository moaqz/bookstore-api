package httpErrors

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	v "github.com/techwithmat/bookery-api/pkg/utils/validation"
)

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func NewRestError(status int, err string, causes interface{}) RestErr {
	log.Println("Causa", causes)
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func Parse(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, ErrNotFound.Error(), nil)
	case errors.Is(err, strconv.ErrSyntax):
		return NewRestError(http.StatusBadRequest, ErrBadQueryParams.Error(), nil)
	case v.IsValidationError(err):
		return NewRestError(http.StatusBadRequest, ErrValidation.Error(), v.ValidatorErrors(err))
	default:
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError.Error(), nil)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return Parse(err).Status(), Parse(err)
}
