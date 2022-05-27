package contract

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"
)

// RecipeService is the interface that recipe service must implement.
type RecipeService interface {
	CreateRecipe(fList []*model.Food) error
}
