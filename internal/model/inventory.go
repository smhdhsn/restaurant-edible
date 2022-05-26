package model

import (
	"time"
)

// Inventory represents the inventories table's model.
type Inventory struct {
	ID          uint32    `gorm:"primaryKey"`
	ComponentID uint32    `gorm:"index;not null"`
	Component   Component `gorm:"constraint:OnDelete:CASCADE"`
	Stock       uint32    `gorm:"not null"`
	BestBefore  time.Time `gorm:"not null"`
	ExpiresAt   time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
