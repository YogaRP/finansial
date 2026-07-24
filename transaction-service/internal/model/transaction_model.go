package model

import (
	"slices"
	"time"

	"github.com/google/uuid"
)

type TrxCategory string

const (
	Lainnya  TrxCategory = "lainnya"
	Pokok    TrxCategory = "pokok"
	Sekunder TrxCategory = "sekunder"
)

var ValidTrxCategories = []TrxCategory{Lainnya, Pokok, Sekunder}

func (c TrxCategory) IsValid() bool {
	return slices.Contains(ValidTrxCategories, c)
}

type Transaction struct {
	ID        uuid.UUID   `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID   `gorm:"type:uuid;not null"`
	Name      string      `gorm:"type:varchar(255);not null"`
	Amount    int64       `gorm:"type:bigint;not null"`
	Category  TrxCategory `gorm:"type:varchar(50);default:'lainnya'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Transaction) TableName() string {
	return "transactions"
}
