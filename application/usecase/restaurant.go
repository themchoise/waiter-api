package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/waiter/back/domain/entity"
	"github.com/waiter/back/domain/repository"
)

type RestaurantUseCase struct {
	restaurantRepo repository.RestaurantRepository
	tableRepo      repository.TableRepository
}

func NewRestaurantUseCase(rr repository.RestaurantRepository, tr repository.TableRepository) *RestaurantUseCase {
	return &RestaurantUseCase{restaurantRepo: rr, tableRepo: tr}
}

type CreateRestaurantInput struct {
	Name string `json:"name" binding:"required"`
	Plan string `json:"plan"`
}

func (uc *RestaurantUseCase) CreateRestaurant(input CreateRestaurantInput) (*entity.Restaurant, error) {
	plan := input.Plan
	if plan == "" {
		plan = "free"
	}

	r := &entity.Restaurant{
		ID:   uuid.New().String(),
		Name: input.Name,
		Plan: plan,
	}

	if err := uc.restaurantRepo.Create(r); err != nil {
		return nil, fmt.Errorf("failed to create restaurant: %w", err)
	}

	return r, nil
}

func (uc *RestaurantUseCase) GetRestaurant(id string) (*entity.Restaurant, error) {
	return uc.restaurantRepo.FindByID(id)
}

type CreateTableInput struct {
	Number       int    `json:"number" binding:"required"`
	RestaurantID string `json:"restaurant_id" binding:"required"`
}

func (uc *RestaurantUseCase) CreateTable(input CreateTableInput) (*entity.Table, error) {
	if _, err := uc.restaurantRepo.FindByID(input.RestaurantID); err != nil {
		return nil, fmt.Errorf("restaurant not found: %w", err)
	}

	t := &entity.Table{
		ID:           uuid.New().String(),
		Number:       input.Number,
		RestaurantID: input.RestaurantID,
		QRCode:       fmt.Sprintf("table-%s-%d", input.RestaurantID, input.Number),
	}

	if err := uc.tableRepo.Create(t); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return t, nil
}

func (uc *RestaurantUseCase) GetTables(restaurantID string) ([]entity.Table, error) {
	return uc.tableRepo.FindByRestaurantID(restaurantID)
}
