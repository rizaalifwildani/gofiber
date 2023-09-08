package routes

import (
	"bitbucket.org/rizaalifofficial/gofiber/app/middlewares"
	"bitbucket.org/rizaalifofficial/gofiber/configs"
	v1 "bitbucket.org/rizaalifofficial/gofiber/routes/v1"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	v1Route := api.Group("/v1")

	/* === OPEN ROUTE === */
	v1.GuestRoute(v1Route, db)

	/* === AUTHENTICATED ROUTE === */
	app.Use(configs.InitJWT())
	/* === MIDDLEWARE === */
	app.Use(middlewares.InitMiddleware)
	v1.AuthRoute(v1Route, db)
	v1.UserRoute(v1Route, db)
	v1.RoleRoute(v1Route, db)
	v1.PermissionRoute(v1Route, db)
}
