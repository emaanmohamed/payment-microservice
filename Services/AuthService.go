package Services

import (
	"golang.org/x/crypto/bcrypt"
	"payment-microservice/Initializers"
	"payment-microservice/Utils"
	"payment-microservice/models"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// just to make the idea chine, i make simple register to be able to login, but take in consideration that it
//should be validation on unique email to avoid duplication

func (authService *AuthService) Register(request Utils.AuthRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := models.User{
		Username: request.Username,
		Password: string(hash),
		Role:     request.Role,
	}
	if err := Initializers.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (authService *AuthService) Login(request Utils.AuthRequest) (string, error) {
	var user models.User
	if err := Initializers.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", err
	}

	token, err := Utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", err
	}
	return token,nil

}
