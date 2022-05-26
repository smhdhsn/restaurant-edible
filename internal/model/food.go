package model

import (
	"time"
)

// Food represents the food table's model.
type Food struct {
	ID         uint32       `gorm:"primaryKey"`
	Title      string       `gorm:"unique;not null"`
	Components []*Component `gorm:"many2many:food_components"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
