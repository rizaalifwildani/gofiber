package repositories

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	BaseRepository
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{BaseRepository{db: db}}
}

func (r *PermissionRepository) FindAllPermission(filters []FilterType) ([]models.Permission, error) {
	data := []models.Permission{}
	err := r.Find(&data, filters, "name")
	return data, err
}

func (r *PermissionRepository) FindPermission(id string) (models.Permission, error) {
	data := models.Permission{}
	err := r.FindOne(&data, id)
	return data, err
}
