package entity

import (
	"time"
)

// Component represents the component model's entity.
type Component struct {
	ID        uint32
	Title     string
	Foods     []*Food
	CreatedAt time.Time
	UpdatedAt time.Time
}
