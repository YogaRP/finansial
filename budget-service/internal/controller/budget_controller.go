package controller

import (
	"github.com/YogaRP/finansial/budget-service/internal/dto"
	"github.com/YogaRP/finansial/budget-service/internal/pkg/response"
	"github.com/YogaRP/finansial/budget-service/internal/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/google/uuid"
)

type BudgetControllerInterface interface {
	CreateBudget(ctx fiber.Ctx) error
	GetBudgetByID(ctx fiber.Ctx) error
	UpdateBudget(ctx fiber.Ctx) error
	GetBudgetByUserID(ctx fiber.Ctx) error
}

type budgetController struct {
	budgetService service.BudgetServiceInterface
}

// CreateBudget implements [BudgetControllerInterface].
func (b *budgetController) CreateBudget(ctx fiber.Ctx) error {
	var req dto.CreateBudgetRequest
	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[BudgetController] Create Budget - 1: %v", err)
		return response.BadRequest(ctx, "Invalid request body")
	}

	if err := req.Validate(); err != nil {
		log.Errorf("[BudgetController] Create Budget - 2: %v", err)
		return response.BadRequest(ctx, err.Error())
	}

	if err := b.budgetService.CreateBudget(ctx.Context(), req); err != nil {
		log.Errorf("[BudgetController] Create Budget - 3: %v", err)
		return response.InternalError(ctx, "Failed to create budget")
	}

	return response.Created(ctx, "Budget created successfully", nil)
}

// GetBudgetByID implements [BudgetControllerInterface].
func (b *budgetController) GetBudgetByID(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	parseId, _ := uuid.Parse(id)

	budget, err := b.budgetService.GetBudgetByID(ctx.Context(), parseId)

	if err != nil {
		log.Errorf("[BudgetController] GetBudgetByID - 1: %v", err)
		return response.InternalError(ctx, err.Error())
	}

	return response.OK(ctx, "Success get budget", budget)
}

// GetBudgetByUserID implements [BudgetControllerInterface].
func (b *budgetController) GetBudgetByUserID(ctx fiber.Ctx) error {
	userId := ctx.Params("user_id")
	parseUserId, _ := uuid.Parse(userId)

	budget, err := b.budgetService.GetBudgetByUserID(ctx.Context(), parseUserId)

	if err != nil {
		log.Errorf("[BudgetController] GetBudgetByUserID - 1: %v", err)
		return response.InternalError(ctx, err.Error())
	}

	return response.OK(ctx, "Success get budget", budget)
}

// UpdateBudget implements [BudgetControllerInterface].
func (b *budgetController) UpdateBudget(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	parseId, _ := uuid.Parse(id)

	var req dto.CreateBudgetRequest

	if err := ctx.Bind().Body(&req); err != nil {
		log.Errorf("[BudgetController] Update Budget - 1: %v", err)
		return response.BadRequest(ctx, "Invalid request body")
	}

	if err := req.Validate(); err != nil {
		log.Errorf("[BudgetController] Update Budget - 2: %v", err)
		return response.BadRequest(ctx, err.Error())
	}

	if err := b.budgetService.UpdateBudget(ctx.Context(), parseId, req); err != nil {
		log.Errorf("[BudgetController] Update Budget - 3: %v", err)
		return response.InternalError(ctx, "Failed to create budget")
	}

	return response.Created(ctx, "Budget created successfully", nil)
}

func NewBudgetController(budgetService service.BudgetServiceInterface) BudgetControllerInterface {
	return &budgetController{
		budgetService,
	}
}
