package contract

import (
	"time"

	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// InventoryService is the interface that inventory service must implement.
type InventoryService interface {
	Buy(uint32, time.Time) error
	Use(*dto.FoodDTO) error
	Recycle(bool, bool) error
}
