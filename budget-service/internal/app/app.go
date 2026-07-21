package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/YogaRP/finansial/budget-service/internal/configs"
	logger "github.com/YogaRP/finansial/budget-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/budget-service/internal/pkg/rabbitmq"
	"github.com/YogaRP/finansial/budget-service/internal/pkg/response"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func RunServer() {
	cfg := configs.NewConfig()

	logger.Init(cfg.App.AppEnv)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			logger.Errorf("Error: %v", err)
			// return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			return response.InternalError(c, "Internal Server Error")
		},
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format: "[${time}] $ip ${status} - ${latency}  ${method}  ${path}\n",
	}))

	// app.Use(middlewareGateway.GatewayAuth())
	container := BuildContainer()
	if container.RabbitClient == nil {
		log.Fatalf("RabbitMQ client is not initialized")
	}

	if err := rabbitmq.StartBudgetConsumer(container.RabbitClient, container.BudgetService); err != nil {
		log.Fatalf("Error starting RabbitMQ consumer: %v", err)
	}

	defer func() {
		if container.RabbitClient != nil {
			_ = container.RabbitClient.Close()
		}
	}()
	SetupRoutes(app, container)

	port := cfg.App.AppPort
	if port == "" {
		port = os.Getenv("APP_PORT")
		if port == "" {
			log.Fatalf("Server port not specified")
		}
	}
	logger.Infof("Starting server on port: %s", port)

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Error during shutdown: %v", err)
	}
	logger.Info("Server shutdown complete")
}
