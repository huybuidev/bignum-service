package numobj

import (
	"bignum-service/domain/bignum"
	"bignum-service/domain/bignum/bignummem"
	"bignum-service/lib/ctxlib"
	"fmt"
	"math/big"
)

// NumObjConfig is an alias for a function that will take in a pointer to an NumObj service and modify it
type NumObjConfig func(nos *Service) error

type Service struct {
	bignumRepo bignum.Repository
}

func NewService(cfgs ...NumObjConfig) (*Service, error) {
	// Create the service
	nos := &Service{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(nos)
		if err != nil {
			return nil, err
		}
	}
	return nos, nil
}

// WithBignumMemRepo applies a given BignumMemory repository to the service
func WithBignumMemRepo(bignumMemRepo *bignummem.BignumMemoryRepo) NumObjConfig {
	// return a function that matches the NumObjConfig alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(nos *Service) error {
		nos.bignumRepo = bignumMemRepo
		return nil
	}
}

// GetNum retrieves a number object with provided name from the storage
func (s *Service) GetNum(ctx ctxlib.Context, name string) (num *bignum.BigNum, err error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	return s.bignumRepo.GetNum(ctx, name)
}

// PutNum puts a number object to the storage
func (s *Service) PutNum(ctx ctxlib.Context, name string, value string) (err error) {
	if name == "" {
		return ErrEmptyName
	}

	// Check if value is a big number
	num, err := bignum.NewFromString(name, value)
	if err != nil {
		return ErrInvalidBigNumber
	}

	return s.bignumRepo.PutNum(ctx, num)
}

// UpdateNum updates a number object to the storage
func (s *Service) UpdateNum(ctx ctxlib.Context, name string, value string) (err error) {
	if name == "" {
		return ErrEmptyName
	}

	// Check if value is a big number
	num, err := bignum.NewFromString(name, value)
	if err != nil {
		return ErrInvalidBigNumber
	}

	return s.bignumRepo.UpdateNum(ctx, num)
}

// DeleteNum deletes a number object with provided name from the storage
func (s *Service) DeleteNum(ctx ctxlib.Context, name string) (err error) {
	if name == "" {
		return ErrEmptyName
	}
	return s.bignumRepo.DeleteNum(ctx, name)
}

// Multiply receives two params which could be a number object with provided name or a big number value,
// then return the multiplied result
func (s *Service) Multiply(ctx ctxlib.Context, num1Str, num2Str string) (result *big.Float, err error) {
	if num1Str == "" || num2Str == "" {
		ctx.Logger.Error().Str("num1Str", num1Str).Str("num2Str", num2Str).Msg("empty input")
		return nil, ErrEmptyInput
	}
	numFloat1, err := bignum.ParseFloat(num1Str)
	if err != nil {
		// Invalid big float number, consider it as the name of an number object
		num1, err := s.GetNum(ctx, num1Str)
		if err != nil {
			return nil, fmt.Errorf("first number object not found")
		}
		numFloat1 = num1.Value()
	}

	numFloat2, err := bignum.ParseFloat(num2Str)
	if err != nil {
		// Invalid big float number, consider it as the name of an number object
		num2, err := s.GetNum(ctx, num2Str)
		if err != nil {
			return nil, fmt.Errorf("second number object not found")
		}
		numFloat2 = num2.Value()
	}

	result = numFloat1.Mul(numFloat1, numFloat2)

	return result, nil
}
