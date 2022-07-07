package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// RecipeServ contains repositories that will be used within this service.
type RecipeServ struct {
	fRepo repositoryContract.FoodRepository
}

// NewRecipeService creates a recipe service with it's dependencies.
func NewRecipeService(fr repositoryContract.FoodRepository) serviceContract.RecipeService {
	return &RecipeServ{fRepo: fr}
}

// CreateRecipe stores couple of sample recipes inside database.
func (s *RecipeServ) Store(fListDTO []*dto.Food) error {
	fListEntity := multipleFoodDTOToEntity(fListDTO)

	err := s.fRepo.BatchInsert(fListEntity)
	if err != nil {
		return errors.Wrap(err, "failed to batch insert foods")
	}

	return nil
}

// multipleFoodDTOToEntity is responsible for transforming a list of food dtos into a list of food entity struct.
func multipleFoodDTOToEntity(fListDTO []*dto.Food) []*entity.Food {
	fListEntity := make([]*entity.Food, len(fListDTO))

	for i, fDTO := range fListDTO {
		cListEntity := make([]*entity.Component, len(fDTO.Components))

		for j, cDTO := range fDTO.Components {
			cListEntity[j] = &entity.Component{
				ID:        cDTO.ID,
				Title:     cDTO.Title,
				CreatedAt: cDTO.CreatedAt,
				UpdatedAt: cDTO.UpdatedAt,
			}
		}

		fListEntity[i] = &entity.Food{
			ID:         fDTO.ID,
			Title:      fDTO.Title,
			Components: cListEntity,
			CreatedAt:  fDTO.CreatedAt,
			UpdatedAt:  fDTO.UpdatedAt,
		}
	}

	return fListEntity
}
