package domain

import (
	"context"

	"github.com/techwithmat/bookstore-api/pkg/utils/pagination"
)

// Book is representing the Book data struct
type Book struct {
	ID           int64   `json:"id" swaggerignore:"true"`
	Title        string  `json:"title" validate:"required"`
	Subtitle     string  `json:"subtitle" validate:"required"`
	AboutTheBook string  `json:"about_the_book" db:"about_the_book" validate:"required"`
	PageCount    int64   `json:"page_count" db:"page_count" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Image        string  `json:"image" validate:"required,url"`
	Language     string  `json:"language" validate:"required"`
	AuthorName   string  `json:"author_name" db:"author_name" validate:"required"`
	AuthorAvatar string  `json:"author_avatar" db:"author_avatar" validate:"required,url"`
	CategoryId   int64   `json:"category_id" db:"category_id" validate:"required"`
}

type Books struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Image      string `json:"image"`
	AuthorName string `json:"author_name" db:"author_name"`
}

// BookUsecase represent the Book's usecases
type BookUseCase interface {
	GetByID(ctx context.Context, id int64) (*Book, error)
	GetByCategory(ctx context.Context, category string, p *pagination.PaginationQuery) ([]Books, error)
	GetAll(ctx context.Context, p *pagination.PaginationQuery) ([]Books, error)
	InsertBook(ctx context.Context, book *Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) error
}

// BookRepository represent the Book's repository contract
type BookRepository interface {
	GetByID(ctx context.Context, id int64) (*Book, error)
	GetAll(ctx context.Context, p *pagination.PaginationQuery) ([]Books, error)
	GetByCategory(ctx context.Context, category string, p *pagination.PaginationQuery) ([]Books, error)
	InsertBook(ctx context.Context, book *Book) (int64, error)
	DeleteBook(ctx context.Context, id int64) error
}
