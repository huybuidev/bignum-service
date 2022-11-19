package bignummem

import (
	"bignum-service/domain/bignum"
	"bignum-service/lib/ctxlib"
	"math/big"
)

type BignumMemoryRepo struct {
	numObjects map[string]*big.Float
}

func NewBignumMemoryRepository() *BignumMemoryRepo {
	return &BignumMemoryRepo{
		numObjects: make(map[string]*big.Float),
	}
}

func (r *BignumMemoryRepo) GetNum(ctx ctxlib.Context, name string) (num *bignum.BigNum, err error) {
	if floatValue, exist := r.numObjects[name]; exist {
		return bignum.New(name, floatValue)
	}
	return nil, ErrNotFound
}

func (r *BignumMemoryRepo) PutNum(ctx ctxlib.Context, num *bignum.BigNum) (err error) {
	r.numObjects[num.Name()] = num.Value()
	return nil
}

func (r *BignumMemoryRepo) UpdateNum(ctx ctxlib.Context, num *bignum.BigNum) (err error) {
	if _, exist := r.numObjects[num.Name()]; !exist {
		return ErrNotFound
	}
	r.numObjects[num.Name()] = num.Value()
	return nil
}

func (r *BignumMemoryRepo) DeleteNum(ctx ctxlib.Context, name string) (err error) {
	if _, exist := r.numObjects[name]; !exist {
		return ErrNotFound
	}
	delete(r.numObjects, name)
	return nil
}
