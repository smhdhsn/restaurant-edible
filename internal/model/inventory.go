package model

import (
	"time"
)

// Inventory represents the inventories table on database.
type Inventory struct {
	ID          uint32    `gorm:"primaryKey"`
	ComponentID uint32    `gorm:"index;not null"`
	Component   Component `gorm:"constraint:OnDelete:CASCADE"`
	Stock       uint32    `gorm:"not null"`
	ExpiresAt   time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// InventoryDTO represents inventory's data transfer object.
type InventoryDTO struct {
	ID          uint32
	ComponentID uint32
	Component   Component
	Stock       uint32
	ExpiresAt   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
