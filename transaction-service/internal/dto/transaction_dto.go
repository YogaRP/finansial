package dto

import (
	"github.com/YogaRP/finansial/transaction-service/internal/model"
	"github.com/google/uuid"
)

type CreateTransactionInput struct {
	UserID   uuid.UUID         `json:"user_id" validate:"required"`
	Name     string            `json:"name" validate:"required,min=3,max=100"`
	Category model.TrxCategory `json:"category" validate:"required,trxcategory"`
	Amount   int               `json:"amount" validate:"required"`
}

type UpdateTransactionInput struct {
	Name     *string `json:"name" validate:"omitempty,min=3,max=100"`
	Category *string `json:"category" validate:"omitempty,trxcategory"`
	Amount   *int    `json:"amount" validate:"omitempty"`
}
