package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// MenuService is the interface that menu service must implement.
type MenuService interface {
	List() ([]*dto.FoodDTO, error)
}
