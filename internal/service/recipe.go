package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	"github.com/smhdhsn/restaurant-menu/internal/repository"
)

// RecipeService contains repositories that will be used within this service.
type RecipeService struct {
	fRepo repository.FoodRepository
}

// NewRecipeService creates a recipe service with it's dependencies.
func NewRecipeService(fRepo repository.FoodRepository) *RecipeService {
	return &RecipeService{fRepo: fRepo}
}

// CreateRecipe stores couple of sample recipes inside database.
func (s *RecipeService) CreateRecipe(fList []*model.Food) (err error) {
	err = s.fRepo.BatchInsert(fList)
	if err != nil {
		return errors.Wrap(err, "failed to batch insert foods")
	}

	return
}
