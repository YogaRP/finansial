package database

import (
	"github.com/YogaRP/finansial/user-service/internal/configs"
	"github.com/YogaRP/finansial/user-service/internal/model"
	"github.com/YogaRP/finansial/user-service/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB, cfg *configs.Config) {

	var userExist model.User
	if err := db.Where("email = ?", "yogarizky51@gmail.com").First(&userExist).Error; err == nil {
		logger.Info("[UserSeeder] SeedUser - 1: Data has been seeded")
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(cfg.App.UserPass), 10)
	if err != nil {
		logger.Errorf("[UserSeeder] SeedUser - 2: %v", err)
	}

	hashedPassword := string(bytes)

	user := model.User{
		Name:     "Yoga Rizky Putra",
		Email:    "yogarizky51@gmail.com",
		Password: hashedPassword,
	}

	if err := db.FirstOrCreate(&user, model.User{Email: "yogarizky51@gmail.com"}).Error; err != nil {
		logger.Errorf("[UserSeeder] SeedUser - 3: %v", err)
	} else {
		logger.Infof("[UserSeeder] SeedUser - 4: %v", "User created successfully")
	}
}
