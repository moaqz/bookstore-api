package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := c.Get("admin").(bool)

		if !isAdmin {
			return echo.NewHTTPError(http.StatusUnauthorized, "Not Authorized")
		}

		return next(c)
	}
}
