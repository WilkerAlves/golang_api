package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wilker/golang_api/config"
	"github.com/wilker/golang_api/controller"
	"github.com/wilker/golang_api/middleware"
	"github.com/wilker/golang_api/repository"
	"github.com/wilker/golang_api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDataBaseConnection()
	userRepository repository.UserRepository = repository.NewUserRespository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRouter := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRouter.GET("/profile", userController.Profile)
		userRouter.PUT("/profile", userController.Update)
	}

	defer config.CloseDataBaseConnection(db)
	r.Run(":8000")
}
