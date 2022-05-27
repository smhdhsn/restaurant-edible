package resource

import (
	eipb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	empb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	erpb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
)

// EdibleResource holds menu resource's handlers.
type EdibleResource struct {
	MenuHandler      empb.MenuServiceServer
	RecipeHandler    erpb.RecipeServiceServer
	InventoryHandler eipb.InventoryServiceServer
}

// NewEdibleResource creates a new menu resource with all handlers within itself.
func NewEdibleResource(mh empb.MenuServiceServer, rh erpb.RecipeServiceServer, ih eipb.InventoryServiceServer) *EdibleResource {
	return &EdibleResource{
		MenuHandler:      mh,
		RecipeHandler:    rh,
		InventoryHandler: ih,
	}
}
