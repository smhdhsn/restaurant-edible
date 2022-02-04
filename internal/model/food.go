package model

import (
	"gorm.io/gorm"
)

// Food represents the food table's model.
type Food struct {
	gorm.Model
	Title       string        `gorm:"not null" validate:"required"`
	Ingredients []*Ingredient `gorm:"many2many:food_ingredients;" validate:"required"`
}
