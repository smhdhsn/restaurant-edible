package service

import (
	"github.com/pkg/errors"
	"github.com/smhdhsn/food/internal/model"
	"github.com/smhdhsn/food/internal/repository"
)

// RecipeService contains repositories that will be used within this service.
type RecipeService struct {
	fRepo repository.FoodRepository
}

// RecipeSchema holds schema for recipe JSON.
type RecipeSchema struct {
	Foods []struct {
		Title      string   `json:"title"`
		Components []string `json:"components"`
	} `json:"foods"`
}

// NewRecipeService creates a recipe service with it's dependencies.
func NewRecipeService(fRepo repository.FoodRepository) *RecipeService {
	return &RecipeService{fRepo: fRepo}
}

// CreateRecipe stores couple of sample recipes inside database.
func (s *RecipeService) CreateRecipe(recipe *RecipeSchema) (err error) {
	cList := make([]*model.Component, 0)
	fList := make([]*model.Food, 0)
	for _, f := range recipe.Foods {
		for _, cTitle := range f.Components {
			cList = append(cList, &model.Component{Title: cTitle})
		}

		fList = append(fList, &model.Food{Title: f.Title, Components: cList})
	}

	err = s.fRepo.BatchInsert(fList)
	if err != nil {
		return errors.Wrap(err, "failed to batch insert foods")
	}

	return
}
