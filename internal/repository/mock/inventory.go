package mock

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/stretchr/testify/mock"
)

// InventoryRepo is inventory repository's mock.
type InventoryRepo struct {
	mock.Mock
}

// UseStocks is a mocked method in mocked inventory repository.
func (r *InventoryRepo) UseStocks(foodID uint) error {
	args := r.Mock.Called(foodID)

	return args.Error(0)
}

// BuyStocks is a mocked method in mocked inventory repository.
func (r *InventoryRepo) BuyStocks(iList []*model.Inventory) error {
	args := r.Mock.Called(iList)

	return args.Error(0)
}
