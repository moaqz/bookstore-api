package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/internal/middleware"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
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
	router.DELETE("/books/:id", handler.GetBookByCategory, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
	router.PATCH("/books/:id", handler.GetBookByCategory, middleware.AuthJWTMiddleware, middleware.AdminMiddleware)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	ctx := c.Request().Context()

	books, err := h.usecase.GetAll(ctx)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)
		return c.JSON(status, apiErr)
	}

	if len(books) == 0 {
		return c.NoContent(http.StatusNoContent)
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
		status, apiErr := httpErrors.ParseErrors(err)
		return c.JSON(status, apiErr)
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) GetBookByCategory(c echo.Context) error {
	ctx := c.Request().Context()
	category := c.Param("category")

	books, err := h.usecase.GetByCategory(ctx, category)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)
		return c.JSON(status, apiErr)
	}

	if len(books) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	ctx := c.Request().Context()
	var book domain.Book

	c.Bind(&book)

	err := h.usecase.InsertBook(ctx, &book)

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)
		return c.JSON(status, apiErr)
	}

	return c.JSON(http.StatusCreated, "Product created")
}
