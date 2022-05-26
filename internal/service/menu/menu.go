package menu

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
)

// MenuService contains repositories that will be used within this service.
type MenuService struct {
	fRepo fRepoContract.FoodRepository
}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService(fRepo fRepoContract.FoodRepository) *MenuService {
	return &MenuService{fRepo: fRepo}
}

// GetFood is responsible for fetching available meals from database.
func (s *MenuService) GetFoods() ([]*model.Food, error) {
	return s.fRepo.GetAvailable()
}
