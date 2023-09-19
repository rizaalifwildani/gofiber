package requests

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
)

type CreateRoleRequest struct {
	DisplayName string              `json:"displayName" validate:"required,min=3,max=35"`
	Permissions []models.Permission `json:"permissions" validate:"required"`
}

type UpdateRoleRequest struct {
	DisplayName string              `json:"displayName" validate:"required,min=3,max=35"`
	Permissions []models.Permission `json:"permissions" validate:"required"`
}
