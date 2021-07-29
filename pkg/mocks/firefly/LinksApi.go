// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"
)

// LinksApi is an autogenerated mock type for the LinksApi type
type LinksApi struct {
	mock.Mock
}

// DeleteLinkType provides a mock function with given fields: ctx, id
func (_m *LinksApi) DeleteLinkType(ctx context.Context, id int32) firefly.ApiDeleteLinkTypeRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDeleteLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDeleteLinkTypeRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeleteLinkTypeRequest)
	}

	return r0
}

// DeleteLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) DeleteLinkTypeExecute(r firefly.ApiDeleteLinkTypeRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeleteLinkTypeRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeleteLinkTypeRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTransactionLink provides a mock function with given fields: ctx, id
func (_m *LinksApi) DeleteTransactionLink(ctx context.Context, id int32) firefly.ApiDeleteTransactionLinkRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDeleteTransactionLinkRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDeleteTransactionLinkRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeleteTransactionLinkRequest)
	}

	return r0
}

// DeleteTransactionLinkExecute provides a mock function with given fields: r
func (_m *LinksApi) DeleteTransactionLinkExecute(r firefly.ApiDeleteTransactionLinkRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeleteTransactionLinkRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeleteTransactionLinkRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLinkType provides a mock function with given fields: ctx, id
func (_m *LinksApi) GetLinkType(ctx context.Context, id int32) firefly.ApiGetLinkTypeRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiGetLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiGetLinkTypeRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetLinkTypeRequest)
	}

	return r0
}

// GetLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) GetLinkTypeExecute(r firefly.ApiGetLinkTypeRequest) (firefly.LinkTypeSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.LinkTypeSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetLinkTypeRequest) firefly.LinkTypeSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.LinkTypeSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetLinkTypeRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetLinkTypeRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTransactionLink provides a mock function with given fields: ctx, id
func (_m *LinksApi) GetTransactionLink(ctx context.Context, id int32) firefly.ApiGetTransactionLinkRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiGetTransactionLinkRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiGetTransactionLinkRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetTransactionLinkRequest)
	}

	return r0
}

// GetTransactionLinkExecute provides a mock function with given fields: r
func (_m *LinksApi) GetTransactionLinkExecute(r firefly.ApiGetTransactionLinkRequest) (firefly.TransactionLinkSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionLinkSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetTransactionLinkRequest) firefly.TransactionLinkSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionLinkSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetTransactionLinkRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetTransactionLinkRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListLinkType provides a mock function with given fields: ctx
func (_m *LinksApi) ListLinkType(ctx context.Context) firefly.ApiListLinkTypeRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListLinkTypeRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListLinkTypeRequest)
	}

	return r0
}

// ListLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) ListLinkTypeExecute(r firefly.ApiListLinkTypeRequest) (firefly.LinkTypeArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.LinkTypeArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListLinkTypeRequest) firefly.LinkTypeArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.LinkTypeArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListLinkTypeRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListLinkTypeRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactionByLinkType provides a mock function with given fields: ctx, id
func (_m *LinksApi) ListTransactionByLinkType(ctx context.Context, id int32) firefly.ApiListTransactionByLinkTypeRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiListTransactionByLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiListTransactionByLinkTypeRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiListTransactionByLinkTypeRequest)
	}

	return r0
}

// ListTransactionByLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) ListTransactionByLinkTypeExecute(r firefly.ApiListTransactionByLinkTypeRequest) (firefly.TransactionArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListTransactionByLinkTypeRequest) firefly.TransactionArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListTransactionByLinkTypeRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListTransactionByLinkTypeRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactionLink provides a mock function with given fields: ctx
func (_m *LinksApi) ListTransactionLink(ctx context.Context) firefly.ApiListTransactionLinkRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListTransactionLinkRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListTransactionLinkRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListTransactionLinkRequest)
	}

	return r0
}

// ListTransactionLinkExecute provides a mock function with given fields: r
func (_m *LinksApi) ListTransactionLinkExecute(r firefly.ApiListTransactionLinkRequest) (firefly.TransactionLinkArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionLinkArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListTransactionLinkRequest) firefly.TransactionLinkArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionLinkArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListTransactionLinkRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListTransactionLinkRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreLinkType provides a mock function with given fields: ctx
func (_m *LinksApi) StoreLinkType(ctx context.Context) firefly.ApiStoreLinkTypeRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStoreLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStoreLinkTypeRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStoreLinkTypeRequest)
	}

	return r0
}

// StoreLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) StoreLinkTypeExecute(r firefly.ApiStoreLinkTypeRequest) (firefly.LinkTypeSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.LinkTypeSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStoreLinkTypeRequest) firefly.LinkTypeSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.LinkTypeSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStoreLinkTypeRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStoreLinkTypeRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreTransactionLink provides a mock function with given fields: ctx
func (_m *LinksApi) StoreTransactionLink(ctx context.Context) firefly.ApiStoreTransactionLinkRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStoreTransactionLinkRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStoreTransactionLinkRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStoreTransactionLinkRequest)
	}

	return r0
}

// StoreTransactionLinkExecute provides a mock function with given fields: r
func (_m *LinksApi) StoreTransactionLinkExecute(r firefly.ApiStoreTransactionLinkRequest) (firefly.TransactionLinkSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionLinkSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStoreTransactionLinkRequest) firefly.TransactionLinkSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionLinkSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStoreTransactionLinkRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStoreTransactionLinkRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateLinkType provides a mock function with given fields: ctx, id
func (_m *LinksApi) UpdateLinkType(ctx context.Context, id int32) firefly.ApiUpdateLinkTypeRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUpdateLinkTypeRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUpdateLinkTypeRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdateLinkTypeRequest)
	}

	return r0
}

// UpdateLinkTypeExecute provides a mock function with given fields: r
func (_m *LinksApi) UpdateLinkTypeExecute(r firefly.ApiUpdateLinkTypeRequest) (firefly.LinkTypeSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.LinkTypeSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdateLinkTypeRequest) firefly.LinkTypeSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.LinkTypeSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdateLinkTypeRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdateLinkTypeRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTransactionLink provides a mock function with given fields: ctx, id
func (_m *LinksApi) UpdateTransactionLink(ctx context.Context, id int32) firefly.ApiUpdateTransactionLinkRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUpdateTransactionLinkRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUpdateTransactionLinkRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdateTransactionLinkRequest)
	}

	return r0
}

// UpdateTransactionLinkExecute provides a mock function with given fields: r
func (_m *LinksApi) UpdateTransactionLinkExecute(r firefly.ApiUpdateTransactionLinkRequest) (firefly.TransactionLinkSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionLinkSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdateTransactionLinkRequest) firefly.TransactionLinkSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionLinkSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdateTransactionLinkRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdateTransactionLinkRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
