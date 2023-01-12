package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := c.Get("admin").(bool)

		if !isAdmin {
			return echo.NewHTTPError(http.StatusForbidden, httpErrors.NewForbiddenError(nil))
		}

		return next(c)
	}
}
