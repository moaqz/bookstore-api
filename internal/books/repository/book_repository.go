package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/techwithmat/bookery-api/internal/domain"
)

type BooksRepo struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) domain.BookRepository {
	return &BooksRepo{
		db: db,
	}
}

func (r *BooksRepo) GetAll(ctx context.Context) ([]domain.Books, error) {
	var books []domain.Books

	err := r.db.SelectContext(ctx, &books, getBooks)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BooksRepo) GetByCategory(ctx context.Context, category string) ([]domain.Books, error) {
	var books []domain.Books

	err := r.db.SelectContext(ctx, &books, getBookByCategory, category)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BooksRepo) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	var book domain.Book

	err := r.db.GetContext(ctx, &book, getBookById, id)

	if err != nil {
		return nil, err
	}

	return &book, nil
}
