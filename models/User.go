package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Active    int       `gorm:"column:active"`
}

func (User) TableName() string {
	return "users"
}
