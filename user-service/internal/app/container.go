package app

import (
	"github.com/YogaRP/finansial/user-service/internal/configs"
	"github.com/YogaRP/finansial/user-service/internal/controller"
	"github.com/YogaRP/finansial/user-service/internal/database"
	"github.com/YogaRP/finansial/user-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/user-service/internal/pkg/rabbitmq"
	"github.com/YogaRP/finansial/user-service/internal/repository"
	"github.com/YogaRP/finansial/user-service/internal/service"
)

type Container struct {
	UserController controller.UserControllerInterface
	RabbitClient   *rabbitmq.Client
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.SetupPostgres(config)
	if err != nil {
		logger.Infof("Failed to connect to database: %v", err)
	}

	rabbitClient, err := rabbitmq.NewClient(config)
	if err != nil {
		logger.Infof("Failed to initialise RabbitMQ client: %v", err)
	}

	database.SeedUser(db.DB, config, rabbitClient)

	userRepo := repository.NewUserRepository(db.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &Container{
		UserController: userController,
		RabbitClient:   rabbitClient,
	}
}
