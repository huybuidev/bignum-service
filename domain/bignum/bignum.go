package bignum

import (
	"bignum-service/entity"
	"errors"
	"fmt"
	"math/big"
)

type BigNum struct {
	numObj *entity.NumObject
}

func NewBigNum(name, value string) (num *BigNum, err error) {
	if name == "" {
		return nil, errors.New("empty name")
	}

	// Check if value is a valid big float
	floatNum, err := ParseFloat(value)
	if err != nil {
		return nil, err
	}

	// Return number object
	numName := name
	return &BigNum{
		numObj: &entity.NumObject{
			Name:  &numName,
			Value: floatNum,
		},
	}, nil
}

func ParseFloat(value string) (floatNum *big.Float, err error) {
	floatNum = &big.Float{}
	_, ok := floatNum.SetString(value)
	if !ok {
		return nil, fmt.Errorf("invalid big float number %s", value)
	}
	return floatNum, nil
}

func (bn *BigNum) Value() *big.Float {
	if bn == nil || bn.numObj == nil {
		return nil
	}
	return bn.numObj.Value
}
