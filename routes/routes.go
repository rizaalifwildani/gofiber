package routes

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	v1 "bitbucket.org/rizaalifofficial/gofiber/routes/v1"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, db *gorm.DB) {
	app.Use(middlewares.Guest)
	api := app.Group("/api")
	v1Route := api.Group("/v1")

	/* === V1 === */
	v1.AuthRoute(v1Route, db)
	v1.UserRoute(v1Route, db)
	v1.RoleRoute(v1Route, db)
	v1.PermissionRoute(v1Route, db)
	v1.BranchRoute(v1Route, db)
}
