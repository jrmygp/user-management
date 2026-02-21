package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jrmygp/user-management/controllers"
)

func NewRouter(userController *controllers.UserController) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	user := router.Group("/user")
	user.POST("/create-user", userController.CreateUser)
	user.GET("/:id", userController.FindUserByID)

	return router
}
