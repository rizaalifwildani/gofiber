package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PermissionResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"displayName,omitempty"`
	Actions     string    `json:"actions,omitempty"`
}

func NewPermissionResponse(ctx *fiber.Ctx, m models.Permission) error {
	data := PermissionResponse{
		ID:          m.ID,
		Name:        m.Name,
		DisplayName: m.DisplayName,
	}
	return SuccessResponse(ctx, data)
}

func NewPermissionCollections(ctx *fiber.Ctx, m []models.Permission) error {
	data := []PermissionResponse{}

	for _, v := range m {
		data = append(data, PermissionResponse{
			ID:          v.ID,
			Name:        v.Name,
			DisplayName: v.DisplayName,
		})
	}
	return SuccessResponse(ctx, data)
}
