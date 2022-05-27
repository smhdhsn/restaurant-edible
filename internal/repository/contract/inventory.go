package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// RecycleReq is the struct responsible for telling the service which items to clean up from inventory.
type RecycleReq struct {
	Finished bool
	Expired  bool
}

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	Buy([]*model.Inventory) error
	Clean(RecycleReq) error
	Use(uint32) error
}
