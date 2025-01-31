package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"gorm.io/gorm"
	"my-service/internal/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) CreateUser(c *gin.Context) (uint, error) {
	user := &models.User{}

	if err := c.ShouldBind(user); err != nil {
		return 0, err
	}

	result := s.db.Create(user)

	if result.Error != nil {
		return user.ID, result.Error
	}

	return user.ID, nil
}

func (s *UserService) GetUser(c *gin.Context) (*models.User, error) {
	userId := c.Param("id")
	var user models.User
	result := s.db.First(&user, userId)

	if result.Error != nil {
		slog.Error(result.Error)
		return nil, result.Error
	}

	return &user, nil
}
