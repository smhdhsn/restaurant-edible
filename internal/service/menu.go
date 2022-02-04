package service

import (
	"github.com/smhdhsn/food/internal/repository"
)

// MenuService contains repositories that will be used within this service.
type MenuService struct {
	foodRepo       repository.FoodRepository
	ingredientRepo repository.IngredientRepository
}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService(foodRepo repository.FoodRepository, ingredientRepo repository.IngredientRepository) *MenuService {
	return &MenuService{foodRepo: foodRepo, ingredientRepo: ingredientRepo}
}
