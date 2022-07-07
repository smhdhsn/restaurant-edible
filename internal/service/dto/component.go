package dto

import (
	"time"
)

// Component represents component's data transfer object.
type Component struct {
	ID        uint32
	Title     string
	Foods     []*Food
	CreatedAt time.Time
	UpdatedAt time.Time
}
