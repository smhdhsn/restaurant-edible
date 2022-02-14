package service

import (
	"github.com/smhdhsn/food/internal/repository"
)

// InventoryService contains repositories that will be used within this service.
type InventoryService struct {
	fRepo repository.FoodRepository
	cRepo repository.ComponentRepository
	iRepo repository.InventoryRepository
}

// NewInventoryService creates an inventory service with it's dependencies.
func NewInventoryService(fRepo repository.FoodRepository, cRepo repository.ComponentRepository, iRepo repository.InventoryRepository) *InventoryService {
	return &InventoryService{fRepo: fRepo, cRepo: cRepo, iRepo: iRepo}
}
