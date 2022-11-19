package numobj

import "errors"

var (
	ErrNotImpl              = errors.New("not implemented yet")
	ErrEmptyName            = errors.New("empty name")
	ErrEmptyInput           = errors.New("empty input")
	ErrInvalidBigNumber     = errors.New("invalid big number")
	ErrNumberObjectNotFound = errors.New("number object not found")
)
