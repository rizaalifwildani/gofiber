package models

import (
	"time"

	"github.com/google/uuid"
)

type MemberBranch struct {
	ID        uuid.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	MemberID  uuid.UUID    `gorm:"type:uuid"`
	BranchID  uuid.UUID    `gorm:"type:uuid"`
	Status    BranchStatus `gorm:"type:ENUM('pending', 'active', 'inactive');default:'pending'"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`

	// Add foreign key constraints
	Branch Branch `gorm:"foreignKey:BranchID;references:ID;constraint:OnDelete:CASCADE:OnUpdate:CASCADE"`
}
