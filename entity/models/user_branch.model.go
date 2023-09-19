package models

import (
	"time"

	"github.com/google/uuid"
)

type UserBranchStatus string

const (
	Pending  UserBranchStatus = "pending"
	Active   UserBranchStatus = "active"
	Inactive UserBranchStatus = "inactive"
)

type UserBranch struct {
	ID        uuid.UUID        `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID        `gorm:"type:uuid"`
	BranchID  uuid.UUID        `gorm:"type:uuid"`
	Status    UserBranchStatus `gorm:"type:ENUM('pending', 'active', 'inactive');default:'pending'"`
	CreatedAt time.Time        `gorm:"autoCreateTime"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime"`

	// Add foreign key constraints
	Branch Branch `gorm:"foreignKey:BranchID;references:ID;constraint:OnDelete:CASCADE:OnUpdate:CASCADE"`
}
