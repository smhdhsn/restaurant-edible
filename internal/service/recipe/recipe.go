package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-menu/internal/model"
	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
)

// RecipeService contains repositories that will be used within this service.
type RecipeService struct {
	fRepo fRepoContract.FoodRepository
}

// NewRecipeService creates a recipe service with it's dependencies.
func NewRecipeService(fRepo fRepoContract.FoodRepository) *RecipeService {
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
