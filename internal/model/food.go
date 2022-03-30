package model

import (
	"gorm.io/gorm"
)

// Food represents the food table's model.
type Food struct {
	gorm.Model
	Title      string       `gorm:"unique;not null"`
	Components []*Component `gorm:"many2many:food_components"`
}
