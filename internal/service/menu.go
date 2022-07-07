package service

import (
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// MenuServ contains repositories that will be used within this service.
type MenuServ struct {
	fRepo repositoryContract.FoodRepository
}

// NewMenuService creates a menu service with it's dependencies.
func NewMenuService(fr repositoryContract.FoodRepository) serviceContract.MenuService {
	return &MenuServ{fRepo: fr}
}

// List is responsible for fetching available meals from database.
func (s *MenuServ) List() ([]*dto.FoodDTO, error) {
	return s.fRepo.GetAvailable()
}
