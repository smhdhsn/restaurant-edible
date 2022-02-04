package service

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
)

// MenuService contains repositories that will be used within this service.
type MenuService struct {
	foodRepo repository.FoodRepository
}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService(foodRepo repository.FoodRepository) *MenuService {
	return &MenuService{foodRepo: foodRepo}
}

func (s *MenuService) GetFoods() ([]*model.Food, error) {
	return s.foodRepo.GetAvailableMeals()
}
