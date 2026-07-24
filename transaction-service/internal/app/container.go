package app

import (
	"github.com/YogaRP/finansial/transaction-service/internal/configs"
	"github.com/YogaRP/finansial/transaction-service/internal/controller"
	"github.com/YogaRP/finansial/transaction-service/internal/database"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/kafka"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/transaction-service/internal/repository"
	"github.com/YogaRP/finansial/transaction-service/internal/service"
)

type Container struct {
	TransactionController controller.TransactionControllerInterface
	KafkaProducer         *kafka.Producer
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	db, err := database.SetupPostgres(config)
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
	}

	kafkaProducer := kafka.NewProducer(config.Kafka.Brokers)
	if err != nil {
		logger.Errorf("Failed to initialise Kafka producer: %v", err)
	}

	// Repositories
	transactionRepo := repository.NewTransactionRepository(db.DB)

	// Services
	transactionService := service.NewTransactionService(transactionRepo)

	// Controllers
	transactionController := controller.NewTransactionController(transactionService)

	return &Container{
		TransactionController: transactionController,
		KafkaProducer:         kafkaProducer,
	}
}
