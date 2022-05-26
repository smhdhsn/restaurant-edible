package service

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-menu/internal/model"

	fRepoContract "github.com/smhdhsn/restaurant-menu/internal/repository/contract/food"
	rServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/recipe"
)

// RecipeServ contains repositories that will be used within this service.
type RecipeServ struct {
	fRepo fRepoContract.FoodRepository
}

// NewRecipeServ creates a recipe service with it's dependencies.
func NewRecipeServ(fRepo fRepoContract.FoodRepository) rServContract.RecipeService {
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
