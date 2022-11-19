package bignumrpc

import (
	"bignum-service/lib/ctxlib"
	"bignum-service/lib/kjsonrpc"
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

func (bns *BigNumRPCService) Create(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "create").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(methodParams.Params) < 2 {
		return errors.New("value must be provided")
	}

	name := methodParams.Params[0]
	valueStr := methodParams.Params[1]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.PutNum(ctx, name, valueStr)
	if err != nil {
		return err
	}

	// No error
	*reply = "created"
	return nil
}

func (bns *BigNumRPCService) Update(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "update").Msg("")

	// First param is the number object's name
	// The number object's value must be provided (the second param)
	if len(methodParams.Params) < 2 {
		return errors.New("value must be provided")
	}

	name := methodParams.Params[0]
	valueStr := methodParams.Params[1]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.UpdateNum(ctx, name, valueStr)
	if err != nil {
		return err
	}

	// No error
	*reply = "updated"
	return nil
}

func (bns *BigNumRPCService) Delete(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "delete").Msg("")

	// First param is the number object's name
	if len(methodParams.Params) < 1 {
		return errors.New("number object name must be provided")
	}

	name := methodParams.Params[0]

	ctx := ctxlib.Background()
	err := bns.numObjSvc.DeleteNum(ctx, name)
	if err != nil {
		return err
	}

	// No error
	*reply = "deleted"
	return nil
}

func (bns *BigNumRPCService) Add(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "multiply").Msg("")

	if len(methodParams.Params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}
	num1Str := methodParams.Params[0]
	num2Str := methodParams.Params[1]

	ctx := ctxlib.Background()
	result, err := bns.numObjSvc.Add(ctx, num1Str, num2Str)
	if err != nil {
		return err
	}
	*reply = result.String()
	return nil
}

func (bns *BigNumRPCService) Subtract(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "multiply").Msg("")

	if len(methodParams.Params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}
	num1Str := methodParams.Params[0]
	num2Str := methodParams.Params[1]

	ctx := ctxlib.Background()
	result, err := bns.numObjSvc.Subtract(ctx, num1Str, num2Str)
	if err != nil {
		return err
	}
	*reply = result.String()
	return nil
}

func (bns *BigNumRPCService) Multiply(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "multiply").Msg("")

	if len(methodParams.Params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}
	num1Str := methodParams.Params[0]
	num2Str := methodParams.Params[1]

	ctx := ctxlib.Background()
	result, err := bns.numObjSvc.Multiply(ctx, num1Str, num2Str)
	if err != nil {
		return err
	}
	*reply = result.String()
	return nil
}

func (bns *BigNumRPCService) Divide(methodParams *kjsonrpc.MethodParam, reply *string) error {
	log.Debug().Interface("methodParams", methodParams).Str("method", "multiply").Msg("")

	if len(methodParams.Params) < 2 {
		return errors.New("the name of two number objects must be provided")
	}
	num1Str := methodParams.Params[0]
	num2Str := methodParams.Params[1]

	ctx := ctxlib.Background()
	result, err := bns.numObjSvc.Divide(ctx, num1Str, num2Str)
	if err != nil {
		return err
	}
	*reply = result.String()
	return nil
}
