package handler

import (
	serviceContract "github.com/smhdhsn/restaurant-menu/internal/service/contract"
)

// RecipeHandler contains services that can be used within recipe handler.
type RecipeHandler struct {
	rServ serviceContract.RecipeService
}

// NewRecipeHandler creates a new menu handler.
func NewRecipeHandler(rs serviceContract.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		rServ: rs,
	}
}
