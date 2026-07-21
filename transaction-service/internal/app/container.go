package app

import (
	"github.com/YogaRP/finansial/transaction-service/internal/configs"
	"github.com/YogaRP/finansial/transaction-service/internal/database"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/rabbitmq"
)

type Container struct {
	// budgetController controller.BudgetControllerInterface
	// BudgetService service.BudgetServiceInterface
	RabbitClient *rabbitmq.Client
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	_, err := database.SetupPostgres(config)
	if err != nil {
		logger.Infof("Failed to connect to database: %v", err)
	}

	rabbitClient, err := rabbitmq.NewClient(config)
	if err != nil {
		logger.Infof("Failed to initialise RabbitMQ client: %v", err)
	}

	// budgetRepo := repository.NewBudgetRepository(db.DB)
	// budgetService := service.NewBudgetService(budgetRepo)
	// budgetController := controller.NewBudgetController(budgetService)

	return &Container{
		// budgetController: budgetController,
		// BudgetService:    budgetService,
		RabbitClient: rabbitClient,
	}
}
