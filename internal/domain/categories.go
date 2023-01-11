package domain

import (
	"context"

	p "github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
}

type CategoryUseCase interface {
	GetAll(ctx context.Context, params *p.PaginationQuery) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	InsertCategory(ctx context.Context, c *Category) (int64, error)
	DeleteCategory(ctx context.Context, id int64) error
	UpdateCategory(ctx context.Context, c *Category) error
}

type CategoryRepository interface {
	GetAll(ctx context.Context, params *p.PaginationQuery) ([]Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	InsertCategory(ctx context.Context, c *Category) (int64, error)
	DeleteCategory(ctx context.Context, id int64) error
	UpdateCategory(ctx context.Context, c *Category) error
}
