package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// OrderService is the interface that order service must implement.
type OrderService interface {
	OrderFood(*model.FoodDTO) error
}
