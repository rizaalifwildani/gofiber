package repositories

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	BaseRepository
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{BaseRepository{db: db, Preload: []string{"Permissions.Permission"}}}
}

func (r *RoleRepository) FindAllRole(filters []FilterType) ([]models.Role, error) {
	data := []models.Role{}
	err := r.Find(&data, filters, "name")
	return data, err
}

func (r *RoleRepository) FindRole(id string) (models.Role, error) {
	data := models.Role{}
	err := r.FindOne(&data, id)
	return data, err
}

func (r *RoleRepository) UpdateRole(model *models.Role) error {
	permissions := []*models.RolePermission{}
	for _, permission := range model.Permissions {
		permissions = append(permissions, &models.RolePermission{
			RoleID:       model.ID,
			PermissionID: permission.PermissionID,
		})
	}

	err := r.UpdateAssociation(*model, "Permissions", permissions)
	return err
}
