package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-edible/internal/service/dto"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	recipeProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
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
	fListDTO := singleRecipeReqToMultipleFoodDTO(req)

	err := s.recipeServ.Store(fListDTO)
	if err != nil {
		log.Error(err)
		return nil, status.Errorf(codes.Internal, "internal server error: %w", err)
	}

	resp := new(recipeProto.RecipeStoreResponse)

	return resp, nil
}

// singleRecipeReqToMultipleFoodDTO is responsible for transforming a list of recipe req to a list of food dto struct.
func singleRecipeReqToMultipleFoodDTO(req *recipeProto.RecipeStoreRequest) []*dto.Food {
	fListDTO := make([]*dto.Food, len(req.Recipes))

	for i, fReq := range req.GetRecipes() {
		cListDTO := make([]*dto.Component, len(fReq.GetComponentTitles()))

		for i, cTitle := range fReq.GetComponentTitles() {
			cListDTO[i] = &dto.Component{Title: cTitle}
		}

		fListDTO[i] = &dto.Food{
			Title:      fReq.GetFoodTitle(),
			Components: cListDTO,
		}
	}

	return fListDTO
}
