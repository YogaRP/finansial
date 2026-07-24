package repository

import (
	"context"
	"time"

	"github.com/YogaRP/finansial/transaction-service/internal/dto"
	"github.com/YogaRP/finansial/transaction-service/internal/model"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	Create(ctx context.Context, transaction *model.Transaction) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	FindAll(ctx context.Context) ([]model.Transaction, error)
	Update(ctx context.Context, id uuid.UUID, transaction *dto.UpdateTransactionInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepositoryInterface {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, transaction *model.Transaction) error {
	return r.db.WithContext(ctx).Create(transaction).Error
}

func (r *transactionRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	var transaction model.Transaction

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&transaction).Error; err != nil {
		logger.Errorf("[transactionRepository] FindByID - 1", err)
		return nil, err
	}

	return &transaction, nil
}

func (r *transactionRepository) FindAll(ctx context.Context) ([]model.Transaction, error) {
	var transaction []model.Transaction

	if err := r.db.WithContext(ctx).Find(&transaction).Error; err != nil {
		logger.Errorf("[transactionRepository] FindAll - 1", err)
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) Update(ctx context.Context, id uuid.UUID, transaction *dto.UpdateTransactionInput) error {
	updates := map[string]any{
		"updated_at": time.Now(),
	}

	if transaction.Name != nil {
		updates["name"] = *transaction.Name
	}
	if transaction.Category != nil {
		updates["category"] = *transaction.Category
	}
	if transaction.Amount != nil {
		updates["amount"] = *transaction.Amount
	}

	result := r.db.WithContext(ctx).
		Model(&model.Transaction{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		logger.Errorf("[TransactionRepository] UpdateTransaction - 1", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logger.Errorf("[TransactionRepository] UpdateTransaction - 2", gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *transactionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Transaction{})

	if result.Error != nil {
		logger.Errorf("[TransactionRepository] DeleteTransaction - 1: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logger.Errorf("[TransactionRepository] DeleteTransaction - 2: %v", gorm.ErrRecordNotFound)
		return gorm.ErrRecordNotFound
	}
	return nil
}
