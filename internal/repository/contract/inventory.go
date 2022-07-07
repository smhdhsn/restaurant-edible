package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	Buy([]*model.InventoryDTO) error
	Use(*model.FoodDTO) error
	Recycle(bool, bool) error
}
