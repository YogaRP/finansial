package database

import (
	"fmt"
	"net/url"

	"github.com/YogaRP/finansial/user-service/internal/configs"
	"github.com/YogaRP/finansial/user-service/internal/model"
	"github.com/YogaRP/finansial/user-service/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func SetupPostgres(cfg *configs.Config) (*Postgres, error) {

	// ini lebih simpel
	// connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	cfg.SqlDB.User, cfg.SqlDB.Password, cfg.SqlDB.Host, cfg.SqlDB.Port, cfg.SqlDB.DBName)

	// Yang ini aman kalau misal di user atau password ada simbol
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.SqlDB.User, cfg.SqlDB.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.SqlDB.Host, cfg.SqlDB.Port),
		Path:   cfg.SqlDB.DBName,
	}

	q := u.Query()
	q.Set("sslmode", "disable") //Kalau prod true atau require
	u.RawQuery = q.Encode()

	db, err := gorm.Open(postgres.Open(u.String()), &gorm.Config{})
	if err != nil {
		logger.Errorf("[Postgres] ConnectionPostgres - 1: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("[Postgres] ConnectionPostgres - 2: %v", err)
		return nil, err
	}

	SeedUser(db, cfg)

	sqlDB.SetMaxIdleConns(cfg.SqlDB.DBMaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.SqlDB.DBMaxOpenConns)

	return &Postgres{DB: db}, nil
}
