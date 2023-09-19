package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RoleRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewRoleRepository(db)
	controller := controllers.NewRoleController(repository)
	route := router.Group("/roles", middlewares.SuperUser())
	route.Post("/", controller.CreateRole)
	route.Get("/", controller.AllRole)
	route.Get("/:id", controller.ShowRole)
	route.Patch("/:id", controller.UpdateRole)
}
