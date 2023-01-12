package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/httpErrors"
	"github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

// @Summary Get a list of categories
// @Description Get a list of categories. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.
// @Tags categories
// @Accept json
// @Produce json
// @Param  page query int false "Page number"
// @Param  size query int false "Size number"
// @Success 200 {array} domain.Category
// @Failure 400 {object} httpErrors.RestError
// @Failure 404 {object} httpErrors.EmptyBody
// @Failure 500	{object}	httpErrors.RestError
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c echo.Context) error {
	ctx := c.Request().Context()
	params, err := pagination.GetPagination(c)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	categories, err := h.usecase.GetAll(ctx, params)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	if len(categories) == 0 {
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	return c.JSON(http.StatusOK, categories)
}

// @Summary Get a category by its id
// @Description Get a specific category object. Id parameter must be an integer.
// @Tags categories
// @Accept json
// @Produce json
// @Param category_id path int true "Category ID"
// @Success 200 {array} domain.Category
// @Failure 400 {object} httpErrors.RestError
// @Failure 404 {object} httpErrors.RestError
// @Failure 500	{object}	httpErrors.RestError
// @Router /categories/{category_id} [get]
func (h *CategoryHandler) GetCategoryById(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	category, err := h.usecase.GetByID(ctx, int64(id))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, category)
}

// @Summary Create a new category
// @Description Create a new category object.
// @Tags categories
// @Accept json
// @Produce json
// @Param	request	body domain.Category	true	"New Category data"
// @Param	Authorization	header	string	true "With the bearer started. Only staff members"
// @Success		201				{object}	domain.Category
// @Failure		400				{object}	httpErrors.RestError
// @Failure		500				{object}	httpErrors.RestError
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	ctx := c.Request().Context()
	var category domain.Category
	c.Bind(&category)

	id, err := h.usecase.InsertCategory(ctx, &category)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	category.ID = id

	return c.JSON(http.StatusCreated, category)
}

// @Summary      Update a category
// @Description  Update a category object by its Id.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param request body domain.Category true "Updated Category data"
// @Param Authorization header string true "With the bearer started. Only staff members"
// @Param category_id path int true "Category ID"
// @Success      200  {object}  domain.Category
// @Failure      400  {object}  httpErrors.RestError
// @Failure      404  {object}  httpErrors.RestError
// @Failure      500  {object}  httpErrors.RestError
// @Router       /categories/{category_id} [patch]
func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	var category domain.Category
	c.Bind(&category)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	category.ID = int64(id)
	err = h.usecase.UpdateCategory(ctx, &category)

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.JSON(http.StatusOK, category)
}

// @Summary		Delete a category
// @Description	Delete a category object.
// @Tags			categories
// @Accept			json
// @Produce		json
// @Param			category_id			path	int		true	"Category ID"
// @Param			Authorization	header	string	true	"With the bearer started. Only staff members"
// @Success		204
// @Failure		404	{object}	httpErrors.RestError
// @Failure   409 {object}  httpErrors.RestError
// @Failure		500	{object}	httpErrors.RestError
// @Router			/categories/{category_id} [delete]
func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	err = h.usecase.DeleteCategory(ctx, int64(id))

	if err != nil {
		return c.JSON(httpErrors.ErrorResponse(err))
	}

	return c.NoContent(http.StatusNoContent)
}
