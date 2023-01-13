package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
	v "github.com/techwithmat/bookery-api/pkg/utils/validation"
)

//	@Summary		Register a new user
//	@Description	Register a user using email, username, password and password confirmation
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.SignUpRequest	true	"Login data: email, password and password confirmation"
//	@Success		201		{object}	domain.TokenResponse
//	@Failure		400		{object}	httpErrors.RestError
//	@Failure		409		{object}	httpErrors.RestError
//	@Failure		500		{object}	httpErrors.RestError
//	@Router			/users/signup [post]
func (h *UserHandler) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.SignUpRequest

	c.Bind(&user)

	// Validate body request
	err := v.ValidateStruct(user)
	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	registerUser, err := h.usecase.RegisterUser(ctx, &user)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, registerUser)
}

//	@Summary		Get an user account data
//	@Description	Get id, email and role from a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int	true	"User ID"
//	@Success		200		{object}	domain.GetUserResponse
//	@Failure		400		{object}	httpErrors.RestError
//	@Failure		404		{object}	httpErrors.RestError
//	@Failure		500		{object}	httpErrors.RestError
//	@Router			/users/{user_id} [get]
func (h *UserHandler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	user, err := h.usecase.GetUserByID(ctx, int64(id))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, user)
}

//	@Summary		Login a user
//	@Description	Login a user using email and password receive a JWT as a response from a successful login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		domain.LoginRequest	true	"Login data: email and password"
//	@Success		200		{object}	domain.TokenResponse
//	@Failure		400		{object}	httpErrors.RestError
//	@Failure		404		{object}	httpErrors.RestError
//	@Failure		500		{object}	httpErrors.RestError
//	@Router			/users/login [post]
func (h *UserHandler) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user domain.LoginRequest
	c.Bind(&user)

	// Validate body request
	if err := v.ValidateStruct(user); err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	loggedUser, err := h.usecase.LoginUser(ctx, &user)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, loggedUser)
}

//	@Summary		Delete current user
//	@Description	Delete the current user account
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user_id			path	int		true	"User ID"
//	@Param			Authorization	header	string	true	"With the bearer started."
//	@Success		204
//	@Failure		401	{object}	httpErrors.RestError
//	@Failure		400	{object}	httpErrors.RestError
//	@Failure		404	{object}	httpErrors.RestError
//	@Failure		500	{object}	httpErrors.RestError
//	@Router			/users/{user_id} [delete]
func (u *UserHandler) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("id").(int64)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	// If claims from JWT token and params are differente raise an error
	if userId != int64(id) {
		return c.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError("You don't have permissions to delete this account."))
	}

	err = u.usecase.DeleteUser(ctx, userId)
	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.NoContent(http.StatusNoContent)
}
