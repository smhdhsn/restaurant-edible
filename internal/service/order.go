package service

import (
	"github.com/smhdhsn/food/internal/repository"
)

// OrderService contains repositories that will be used within this service.
type OrderService struct {
	fRepo repository.FoodRepository
	sRepo repository.StockRepository
	iRepo repository.IngredientRepository
}

// NewOrderService creates an order service with it's dependencies.
func NewOrderService(fRepo repository.FoodRepository, sRepo repository.StockRepository, iRepo repository.IngredientRepository) *OrderService {
	return &OrderService{fRepo: fRepo, sRepo: sRepo, iRepo: iRepo}
}

// GetFood is responsible for fetching available meals from database.
func (s *OrderService) OrderFood(foodID uint) (bool, error) {
	foods, err := s.fRepo.GetAvailableMeals()
	if err != nil {
		return false, err
	}

	var isAvailable bool
	for _, f := range foods {
		if f.ID == foodID {
			isAvailable = true
		}
	}

	if !isAvailable {
		return false, repository.ErrNotAvailable
	}

	ingrds, err := s.iRepo.GetFoodIngredients(foodID)
	if err != nil {
		return false, err
	}

	ingrdIDs := make([]uint, 0)
	for _, i := range ingrds {
		ingrdIDs = append(ingrdIDs, i.ID)
	}

	err = s.sRepo.UseIngredients(ingrdIDs)
	if err != nil {
		return false, repository.ErrNotAvailable
	}

	return true, nil
}
