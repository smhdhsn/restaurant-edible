package model

import (
	"time"

	"gorm.io/gorm"
)

// Inventory represents the inventories table's model.
type Inventory struct {
	gorm.Model
	ComponentID uint      `gorm:"index;not null"`
	Component   Component `gorm:"constraint:OnDelete:CASCADE"`
	Stock       uint      `gorm:"not null"`
	BestBefore  time.Time `gorm:"not null"`
	ExpiresAt   time.Time `gorm:"not null"`
}
