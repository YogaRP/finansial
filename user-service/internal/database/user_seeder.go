package database

import (
	"github.com/YogaRP/finansial/user-service/internal/configs"
	"github.com/YogaRP/finansial/user-service/internal/model"
	"github.com/YogaRP/finansial/user-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/user-service/internal/pkg/rabbitmq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB, cfg *configs.Config, rabbitClient *rabbitmq.Client) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(cfg.App.UserPass), 10)
	if err != nil {
		logger.Errorf("[UserSeeder] SeedUser - 1: %v", err)
	}

	hashedPassword := string(bytes)

	user := model.User{
		Name:     "Yoga Rizky Putra",
		Email:    "yogarizky51@gmail.com",
		Password: hashedPassword,
	}

	result := db.FirstOrCreate(&user, model.User{Email: "yogarizky51@gmail.com"})
	if result.Error != nil {
		logger.Errorf("[UserSeeder] SeedUser - 2: %v", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		logger.Info("[UserSeeder] SeedUser - 3: Data has been seeded")
		return
	}

	if rabbitClient == nil {
		logger.Warn("[UserSeeder] SeedUser - RabbitMQ client not available, skipping budget publish")
		return
	}

	budgets := []struct {
		Amount int64
		Period string
	}{
		{
			Amount: 1000000,
			Period: "monthly",
		},
		{
			Amount: 250000,
			Period: "weekly",
		},
		{
			Amount: 35000,
			Period: "weekly",
		},
	}

	for _, b := range budgets {
		if err := rabbitClient.PublishCreateBudget(user.ID.String(), uint(b.Amount), b.Period); err != nil {
			logger.Errorf("[UserSeeder] SeedUser - 4: failed publishing budget creation: %v", err)
		} else {
			logger.Infof("[UserSeeder] SeedUser - 5: budget creation published for user %s", user.Email)
		}
	}

}
