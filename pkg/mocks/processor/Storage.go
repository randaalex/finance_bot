// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	db "github.com/randaalex/finance_bot/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// CreateMapping provides a mock function with given fields: ctx, arg
func (_m *Storage) CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Mapping
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateMappingParams) db.Mapping); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Mapping)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, db.CreateMappingParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMappingByTransactionHash provides a mock function with given fields: ctx, transactionHash
func (_m *Storage) GetMappingByTransactionHash(ctx context.Context, transactionHash string) (db.Mapping, error) {
	ret := _m.Called(ctx, transactionHash)

	var r0 db.Mapping
	if rf, ok := ret.Get(0).(func(context.Context, string) db.Mapping); ok {
		r0 = rf(ctx, transactionHash)
	} else {
		r0 = ret.Get(0).(db.Mapping)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, transactionHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
