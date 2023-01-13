package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/internal/middleware"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func NewUserHandler(router *echo.Group, usecase domain.UserUseCase) {
	handler := &UserHandler{
		usecase: usecase,
	}

	router.GET("/users/:id", handler.GetUserByID)
	router.POST("/users/signup", handler.RegisterUser)
	router.POST("/users/login", handler.LoginUser)
	router.DELETE("/users/:id", handler.DeleteUser, middleware.AuthJWTMiddleware)
}
