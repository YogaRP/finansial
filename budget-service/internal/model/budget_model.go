package model

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID
	Limit     uint
	Period    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
