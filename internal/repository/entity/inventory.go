package entity

import (
	"time"
)

// Inventory represents the inventory repository's entity.
type Inventory struct {
	ID          uint32
	ComponentID uint32
	Component   Component
	Stock       uint32
	ExpiresAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
