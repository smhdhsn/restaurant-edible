package resource

import (
	"github.com/smhdhsn/restaurant-edible/internal/server/handler"
)

// EdibleResource holds menu resource's handlers.
type EdibleResource struct {
	MenuHandler      *handler.MenuHandler
	RecipeHandler    *handler.RecipeHandler
	InventoryHandler *handler.InventoryHandler
}

// NewEdibleResource creates a new menu resource with all handlers within itself.
func NewEdibleResource(mh *handler.MenuHandler, rh *handler.RecipeHandler, ih *handler.InventoryHandler) *EdibleResource {
	return &EdibleResource{
		MenuHandler:      mh,
		RecipeHandler:    rh,
		InventoryHandler: ih,
	}
}
