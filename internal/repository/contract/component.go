package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/service/dto"
)

// ComponentRepository is the interface representing component repository or it's mock.
type ComponentRepository interface {
	GetUnavailable() ([]*dto.ComponentDTO, error)
}
