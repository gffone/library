package library

import (
	"context"
	"library/internal/entity"

	"github.com/google/uuid"
)

func (l *libraryImpl) RegisterBook(ctx context.Context, name string, authorIDs []string) (entity.Book, error) {
	return l.booksRepository.CreateBook(ctx, entity.Book{
		ID:        uuid.New().String(),
		Name:      name,
		AuthorIDs: authorIDs,
	})
}

func (l *libraryImpl) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	return l.booksRepository.GetBook(ctx, bookID)
}
