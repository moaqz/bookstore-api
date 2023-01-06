package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/booki-api/internal/domain"
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
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	ctx := c.Request().Context()

	books, err := h.usecase.GetAll(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookById(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	book, err := h.usecase.GetByID(ctx, int64(id))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBookByCategory(c echo.Context) error {
	ctx := c.Request().Context()
	category := c.Param("category")

	books, err := h.usecase.GetByCategory(ctx, category)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, books)
}