package v1

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/controllers"
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/entity/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func MemberRoute(router fiber.Router, db *gorm.DB) {
	repository := repositories.NewMemberRepository(db)
	controller := controllers.NewMemberController(repository)
	route := router.Group("/members")
	route.Post("/", middlewares.BasicUser(), controller.CreateMember)
	route.Get("/", middlewares.BasicUser(), controller.AllMember)
	route.Get("/:id", middlewares.BasicUser(), controller.ShowMember)
	route.Put("/:id", middlewares.BasicUser(), controller.UpdateMember)
}
