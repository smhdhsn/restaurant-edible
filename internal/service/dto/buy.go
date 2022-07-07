package dto

import (
	"time"
)

// Buy represents buy's data transfer object.
type Buy struct {
	Stock     uint32
	ExpiresAt time.Time
}
