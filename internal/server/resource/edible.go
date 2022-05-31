package resource

import (
	eipb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	empb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	erpb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
)

// EdibleResource holds menu resource's handlers.
type EdibleResource struct {
	InventoryHandler eipb.EdibleInventoryServiceServer
	RecipeHandler    erpb.EdibleRecipeServiceServer
	MenuHandler      empb.EdibleMenuServiceServer
}

// NewEdibleResource creates a new menu resource with all handlers within itself.
func NewEdibleResource(ih eipb.EdibleInventoryServiceServer, rh erpb.EdibleRecipeServiceServer, mh empb.EdibleMenuServiceServer) *EdibleResource {
	return &EdibleResource{
		InventoryHandler: ih,
		RecipeHandler:    rh,
		MenuHandler:      mh,
	}
}
