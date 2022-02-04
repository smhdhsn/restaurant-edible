package model

import (
	"gorm.io/gorm"
)

// Ingredient represents the ingredient table's model.
type Ingredient struct {
	gorm.Model
	Title string  `gorm:"not null" validate:"required"`
	Foods []*Food `gorm:"many2many:food_ingredients;" validate:"required"`
}
