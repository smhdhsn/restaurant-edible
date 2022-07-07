package dto

import (
	"time"
)

// ComponentDTO represents component's data transfer object.
type ComponentDTO struct {
	ID        uint32
	Title     string
	Foods     []*FoodDTO
	CreatedAt time.Time
	UpdatedAt time.Time
}
