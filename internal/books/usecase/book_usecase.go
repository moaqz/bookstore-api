package usecase

import (
	"context"

	"github.com/techwithmat/bookery-api/internal/domain"
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

func (u *bookUsecase) GetByCategory(ctx context.Context, category string) ([]domain.Books, error) {
	books, err := u.bookRepo.GetByCategory(ctx, category)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) GetAll(ctx context.Context) ([]domain.Books, error) {
	books, err := u.bookRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookUsecase) InsertBook(ctx context.Context, book *domain.Book) error {
	validationErrors := validation.ValidateStruct(book)

	if validationErrors != nil {
		return validationErrors
	}

	err := u.bookRepo.InsertBook(ctx, book)

	return err
}

func (u *bookUsecase) DeleteBook(ctx context.Context, book *domain.Book) error {
	return nil
}

func (u *bookUsecase) UpdateBook(ctx context.Context, book *domain.Book) error {
	return nil
}
