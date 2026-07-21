package app

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App, container *Container) {
	api := app.Group("/api/v1")
	budgetRoutes := api.Group("/budget")

	budgetRoutes.Post("/", container.budgetController.CreateBudget)
	budgetRoutes.Get("/:id", container.budgetController.GetBudgetByID)
	budgetRoutes.Get("/user/:user_id", container.budgetController.GetBudgetByUserID)
	budgetRoutes.Put("/:id", container.budgetController.UpdateBudget)
}
