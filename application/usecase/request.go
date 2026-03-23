package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/waiter/back/domain/entity"
	"github.com/waiter/back/domain/repository"
)

type RequestUseCase struct {
	requestRepo repository.RequestRepository
	tableRepo   repository.TableRepository
	notifier    EventNotifier
}

// EventNotifier sends real-time events to connected clients.
type EventNotifier interface {
	Notify(restaurantID string, event any)
}

func NewRequestUseCase(rr repository.RequestRepository, tr repository.TableRepository, n EventNotifier) *RequestUseCase {
	return &RequestUseCase{requestRepo: rr, tableRepo: tr, notifier: n}
}

type CreateRequestInput struct {
	TableID string             `json:"table_id" binding:"required"`
	Type    entity.RequestType `json:"type" binding:"required"`
}

func (uc *RequestUseCase) CreateRequest(input CreateRequestInput) (*entity.Request, error) {
	if input.Type != entity.CallWaiter && input.Type != entity.AskBill && input.Type != entity.AskHelp {
		return nil, fmt.Errorf("invalid request type: %s", input.Type)
	}

	table, err := uc.tableRepo.FindByID(input.TableID)
	if err != nil {
		return nil, fmt.Errorf("table not found: %w", err)
	}

	req := &entity.Request{
		ID:      uuid.New().String(),
		TableID: input.TableID,
		Type:    input.Type,
		Status:  entity.Pending,
	}

	if err := uc.requestRepo.Create(req); err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	uc.notifier.Notify(table.RestaurantID, map[string]any{
		"event":   "new_request",
		"request": req,
	})

	return req, nil
}

func (uc *RequestUseCase) GetActiveRequests(restaurantID string) ([]entity.Request, error) {
	return uc.requestRepo.FindActiveByRestaurantID(restaurantID)
}

func (uc *RequestUseCase) CompleteRequest(id string) error {
	req, err := uc.requestRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("request not found: %w", err)
	}

	if req.Status == entity.Done {
		return fmt.Errorf("request already completed")
	}

	if err := uc.requestRepo.UpdateStatus(id, entity.Done); err != nil {
		return fmt.Errorf("failed to update request: %w", err)
	}

	table, err := uc.tableRepo.FindByID(req.TableID)
	if err == nil {
		uc.notifier.Notify(table.RestaurantID, map[string]any{
			"event":      "request_completed",
			"request_id": id,
		})
	}

	return nil
}

func (uc *RequestUseCase) GetTableStatus(tableID string) ([]entity.Request, error) {
	return uc.requestRepo.FindByTableID(tableID)
}
