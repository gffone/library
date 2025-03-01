package library

import (
	"context"
	"library/internal/entity"
	"library/internal/usecase/repository"

	"go.uber.org/zap"
)

type AuthorUseCase interface {
	RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error)
}

type BooksUseCase interface {
	RegisterBook(ctx context.Context, name string, authorIDs []string) (entity.Book, error)
	GetBook(ctx context.Context, bookID string) (entity.Book, error)
}

var _ AuthorUseCase = (*libraryImpl)(nil)
var _ BooksUseCase = (*libraryImpl)(nil)

type libraryImpl struct {
	logger           *zap.Logger
	authorRepository repository.AuthorRepository
	booksRepository  repository.BooksRepository
}

func New(
	logger *zap.Logger,
	authorRepository repository.AuthorRepository,
	booksRepository repository.BooksRepository,
) *libraryImpl {
	return &libraryImpl{
		logger:           logger,
		authorRepository: authorRepository,
		booksRepository:  booksRepository,
	}
}
