package resource

import (
	"github.com/smhdhsn/restaurant-menu/internal/server/handler"
)

// MenuResource holds menu resource's handlers.
type MenuResource struct {
	SourceHandler *handler.MenuSourceHandler
}

// NewMenuResource creates a new menu resource with all handlers within itself.
func NewMenuResource(source *handler.MenuSourceHandler) *MenuResource {
	return &MenuResource{
		SourceHandler: source,
	}
}
