package main

import (
	"log"
	"os"

	"bitbucket.org/rizaalifofficial/gofiber/configs"
	"bitbucket.org/rizaalifofficial/gofiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(configs.InitFiberConfig())

	/* CORS */
	app.Use(configs.InitCors())

	/* INITIALIZED */
	configs.InitEnvirontment()
	db, err := configs.InitDB()
	if err != nil {
		log.Fatal("Failed to connect DB.\n", err)
	}
	routes.InitRouter(app, db)

	app.Listen(":" + os.Getenv("PORT"))
}
