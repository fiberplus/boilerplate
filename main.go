package main

import (
	"boilerplate/config"
	"boilerplate/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// load .env
	config.InitConfig()

	app := fiber.New()
	routes.Registery(app)

	log.Fatal(app.Listen(":" + os.Getenv("HOST_PORT")))
}
