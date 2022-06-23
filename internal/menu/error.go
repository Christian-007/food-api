package menu

import "errors"

var (
	ErrNotFound = errors.New("menu not found")
	ErrDuplicateIdentifier = errors.New("menu duplicate identifier")
)

