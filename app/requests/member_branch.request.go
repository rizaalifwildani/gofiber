package requests

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/google/uuid"
)

type CreateMemberBranchRequest struct {
	BranchID uuid.UUID           `json:"branchId" validate:"required"`
	Status   models.BranchStatus `json:"status" validate:"required,branchStatus"`
}

type UpdateMemberBranchRequest struct {
	BranchID uuid.UUID           `json:"branchId" validate:"required"`
	Status   models.BranchStatus `json:"status" validate:"required,branchStatus"`
}
