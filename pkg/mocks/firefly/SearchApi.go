// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"
)

// SearchApi is an autogenerated mock type for the SearchApi type
type SearchApi struct {
	mock.Mock
}

// SearchAccounts provides a mock function with given fields: ctx
func (_m *SearchApi) SearchAccounts(ctx context.Context) firefly.ApiSearchAccountsRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiSearchAccountsRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiSearchAccountsRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiSearchAccountsRequest)
	}

	return r0
}

// SearchAccountsExecute provides a mock function with given fields: r
func (_m *SearchApi) SearchAccountsExecute(r firefly.ApiSearchAccountsRequest) (firefly.AccountArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AccountArray
	if rf, ok := ret.Get(0).(func(firefly.ApiSearchAccountsRequest) firefly.AccountArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AccountArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiSearchAccountsRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiSearchAccountsRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SearchTransactions provides a mock function with given fields: ctx
func (_m *SearchApi) SearchTransactions(ctx context.Context) firefly.ApiSearchTransactionsRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiSearchTransactionsRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiSearchTransactionsRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiSearchTransactionsRequest)
	}

	return r0
}

// SearchTransactionsExecute provides a mock function with given fields: r
func (_m *SearchApi) SearchTransactionsExecute(r firefly.ApiSearchTransactionsRequest) (firefly.TransactionArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionArray
	if rf, ok := ret.Get(0).(func(firefly.ApiSearchTransactionsRequest) firefly.TransactionArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiSearchTransactionsRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiSearchTransactionsRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
