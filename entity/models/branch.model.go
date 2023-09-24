package models

import (
	"time"

	"github.com/google/uuid"
)

type BranchStatus string

const (
	Pending  BranchStatus = "pending"
	Active   BranchStatus = "active"
	Inactive BranchStatus = "inactive"
)

type Branch struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;size:35;not null"`
	Code        string    `gorm:"unique;size:15;not null"`
	Address     string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
