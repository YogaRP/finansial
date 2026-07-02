package app

import (
	"github.com/YogaRP/finansial/budget-service/internal/configs"
	"github.com/YogaRP/finansial/budget-service/internal/database"
	"github.com/YogaRP/finansial/budget-service/internal/pkg/logger"
)

type Container struct {
}

func BuildContainer() *Container {
	config := configs.NewConfig()
	_, err := database.SetupPostgres(config)
	if err != nil {
		logger.Infof("Failed to connect to database: %v", err)
	}

	return &Container{}
}
