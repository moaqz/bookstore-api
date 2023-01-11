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

	router.GET("/user/:id", handler.GetUserByID)
	router.POST("/user/signup", handler.RegisterUser)
	router.POST("/user/login", handler.LoginUser)
	router.DELETE("/user/unregister", handler.DeleteUser, middleware.AuthJWTMiddleware)
}
