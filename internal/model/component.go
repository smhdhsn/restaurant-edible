package model

import (
	"gorm.io/gorm"
)

// Component represents the component table's model.
type Component struct {
	gorm.Model
	Title string  `gorm:"unique,not null" validate:"required"`
	Foods []*Food `gorm:"many2many:food_components;" validate:"required"`
}
