package main

import (
	"github.com/jrmygp/user-management/config"
	"github.com/jrmygp/user-management/controllers"
	userRepo "github.com/jrmygp/user-management/repositories/user"
	userService "github.com/jrmygp/user-management/services/user"
)

func main() {
	db := config.DatabaseConnection()

	userRepository := userRepo.NewRepository(db)
	userService := userService.NewService(userRepository)
	userController := controllers.NewUserController(userService)

	router := config.NewRouter(userController)

	router.Run(":8080")
}
