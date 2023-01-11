package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
	"github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	ctx := c.Request().Context()
	params, err := pagination.GetPagination(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "An internal error occurred")
	}

	books, err := h.usecase.GetAll(ctx, params)

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

	params, err := pagination.GetPagination(c)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "An internal error occurred")
	}

	books, err := h.usecase.GetByCategory(ctx, category, params)

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

func (h *BookHandler) DeleteBook(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "An internal error occurred")
	}

	err = h.usecase.DeleteBook(ctx, int64(id))

	if err != nil {
		status, apiErr := httpErrors.ParseErrors(err)
		return c.JSON(status, apiErr)
	}

	return c.NoContent(http.StatusNoContent)
}
