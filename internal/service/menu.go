package service

import (
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
)

// MenuService contains repositories that will be used within this service.
type MenuService struct {
	fRepo repository.FoodRepository
}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService(fRepo repository.FoodRepository) *MenuService {
	return &MenuService{fRepo: fRepo}
}

// GetFood is responsible for fetching available meals from database.
func (s *MenuService) GetFoods() ([]*model.Food, error) {
	return s.fRepo.GetAvailable()
}
