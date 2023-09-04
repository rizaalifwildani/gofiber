package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	DisplayName string               `json:"displayName"`
	Permissions []PermissionResponse `json:"permissions"`
}

func NewRoleResponse(ctx *fiber.Ctx, m models.Role) error {
	permissions := []PermissionResponse{}
	for _, permission := range m.Permissions {
		permissions = append(permissions, PermissionResponse{
			ID:          permission.PermissionID,
			Name:        permission.Permission.Name,
			DisplayName: permission.Permission.DisplayName,
		})
	}
	data := RoleResponse{
		ID:          m.ID,
		Name:        m.Name,
		DisplayName: m.DisplayName,
		Permissions: permissions,
	}
	return SuccessResponse(ctx, data)
}

func NewRoleCollections(ctx *fiber.Ctx, m []models.Role) error {
	data := []RoleResponse{}

	for _, v := range m {
		permissions := []PermissionResponse{}
		for _, permission := range v.Permissions {
			permissions = append(permissions, PermissionResponse{
				ID:          permission.PermissionID,
				Name:        permission.Permission.Name,
				DisplayName: permission.Permission.DisplayName,
			})
		}
		data = append(data, RoleResponse{
			ID:          v.ID,
			Name:        v.Name,
			DisplayName: v.DisplayName,
			Permissions: permissions,
		})
	}
	return SuccessResponse(ctx, data)
}
