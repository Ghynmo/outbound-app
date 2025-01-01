package user

import (
	"time"
)

// Database Migration
type User struct {
	ID         string `gorm:"primaryKey"`
	Fullname   string
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Role       string `gorm:"default:user"`
	ProfileImg string
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  time.Time `gorm:"default:null"`
}
