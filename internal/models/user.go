package models

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

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

func (u *User) GetAge() (int, int, int) {
	if u.Birthday == nil {
		return 0, 0, 0
	}

	now := time.Now()
	years := now.Year() - u.Birthday.Year()
	months := now.Month() - u.Birthday.Month()
	days := now.Day() - u.Birthday.Day()

	if days < 0 {
		months--
		days += time.Date(now.Year(), now.Month(), 0, 0, 0, 0, 0, time.UTC).Day()
	}

	if months < 0 {
		years--
		months += 12
	}

	return years, int(months), days
}

func (u *User) CheckPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
