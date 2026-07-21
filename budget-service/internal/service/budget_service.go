package service

import (
	"context"

	"github.com/YogaRP/finansial/budget-service/internal/dto"
	"github.com/YogaRP/finansial/budget-service/internal/model"
	"github.com/YogaRP/finansial/budget-service/internal/repository"
	"github.com/google/uuid"
)

type BudgetServiceInterface interface {
	CreateBudget(ctx context.Context, data dto.CreateBudgetRequest) error
	GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error)
	UpdateBudget(ctx context.Context, id uuid.UUID, budget dto.CreateBudgetRequest) error
	GetBudgetByUserID(ctx context.Context, userID uuid.UUID) (*model.Budget, error)
}

type budgetService struct {
	budgetRepository repository.BudgetRepositoryInterface
}

// CreateBudget implements [BudgetServiceInterface].
func (b *budgetService) CreateBudget(ctx context.Context, data dto.CreateBudgetRequest) error {
	budget := &model.Budget{
		UserID: data.UserID,
		Limit:  data.Limit,
		Period: data.Period,
	}
	return b.budgetRepository.CreateBudget(ctx, budget)
}

// GetBudgetByID implements [BudgetServiceInterface].
func (b *budgetService) GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error) {
	return b.budgetRepository.GetBudgetByID(ctx, id)
}

// GetBudgetByUserID implements [BudgetServiceInterface].
func (b *budgetService) GetBudgetByUserID(ctx context.Context, userID uuid.UUID) (*model.Budget, error) {
	return b.budgetRepository.GetBudgetByUserID(ctx, userID)
}

// UpdateBudget implements [BudgetServiceInterface].
func (b *budgetService) UpdateBudget(ctx context.Context, id uuid.UUID, data dto.CreateBudgetRequest) error {
	budget := &model.Budget{
		UserID: data.UserID,
		Limit:  data.Limit,
		Period: data.Period,
	}
	return b.budgetRepository.UpdateBudget(ctx, id, budget)
}

func NewBudgetService(budgetRepository repository.BudgetRepositoryInterface) BudgetServiceInterface {
	return &budgetService{
		budgetRepository,
	}
}
