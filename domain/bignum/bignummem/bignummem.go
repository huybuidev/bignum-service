package bignummem

import (
	"bignum-service/domain/bignum"
	"bignum-service/lib/ctxlib"
	"math/big"
	"sync"
)

type BignumMemoryRepo struct {
	numObjects map[string]*big.Float
	mutex      sync.Mutex
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
	// Just make sure the internal map is not nil, should not happen if using the factory function
	if r.numObjects == nil {
		r.mutex.Lock()
		r.numObjects = make(map[string]*big.Float)
		r.mutex.Unlock()
	}

	// Make sure the number object does not exist
	if _, exist := r.numObjects[num.Name()]; exist {
		return ErrAlreadyExist
	}

	r.mutex.Lock()
	r.numObjects[num.Name()] = num.Value()
	r.mutex.Unlock()

	return nil
}

func (r *BignumMemoryRepo) UpdateNum(ctx ctxlib.Context, num *bignum.BigNum) (err error) {
	if _, exist := r.numObjects[num.Name()]; !exist {
		return ErrNotFound
	}
	r.mutex.Lock()
	r.numObjects[num.Name()] = num.Value()
	r.mutex.Unlock()
	return nil
}

func (r *BignumMemoryRepo) DeleteNum(ctx ctxlib.Context, name string) (err error) {
	if _, exist := r.numObjects[name]; !exist {
		return ErrNotFound
	}
	delete(r.numObjects, name)
	return nil
}
