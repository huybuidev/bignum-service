package bignumrpc

import (
	"bignum-service/lib/ctxlib"
	"bignum-service/service/numobj"
	"errors"

	"github.com/rs/zerolog/log"
)

// BigNumRPCService is the RPC service placeholder for the Big Number Service
type BigNumRPCService struct {
	numObjSvc *numobj.Service
}

func New(numObjSvc *numobj.Service) *BigNumRPCService {
	return &BigNumRPCService{
		numObjSvc: numObjSvc,
	}
}

func (bns *BigNumRPCService) Create(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "create").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(params) < 2 {
		return errors.New("value must be provided")
	}

	name := params[0]
	valueStr := params[1]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.PutNum(ctx, name, valueStr)
	if err != nil {
		return err
	}

	// No error
	*reply = "created"
	return nil
}

func (bns *BigNumRPCService) Update(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "update").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(params) < 2 {
		return errors.New("value must be provided")
	}

	name := params[0]
	valueStr := params[1]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.UpdateNum(ctx, name, valueStr)
	if err != nil {
		return err
	}

	// No error
	*reply = "updated"
	return nil
}

func (bns *BigNumRPCService) Delete(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "delete").Msg("")

	// First param is the number object's name
	if len(params) < 1 {
		return errors.New("number object name must be provided")
	}

	name := params[0]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.DeleteNum(ctx, name)
	if err != nil {
		return err
	}

	// No error
	*reply = "deleted"
	return nil
}

func (bns *BigNumRPCService) Multiply(params []string, reply *string) error {
	log.Debug().Interface("params", params).Str("method", "multiply").Msg("")

	if len(params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}
	num1Str := params[0]
	num2Str := params[1]

	ctx := ctxlib.Background()
	result, err := bns.numObjSvc.Multiply(ctx, num1Str, num2Str)
	if err != nil {
		return err
	}
	*reply = result.String()
	return nil
}
