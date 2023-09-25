package responses

import (
	"github.com/google/uuid"
)

type MemberOccupationResponse struct {
	ID         uuid.UUID `json:"id"`
	Company    string    `json:"company"`
	Department string    `json:"department"`
	Address    string    `json:"address"`
	PostalCode string    `json:"postalCode"`
	Phone      string    `json:"phone"`
	Fax        string    `json:"fax"`
	Email      string    `json:"email"`
}
