package requests

import "bitbucket.org/rizaalifofficial/gofiber/entity/models"

type CreateUserRequest struct {
	Username  string `json:"username" validate:"required"`
	Phone     string `json:"phone" validate:"numeric,min=9,max=15"`
	Email     string `json:"email" validate:"email,min=5,max=50"`
	FirstName string `json:"firstName" validate:"alpha,max=20"`
	LastName  string `json:"lastName" validate:"alpha,max=30"`
	Password  string `json:"password" validate:"required"`

	Roles []models.UserRole `json:"roles" validate:"required"`
}

type UpdateUserRequest struct {
	Phone     string `json:"phone" validate:"numeric,min=9,max=15"`
	Email     string `json:"email" validate:"email,min=5,max=50"`
	FirstName string `json:"firstName" validate:"alpha,max=20"`
	LastName  string `json:"lastName" validate:"alpha,max=30"`
	Password  string `json:"password" validate:"required"`

	Roles []models.UserRole `json:"roles" validate:"required"`
}
