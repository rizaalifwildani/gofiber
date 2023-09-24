package repositories

import (
	"fmt"

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
	model := []models.Member{}
	query := r.db.Model(&model)
	for _, v := range filters {
		if v.Value != "" {
			if v.Key == "branch" {
				query.
					Joins("JOIN member_branches ON members.id = member_branches.member_id").
					Joins("JOIN branches ON member_branches.branch_id = branches.id").
					Where("LOWER(branches.name) ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
			} else {
				query.Where("LOWER("+v.Key+")"+" ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
			}
		}
	}
	query.Order(fmt.Sprintf("%v DESC", "created_at"))
	err := query.Find(&model)
	return model, err.Error
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
