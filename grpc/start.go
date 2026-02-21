package grpc

import (
	"log"
	"net"

	userpb "github.com/jrmygp/contracts/proto/userpb"
	"github.com/jrmygp/user-management/services/user"
	"google.golang.org/grpc"
)

func StartGRPC(userService user.UserService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, &UserServer{
		Service: userService,
	})

	log.Println("User gRPC server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
