package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// InventoryRepository is the interface representing inventory repository or it's mock.
type InventoryRepository interface {
	Buy([]*dto.InventoryDTO) error
	Use(*dto.FoodDTO) error
	Recycle(bool, bool) error
}
