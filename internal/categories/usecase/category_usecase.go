package usecase

import (
	"context"

	"github.com/techwithmat/bookery-api/internal/domain"
	p "github.com/techwithmat/bookery-api/pkg/utils/pagination"
)

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUseCase(categoryRepo domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (u *categoryUseCase) GetAll(ctx context.Context, params *p.PaginationQuery) ([]domain.Category, error) {
	categories, err := u.categoryRepo.GetAll(ctx, params)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (u *categoryUseCase) GetByID(ctx context.Context, id int64) (*domain.Category, error) {
	category, err := u.categoryRepo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *categoryUseCase) InsertCategory(ctx context.Context, c *domain.Category) (int64, error) {
	id, err := u.categoryRepo.InsertCategory(ctx, c)

	return id, err
}

func (u *categoryUseCase) DeleteCategory(ctx context.Context, id int64) error {
	err := u.categoryRepo.DeleteCategory(ctx, id)

	return err
}

func (u *categoryUseCase) UpdateCategory(ctx context.Context, c *domain.Category) error {
	err := u.categoryRepo.UpdateCategory(ctx, c)

	return err
}
