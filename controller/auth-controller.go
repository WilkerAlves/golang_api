package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// here enter the service you need
	// this is where you put your service
}

func (a *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"hello login",
	})
}

func (a *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message":"hello register",
	})
}

//NewAuthController creates a new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

