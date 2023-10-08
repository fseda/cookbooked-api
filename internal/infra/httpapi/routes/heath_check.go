package routes

import "github.com/gofiber/fiber/v2"

func loadHealthCheck(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})
}
