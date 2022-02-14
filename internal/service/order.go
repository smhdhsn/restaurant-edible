package service

import (
	"github.com/pkg/errors"
	"github.com/smhdhsn/food/internal/repository"
)

// OrderService contains repositories that will be used within this service.
type OrderService struct {
	fRepo repository.FoodRepository
	iRepo repository.InventoryRepository
}

// NewOrderService creates an order service with it's dependencies.
func NewOrderService(fRepo repository.FoodRepository, iRepo repository.InventoryRepository) *OrderService {
	return &OrderService{fRepo: fRepo, iRepo: iRepo}
}

// GetFood is responsible for fetching available meals from database.
func (s *OrderService) OrderFood(foodID uint) (bool, error) {
	foods, err := s.fRepo.GetAvailableMeals()
	if err != nil {
		return false, errors.Wrap(err, "failed to get available foods")
	}

	var isAvailable bool
	for _, f := range foods {
		if f.ID == foodID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return false, errors.New("requested order cannot be fulfilled because of the lack of components")
	}

	err = s.iRepo.UseComponents(foodID)
	if err != nil {
		return false, errors.Wrap(err, "failed to use components")
	}

	return true, nil
}
