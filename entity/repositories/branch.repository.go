package repositories

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"gorm.io/gorm"
)

type BranchRepository struct {
	BaseRepository
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{BaseRepository{db: db}}
}

func (r *BranchRepository) FindAllBranch(filters []FilterType) ([]models.Branch, error) {
	data := []models.Branch{}
	err := r.Find(&data, filters, "name")
	return data, err
}

func (r *BranchRepository) FindBranch(id string) (models.Branch, error) {
	data := models.Branch{}
	err := r.FindOne(&data, id)
	return data, err
}
