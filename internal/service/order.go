package service

import (
	"github.com/pkg/errors"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// OrderServ contains repositories that will be used within this service.
type OrderServ struct {
	fRepo repositoryContract.FoodRepository
	iRepo repositoryContract.InventoryRepository
}

// NewOrderService creates an order service with it's dependencies.
func NewOrderService(fr repositoryContract.FoodRepository, ir repositoryContract.InventoryRepository) serviceContract.OrderService {
	return &OrderServ{
		fRepo: fr,
		iRepo: ir,
	}
}

// GetFood is responsible for fetching available meals from database.
func (s *OrderServ) OrderFood(foodID uint32) (bool, error) {
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
