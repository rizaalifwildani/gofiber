package routes

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewAuthRepository(db)
	controller := controllers.NewAuthController(repository)
	route := router.Group("/auth")
	route.Post("/login", controller.Login)
}
