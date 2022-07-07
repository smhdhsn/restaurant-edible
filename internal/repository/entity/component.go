package entity

import (
	"time"
)

// Component represents the component repository's entity.
type Component struct {
	ID        uint32
	Title     string
	Foods     []*Food
	CreatedAt time.Time
	UpdatedAt time.Time
}
