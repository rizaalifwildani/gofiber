package repositories

import (
	"fmt"

	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type MemberRepository struct {
	BaseRepository
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{BaseRepository{db: db, Preload: []string{"Branches.Branch", "Occupation"}}}
}

func (r *MemberRepository) FindAllMember(c *fiber.Ctx, filters []FilterType) (paginate.Page, error) {
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
	pg := paginate.New()
	result := pg.With(query).Request(c.Request()).Response(&model)
	return result, err.Error
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

		// UPDATE OCCUPATION
		occupation := models.MemberOccupation{}
		findOccupation := tx.Model(&occupation).Where("member_id = ?", model.ID).First(&occupation).Error
		if findOccupation == nil {
			occupation.Company = model.Occupation.Company
			occupation.Department = model.Occupation.Department
			occupation.PostalCode = model.Occupation.PostalCode
			occupation.Phone = model.Occupation.Phone
			occupation.Fax = model.Occupation.Fax
			occupation.Email = model.Occupation.Email
			occupationErr := r.Update(&occupation, occupation.ID.String())
			if occupationErr != nil {
				return occupationErr
			}
		}

		return nil
	})
}
