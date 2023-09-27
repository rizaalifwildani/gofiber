package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `gorm:"unique;size:25;not null"`
	Phone     string    `gorm:"size:15"`
	Email     string    `gorm:"size:50"`
	FirstName string    `gorm:"size:20"`
	LastName  string    `gorm:"size:30"`
	RegNumber string    `gorm:"unique;size:50;default:null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Roles    []UserRole   `gorm:"foreignKey:UserID"`
	Branches []UserBranch `gorm:"foreignKey:UserID"`
}
