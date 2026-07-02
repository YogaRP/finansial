package model

type Budget struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Limit     int
	Period    string
	CreatedAt int64
	UpdatedAt int64
}
