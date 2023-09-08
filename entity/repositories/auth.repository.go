package repositories

import (
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/static"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRepository struct {
	BaseRepository
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{BaseRepository{db: db, Preload: []string{"User"}}}
}

func (r *AuthRepository) Logout(ctx *fiber.Ctx) error {
	_, jwt, claims, ok := utils.ClaimsJWT(ctx)

	if ok && jwt.Valid {
		// FIND USER AUTH
		auth := models.UserAuth{}
		err := r.db.Where("user_id", claims.ID).First(&auth).Error
		if err != nil {
			return err
		}

		auth.Token = ""
		auth.ExpiredAt = nil
		updateAuthErr := r.Update(&auth, auth.ID.String())
		if updateAuthErr != nil {
			return updateAuthErr
		}

		redis := utils.SetRedis(static.REDIS_TOKEN, "", 0)
		if redis != nil {
			return redis
		}
	}

	return nil
}
