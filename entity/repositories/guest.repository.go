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

type GuestRepository struct {
	BaseRepository
}

func NewGuestRepository(db *gorm.DB) *GuestRepository {
	return &GuestRepository{BaseRepository{db: db, Preload: []string{"User"}}}
}

func (r *GuestRepository) Login(username string, password string) (*models.UserAuth, error) {
	auth := models.UserAuth{}
	user := models.User{}

	findUser := r.db.Where("username", username).Preload("Roles.Role").First(&user).Error
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
			ID:          role.RoleID,
			Name:        role.Role.Name,
			DisplayName: role.Role.DisplayName,
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
