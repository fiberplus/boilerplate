package routes

import (
	"boilerplate/database"
	"boilerplate/pkg/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Registery - Routes And Repo Registery
func Registery(app *fiber.App) {

	dbMysql := database.GetDB()
	fmt.Println("DB connection established")

	userRepo := user.NewRepo(dbMysql)
	userService := user.NewService(userRepo)

	// TODO: controller auto gen
	app.Get("/", func(c *fiber.Ctx) error {
		data, _ := userService.Find(2)
		fmt.Println("Test Response", data.ID)
		return c.SendString("Test Response")
	})


	app.Use(func(c *fiber.Ctx) error {
		return c.Status(401).SendString("Page not found " + c.BaseURL())
	})

}
