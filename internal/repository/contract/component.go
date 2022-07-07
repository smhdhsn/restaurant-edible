package contract

import (
	"github.com/smhdhsn/restaurant-edible/internal/repository/entity"
)

// ComponentRepository is the interface representing component repository or it's mock.
type ComponentRepository interface {
	GetUnavailable() ([]*entity.Component, error)
}
