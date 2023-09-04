package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAuth struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Password  string
	Token     string
	ExpiredAt *time.Time `gorm:"default:NULL"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`

	// Add foreign key constraints
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE:OnUpdate:CASCADE"`
}
