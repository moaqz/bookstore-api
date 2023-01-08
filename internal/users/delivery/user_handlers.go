package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
	"github.com/techwithmat/bookery-api/pkg/utils/jwtToken"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func NewUserHandler(router *echo.Group, usecase domain.UserUseCase) {
	handler := &UserHandler{
		usecase: usecase,
	}

	router.GET("/user/:id", handler.GetUserByID)
	router.POST("/user/signup", handler.RegisterUser)
	router.POST("/user/login", handler.LoginUser)
	router.DELETE("/user/:id", handler.DeleteUser)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.SignUpRequest

	c.Bind(&user)

	err := h.usecase.RegisterUser(ctx, &user)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	token, err := jwtToken.GenerateJWT(user.Email)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, domain.TokenResponse{
		Message: "User created",
		Token:   token,
	})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user, err := h.usecase.GetUserByID(ctx, int64(id))

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.LoginRequest

	c.Bind(&user)

	err := h.usecase.LoginUser(ctx, &user)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	token, err := jwtToken.GenerateJWT(user.Email)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, domain.TokenResponse{
		Message: "Login succesful",
		Token:   token,
	})
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = u.usecase.DeleteUser(ctx, int64(id))

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	return c.NoContent(http.StatusNoContent)
}
