package entity

import (
	"time"
)

// Food represents the food repository's entity.
type Food struct {
	ID         uint32
	Title      string
	Components []*Component
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
