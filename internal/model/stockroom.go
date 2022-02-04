package model

import (
	"time"

	"gorm.io/gorm"
)

// Stockroom represents the stockroom table's model.
type Stockroom struct {
	gorm.Model
	Ingredient Ingredient `gorm:"constraint:OnDelete:CASCADE" validate:"required"`
	Stock      uint       `gorm:"not null" validate:"required,gte=0"`
	BestBefore time.Time  `gorm:"not null" validate:"required,datetime"`
	ExpiresAt  time.Time  `gorm:"not null" validate:"required,datetime"`
}
