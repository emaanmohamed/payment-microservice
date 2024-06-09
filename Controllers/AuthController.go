package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Services "payment-microservice/Services"
	"payment-microservice/Utils"
)

type AuthController struct {
	AuthService *Services.AuthService
}

func NewAuthController(AuthService *Services.AuthService) *AuthController {
	return &AuthController{
		AuthService: AuthService,
	}
}

func (authController AuthController) Register(c *gin.Context) {
	var request Utils.AuthRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := authController.AuthService.Register(request)
	if err != nil {
		Utils.RespondWithError(c, 400, "Error registering user")
	}

	Utils.RespondWithJSON(c, 200, gin.H{"message": "User registered successfully"})
}

func (authController *AuthController) Login(c *gin.Context) {
	var loginRequest Utils.AuthRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		Utils.RespondWithError(c, 400, "Invalid request")
		return
	}

	token, err := authController.AuthService.Login(loginRequest)
	if err != nil {
		Utils.RespondWithError(c, 400, "Invalid credentials")
		return
	}
	Utils.RespondWithJSON(c, 200, gin.H{"token": token})

}
