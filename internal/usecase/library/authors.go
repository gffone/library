package library

import (
	"context"
	"library/internal/entity"

	"github.com/google/uuid"
)

func (l *libraryImpl) RegisterAuthor(ctx context.Context, authorName string) (entity.Author, error) {
	author, err := l.authorRepository.CreateAuthor(ctx, entity.Author{
		ID:   uuid.New().String(),
		Name: authorName,
	})

	if err != nil {
		return entity.Author{}, err
	}

	return author, nil
}
