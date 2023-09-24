package requests

import "bitbucket.org/rizaalifofficial/gofiber/entity/models"

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,max=25"`
	Phone     string `json:"phone" validate:"phoneNumberOrEmpty"`
	Email     string `json:"email" validate:"emailOrEmpty"`
	FirstName string `json:"firstName" validate:"required,max=20"`
	LastName  string `json:"lastName" validate:"max=30"`
	RegNumber string `json:"regNumber" validate:"max=50"`
	Password  string `json:"password" validate:"required"`

	Roles    []models.Role       `json:"roles" validate:"required"`
	Branches []models.UserBranch `json:"branches,omitempty"`
}

type UpdateUserRequest struct {
	Username  string `json:"username" validate:"max=25"`
	Phone     string `json:"phone" validate:"phoneNumberOrEmpty"`
	Email     string `json:"email" validate:"emailOrEmpty"`
	FirstName string `json:"firstName" validate:"required,max=20"`
	LastName  string `json:"lastName" validate:"max=30"`
	RegNumber string `json:"regNumber" validate:"max=50"`
	Password  string `json:"password" validate:"required"`

	Roles    []models.Role       `json:"roles" validate:"required"`
	Branches []models.UserBranch `json:"branches,omitempty"`
}

type UpdateProfileRequest struct {
	Phone     string `json:"phone" validate:"phoneNumberOrEmpty"`
	Email     string `json:"email" validate:"emailOrEmpty"`
	FirstName string `json:"firstName" validate:"required;max=20"`
	LastName  string `json:"lastName" validate:"max=30"`
}
