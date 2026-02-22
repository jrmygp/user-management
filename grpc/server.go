package grpc

import (
	"context"

	userpb "github.com/jrmygp/contracts/proto/userpb"
	"github.com/jrmygp/user-management/requests"
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

func (s *UserServer) EditUser(
	ctx context.Context,
	req *userpb.EditUserRequest,
) (*userpb.EditUserResponse, error) {

	updatedUser, err := s.Service.EditUser(requests.EditUserRequest{
		UserId:       int(req.UserId),
		BalanceDelta: int(req.IncomingBalance),
	})
	if err != nil {
		return nil, err
	}

	return &userpb.EditUserResponse{
		Id:       int32(updatedUser.ID),
		Username: updatedUser.Username,
		Balance:  int32(updatedUser.Balance),
	}, nil
}
