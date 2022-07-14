package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
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
func (s *MenuServ) List() ([]*dto.Food, error) {
	fListEntity, err := s.fRepo.GetAvailable()
	if err != nil {
		return nil, errors.Wrap(err, "error on calling get available on food repository")
	}

	fListDTO := multipleFoodEntityToDTO(fListEntity)

	return fListDTO, nil
}

// multipleFoodEntityToDTO is responsible for transforming a list of food entity to food dto struct.
func multipleFoodEntityToDTO(fListEntity []*entity.Food) []*dto.Food {
	fListDTO := make([]*dto.Food, len(fListEntity))

	for i, fEntity := range fListEntity {
		cListDTO := make([]*dto.Component, len(fEntity.Components))

		for j, cEntity := range fEntity.Components {
			cListDTO[j] = &dto.Component{
				ID:        cEntity.ID,
				Title:     cEntity.Title,
				CreatedAt: cEntity.CreatedAt,
				UpdatedAt: cEntity.UpdatedAt,
			}
		}

		fListDTO[i] = &dto.Food{
			ID:         fEntity.ID,
			Title:      fEntity.Title,
			Components: cListDTO,
			CreatedAt:  fEntity.CreatedAt,
			UpdatedAt:  fEntity.UpdatedAt,
		}
	}

	return fListDTO
}
