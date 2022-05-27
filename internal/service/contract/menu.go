package contract

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"
)

// MenuService is the interface that menu service must implement.
type MenuService interface {
	List() ([]*model.Food, error)
}
