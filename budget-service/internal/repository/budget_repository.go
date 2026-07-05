package repository

import (
	"github.com/YogaRP/finansial/budget-service/internal/model"
	"gorm.io/gorm"
)

type BudgetRepositoryInterface interface {
	CreateBudget(budget *model.Budget) error
	GetBudgetByID(id string) (*model.Budget, error)
	UpdateBudget(budget *model.Budget) error
	GetBudgetsByUserID(userID string) (model.Budget, error)
}

type budgetRepository struct {
	db *gorm.DB
}

// CreateBudget implements [BudgetRepositoryInterface].
func (b *budgetRepository) CreateBudget(budget *model.Budget) error {
	panic("unimplemented")
}

// GetBudgetByID implements [BudgetRepositoryInterface].
func (b *budgetRepository) GetBudgetByID(id string) (*model.Budget, error) {
	panic("unimplemented")
}

// GetBudgetsByUserID implements [BudgetRepositoryInterface].
func (b *budgetRepository) GetBudgetsByUserID(userID string) (model.Budget, error) {
	panic("unimplemented")
}

// UpdateBudget implements [BudgetRepositoryInterface].
func (b *budgetRepository) UpdateBudget(budget *model.Budget) error {
	panic("unimplemented")
}

func newBudgetRepository(db *gorm.DB) BudgetRepositoryInterface {
	return &budgetRepository{
		db: db,
	}
}
