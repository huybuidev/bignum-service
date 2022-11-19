package entity

import "math/big"

// NumObject represents a number object
// which has a unique name and the value is a big float number type
type NumObject struct {
	Name  *string    `json:"name"`
	Value *big.Float `json:"value"`
}
