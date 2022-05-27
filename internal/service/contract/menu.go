package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// MenuService is the interface that menu service must implement.
type MenuService interface {
	List() ([]*model.FoodDTO, error)
}
