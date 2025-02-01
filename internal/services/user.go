package services

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"gorm.io/gorm"
	"my-service/internal/models"
	"my-service/internal/requests"
	"strconv"
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
		slog.Error(result.Error)
		return user.ID, result.Error
	}

	slog.WithData(slog.M{"id": user.ID}).Info("user created successfully")
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

func (s *UserService) UpdateUser(c *gin.Context) error {
	userId := c.Param("id")
	userRequest := &requests.User{}

	if err := c.ShouldBind(userRequest); err != nil {
		return err
	}

	id, _ := strconv.Atoi(userId)
	result := s.db.Model(&models.User{ID: uint(id)}).Update("name", userRequest.Name)

	if result.Error != nil {
		slog.Error(result.Error)
		return result.Error
	}

	slog.WithData(slog.M{"id": id}).Info("user updated successfully")
	return nil
}
