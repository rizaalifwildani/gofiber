package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GuestRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewGuestRepository(db)
	controller := controllers.NewGuestController(repository)
	route := router.Group("/guest")
	route.Post("/login", controller.Login)
}
