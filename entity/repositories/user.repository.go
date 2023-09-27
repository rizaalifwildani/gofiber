package repositories

import (
	"errors"
	"fmt"

	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository{
			db:      db,
			Preload: []string{"Roles.Role.Permissions.Permission", "Branches.Branch"},
		},
	}
}

func (r *UserRepository) CreateUser(model *models.User, authModel *models.UserAuth) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model).Error
		if err != nil {
			return err
		}

		// === USER AUTH === //
		password, err := utils.GeneratePassword(model.ID.String() + authModel.Password)
		if err != nil {
			passwordError := errors.New("invalid password")
			return passwordError
		}
		userAuth := models.UserAuth{
			UserID:   model.ID,
			Password: password,
		}
		userAuthErr := tx.Create(&userAuth).Error
		if userAuthErr != nil {
			return userAuthErr
		}

		return nil
	})
}

func (r *UserRepository) FindAllUser(c *fiber.Ctx, filters []FilterType) (paginate.Page, error) {
	model := []models.User{}
	query := r.db.Model(&model)
	query.
		Preload("Roles.Role").
		Preload("Branches.Branch").
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Joins("JOIN user_branches ON user_branches.user_id = users.id").
		Joins("JOIN branches ON branches.id = user_branches.branch_id")

	for _, v := range filters {
		if v.Value != "" {
			if v.Key == "branch" {
				query.Where("LOWER(branches.name) ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
			} else if v.Key == "role" {
				query.Where("LOWER(roles.display_name) ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
			} else {
				query.Where("LOWER("+v.Key+")"+" ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
			}
		}
	}

	query.Where("roles.name != ?", "root")
	query.Order(fmt.Sprintf("%v DESC", "first_name"))
	err := query.Find(&model)
	pg := paginate.New()
	result := pg.With(query).Request(c.Request()).Response(&model)
	return result, err.Error
}

func (r *UserRepository) FindUser(id string) (models.User, error) {
	data := models.User{}
	err := r.FindOne(&data, id)
	return data, err
}

func (r *UserRepository) UpdateUser(model *models.User, authModel *models.UserAuth) error {
	roles := []*models.UserRole{}
	for _, role := range model.Roles {
		roles = append(roles, &models.UserRole{
			UserID: model.ID,
			RoleID: role.RoleID,
		})
	}

	branches := []*models.UserBranch{}
	for _, branch := range model.Branches {
		branches = append(branches, &models.UserBranch{
			UserID:   model.ID,
			BranchID: branch.BranchID,
			Status:   branch.Status,
		})
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		// === USER DETAIL === //
		userErr := r.Update(*model, model.ID.String())
		if userErr != nil {
			return userErr
		}

		// === USER ROLE === //
		if len(roles) > 0 {
			roleErr := r.UpdateAssociation(&model, "Roles", roles)
			if roleErr != nil {
				return roleErr
			}
		}

		// === USER BRANCH === //
		if len(branches) > 0 {
			branchErr := r.UpdateAssociation(&model, "Branches", branches)
			if branchErr != nil {
				return branchErr
			}
		}

		// === USER AUTH === //
		if authModel != nil {
			password, err := utils.GeneratePassword(model.ID.String() + authModel.Password)
			if err != nil {
				passwordError := errors.New("invalid password")
				return passwordError
			}
			userAuth := models.UserAuth{
				ID:        authModel.ID,
				UserID:    authModel.UserID,
				Password:  password,
				Token:     authModel.Token,
				ExpiredAt: authModel.ExpiredAt,
			}
			userAuthErr := tx.Model(&userAuth).Where("id = ?", userAuth.ID).Updates(userAuth).Error
			if userAuthErr != nil {
				return userAuthErr
			}
		}

		return nil
	})
}
