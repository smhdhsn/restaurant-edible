package handler

import (
	"context"

	eipb "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	serviceContract "github.com/smhdhsn/restaurant-edible/internal/service/contract"
)

// InventoryHandler contains services that can be used within inventory handler.
type InventoryHandler struct {
	inventoryServ serviceContract.InventoryService
}

// NewInventoryHandler creates a new menu handler.
func NewInventoryHandler(is serviceContract.InventoryService) eipb.InventoryServiceServer {
	return &InventoryHandler{
		inventoryServ: is,
	}
}

// Recycle is responsible for recycling finished and/or expired items from inventory.
func (s *InventoryHandler) Recycle(ctx context.Context, req *eipb.InventoryRecycleRequest) (*eipb.InventoryRecycleResponse, error) {
	return nil, nil
}

// Use is responsible for decreasing item's inside from inventory.
func (s *InventoryHandler) Use(ctx context.Context, req *eipb.InventoryUseRequest) (*eipb.InventoryUseResponse, error) {
	return nil, nil
}

// Buy is responsible for increasing item's stock inside inventory.
func (s *InventoryHandler) Buy(ctx context.Context, req *eipb.InventoryBuyRequest) (*eipb.InventoryBuyResponse, error) {
	return nil, nil
}
