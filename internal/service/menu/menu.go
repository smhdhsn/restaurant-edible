package menu

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
	mServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/menu"
)

// MenuServ contains repositories that will be used within this service.
type MenuServ struct {
	fRepo fRepoContract.FoodRepository
}

// NewMenuServ creates a menu service with it's dependencies.
func NewMenuServ(fRepo fRepoContract.FoodRepository) mServContract.MenuService {
	return &MenuServ{fRepo: fRepo}
}

// GetFood is responsible for fetching available meals from database.
func (s *MenuServ) GetFoods() ([]*model.Food, error) {
	return s.fRepo.GetAvailable()
}
