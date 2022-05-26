package menu

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"
)

// MenuService is the interface that menu service must implement.
type MenuService interface {
	GetFoods() ([]*model.Food, error)
}
