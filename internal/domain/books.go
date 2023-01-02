package domain

import (
	"context"
)

// Book is representing the Book data struct
type Book struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Subtitle     string  `json:"subtitle"`
	AboutTheBook string  `json:"about_the_book" db:"about_the_book"`
	PageCount    int     `json:"page_count" db:"page_count"`
	Price        float64 `json:"price"`
	Image        string  `json:"image"`
	Language     string  `json:"language"`
	AuthorName   string  `json:"author_name" db:"author_name"`
	AuthorAvatar string  `json:"author_avatar" db:"author_avatar"`
	CategoryId   int     `json:"category_id" db:"category_id"`
}

type Books struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Image      string `json:"image"`
	AuthorName string `json:"author_name" db:"author_name"`
}

// BookUsecase represent the Book's usecases
type BookUseCase interface {
	GetByID(ctx context.Context, id int64) (*Book, error)
	GetByCategory(ctx context.Context, category string) ([]Books, error)
	GetAll(ctx context.Context) ([]Books, error)
}

// BookRepository represent the Book's repository contract
type BookRepository interface {
	GetByID(ctx context.Context, id int64) (*Book, error)
	GetAll(ctx context.Context) ([]Books, error)
	GetByCategory(ctx context.Context, category string) ([]Books, error)
}
