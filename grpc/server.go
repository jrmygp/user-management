package grpc

import (
	"context"

	userpb "github.com/jrmygp/contracts/proto/userpb"
	"github.com/jrmygp/user-management/services/user"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	Service user.UserService
}

func (s *UserServer) GetUserByID(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userData, err := s.Service.GetUserByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		Id:       int32(userData.ID),
		Username: userData.Username,
	}, nil
}
