package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilker/golang_api/config"
	"github.com/wilker/golang_api/controller"
	"github.com/wilker/golang_api/repository"
	"github.com/wilker/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDataBaseConnection()
	userRepository repository.UserRepository = repository.NewUserRespository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
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
