package model

import (
	"time"

	"gorm.io/gorm"
)

// Stock represents the Stock table's model.
type Stock struct {
	gorm.Model
	IngredientID uint       `gorm:"not null index" validate:"required"`
	Ingredient   Ingredient `gorm:"constraint:OnDelete:CASCADE" validate:"required"`
	Stock        uint       `gorm:"not null" validate:"required,gte=0"`
	BestBefore   time.Time  `gorm:"not null" validate:"required,datetime"`
	ExpiresAt    time.Time  `gorm:"not null" validate:"required,datetime"`
}
