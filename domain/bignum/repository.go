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
	// GetNum retrieves a number object with provided name from the storage.
	// It will return error if the object does not exist
	GetNum(ctx ctxlib.Context, name string) (num *BigNum, err error)
	// PutNum puts a number object to the storage
	PutNum(ctx ctxlib.Context, num *BigNum) (err error)
	// UpdateNum updates a number object in the storage
	// It will return error if the object does not exist
	UpdateNum(ctx ctxlib.Context, num *BigNum) (err error)
	// DeleteNum deletes a number object with provided name from the storage
	// It will return error if the object does not exist
	DeleteNum(ctx ctxlib.Context, name string) (err error)
}
