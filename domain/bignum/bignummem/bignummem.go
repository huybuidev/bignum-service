package bignummem

import (
	"bignum-service/domain/bignum"
	"bignum-service/lib/ctxlib"
)

type BignumMemoryRepo struct{}

func NewBignumMemoryRepository() *BignumMemoryRepo {
	return &BignumMemoryRepo{}
}

func (r *BignumMemoryRepo) GetNum(ctx ctxlib.Context, name string) (num *bignum.BigNum, err error) {
	return nil, ErrNotImpl
}

func (r *BignumMemoryRepo) PutNum(ctx ctxlib.Context, num *bignum.BigNum) (err error) {
	return ErrNotImpl
}

func (r *BignumMemoryRepo) DeleteNum(ctx ctxlib.Context, name string) (err error) {
	return ErrNotImpl
}
