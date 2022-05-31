package model

import (
	"time"
)

// Component represents the component table's model.
type Component struct {
	ID        uint32  `gorm:"primaryKey"`
	Title     string  `gorm:"unique;not null"`
	Foods     []*Food `gorm:"many2many:food_components"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ComponentDTO represents component's data transfer object.
type ComponentDTO struct {
	ID        uint32
	Title     string
	Foods     FoodListDTO
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ComponentListDTO represents a list of ComponentDTOs.
type ComponentListDTO []*ComponentDTO
