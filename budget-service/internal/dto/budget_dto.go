package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateBudgetRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Limit  uint      `json:"limit" validate:"required"`
	Period string    `json:"period" validate:"required"`
}

func (c *CreateBudgetRequest) Validate() error {
	v := validator.New()
	return v.Struct(c)
}
