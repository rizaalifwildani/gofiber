package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PermissionRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewPermissionRepository(db)
	controller := controllers.NewPermissionController(repository)
	route := router.Group("/permissions")
	route.Post("/", controller.CreatePermission)
	route.Get("/", controller.AllPermission)
	route.Get("/:id", controller.ShowPermission)
	route.Patch("/:id", controller.UpdatePermission)
}
