package repository

import "github.com/smhdhsn/food/internal/model"

// FoodRepository is the interface representing food repository or it's mock.
type FoodRepository interface {
	GetAvailableMeals() ([]*model.Food, error)
}
