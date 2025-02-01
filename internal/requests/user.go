package requests

import (
	"my-service/internal/models"
	"time"
)

type User struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name" binding:"required"`
	Birthday *time.Time `json:"birthday"`
}

func (u User) ToModel() *models.User {
	return &models.User{
		ID:       u.ID,
		Name:     u.Name,
		Birthday: u.Birthday,
	}
}
