package Auth

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Controllers"
	Services "payment-microservice/Services"
)

var (
	AuthService    = Services.NewAuthService()
	AuthController = Controllers.NewAuthController(AuthService)
)

func SetUpAuthRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/Auth")
	{
		authGroup.POST("/Register", AuthController.Register)
		authGroup.POST("/Login", AuthController.Login)
	}

}
