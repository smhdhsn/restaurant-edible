package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// FoodRepository is the interface representing food repository or it's mock.
type FoodRepository interface {
	GetAvailable() ([]*model.Food, error)
	BatchInsert([]*model.Food) error
}
