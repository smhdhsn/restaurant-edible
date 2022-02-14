package model

import (
	"time"

	"gorm.io/gorm"
)

// Inventory represents the inventories table's model.
type Inventory struct {
	gorm.Model
	ComponentID uint      `gorm:"not null,index" validate:"required"`
	Component   Component `gorm:"constraint:OnDelete:CASCADE" validate:"required"`
	Stock       uint      `gorm:"not null" validate:"required,gte=0"`
	BestBefore  time.Time `gorm:"not null" validate:"required,datetime"`
	ExpiresAt   time.Time `gorm:"not null" validate:"required,datetime"`
}
