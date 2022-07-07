package model

import (
	"time"
)

// InventoryDTO represents inventory's data transfer object.
type InventoryDTO struct {
	ID          uint32
	ComponentID uint32
	Component   ComponentDTO
	Stock       uint32
	ExpiresAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
