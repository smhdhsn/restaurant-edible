package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// RecipeService is the interface that recipe service must implement.
type RecipeService interface {
	Store([]*dto.Food) error
}
