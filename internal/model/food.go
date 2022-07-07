package model

import (
	"time"
)

// FoodDTO represents food's data transfer object.
type FoodDTO struct {
	ID         uint32
	Title      string
	Components []*ComponentDTO
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
