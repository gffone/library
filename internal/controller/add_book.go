package controller

import (
	"context"
	"library/generated/api/library"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *implementation) AddBook(ctx context.Context, req *library.AddBookRequest) (*library.AddBookResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	book, err := i.booksUseCase.RegisterBook(ctx, req.GetName(), req.GetAuthorId())

	if err != nil {
		return nil, i.convertErr(err)
	}

	return &library.AddBookResponse{
		Book: &library.Book{
			Id:       book.ID,
			Name:     book.Name,
			AuthorId: book.AuthorIDs,
		},
	}, nil
}
