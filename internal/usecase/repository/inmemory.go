package repository

import (
	"context"
	"library/internal/entity"
	"sync"
)

var _ AuthorRepository = (*inMemoryImpl)(nil)
var _ BooksRepository = (*inMemoryImpl)(nil)

type inMemoryImpl struct {
	authorsMx *sync.RWMutex
	authors   map[string]*entity.Author

	booksMx *sync.RWMutex
	books   map[string]*entity.Book
}

func NewInMemoryImpl() *inMemoryImpl {
	return &inMemoryImpl{
		authorsMx: &sync.RWMutex{},
		authors:   make(map[string]*entity.Author),
		booksMx:   &sync.RWMutex{},
		books:     make(map[string]*entity.Book),
	}
}

func (i *inMemoryImpl) CreateAuthor(ctx context.Context, author entity.Author) (entity.Author, error) {
	i.authorsMx.Lock()
	defer i.authorsMx.Unlock()

	if _, ok := i.authors[author.ID]; ok {
		return entity.Author{}, entity.ErrAuthorAlreadyExists
	}
	i.authors[author.ID] = &author

	return author, nil
}

func (i *inMemoryImpl) CreateBook(ctx context.Context, book entity.Book) (entity.Book, error) {
	i.booksMx.Lock()
	defer i.booksMx.Unlock()
	if _, ok := i.books[book.ID]; ok {
		return entity.Book{}, entity.ErrBookAlreadyExists
	}
	i.books[book.ID] = &book
	return book, nil
}

func (i *inMemoryImpl) GetBook(ctx context.Context, bookID string) (entity.Book, error) {
	i.booksMx.Lock()
	defer i.booksMx.Unlock()

	if book, ok := i.books[bookID]; !ok {
		return entity.Book{}, entity.ErrBookNotFound
	} else {
		return *book, nil
	}
}
