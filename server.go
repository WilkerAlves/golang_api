package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilker/golang_api/config"
	"github.com/wilker/golang_api/controller"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDataBaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	defer config.CloseDataBaseConnection(db)
	r.Run(":8000")
}
