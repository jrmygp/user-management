package main

import (
	"github.com/jrmygp/user-management/config"
	"github.com/jrmygp/user-management/controllers"
	usergrpc "github.com/jrmygp/user-management/grpc"
	userRepo "github.com/jrmygp/user-management/repositories/user"
	userService "github.com/jrmygp/user-management/services/user"
)

func main() {
	db := config.DatabaseConnection()

	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)
	userController := controllers.NewUserController(userService)

	router := config.NewRouter(userController)

	go usergrpc.StartGRPC(userService)

	router.Run(":8080")
}
