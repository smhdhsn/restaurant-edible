package dto

import (
	"time"
)

// Food represents food's data transfer object.
type Food struct {
	ID         uint32
	Title      string
	Components []*Component
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
