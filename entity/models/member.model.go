package models

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID             uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Phone          string    `gorm:"unique;size:15"`
	Email          string    `gorm:"unique;size:50"`
	FirstName      string    `gorm:"size:20"`
	LastName       string    `gorm:"size:30"`
	IdentityNumber string    `gorm:"unique;size:15;not null"`
	PlaceOfBirth   string    `gorm:"size:30;not null"`
	Birthdate      time.Time `gorm:"type:date;not null"`
	Gender         string    `gorm:"type:ENUM('male', 'female');not null"`
	Nationality    string    `gorm:"size:30;not null"`
	Address        string    `gorm:"type:text;not null"`
	PostalCode     string    `gorm:"size:6;not null"`
	HomePhone      string    `gorm:"size:15;not null"`
	OfficePhone    string    `gorm:"size:15;not null"`
	Education      string    `gorm:"type:text;not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	Occupation MemberOccupation `gorm:"foreignKey:MemberID"`
	Branches   []MemberBranch   `gorm:"foreignKey:MemberID"`
}
