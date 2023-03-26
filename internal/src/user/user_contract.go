package user

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/userentity"
)

type Service interface {
	InsertNewUser(ctx context.Context, payload *userentity.UserRequest) error
	FindUser(ctx context.Context, userID string) (*userentity.UserData, error)
	// GetListUsers(ctx context.Context, sortBy, orderBy string, perPage, page int) (*[]userentity.Users, int64, error)
	// UpdateUser(ctx context.Context, payload *userentity.UserData) (*userentity.UserData, error)
	// DeleteDataUser(ctx context.Context, userID string) error
}
