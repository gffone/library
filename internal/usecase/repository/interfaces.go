package repository

import (
	"context"
	"library/internal/entity"
)

type AuthorRepository interface {
	CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error)
}

type BooksRepository interface {
	CreateBook(ctx context.Context, book entity.Book) (entity.Book, error)
	GetBook(ctx context.Context, bookID string) (entity.Book, error)
}
