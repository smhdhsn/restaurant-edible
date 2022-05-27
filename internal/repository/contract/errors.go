package contract

import (
	"errors"
)

// This block holds common errors that might happen within repository.
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
)
