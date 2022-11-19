// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	bignum "bignum-service/domain/bignum"
	ctxlib "bignum-service/lib/ctxlib"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteNum provides a mock function with given fields: ctx, name
func (_m *Repository) DeleteNum(ctx ctxlib.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctxlib.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetNum provides a mock function with given fields: ctx, name
func (_m *Repository) GetNum(ctx ctxlib.Context, name string) (*bignum.BigNum, error) {
	ret := _m.Called(ctx, name)

	var r0 *bignum.BigNum
	if rf, ok := ret.Get(0).(func(ctxlib.Context, string) *bignum.BigNum); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bignum.BigNum)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ctxlib.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutNum provides a mock function with given fields: ctx, num
func (_m *Repository) PutNum(ctx ctxlib.Context, num *bignum.BigNum) error {
	ret := _m.Called(ctx, num)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctxlib.Context, *bignum.BigNum) error); ok {
		r0 = rf(ctx, num)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateNum provides a mock function with given fields: ctx, num
func (_m *Repository) UpdateNum(ctx ctxlib.Context, num *bignum.BigNum) error {
	ret := _m.Called(ctx, num)

	var r0 error
	if rf, ok := ret.Get(0).(func(ctxlib.Context, *bignum.BigNum) error); ok {
		r0 = rf(ctx, num)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}