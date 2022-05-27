package mock

import (
	"github.com/stretchr/testify/mock"

	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// FoodRepo is food repository's mock.
type FoodRepo struct {
	mock.Mock
}

// GetAvailable is a mocked method in mocked food repository.
func (r *FoodRepo) GetAvailable() ([]*model.Food, error) {
	args := r.Mock.Called()

	if v := args.Get(0); v != nil {
		return v.([]*model.Food), nil
	}

	return nil, args.Error(1)
}

// BatchInert is a mocked method in mocked food repository.
func (r *FoodRepo) BatchInsert(fList []*model.Food) error {
	args := r.Mock.Called(fList)

	return args.Error(0)
}
