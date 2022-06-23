package restaurant

import "errors"

var (
	ErrNotFound = errors.New("restaurant not found")
	ErrDuplicateIdentifier = errors.New("restaurant duplicate identifier")
)

