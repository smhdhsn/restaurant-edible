package repository

import (
	"github.com/smhdhsn/food/internal/model"
)

// IngredientRepository is the interface representing ingredient repository or it's mock.
type IngredientRepository interface {
	GetFoodIngredients(uint) ([]*model.Ingredient, error)
}
