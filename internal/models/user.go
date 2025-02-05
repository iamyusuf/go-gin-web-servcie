package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"type:varchar(100);not null" binding:"required"`
	Email       string         `json:"email" gorm:"type:varchar(100);not null" binding:"required,email"`
	Password    string         `json:"password" gorm:"type:varchar(60);not null" binding:"required"`
	Birthday    *time.Time     `json:"birthday"`
	ActivatedAt sql.NullTime   `json:"activated_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) HashPassword() error {
	if bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.Password = string(bytes)
		return nil
	}
}
