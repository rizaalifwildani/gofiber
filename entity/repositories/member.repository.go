package repositories

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"gorm.io/gorm"
)

type MemberRepository struct {
	BaseRepository
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{BaseRepository{db: db, Preload: []string{"Branches.Branch"}}}
}

func (r *MemberRepository) FindAllMember(filters []FilterType) ([]models.Member, error) {
	data := []models.Member{}
	err := r.Find(&data, filters, "created_at")
	return data, err
}

func (r *MemberRepository) FindMember(id string) (models.Member, error) {
	data := models.Member{}
	err := r.FindOne(&data, id)
	return data, err
}

func (r *MemberRepository) UpdateMember(model *models.Member) error {
	branches := []*models.MemberBranch{}
	for _, branch := range model.Branches {
		branches = append(branches, &models.MemberBranch{
			MemberID: model.ID,
			BranchID: branch.BranchID,
			Status:   branch.Status,
		})
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		roleErr := r.Update(&model, model.ID.String())
		if roleErr != nil {
			return roleErr
		}

		if len(branches) > 0 {
			roleErr := r.UpdateAssociation(&model, "Branches", branches)
			if roleErr != nil {
				return roleErr
			}
		}

		return nil
	})
}
