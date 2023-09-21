package responses

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID        `json:"id"`
	Username  string           `json:"username"`
	Phone     string           `json:"phone"`
	Email     string           `json:"email"`
	FirstName string           `json:"firstName"`
	LastName  string           `json:"lastName"`
	RegNumber string           `json:"regNumber"`
	Roles     []RoleResponse   `json:"roles,omitempty"`
	Branches  []BranchResponse `json:"branches,omitempty"`
}

func NewUserResponse(ctx *fiber.Ctx, m models.User) error {
	roles := []RoleResponse{}
	for _, role := range m.Roles {
		permissions := []PermissionResponse{}
		for _, permission := range role.Permissions {
			permissions = append(permissions, PermissionResponse{
				ID:          permission.PermissionID,
				Name:        permission.Permission.Name,
				DisplayName: permission.Permission.DisplayName,
				Actions:     permission.Actions,
			})
		}
		roles = append(roles, RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			DisplayName: role.DisplayName,
			Permissions: permissions,
		})
	}
	branches := []BranchResponse{}
	for _, branch := range m.Branches {
		branches = append(branches, BranchResponse{
			ID:          branch.BranchID,
			Name:        branch.Branch.Name,
			Code:        branch.Branch.Code,
			Address:     branch.Branch.Address,
			Description: branch.Branch.Description,
			Status:      branch.Status,
		})
	}
	data := UserResponse{
		ID:        m.ID,
		Username:  m.Username,
		Phone:     m.Phone,
		Email:     m.Email,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		RegNumber: m.RegNumber,
		Roles:     roles,
		Branches:  branches,
	}
	return SuccessResponse(ctx, data)
}

func NewUserCollections(ctx *fiber.Ctx, m []models.User) error {
	data := []UserResponse{}

	for _, v := range m {
		data = append(data, UserResponse{
			ID:        v.ID,
			Username:  v.Username,
			Phone:     v.Phone,
			Email:     v.Email,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			RegNumber: v.RegNumber,
		})
	}

	return SuccessResponse(ctx, data)
}
