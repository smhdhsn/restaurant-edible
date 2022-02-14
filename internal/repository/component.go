package repository

import "github.com/smhdhsn/food/internal/model"

// ComponentRepository is the interface representing component repository or it's mock.
type ComponentRepository interface {
	GetUnavailable() ([]*model.Component, error)
}
