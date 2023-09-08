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

	/* ENVIRONTMENT */
	configs.InitEnvirontment()

	/* REDIS */
	configs.InitRedis()

	/* CORS */
	app.Use(configs.InitCors())

	/* DATABASE */
	db, err := configs.InitDB()
	if err != nil {
		log.Fatal("failed to connect db.\n", err)
	}

	/* ROUTER */
	routes.InitRouter(app, db)

	app.Listen(":" + os.Getenv("PORT"))
}
