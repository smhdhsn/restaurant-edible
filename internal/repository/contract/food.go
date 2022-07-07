package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// FoodRepository is the interface representing food repository or it's mock.
type FoodRepository interface {
	GetAvailable() ([]*dto.FoodDTO, error)
	BatchInsert([]*dto.FoodDTO) error
}
