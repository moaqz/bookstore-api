package pagination

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	defaultSize = 10
	defaultPage = 0
)

type PaginationQuery struct {
	Size int `json:"size,omitempty"`
	Page int `json:"page,omitempty"`
}

// Set page
func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Page = defaultPage
		return nil
	}

	n, err := strconv.Atoi(pageQuery)

	if err != nil {
		return err
	}

	q.Page = n

	return nil
}

func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}

	q.Size = n

	return nil
}

func GetPagination(c echo.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}

	if err := q.SetPage(c.QueryParam("page")); err != nil {
		return nil, err
	}

	if err := q.SetSize(c.QueryParam("size")); err != nil {
		return nil, err
	}

	return q, nil
}
