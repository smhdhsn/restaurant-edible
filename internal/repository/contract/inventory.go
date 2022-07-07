package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
)

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	Recycle(*entity.Recycle) error
	Buy([]*entity.Inventory) error
	Use(*entity.Food) error
}
