package repositories

import (
	"errors"
	"net/http"
	"os"
	"time"

	"bitbucket.org/rizaalifofficial/gofiber/app/responses"
	"bitbucket.org/rizaalifofficial/gofiber/configs"
	"bitbucket.org/rizaalifofficial/gofiber/entity/models"
	"bitbucket.org/rizaalifofficial/gofiber/static"
	"bitbucket.org/rizaalifofficial/gofiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	BaseRepository
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{BaseRepository{db: db, Preload: []string{"User"}}}
}

func (r *AuthRepository) Login(username string, password string) (*models.UserAuth, error) {
	auth := models.UserAuth{}
	user := models.User{}

	findUser := r.db.Where("username", username).Preload("Roles").First(&user).Error
	if findUser != nil {
		return nil, findUser
	}

	findAuth := r.db.Where("user_id", user.ID).First(&auth).Error
	if findAuth != nil {
		return nil, findAuth
	}

	password, passwordError := utils.GeneratePassword(user.ID.String() + password)
	if passwordError != nil {
		return nil, passwordError
	}

	if password != auth.Password {
		return nil, errors.New("invalid password")
	}

	exp := time.Now().Add(time.Hour * 24)
	roles := []responses.RoleResponse{}
	for _, role := range user.Roles {
		roles = append(roles, responses.RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			DisplayName: role.DisplayName,
		})
	}
	claims := configs.JWTConfig{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
		ID:    user.ID,
		Roles: roles,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, tokenError := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if tokenError != nil {
		tokenError = errors.New(
			http.StatusText(
				fiber.StatusUnauthorized,
			),
		)
		return nil, tokenError
	}
	auth.Token = signedToken
	auth.ExpiredAt = &exp
	authErr := r.Update(&auth, auth.ID.String())
	if authErr != nil {
		return nil, authErr
	}

	// Save token to Redis 1 day
	err := utils.SetRedis(static.REDIS_TOKEN, signedToken, time.Hour*24)
	if err != nil {
		return nil, err
	}

	return &auth, nil
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
