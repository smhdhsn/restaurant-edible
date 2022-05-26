package service

import (
	"github.com/pkg/errors"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
	iRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/inventory"
)

// OrderService contains repositories that will be used within this service.
type OrderService struct {
	fRepo fRepoContract.FoodRepository
	iRepo iRepoContract.InventoryRepository
}

// NewOrderService creates an order service with it's dependencies.
func NewOrderService(fRepo fRepoContract.FoodRepository, iRepo iRepoContract.InventoryRepository) *OrderService {
	return &OrderService{
		fRepo: fRepo,
		iRepo: iRepo,
	}
}

// GetFood is responsible for fetching available meals from database.
func (s *OrderService) OrderFood(foodID uint32) (bool, error) {
	foods, err := s.fRepo.GetAvailable()
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

	err = s.iRepo.Use(foodID)
	if err != nil {
		return false, errors.Wrap(err, "failed to use components")
	}

	return true, nil
}
