package contract

import (
	"errors"

	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// This section holds errors that might happen within service layer.
var (
	ErrLackOfComponents = errors.New("lack_of_components")
)

// InventoryService is the interface that inventory service must implement.
type InventoryService interface {
	Recycle(*dto.Recycle) error
	Use(*dto.Food) error
	Buy(*dto.Buy) error
}
