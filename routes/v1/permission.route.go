package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PermissionRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewPermissionRepository(db)
	controller := controllers.NewPermissionController(repository)
	route := router.Group("/permissions", middlewares.SuperUser())
	route.Post("/", controller.CreatePermission)
	route.Get("/", controller.AllPermission)
	route.Get("/:id", controller.ShowPermission)
	route.Put("/:id", controller.UpdatePermission)
}
