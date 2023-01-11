package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/techwithmat/bookery-api/internal/domain"
	p "github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) domain.CategoryRepository {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) GetAll(ctx context.Context, params *p.PaginationQuery) ([]domain.Category, error) {
	var categories []domain.Category

	err := r.db.SelectContext(ctx, &categories, getCategoriesQuery, params.Size, params.Page*params.Size)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepo) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	var category domain.Category

	err := r.db.GetContext(ctx, &category, getCategoryByIdQuery, id)

	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepo) InsertCategory(ctx context.Context, c *domain.Category) (int64, error) {
	var id int64

	err := r.db.QueryRowContext(ctx, insertCategoryQuery, c.Name).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *CategoryRepo) DeleteCategory(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, deleteCategoryQuery, id)

	return err
}

func (r *CategoryRepo) UpdateCategory(ctx context.Context, c *domain.Category) error {
	_, err := r.db.NamedExecContext(ctx, updateCategoryQuery, c)

	return err
}
