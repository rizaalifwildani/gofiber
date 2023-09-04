package models

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;size:25;not null"`
	DisplayName string    `gorm:"unique;size:35"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
