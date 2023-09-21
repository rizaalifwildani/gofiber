package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewAuthRepository(db)
	controller := controllers.NewAuthController(repository)
	route := router.Group("/auth", middlewares.BasicUser())
	route.Post("/login", controller.Login)
	route.Post("/logout", controller.Logout)
	route.Put("/password", controller.ChangePassword)
}
