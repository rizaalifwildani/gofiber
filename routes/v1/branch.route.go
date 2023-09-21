package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BranchRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewBranchRepository(db)
	controller := controllers.NewBranchController(repository)
	route := router.Group("/branches")
	route.Post("/", middlewares.BasicUser(), controller.CreateBranch)
	route.Get("/", controller.AllBranch)
	route.Get("/:id", middlewares.BasicUser(), controller.ShowBranch)
	route.Put("/:id", middlewares.BasicUser(), controller.UpdateBranch)
}
