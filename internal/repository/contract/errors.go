package contract

import (
	"errors"
)

// This block holds common errors that might happen within repository.
var (
	ErrRecordNotFound = errors.New("record_not_found")
	ErrDuplicateEntry = errors.New("duplicate_entry")
	ErrEmptyResult    = errors.New("empty_result")
)
