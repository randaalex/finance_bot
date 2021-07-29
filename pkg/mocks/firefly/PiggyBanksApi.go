// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"
)

// PiggyBanksApi is an autogenerated mock type for the PiggyBanksApi type
type PiggyBanksApi struct {
	mock.Mock
}

// DeletePiggyBank provides a mock function with given fields: ctx, id
func (_m *PiggyBanksApi) DeletePiggyBank(ctx context.Context, id int32) firefly.ApiDeletePiggyBankRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDeletePiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDeletePiggyBankRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeletePiggyBankRequest)
	}

	return r0
}

// DeletePiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) DeletePiggyBankExecute(r firefly.ApiDeletePiggyBankRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeletePiggyBankRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeletePiggyBankRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPiggyBank provides a mock function with given fields: ctx, id
func (_m *PiggyBanksApi) GetPiggyBank(ctx context.Context, id int32) firefly.ApiGetPiggyBankRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiGetPiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiGetPiggyBankRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetPiggyBankRequest)
	}

	return r0
}

// GetPiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) GetPiggyBankExecute(r firefly.ApiGetPiggyBankRequest) (firefly.PiggyBankSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.PiggyBankSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetPiggyBankRequest) firefly.PiggyBankSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.PiggyBankSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetPiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetPiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAttachmentByPiggyBank provides a mock function with given fields: ctx, id
func (_m *PiggyBanksApi) ListAttachmentByPiggyBank(ctx context.Context, id int32) firefly.ApiListAttachmentByPiggyBankRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiListAttachmentByPiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiListAttachmentByPiggyBankRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiListAttachmentByPiggyBankRequest)
	}

	return r0
}

// ListAttachmentByPiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) ListAttachmentByPiggyBankExecute(r firefly.ApiListAttachmentByPiggyBankRequest) (firefly.AttachmentArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AttachmentArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListAttachmentByPiggyBankRequest) firefly.AttachmentArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AttachmentArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListAttachmentByPiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListAttachmentByPiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListEventByPiggyBank provides a mock function with given fields: ctx, id
func (_m *PiggyBanksApi) ListEventByPiggyBank(ctx context.Context, id int32) firefly.ApiListEventByPiggyBankRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiListEventByPiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiListEventByPiggyBankRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiListEventByPiggyBankRequest)
	}

	return r0
}

// ListEventByPiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) ListEventByPiggyBankExecute(r firefly.ApiListEventByPiggyBankRequest) (firefly.PiggyBankEventArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.PiggyBankEventArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListEventByPiggyBankRequest) firefly.PiggyBankEventArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.PiggyBankEventArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListEventByPiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListEventByPiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListPiggyBank provides a mock function with given fields: ctx
func (_m *PiggyBanksApi) ListPiggyBank(ctx context.Context) firefly.ApiListPiggyBankRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListPiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListPiggyBankRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListPiggyBankRequest)
	}

	return r0
}

// ListPiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) ListPiggyBankExecute(r firefly.ApiListPiggyBankRequest) (firefly.PiggyBankArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.PiggyBankArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListPiggyBankRequest) firefly.PiggyBankArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.PiggyBankArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListPiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListPiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StorePiggyBank provides a mock function with given fields: ctx
func (_m *PiggyBanksApi) StorePiggyBank(ctx context.Context) firefly.ApiStorePiggyBankRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStorePiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStorePiggyBankRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStorePiggyBankRequest)
	}

	return r0
}

// StorePiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) StorePiggyBankExecute(r firefly.ApiStorePiggyBankRequest) (firefly.PiggyBankSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.PiggyBankSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStorePiggyBankRequest) firefly.PiggyBankSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.PiggyBankSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStorePiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStorePiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePiggyBank provides a mock function with given fields: ctx, id
func (_m *PiggyBanksApi) UpdatePiggyBank(ctx context.Context, id int32) firefly.ApiUpdatePiggyBankRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUpdatePiggyBankRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUpdatePiggyBankRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdatePiggyBankRequest)
	}

	return r0
}

// UpdatePiggyBankExecute provides a mock function with given fields: r
func (_m *PiggyBanksApi) UpdatePiggyBankExecute(r firefly.ApiUpdatePiggyBankRequest) (firefly.PiggyBankSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.PiggyBankSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdatePiggyBankRequest) firefly.PiggyBankSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.PiggyBankSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdatePiggyBankRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdatePiggyBankRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}