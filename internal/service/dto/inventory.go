package dto

import (
	"time"
)

// Inventory represents inventory's data transfer object.
type Inventory struct {
	ID          uint32
	ComponentID uint32
	Component   Component
	Stock       uint32
	ExpiresAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
