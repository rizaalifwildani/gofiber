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
	route := router.Group("/users")
	route.Get("/profile", middlewares.CurrentUser(), controller.ProfileUser)
	route.Patch("/profile", middlewares.CurrentUser(), controller.UpdateProfile)
	route.Post("/", middlewares.SuperUser(), controller.CreateUser)
	route.Get("/", middlewares.SuperUser(), controller.AllUser)
	route.Get("/:id", middlewares.SuperUser(), controller.ShowUser)
	route.Patch("/:id", middlewares.SuperUser(), controller.UpdateUser)
}
