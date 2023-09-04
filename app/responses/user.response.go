package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	Username  string         `json:"username"`
	Phone     string         `json:"phone"`
	Email     string         `json:"email"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Roles     []RoleResponse `json:"roles"`
}

func NewUserResponse(ctx *fiber.Ctx, m models.User) error {
	roles := []RoleResponse{}
	for _, role := range m.Roles {
		permissions := []PermissionResponse{}
		for _, permission := range role.Role.Permissions {
			permissions = append(permissions, PermissionResponse{
				ID:          permission.PermissionID,
				Name:        permission.Permission.Name,
				DisplayName: permission.Permission.DisplayName,
			})
		}
		roles = append(roles, RoleResponse{
			ID:          role.RoleID,
			Name:        role.Role.Name,
			DisplayName: role.Role.DisplayName,
			Permissions: permissions,
		})
	}
	data := UserResponse{
		ID:        m.ID,
		Username:  m.Username,
		Phone:     m.Phone,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Roles:     roles,
	}
	return SuccessResponse(ctx, data)
}

func NewUserCollections(ctx *fiber.Ctx, m []models.User) error {
	data := []UserResponse{}

	for _, v := range m {
		roles := []RoleResponse{}
		for _, role := range v.Roles {
			permissions := []PermissionResponse{}
			for _, permission := range role.Role.Permissions {
				permissions = append(permissions, PermissionResponse{
					ID:          permission.PermissionID,
					Name:        permission.Permission.Name,
					DisplayName: permission.Permission.DisplayName,
				})
			}
			roles = append(roles, RoleResponse{
				ID:          role.RoleID,
				Name:        role.Role.Name,
				DisplayName: role.Role.DisplayName,
				Permissions: permissions,
			})
		}
		data = append(data, UserResponse{
			ID:        v.ID,
			Username:  v.Username,
			Phone:     v.Phone,
			Email:     v.Email,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Roles:     roles,
		})
	}
	return SuccessResponse(ctx, data)
}
