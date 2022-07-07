package resource

import (
	inventoryProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	menuProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	recipeProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
)

// EdibleResource holds menu resource's handlers.
type EdibleResource struct {
	InventoryHandler inventoryProto.EdibleInventoryServiceServer
	RecipeHandler    recipeProto.EdibleRecipeServiceServer
	MenuHandler      menuProto.EdibleMenuServiceServer
}

// NewEdibleResource creates a new menu resource with all handlers within itself.
func NewEdibleResource(
	ih inventoryProto.EdibleInventoryServiceServer,
	rh recipeProto.EdibleRecipeServiceServer,
	mh menuProto.EdibleMenuServiceServer,
) *EdibleResource {
	return &EdibleResource{
		InventoryHandler: ih,
		RecipeHandler:    rh,
		MenuHandler:      mh,
	}
}
