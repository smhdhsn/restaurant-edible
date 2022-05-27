package handler

import (
	"context"
	"errors"

	"github.com/smhdhsn/restaurant-edible/internal/model"
	erpb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RecipeHandler contains services that can be used within recipe handler.
type RecipeHandler struct {
	recipeServ serviceContract.RecipeService
}

// NewRecipeHandler creates a new menu handler.
func NewRecipeHandler(rs serviceContract.RecipeService) erpb.RecipeServiceServer {
	return &RecipeHandler{
		recipeServ: rs,
	}
}

// Store is responsible for storing item's recipe inside database.
func (s *RecipeHandler) Store(ctx context.Context, req *erpb.RecipeStoreRequest) (*erpb.RecipeStoreResponse, error) {
	fList := make([]*model.FoodDTO, len(req.Recipes))
	for i, f := range req.GetRecipes() {
		cList := make([]*model.ComponentDTO, len(f.GetComponentTitles()))
		for i, cTitle := range f.GetComponentTitles() {
			cList[i] = &model.ComponentDTO{Title: cTitle}
		}

		fList[i] = &model.FoodDTO{
			Title:      f.GetFoodTitle(),
			Components: cList,
		}
	}

	err := s.recipeServ.Store(fList)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrDuplicateEntry) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(erpb.RecipeStoreResponse)

	return resp, nil
}
