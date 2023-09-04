package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	RoleID    uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Add foreign key constraints
	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE:OnUpdate:CASCADE"`
}
