package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
)

// MenuService is the interface that menu service must implement.
type MenuService interface {
	List() ([]*entity.Food, error)
}
