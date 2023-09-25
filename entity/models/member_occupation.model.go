package models

import (
	"time"

	"github.com/google/uuid"
)

type MemberOccupation struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	MemberID   uuid.UUID `gorm:"type:uuid"`
	Company    string    `gorm:"size:50;not null"`
	Department string    `gorm:"size:50;not null"`
	Address    string    `gorm:"type:text;not null"`
	PostalCode string    `gorm:"size:6;not null"`
	Phone      string    `gorm:"size:15"`
	Fax        string    `gorm:"size:20"`
	Email      string    `gorm:"size:50"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
