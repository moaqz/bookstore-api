package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookstore-api/pkg/utils/httpErrors"
	"github.com/techwithmat/bookstore-api/pkg/utils/jwtToken"
)

func AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the JWT from the request
		AuthHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(AuthHeader, "Bearer") {
			return echo.NewHTTPError(http.StatusBadRequest, httpErrors.NewBadRequestError("Authorization header must be in the format of Bearer [token]"))
		}

		tokenString := strings.TrimPrefix(AuthHeader, "Bearer ")

		// Validate the JWT
		claims, err := jwtToken.GetClaimsFromToken(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, httpErrors.NewUnauthorizedError("Invalid token or token has expired"))
		}

		// set claims to the context
		c.Set("admin", claims["is_staff"])
		c.Set("id", int64(claims["id"].(float64)))

		// Call the next handler
		return next(c)
	}
}
