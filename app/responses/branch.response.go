package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BranchResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Code        string              `json:"code"`
	Address     string              `json:"address,omitempty"`
	Description string              `json:"description,omitempty"`
	Status      models.BranchStatus `json:"status,omitempty"`
}

func NewBranchResponse(ctx *fiber.Ctx, m models.Branch) error {
	data := BranchResponse{
		ID:          m.ID,
		Name:        m.Name,
		Code:        m.Code,
		Address:     m.Address,
		Description: m.Description,
	}
	return SuccessResponse(ctx, data)
}

func NewBranchCollections(ctx *fiber.Ctx, m []models.Branch) error {
	data := []BranchResponse{}

	for _, v := range m {
		data = append(data, BranchResponse{
			ID:          v.ID,
			Name:        v.Name,
			Code:        v.Code,
			Address:     v.Address,
			Description: v.Description,
		})
	}
	return SuccessResponse(ctx, data)
}
