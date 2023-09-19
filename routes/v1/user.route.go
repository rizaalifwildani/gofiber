package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewUserRepository(db)
	controller := controllers.NewUserController(repository)
	route := router.Group("/users", middlewares.SuperUser())
	route.Post("/", controller.CreateUser)
	route.Get("/", controller.AllUser)
	route.Get("/:id", controller.ShowUser)
	route.Patch("/:id", controller.UpdateUser, middlewares.CurrentUser())
}
