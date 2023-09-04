package repositories

import (
	"errors"
	"log"

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
			Preload: []string{"Roles.Role.Permissions.Permission"},
		},
	}
}

func (r *UserRepository) CreateUser(model *models.User, authModel *models.UserAuth) error {
	log.Println(model)
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
	data := []models.User{}
	err := r.Find(&data, filters, "first_name")
	return data, err
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

	return r.db.Transaction(func(tx *gorm.DB) error {
		err := r.UpdateAssociation(*model, "Roles", roles)

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

		return nil
	})
}
