package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-menu/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-menu/internal/service/contract"
)

// RecipeServ contains repositories that will be used within this service.
type RecipeServ struct {
	fRepo repositoryContract.FoodRepository
}

// NewRecipeService creates a recipe service with it's dependencies.
func NewRecipeService(fRepo repositoryContract.FoodRepository) serviceContract.RecipeService {
	return &RecipeServ{fRepo: fRepo}
}

// CreateRecipe stores couple of sample recipes inside database.
func (s *RecipeServ) CreateRecipe(fList []*model.Food) (err error) {
	err = s.fRepo.BatchInsert(fList)
	if err != nil {
		return errors.Wrap(err, "failed to batch insert foods")
	}

	return
}
