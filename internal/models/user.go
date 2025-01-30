package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" binding:"required"`
	Email       *string        `json:"email" binding:"required,email"`
	Birthday    *time.Time     `json:"birthday"`
	ActivatedAt sql.NullTime   `json:"activated_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
