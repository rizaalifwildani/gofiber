package repositories

import (
	"errors"
	"fmt"

	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository{
			db:      db,
			Preload: []string{"Roles.Permissions.Permission", "Branches.Branch"},
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

func (r *UserRepository) FindAllUser(filters []FilterType) ([]models.User, error) {
	model := []models.User{}
	query := r.db.Model(&model)
	query.
		Joins("JOIN user_roles ON users.id = user_roles.user_id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("roles.name != ?", "root")
	for _, v := range filters {
		if v.Value != "" {
			query.Where("LOWER("+v.Key+")"+" ILIKE ?", fmt.Sprintf("%%%s%%", v.Value))
		}
	}
	query.Order(fmt.Sprintf("%v DESC", "first_name"))
	err := query.Find(&model)
	return model, err.Error
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
			RoleID: role.ID,
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
