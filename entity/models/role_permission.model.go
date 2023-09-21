package models

import (
	"time"

	"github.com/google/uuid"
)

type RolePermission struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RoleID       uuid.UUID `gorm:"type:uuid"`
	PermissionID uuid.UUID `gorm:"type:uuid"`
	Actions      string    `gorm:"type:text;default:'create, read, update, delete'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	// Add foreign key constraints
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnDelete:CASCADE:OnUpdate:CASCADE"`
}
