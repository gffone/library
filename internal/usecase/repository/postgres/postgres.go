package postgres

import (
	"context"
	"database/sql"
	"errors"
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

func (p *postgresRepository) CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	panic("implement me")
}

func (p *postgresRepository) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return entity.Book{}, err
	}
	defer tx.Rollback(ctx)

	const queryBook = `
INSERT INTO book (id, name)
VALUES ($1, $2)
RETURNING created_at, updated_at
`
	result := entity.Book{
		ID:        book.ID,
		Name:      book.Name,
		AuthorIDs: book.AuthorIDs,
	}

	err = tx.QueryRow(ctx, queryBook, book.ID, book.Name).Scan(&result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return entity.Book{}, err
	}

	const queryAuthor = `
INSERT INTO author_books (author_id, book_id)
VALUES ($1, $2)
`
	for _, authorID := range book.AuthorIDs {
		_, err = tx.Exec(ctx, queryAuthor, authorID, book.ID)
	}

	if err = tx.Commit(ctx); err != nil {
		return entity.Book{}, err
	}

	return result, nil
}

func (p *postgresRepository) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	const query = `
SELECT id, name, created_at, updated_at
FROM book
WHERE id = $1
`
	var book entity.Book
	err := p.db.QueryRow(ctx, query, bookID).Scan(&book.ID, &book.Name, &book.CreatedAt, &book.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return entity.Book{}, entity.ErrBookNotFound
	}

	if err != nil {
		return entity.Book{}, err
	}

	const queryAuthors = `
SELECT author_id
FROM author_books
WHERE book_id = $1
`
	rows, err := p.db.Query(ctx, queryAuthors, bookID)

	if err != nil {
		return entity.Book{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var authorID string

		if err := rows.Scan(&authorID); err != nil {
			return entity.Book{}, err
		}

		book.AuthorIDs = append(book.AuthorIDs, authorID)
	}

	return book, nil
}
