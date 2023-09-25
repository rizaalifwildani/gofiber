package requests

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
)

type CreateMemberRequest struct {
	Phone          string `json:"phone" validate:"phoneNumberOrEmpty"`
	Email          string `json:"email" validate:"emailOrEmpty"`
	FirstName      string `json:"firstName" validate:"required,max=20"`
	LastName       string `json:"lastName" validate:"max=30"`
	IdentityNumber string `json:"identityNumber" validate:"required,max=15"`
	PlaceOfBirth   string `json:"placeOfBirth" validate:"required,max=30"`
	Birthdate      string `json:"birthDate" validate:"birthDate"`
	Gender         string `json:"gender" validate:"gender"`
	Nationality    string `json:"nationality" validate:"required,max=30"`
	Address        string `json:"address" validate:"required"`
	PostalCode     string `json:"postalCode" validate:"required,max=6"`
	HomePhone      string `json:"homePhone" validate:"required,max=15"`
	OfficePhone    string `json:"officePhone" validate:"required,max=15"`
	Education      string `json:"education" validate:"required"`

	Occupation CreateMemberOccupationRequest `json:"occupation" validate:"required"`
	Branches   []models.MemberBranch         `json:"branches" validate:"required"`
}

type UpdateMemberRequest struct {
	Phone          string `json:"phone" validate:"phoneNumberOrEmpty"`
	Email          string `json:"email" validate:"emailOrEmpty"`
	FirstName      string `json:"firstName" validate:"required,max=20"`
	LastName       string `json:"lastName" validate:"max=30"`
	IdentityNumber string `json:"identityNumber" validate:"required,max=15"`
	PlaceOfBirth   string `json:"placeOfBirth" validate:"required,max=30"`
	Birthdate      string `json:"birthDate" validate:"birthDate"`
	Gender         string `json:"gender" validate:"gender"`
	Nationality    string `json:"nationality" validate:"required,max=30"`
	Address        string `json:"address" validate:"required"`
	PostalCode     string `json:"postalCode" validate:"required,max=6"`
	HomePhone      string `json:"homePhone" validate:"required,max=15"`
	OfficePhone    string `json:"officePhone" validate:"required,max=15"`
	Education      string `json:"education" validate:"required"`

	Occupation UpdateMemberOccupationRequest `json:"occupation" validate:"required"`
	Branches   []models.MemberBranch         `json:"branches" validate:"required"`
}
