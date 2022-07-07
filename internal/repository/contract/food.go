package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
)

// FoodRepository is the interface representing food repository or it's mock.
type FoodRepository interface {
	GetAvailable() ([]*entity.Food, error)
	BatchInsert([]*entity.Food) error
}
