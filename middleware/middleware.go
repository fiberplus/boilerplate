package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// General Middleware

// Init - validation logger swagger recover cors
func Init(app *fiber.App) {

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())
}

// APITokenVerify - verify token for api
func APITokenVerify(c *fiber.Ctx) error {
	if os.Getenv("SECRET") == c.Get("Authorization") {
		return c.Next()
	}
	return c.Status(401).JSON(fiber.Map{
		"Error": "Unauthorized, Please Provide Valid Token",
	})
}

func isLocalENV() bool {
	return os.Getenv("DYNAMIC_ENV_VAR") == ""
}
