package resource

import (
	"github.com/smhdhsn/restaurant-menu/internal/server/handler"
)

// MenuResource holds menu resource's handlers.
type MenuResource struct {
	MenuHandler      *handler.MenuHandler
	RecipeHandler    *handler.RecipeHandler
	InventoryHandler *handler.InventoryHandler
}

// NewMenuResource creates a new menu resource with all handlers within itself.
func NewMenuResource(mh *handler.MenuHandler, rh *handler.RecipeHandler, ih *handler.InventoryHandler) *MenuResource {
	return &MenuResource{
		MenuHandler:      mh,
		RecipeHandler:    rh,
		InventoryHandler: ih,
	}
}
