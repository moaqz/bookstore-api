package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/internal/middleware"
)

type BookHandler struct {
	usecase domain.BookUseCase
}

func NewBookHandler(router *echo.Group, usecase domain.BookUseCase) {
	handler := &BookHandler{
		usecase: usecase,
	}

	router.GET("/books", handler.GetAllBooks)
	router.GET("/books/:id", handler.GetBookById)
	router.GET("/books/category/:category", handler.GetBookByCategory)
	router.POST("/books", handler.CreateBook, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
	router.DELETE("/books/:id", handler.DeleteBook, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
}
