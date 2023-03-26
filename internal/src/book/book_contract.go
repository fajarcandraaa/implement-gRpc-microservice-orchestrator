package book

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/bookentity"
)

type Service interface {
	AddNewBook(ctx context.Context, payload *bookentity.BookRequest) error
	FindBook(ctx context.Context, bookID string) (*bookentity.Book, error)
}
