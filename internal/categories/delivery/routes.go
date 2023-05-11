package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookstore-api/internal/domain"
	"github.com/techwithmat/bookstore-api/internal/middleware"
)

type CategoryHandler struct {
	usecase domain.CategoryUseCase
}

func NewBookHandler(router *echo.Group, usecase domain.CategoryUseCase) {
	handler := &CategoryHandler{
		usecase: usecase,
	}

	router.GET("/categories", handler.GetCategories)
	router.GET("/categories/:id", handler.GetCategoryById)
	router.POST("/categories", handler.CreateCategory, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
	router.PATCH("/categories/:id", handler.UpdateCategory, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
	router.DELETE("/categories/:id", handler.DeleteCategory, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
}
