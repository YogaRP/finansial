package app

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App, container *Container) {
	api := app.Group("/api/v1")
	transactionRoutes := api.Group("/transaction")

	transactionRoutes.Post("/", container.TransactionController.Create)
	transactionRoutes.Get("/:id", container.TransactionController.GetByID)
	transactionRoutes.Get("/", container.TransactionController.GetAll)
	transactionRoutes.Put("/:id", container.TransactionController.Update)
	transactionRoutes.Delete("/:id", container.TransactionController.Delete)
}
