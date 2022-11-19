package bignummem

import "errors"

var (
	ErrNotImpl      = errors.New("not implemented yet")
	ErrNotFound     = errors.New("number object not found")
	ErrAlreadyExist = errors.New("number object already exists")
)
