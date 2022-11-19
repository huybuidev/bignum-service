package bignum

import (
	"errors"

	"bignum-service/lib/ctxlib"
)

var (
	// ErrNotImpl -
	ErrNotImpl = errors.New("not implemented yet")
)

// Repository interfaces for ac aggregate
type Repository interface {
	// GetNum retrieves a number object with provided name from the storage
	GetNum(ctx ctxlib.Context, name string) (num *BigNum, err error)
	// PutNum puts a number object to the storage
	PutNum(ctx ctxlib.Context, num *BigNum) (err error)
	// DeleteNum deletes a number object with provided name from the storage
	DeleteNum(ctx ctxlib.Context, name string) (err error)
}
