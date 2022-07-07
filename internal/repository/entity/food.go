package entity

import (
	"time"
)

// Food represents the food model's entity.
type Food struct {
	ID         uint32
	Title      string
	Components []*Component
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
