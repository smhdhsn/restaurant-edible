package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/model"
)

// ComponentRepository is the interface representing component repository or it's mock.
type ComponentRepository interface {
	GetUnavailable() (model.ComponentListDTO, error)
}
