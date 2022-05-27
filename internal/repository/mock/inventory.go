package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/smhdhsn/restaurant-edible/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
)

// InventoryRepo is inventory repository's mock.
type InventoryRepo struct {
	mock.Mock
}

// Use is a mocked method in mocked inventory repository.
func (r *InventoryRepo) Use(foodID uint32) error {
	args := r.Mock.Called(foodID)

	return args.Error(0)
}

// Buy is a mocked method in mocked inventory repository.
func (r *InventoryRepo) Buy(iList []*model.Inventory) error {
	args := r.Mock.Called(iList)

	return args.Error(0)
}

// Clean is a mocked method in mocked inventory repository.
func (r *InventoryRepo) Clean(req repositoryContract.RecycleReq) error {
	args := r.Mock.Called(req)

	return args.Error(0)
}
