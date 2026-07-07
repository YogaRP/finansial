package model

import "github.com/google/uuid"

type Budget struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid()"`
	UserID    string
	Limit     int
	Period    string
	CreatedAt int64
	UpdatedAt int64
}
