package contract

import (
	"github.com/smhdhsn/restaurant-menu/internal/model"
)

// ComponentRepository is the interface representing component repository or it's mock.
type ComponentRepository interface {
	GetUnavailable() ([]*model.Component, error)
}
