package usecase

import (
	"context"

	"github.com/techwithmat/bookery-api/internal/domain"
	"github.com/techwithmat/bookery-api/pkg/utils/pagination"
	"github.com/techwithmat/bookery-api/pkg/utils/validation"
)

// Book UseCase
type bookUsecase struct {
	bookRepo domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) domain.BookUseCase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (u *bookUsecase) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := u.bookRepo.GetByID(ctx, id)

	if err != nil {
		return &domain.Book{}, err
	}

	return book, nil
}

func (u *bookUsecase) GetByCategory(ctx context.Context, category string, p *pagination.PaginationQuery) ([]domain.Books, error) {
	books, err := u.bookRepo.GetByCategory(ctx, category, p)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) GetAll(ctx context.Context, p *pagination.PaginationQuery) ([]domain.Books, error) {
	books, err := u.bookRepo.GetAll(ctx, p)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) InsertBook(ctx context.Context, book *domain.Book) (int64, error) {
	validationErrors := validation.ValidateStruct(book)

	if validationErrors != nil {
		return 0, validationErrors
	}

	id, err := u.bookRepo.InsertBook(ctx, book)

	return id, err
}

func (u *bookUsecase) DeleteBook(ctx context.Context, id int64) error {
	_, err := u.bookRepo.GetByID(ctx, id)

	if err != nil {
		return err
	}

	err = u.bookRepo.DeleteBook(ctx, id)

	return err
}

func (u *bookUsecase) UpdateBook(ctx context.Context, book *domain.Book) error {
	return nil
}
