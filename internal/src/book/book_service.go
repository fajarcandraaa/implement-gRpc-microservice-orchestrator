package book

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers/errorcodehandling"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/bookentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
)

type service struct {
	book *servicecontract.GrpcContract
	err  *errorcodehandling.CodeError
}

func NewService(book *servicecontract.GrpcContract) *service {
	return &service{
		book: book,
	}
}

func (s *service) AddNewBook(ctx context.Context, payload *bookentity.BookRequest) error {
	book := &pb.CreateBookRequest{
		Title:  payload.Title,
		Author: payload.Author,
	}

	_, err := s.book.Book.ServiceInsertNewBook(ctx, book)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindBook(ctx context.Context, bookID string) (*bookentity.Book, error) {
	payload := &pb.FindBookByIdRequest{
		Id: bookID,
	}

	book, err := s.book.Book.ServiceFindBookById(ctx, payload)
	if err != nil {
		return nil, err
	}

	result := &bookentity.Book{
		ID:        book.Data.Id,
		Title:     book.Data.Title,
		Author:    book.Data.Author,
		CreatedAt: book.Data.CreatedAt,
		UpdatedAt: book.Data.UpdatedAt,
	}

	return result, nil
}
