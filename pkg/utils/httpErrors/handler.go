package httpErrors

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	v "github.com/techwithmat/bookstore-api/pkg/utils/validation"
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
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  ErrUnauthorized.Error(),
		ErrCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  ErrForbidden.Error(),
		ErrCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ErrBadRequest.Error(),
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
	case strings.Contains(err.Error(), "foreign key constraint"):
		return NewRestError(http.StatusConflict, ErrConflict.Error(), err.Error())
	case strings.Contains(err.Error(), "key value violates unique constraint"):
		return NewRestError(http.StatusConflict, ErrEmailAlreadyExists.Error(), err.Error())
	default:
		return NewRestError(http.StatusInternalServerError, ErrInternalServerError.Error(), nil)
	}
}

func ErrorResponse(err error) (int, interface{}) {
	return Parse(err).Status(), Parse(err)
}
