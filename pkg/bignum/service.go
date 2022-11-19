package bignum

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/rs/zerolog/log"
)

// BigNumService is the RPC service placeholder for the Big Number Service
type BigNumService struct{}

func (bns *BigNumService) Create(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "create").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(params) < 2 {
		return errors.New("value must be provided")
	}

	a := &big.Float{}
	_, ok := a.SetString(params[1])
	if !ok {
		return fmt.Errorf("invalid big float number %s", params[1])
	}

	*reply = "created"
	return nil
}

func (bns *BigNumService) Update(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "update").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(params) < 2 {
		return errors.New("value must be provided")
	}

	a := &big.Float{}
	_, ok := a.SetString(params[1])
	if !ok {
		return fmt.Errorf("invalid big float number %s", params[1])
	}

	*reply = "updated"
	return nil
}

func (bns *BigNumService) Delete(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "delete").Msg("")

	// First param is the number object's name
	if len(params) < 1 {
		return errors.New("number object name must be provided")
	}

	*reply = "deleted"
	return nil
}

func (bns *BigNumService) Multiply(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "multiply").Msg("")

	if len(params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}

	*reply = "multiplied"
	return nil
}