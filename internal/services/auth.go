package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"my-service/internal/requests"
	"time"
)

type AuthService struct {
	userService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (s *AuthService) Login(c *gin.Context) (string, error) {
	var login requests.Login

	if err := c.ShouldBind(&login); err != nil {
		return "", err
	}

	user, err := s.userService.GetUserByEmail(login.Email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if err != nil {
		return "", err
	}

	return generateToken(user.ID)
}

func generateToken(userId uint) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
		"iat":    time.Now().Unix(),                     // Issued at
	})

	return claims.SignedString([]byte("my_secret_key"))
}
