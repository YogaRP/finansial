package service

import (
	"context"
	"errors"

	"github.com/YogaRP/finansial/transaction-service/internal/dto"
	"github.com/YogaRP/finansial/transaction-service/internal/model"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/transaction-service/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionServiceInterface interface {
	Create(ctx context.Context, input *dto.CreateTransactionInput) (*model.Transaction, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error)
	GetAll(ctx context.Context) ([]model.Transaction, error)
	Update(ctx context.Context, id uuid.UUID, input *dto.UpdateTransactionInput) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type transactionService struct {
	repo repository.TransactionRepositoryInterface
}

func NewTransactionService(repo repository.TransactionRepositoryInterface) TransactionServiceInterface {
	return &transactionService{repo: repo}
}

func (s *transactionService) Create(ctx context.Context, input *dto.CreateTransactionInput) (*model.Transaction, error) {

	transaction := &model.Transaction{
		UserID:   uuid.MustParse(input.UserID.String()),
		Name:     input.Name,
		Category: model.TrxCategory(input.Category),
		Amount:   int64(input.Amount),
	}

	if err := s.repo.Create(ctx, transaction); err != nil {
		logger.Errorf("[TransactionService] Create - 1: %v", err)
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetByID(ctx context.Context, id uuid.UUID) (*model.Transaction, error) {
	transaction, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionService] GetByID - 1: %v", err)
			return nil, err
		}
		logger.Errorf("[TransactionService] GetByID - 1: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (s *transactionService) GetAll(ctx context.Context) ([]model.Transaction, error) {
	transactions, err := s.repo.FindAll(ctx)
	if err != nil {
		logger.Errorf("[TransactionService] GetAll - 1: %v", err)
		return nil, err
	}
	return transactions, nil
}

func (s *transactionService) Update(ctx context.Context, id uuid.UUID, input *dto.UpdateTransactionInput) error {
	if err := s.repo.Update(ctx, id, input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionService] Update - 1: %v", err)
			return err
		}
		logger.Errorf("[TransactionService] Update - 2: %v", err)
		return err
	}

	return nil
}

func (s *transactionService) Delete(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionService] Delete - 1: %v", err)
			return err
		}
		logger.Errorf("[TransactionService] Delete - 2: %v", err)
		return err
	}
	return nil
}
