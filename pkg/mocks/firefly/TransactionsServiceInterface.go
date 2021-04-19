// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"
	mock "github.com/stretchr/testify/mock"
)

// TransactionsServiceInterface is an autogenerated mock type for the TransactionsServiceInterface type
type TransactionsServiceInterface struct {
	mock.Mock
}

// CreateTransaction provides a mock function with given fields: ctx, request
func (_m *TransactionsServiceInterface) CreateTransaction(ctx context.Context, request firefly.CreateTransactionReq) (*firefly.Transaction, *firefly.Response, error) {
	ret := _m.Called(ctx, request)

	var r0 *firefly.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, firefly.CreateTransactionReq) *firefly.Transaction); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*firefly.Transaction)
		}
	}

	var r1 *firefly.Response
	if rf, ok := ret.Get(1).(func(context.Context, firefly.CreateTransactionReq) *firefly.Response); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*firefly.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, firefly.CreateTransactionReq) error); ok {
		r2 = rf(ctx, request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
