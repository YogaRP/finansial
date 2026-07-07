package service

import (
	"context"

	"github.com/YogaRP/finansial/budget-service/internal/model"
	"github.com/YogaRP/finansial/budget-service/internal/repository"
	"github.com/google/uuid"
)

type BudgetServiceInterface interface {
	CreateBudget(ctx context.Context, data *model.Budget) error
	GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error)
	UpdateBudget(ctx context.Context, id uuid.UUID, budget *model.Budget) error
	GetBudgetByUserID(ctx context.Context, userID string) (*model.Budget, error)
}

type budgetService struct {
	budgetRepository repository.BudgetRepositoryInterface
}

// CreateBudget implements [BudgetServiceInterface].
func (b *budgetService) CreateBudget(ctx context.Context, data *model.Budget) error {
	return b.budgetRepository.CreateBudget(ctx, data)
}

// GetBudgetByID implements [BudgetServiceInterface].
func (b *budgetService) GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error) {
	return b.budgetRepository.GetBudgetByID(ctx, id)
}

// GetBudgetByUserID implements [BudgetServiceInterface].
func (b *budgetService) GetBudgetByUserID(ctx context.Context, userID string) (*model.Budget, error) {
	return b.budgetRepository.GetBudgetByUserID(ctx, userID)
}

// UpdateBudget implements [BudgetServiceInterface].
func (b *budgetService) UpdateBudget(ctx context.Context, id uuid.UUID, budget *model.Budget) error {
	return b.budgetRepository.UpdateBudget(ctx, id, budget)
}

func NewBudgetService() BudgetServiceInterface {
	return &budgetService{}
}
