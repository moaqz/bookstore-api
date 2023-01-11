package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/internal/middleware"
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
	router.DELETE("/user/unregister", handler.DeleteUser, middleware.AuthJWTMiddleware)
}

// @Summary      Register a new user
// @Description  Register a user using email, username, password and password confirmation
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body domain.SignUpRequest true "Login data: email, password and password confirmation"
// @Success      201  {object}  domain.TokenResponse
// @Router       /user/signup [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.SignUpRequest

	c.Bind(&user)

	err := h.usecase.RegisterUser(ctx, &user)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	token, err := jwtToken.GenerateJWT(user.Email, false)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, domain.TokenResponse{
		Message: "User created",
		Token:   token,
	})
}

// @Summary      Get an user account data
// @Description  Get id, username, email, first name, last name and bio from a user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user_id path string true "User ID"
// @Success      200  {object}  domain.GetUserResponse
// @Router       /user/{user_id} [get]
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

// @Summary      Login a user
// @Description  Login a user using email and password receive a JWT as a response from a successful login
// @Tags         users
// @Accept       json
// @Produce      json
// @Param request body domain.LoginRequest true "Login data: email and password"
// @Success      200  {object}  domain.TokenResponse
// @Router       /user/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.LoginRequest

	c.Bind(&user)

	u, err := h.usecase.GetUser(ctx, &user)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	token, err := jwtToken.GenerateJWT(u.Email, u.IsStaff)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, domain.TokenResponse{
		Message: "Login succesful",
		Token:   token,
	})
}

// @Summary      Delete current user
// @Description  Delete the current user account
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user_id path string true "User ID"
// @Param Authorization header string true "With the bearer started."
// @Success      204
// @Router       /user/{user_id} [delete]
func (u *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	email := c.Get("email").(string)

	var user domain.UnregisterRequest

	c.Bind(&user)
	user.Email = email

	err := u.usecase.DeleteUser(ctx, &user)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)

		return c.JSON(status, apiErr)
	}

	return c.NoContent(http.StatusNoContent)
}
