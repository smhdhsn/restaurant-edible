package http

import (
	"fmt"
	"net/http"

	"github.com/smhdhsn/food/internal/service"
)

// MenuHandler contains services that can be used within menu handler.
type MenuHandler struct {
	menuService *service.MenuService
}

// NewMenuHandler creates a new menu handler.
func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

// GetMenu is responsible for getting menu with available food.
func (h *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You're here.")
}
