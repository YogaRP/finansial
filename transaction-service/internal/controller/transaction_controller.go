package controller

import (
	"errors"

	"github.com/YogaRP/finansial/transaction-service/internal/dto"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/logger"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/response"
	"github.com/YogaRP/finansial/transaction-service/internal/pkg/validator"
	"github.com/YogaRP/finansial/transaction-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionControllerInterface interface {
	Create(c fiber.Ctx) error
	GetByID(c fiber.Ctx) error
	GetAll(c fiber.Ctx) error
	Update(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
}

type transactionController struct {
	service service.TransactionServiceInterface
}

func (ctrl *transactionController) Create(c fiber.Ctx) error {
	var input dto.CreateTransactionInput

	if err := c.Bind().Body(&input); err != nil {
		logger.Errorf("[TransactionController] Create - 1: %v", err)
		return response.BadRequest(c, "invalid request body")
	}

	if err := validator.Validate(&input); err != nil {
		logger.Errorf("[TransactionController] Create - 2: %v", err)
		return response.BadRequest(c, err.Error())
	}

	transaction, err := ctrl.service.Create(c.Context(), &input)
	if err != nil {
		logger.Errorf("[TransactionController] Create - 2: %v", err)
		return response.BadRequest(c, "failed to create transaction")
	}

	return response.Created(c, "transaction created", transaction)
}

func (ctrl *transactionController) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Errorf("[TransactionController] GetByID - 1: %v", err)
		return response.BadRequest(c, "invalid id")
	}

	transaction, err := ctrl.service.GetByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionController] GetByID - 2: %v", err)
			return response.NotFound(c, "transaction not found")
		}
		logger.Errorf("[TransactionController] GetByID - 3: %v", err)
		return response.InternalError(c, "failed to get transaction")
	}

	return response.OK(c, "success", transaction)
}

func (ctrl *transactionController) GetAll(c fiber.Ctx) error {
	transactions, err := ctrl.service.GetAll(c.Context())
	if err != nil {
		logger.Errorf("[TransactionController] GetAll - 1: %v", err)
		return response.InternalError(c, "failed to get transactions")
	}

	return response.OK(c, "success", transactions)
}

func (ctrl *transactionController) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Errorf("[TransactionController] Update - 1: %v", err)
		return response.BadRequest(c, "invalid id")
	}

	var input dto.UpdateTransactionInput
	if err := c.Bind().Body(&input); err != nil {
		logger.Errorf("[TransactionController] Update - 2: %v", err)
		return response.BadRequest(c, "invalid request body")
	}

	if err := ctrl.service.Update(c.Context(), id, &input); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionController] Update - 3: %v", err)
			return response.NotFound(c, "transaction not found")
		}
		logger.Errorf("[TransactionController] Update - 4: %v", err)
		return response.InternalError(c, "failed to update transaction")
	}

	return response.OK(c, "transaction updated", nil)
}

func (ctrl *transactionController) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		logger.Errorf("[TransactionController] Delete - 1: %v", err)
		return response.BadRequest(c, "invalid id")
	}

	if err := ctrl.service.Delete(c.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Errorf("[TransactionController] Delete - 2: %v", err)
			return response.NotFound(c, "transaction not found")
		}
		logger.Errorf("[TransactionController] Delete - 3: %v", err)
		return response.InternalError(c, "failed to delete transaction")
	}

	return response.OK(c, "transaction deleted", nil)
}

func NewTransactionController(service service.TransactionServiceInterface) TransactionControllerInterface {
	return &transactionController{service: service}
}
