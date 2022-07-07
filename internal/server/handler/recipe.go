package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	recipeProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
	repositoryContract "github.com/smhdhsn/restaurant-edible/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// RecipeHandler contains services that can be used within recipe handler.
type RecipeHandler struct {
	recipeServ serviceContract.RecipeService
}

// NewRecipeHandler creates a new menu handler.
func NewRecipeHandler(rs serviceContract.RecipeService) recipeProto.EdibleRecipeServiceServer {
	return &RecipeHandler{
		recipeServ: rs,
	}
}

// Store is responsible for storing item's recipe inside database.
func (s *RecipeHandler) Store(ctx context.Context, req *recipeProto.RecipeStoreRequest) (*recipeProto.RecipeStoreResponse, error) {
	fList := make([]*dto.FoodDTO, len(req.Recipes))
	for i, f := range req.GetRecipes() {
		cList := make([]*dto.ComponentDTO, len(f.GetComponentTitles()))
		for i, cTitle := range f.GetComponentTitles() {
			cList[i] = &dto.ComponentDTO{Title: cTitle}
		}

		fList[i] = &dto.FoodDTO{
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

	resp := new(recipeProto.RecipeStoreResponse)

	return resp, nil
}
