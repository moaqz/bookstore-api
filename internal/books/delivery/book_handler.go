package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
	"github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

// @Summary		Get a list of books
// @Description	Get a list of books. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags			books
// @Accept			json
// @Produce		json
// @Param			page	query		int	false	"Page number"
// @Param			size	query		int	false	"Size number"
// @Success		200		{array}		domain.Books
// @Failure		400		{object}	httpErrors.RestError
// @Failure		404		{object}	httpErrors.EmptyBody
// @Failure		500		{object}	httpErrors.RestError
// @Router			/books [get]
func (h *BookHandler) GetAllBooks(c echo.Context) error {
	ctx := c.Request().Context()
	params, err := pagination.GetPagination(c)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	books, err := h.usecase.GetAll(ctx, params)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	if len(books) == 0 {
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	return c.JSON(http.StatusOK, books)
}

// @Summary		Get a book by its id
// @Description	Get a specific book object. Id parameter must be an integer.
// @Tags			books
// @Accept			json
// @Produce		json
// @Param			book_id	path		int	true	"Book ID"
// @Success		200		{object}	domain.Book
// @Failure		400		{object}	httpErrors.RestError
// @Failure		404		{object}	httpErrors.RestError
// @Failure		500		{object}	httpErrors.RestError
// @Router			/books/{book_id} [get]
func (h *BookHandler) GetBookById(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	book, err := h.usecase.GetByID(ctx, int64(id))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, book)
}

// @Summary		Get a list of books by category
// @Description	Get a list of books by category. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags			books
// @Accept			json
// @Produce		json
// @Param			category_name	path		string	true	"Category name"
// @Param			page			query		int		false	"Page number"
// @Param			size			query		int		false	"Size number"
// @Success		200				{object}	domain.Books
// @Failure		400				{object}	httpErrors.RestError
// @Failure		404				{object}	httpErrors.EmptyBody
// @Failure		500				{object}	httpErrors.RestError
// @Router			/books/category/{category_name} [get]
func (h *BookHandler) GetBookByCategory(c echo.Context) error {
	ctx := c.Request().Context()
	category := c.Param("category")
	params, err := pagination.GetPagination(c)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	books, err := h.usecase.GetByCategory(ctx, category, params)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	if len(books) == 0 {
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	return c.JSON(http.StatusOK, books)
}

// @Summary		Create a new book
// @Description	Create a book object.
// @Tags			books
// @Accept			json
// @Produce		json
// @Param			request			body		domain.Book	true	"New Book data"
// @Param			Authorization	header		string		true	"With the bearer started. Only staff members"
// @Success		201				{object}	domain.Book
// @Failure		400				{object}	httpErrors.RestError
// @Failure		500				{object}	httpErrors.RestError
// @Router			/books [post]
func (h *BookHandler) CreateBook(c echo.Context) error {
	ctx := c.Request().Context()
	var book domain.Book
	c.Bind(&book)

	id, err := h.usecase.InsertBook(ctx, &book)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	book.ID = id

	return c.JSON(http.StatusCreated, book)
}

// @Summary		Delete a book
// @Description	Delete a book object.
// @Tags			books
// @Accept			json
// @Produce		json
// @Param			book_id			path	int		true	"Book ID"
// @Param			Authorization	header	string	true	"With the bearer started. Only staff members"
// @Success		204
// @Failure		404	{object}	httpErrors.RestError
// @Failure		500	{object}	httpErrors.RestError
// @Router			/books/{book_id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	err = h.usecase.DeleteBook(ctx, int64(id))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.NoContent(http.StatusNoContent)
}
