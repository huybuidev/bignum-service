package numobj

import (
	"bignum-service/domain/bignum"
	"bignum-service/domain/bignum/bignummem"
	"bignum-service/lib/ctxlib"
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

// getNumberValue receives a string and first check if it is a valid big float number.
// If yes -> returns the number in *big.Float type
// If no  -> consider this string is the name of a number object.
//
//		 It will use bignum repo to retrieve the value of this number object and return.
//	  If no number object found, return error
func (s *Service) getNumberValue(ctx ctxlib.Context, numStr string) (floatValue *big.Float, err error) {
	floatValue, err = bignum.ParseFloat(numStr)
	if err != nil {
		// Invalid big float number, consider it as the name of an number object
		bigNum, err := s.GetNum(ctx, numStr)
		if err != nil {
			return nil, ErrNumberObjectNotFound
		}
		floatValue = bigNum.Value()
	}
	return floatValue, nil
}

func (s *Service) compute(ctx ctxlib.Context, num1Str, num2Str string, computeFunc func(*big.Float, *big.Float) *big.Float) (result *big.Float, err error) {
	lg := ctx.Logger.With().Str("num1Str", num1Str).Str("num2Str", num2Str).Logger()
	if num1Str == "" || num2Str == "" {
		lg.Error().Msg("empty input")
		return nil, ErrEmptyInput
	}

	numFloat1, err := s.getNumberValue(ctx, num1Str)
	if err != nil {
		lg.Err(err).Msg("could not get value for num1Str")
		return nil, err
	}
	numFloat2, err := s.getNumberValue(ctx, num2Str)
	if err != nil {
		lg.Err(err).Msg("could not get value for num2Str")
		return nil, err
	}

	result = computeFunc(numFloat1, numFloat2)

	lg.Info().Interface("result", result).Msg("track result")

	return result, nil
}

// Add receives two params which could be a number object with provided name or a big number value,
// then return the sum result
func (s *Service) Add(ctx ctxlib.Context, num1Str, num2Str string) (result *big.Float, err error) {
	return s.compute(ctx, num1Str, num2Str, new(big.Float).Add)
}

// Subtract receives two params which could be a number object with provided name or a big number value,
// then return the subtracted result
func (s *Service) Subtract(ctx ctxlib.Context, num1Str, num2Str string) (result *big.Float, err error) {
	return s.compute(ctx, num1Str, num2Str, new(big.Float).Sub)
}

// Multiply receives two params which could be a number object with provided name or a big number value,
// then return the multiplied result
func (s *Service) Multiply(ctx ctxlib.Context, num1Str, num2Str string) (result *big.Float, err error) {
	return s.compute(ctx, num1Str, num2Str, new(big.Float).Mul)
}

// Divide receives two params which could be a number object with provided name or a big number value,
// then return the divided result
func (s *Service) Divide(ctx ctxlib.Context, num1Str, num2Str string) (result *big.Float, err error) {
	return s.compute(ctx, num1Str, num2Str, new(big.Float).Quo)
}
