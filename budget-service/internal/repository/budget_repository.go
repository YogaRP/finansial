package repository

import (
	"context"
	"time"

	"github.com/YogaRP/finansial/budget-service/internal/model"
	"github.com/YogaRP/finansial/budget-service/internal/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepositoryInterface interface {
	CreateBudget(ctx context.Context, data *model.Budget) error
	GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error)
	UpdateBudget(ctx context.Context, id uuid.UUID, budget *model.Budget) error
	GetBudgetByUserID(ctx context.Context, userID uuid.UUID) (*model.Budget, error)
}

type budgetRepository struct {
	db *gorm.DB
}

// CreateBudget implements [BudgetRepositoryInterface].
func (b *budgetRepository) CreateBudget(ctx context.Context, data *model.Budget) error {
	return b.db.WithContext(ctx).Create(data).Error
}

// GetBudgetByID implements [BudgetRepositoryInterface].
func (b *budgetRepository) GetBudgetByID(ctx context.Context, id uuid.UUID) (*model.Budget, error) {
	var budget model.Budget

	if err := b.db.WithContext(ctx).Where("id = ?", id).First(&budget).Error; err != nil {
		logger.Errorf("[BudgetRepository] GetBudgetByID - 1", err)
		return nil, err
	}

	return &budget, nil
}

// GetBudgetsByUserID implements [BudgetRepositoryInterface].
func (b *budgetRepository) GetBudgetByUserID(ctx context.Context, userID uuid.UUID) (*model.Budget, error) {
	var budget model.Budget

	if err := b.db.WithContext(ctx).Where("user_id = ?", userID).First(&budget).Error; err != nil {
		logger.Errorf("[BudgetRepository] GetBudgetsByUserID - 1", err)
		return nil, err
	}

	return &budget, nil
}

// UpdateBudget implements [BudgetRepositoryInterface].
func (b *budgetRepository) UpdateBudget(ctx context.Context, id uuid.UUID, budget *model.Budget) error {
	updates := map[string]any{
		"limit":      budget.Limit,
		"period":     budget.Period,
		"updated_at": time.Now(),
	}

	result := b.db.WithContext(ctx).Model(&model.Budget{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		logger.Errorf("[BudgetRepository] UpdateBudget - 1", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logger.Errorf("[BudgetRepository] UpdateBudget - 2", gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}

func NewBudgetRepository(db *gorm.DB) BudgetRepositoryInterface {
	return &budgetRepository{
		db: db,
	}
}
