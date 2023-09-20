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
	route := router.Group("/auth")
	route.Post("/login", controller.Login)
	route.Post("/logout", middlewares.BasicUser(), controller.Logout)
	route.Patch("/password", middlewares.CurrentUser(), controller.ChangePassword)
}
