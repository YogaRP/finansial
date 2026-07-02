package app

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App, container *Container) {
	api := app.Group("/api/v1")
	users := api.Group("/users")
	users.Get("/email/:email", container.UserController.GetUserByEmail)
	users.Get("/:id", container.UserController.GetUserByID)
}
