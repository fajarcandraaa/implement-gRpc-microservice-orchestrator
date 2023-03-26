package user

import (
	"context"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers/errorcodehandling"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers/unique"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/userentity"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
	"github.com/fajarcandraaa/implement-gRpc-microservice-protobank/grpc/pb"
)

type service struct {
	user *servicecontract.GrpcContract
	err  *errorcodehandling.CodeError
}

func NewService(user *servicecontract.GrpcContract) *service {
	return &service{
		user: user,
	}
}

// InsertNewUser represents algorithm to register new user
func (s *service) InsertNewUser(ctx context.Context, payload *userentity.UserRequest) error {

	err := userentity.UserRequestValidate(payload)
	if err != nil {
		return err
	}

	user := &pb.CreateUserRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Username: payload.Username,
		Password: payload.Password,
	}

	_, err = s.user.User.ServiceRegisterUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// FindUser represents algorithm to find user by id
func (s *service) FindUser(ctx context.Context, userID string) (*userentity.UserData, error) {
	if err := unique.ValidateUUID(userID); err != nil {
		return nil, entity.ErrUserNotExist
	}

	payload := &pb.FindUserByIdRequest{
		Id: userID,
	}

	user, err := s.user.User.ServiceFindUserById(ctx, payload)
	if err != nil {
		return nil, err
	}

	usrData := &userentity.UserData{
		ID:        userID,
		Name:      user.Data.Name,
		Email:     user.Data.Email,
		Username:  user.Data.Username,
		CreatedAt: user.Data.CreatedAt,
		UpdatedAt: user.Data.UpdatedAt,
	}

	return usrData, nil
}
