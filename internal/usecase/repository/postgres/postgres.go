package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"library/internal/entity"
	"library/internal/usecase/repository"
)

var _ repository.AuthorRepository = (*postgresRepository)(nil)
var _ repository.BooksRepository = (*postgresRepository)(nil)

type postgresRepository struct {
	db *pgxpool.Pool
}

func NewPostgresRepository(db *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		db: db,
	}
}

func (i *postgresRepository) CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	panic("implement me")
}

func (i *postgresRepository) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	panic("implement me")
}

func (i *postgresRepository) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	panic("implement me")
}
