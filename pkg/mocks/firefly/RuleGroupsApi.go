// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"
)

// RuleGroupsApi is an autogenerated mock type for the RuleGroupsApi type
type RuleGroupsApi struct {
	mock.Mock
}

// DeleteRuleGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) DeleteRuleGroup(ctx context.Context, id int32) firefly.ApiDeleteRuleGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiDeleteRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDeleteRuleGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeleteRuleGroupRequest)
	}

	return r0
}

// DeleteRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) DeleteRuleGroupExecute(r firefly.ApiDeleteRuleGroupRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeleteRuleGroupRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeleteRuleGroupRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FireRuleGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) FireRuleGroup(ctx context.Context, id int32) firefly.ApiFireRuleGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiFireRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiFireRuleGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiFireRuleGroupRequest)
	}

	return r0
}

// FireRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) FireRuleGroupExecute(r firefly.ApiFireRuleGroupRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiFireRuleGroupRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiFireRuleGroupRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuleGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) GetRuleGroup(ctx context.Context, id int32) firefly.ApiGetRuleGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiGetRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiGetRuleGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetRuleGroupRequest)
	}

	return r0
}

// GetRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) GetRuleGroupExecute(r firefly.ApiGetRuleGroupRequest) (firefly.RuleGroupSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleGroupSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetRuleGroupRequest) firefly.RuleGroupSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleGroupSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetRuleGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetRuleGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRuleByGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) ListRuleByGroup(ctx context.Context, id int32) firefly.ApiListRuleByGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiListRuleByGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiListRuleByGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiListRuleByGroupRequest)
	}

	return r0
}

// ListRuleByGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) ListRuleByGroupExecute(r firefly.ApiListRuleByGroupRequest) (firefly.RuleArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListRuleByGroupRequest) firefly.RuleArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListRuleByGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListRuleByGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRuleGroup provides a mock function with given fields: ctx
func (_m *RuleGroupsApi) ListRuleGroup(ctx context.Context) firefly.ApiListRuleGroupRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListRuleGroupRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListRuleGroupRequest)
	}

	return r0
}

// ListRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) ListRuleGroupExecute(r firefly.ApiListRuleGroupRequest) (firefly.RuleGroupArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleGroupArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListRuleGroupRequest) firefly.RuleGroupArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleGroupArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListRuleGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListRuleGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreRuleGroup provides a mock function with given fields: ctx
func (_m *RuleGroupsApi) StoreRuleGroup(ctx context.Context) firefly.ApiStoreRuleGroupRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStoreRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStoreRuleGroupRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStoreRuleGroupRequest)
	}

	return r0
}

// StoreRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) StoreRuleGroupExecute(r firefly.ApiStoreRuleGroupRequest) (firefly.RuleGroupSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleGroupSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStoreRuleGroupRequest) firefly.RuleGroupSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleGroupSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStoreRuleGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStoreRuleGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TestRuleGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) TestRuleGroup(ctx context.Context, id int32) firefly.ApiTestRuleGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiTestRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiTestRuleGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiTestRuleGroupRequest)
	}

	return r0
}

// TestRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) TestRuleGroupExecute(r firefly.ApiTestRuleGroupRequest) (firefly.TransactionArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionArray
	if rf, ok := ret.Get(0).(func(firefly.ApiTestRuleGroupRequest) firefly.TransactionArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiTestRuleGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiTestRuleGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateRuleGroup provides a mock function with given fields: ctx, id
func (_m *RuleGroupsApi) UpdateRuleGroup(ctx context.Context, id int32) firefly.ApiUpdateRuleGroupRequest {
	ret := _m.Called(ctx, id)

	var r0 firefly.ApiUpdateRuleGroupRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiUpdateRuleGroupRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdateRuleGroupRequest)
	}

	return r0
}

// UpdateRuleGroupExecute provides a mock function with given fields: r
func (_m *RuleGroupsApi) UpdateRuleGroupExecute(r firefly.ApiUpdateRuleGroupRequest) (firefly.RuleGroupSingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleGroupSingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdateRuleGroupRequest) firefly.RuleGroupSingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleGroupSingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdateRuleGroupRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdateRuleGroupRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
