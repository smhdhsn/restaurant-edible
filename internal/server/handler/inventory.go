package handler

import (
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// InventoryHandler contains services that can be used within inventory handler.
type InventoryHandler struct {
	iServ serviceContract.InventoryService
}

// NewInventoryHandler creates a new menu handler.
func NewInventoryHandler(is serviceContract.InventoryService) *InventoryHandler {
	return &InventoryHandler{
		iServ: is,
	}
}
