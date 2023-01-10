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

	err := r.db.SelectContext(ctx, &books, getBooksQuery)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BooksRepo) GetByCategory(ctx context.Context, category string) ([]domain.Books, error) {
	var books []domain.Books

	err := r.db.SelectContext(ctx, &books, getBookByCategoryQuery, category)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BooksRepo) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	var book domain.Book

	err := r.db.GetContext(ctx, &book, getBookByIdQuery, id)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BooksRepo) InsertBook(ctx context.Context, book *domain.Book) error {
	_, err := r.db.ExecContext(ctx, InsertBookQuery,
		book.Title,
		book.Subtitle,
		book.AboutTheBook,
		book.PageCount,
		book.Price,
		book.Image,
		book.Language,
		book.AuthorName,
		book.AuthorAvatar,
		book.CategoryId,
	)

	return err
}

func (r *BooksRepo) DeleteBook(ctx context.Context, id int64) error {
	return nil
}

func (r *BooksRepo) UpdateBook(ctx context.Context, id int64) error {
	return nil
}
