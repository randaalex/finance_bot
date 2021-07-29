// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	firefly "github.com/randaalex/finance_bot/pkg/firefly"

	mock "github.com/stretchr/testify/mock"
)

// CurrenciesApi is an autogenerated mock type for the CurrenciesApi type
type CurrenciesApi struct {
	mock.Mock
}

// DefaultCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) DefaultCurrency(ctx context.Context, code string) firefly.ApiDefaultCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiDefaultCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiDefaultCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiDefaultCurrencyRequest)
	}

	return r0
}

// DefaultCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) DefaultCurrencyExecute(r firefly.ApiDefaultCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiDefaultCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiDefaultCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiDefaultCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeleteCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) DeleteCurrency(ctx context.Context, code string) firefly.ApiDeleteCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiDeleteCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiDeleteCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiDeleteCurrencyRequest)
	}

	return r0
}

// DeleteCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) DeleteCurrencyExecute(r firefly.ApiDeleteCurrencyRequest) (*http.Response, error) {
	ret := _m.Called(r)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(firefly.ApiDeleteCurrencyRequest) *http.Response); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(firefly.ApiDeleteCurrencyRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) DisableCurrency(ctx context.Context, code int32) firefly.ApiDisableCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiDisableCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, int32) firefly.ApiDisableCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiDisableCurrencyRequest)
	}

	return r0
}

// DisableCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) DisableCurrencyExecute(r firefly.ApiDisableCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiDisableCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiDisableCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiDisableCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnableCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) EnableCurrency(ctx context.Context, code string) firefly.ApiEnableCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiEnableCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiEnableCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiEnableCurrencyRequest)
	}

	return r0
}

// EnableCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) EnableCurrencyExecute(r firefly.ApiEnableCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiEnableCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiEnableCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiEnableCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) GetCurrency(ctx context.Context, code string) firefly.ApiGetCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiGetCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiGetCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetCurrencyRequest)
	}

	return r0
}

// GetCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) GetCurrencyExecute(r firefly.ApiGetCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDefaultCurrency provides a mock function with given fields: ctx
func (_m *CurrenciesApi) GetDefaultCurrency(ctx context.Context) firefly.ApiGetDefaultCurrencyRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiGetDefaultCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiGetDefaultCurrencyRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiGetDefaultCurrencyRequest)
	}

	return r0
}

// GetDefaultCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) GetDefaultCurrencyExecute(r firefly.ApiGetDefaultCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiGetDefaultCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiGetDefaultCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiGetDefaultCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAccountByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListAccountByCurrency(ctx context.Context, code string) firefly.ApiListAccountByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListAccountByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListAccountByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListAccountByCurrencyRequest)
	}

	return r0
}

// ListAccountByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListAccountByCurrencyExecute(r firefly.ApiListAccountByCurrencyRequest) (firefly.AccountArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AccountArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListAccountByCurrencyRequest) firefly.AccountArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AccountArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListAccountByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListAccountByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListAvailableBudgetByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListAvailableBudgetByCurrency(ctx context.Context, code string) firefly.ApiListAvailableBudgetByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListAvailableBudgetByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListAvailableBudgetByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListAvailableBudgetByCurrencyRequest)
	}

	return r0
}

// ListAvailableBudgetByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListAvailableBudgetByCurrencyExecute(r firefly.ApiListAvailableBudgetByCurrencyRequest) (firefly.AvailableBudgetArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.AvailableBudgetArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListAvailableBudgetByCurrencyRequest) firefly.AvailableBudgetArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.AvailableBudgetArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListAvailableBudgetByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListAvailableBudgetByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListBillByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListBillByCurrency(ctx context.Context, code string) firefly.ApiListBillByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListBillByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListBillByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListBillByCurrencyRequest)
	}

	return r0
}

// ListBillByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListBillByCurrencyExecute(r firefly.ApiListBillByCurrencyRequest) (firefly.BillArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.BillArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListBillByCurrencyRequest) firefly.BillArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.BillArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListBillByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListBillByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListBudgetLimitByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListBudgetLimitByCurrency(ctx context.Context, code string) firefly.ApiListBudgetLimitByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListBudgetLimitByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListBudgetLimitByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListBudgetLimitByCurrencyRequest)
	}

	return r0
}

// ListBudgetLimitByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListBudgetLimitByCurrencyExecute(r firefly.ApiListBudgetLimitByCurrencyRequest) (firefly.BudgetLimitArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.BudgetLimitArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListBudgetLimitByCurrencyRequest) firefly.BudgetLimitArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.BudgetLimitArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListBudgetLimitByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListBudgetLimitByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListCurrency provides a mock function with given fields: ctx
func (_m *CurrenciesApi) ListCurrency(ctx context.Context) firefly.ApiListCurrencyRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiListCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiListCurrencyRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiListCurrencyRequest)
	}

	return r0
}

// ListCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListCurrencyExecute(r firefly.ApiListCurrencyRequest) (firefly.CurrencyArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencyArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListCurrencyRequest) firefly.CurrencyArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencyArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListExchangeRateByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListExchangeRateByCurrency(ctx context.Context, code string) firefly.ApiListExchangeRateByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListExchangeRateByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListExchangeRateByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListExchangeRateByCurrencyRequest)
	}

	return r0
}

// ListExchangeRateByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListExchangeRateByCurrencyExecute(r firefly.ApiListExchangeRateByCurrencyRequest) (firefly.ExchangeRateArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.ExchangeRateArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListExchangeRateByCurrencyRequest) firefly.ExchangeRateArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.ExchangeRateArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListExchangeRateByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListExchangeRateByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRecurrenceByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListRecurrenceByCurrency(ctx context.Context, code string) firefly.ApiListRecurrenceByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListRecurrenceByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListRecurrenceByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListRecurrenceByCurrencyRequest)
	}

	return r0
}

// ListRecurrenceByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListRecurrenceByCurrencyExecute(r firefly.ApiListRecurrenceByCurrencyRequest) (firefly.RecurrenceArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RecurrenceArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListRecurrenceByCurrencyRequest) firefly.RecurrenceArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RecurrenceArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListRecurrenceByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListRecurrenceByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListRuleByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListRuleByCurrency(ctx context.Context, code string) firefly.ApiListRuleByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListRuleByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListRuleByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListRuleByCurrencyRequest)
	}

	return r0
}

// ListRuleByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListRuleByCurrencyExecute(r firefly.ApiListRuleByCurrencyRequest) (firefly.RuleArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.RuleArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListRuleByCurrencyRequest) firefly.RuleArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.RuleArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListRuleByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListRuleByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListTransactionByCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) ListTransactionByCurrency(ctx context.Context, code string) firefly.ApiListTransactionByCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiListTransactionByCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiListTransactionByCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiListTransactionByCurrencyRequest)
	}

	return r0
}

// ListTransactionByCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) ListTransactionByCurrencyExecute(r firefly.ApiListTransactionByCurrencyRequest) (firefly.TransactionArray, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.TransactionArray
	if rf, ok := ret.Get(0).(func(firefly.ApiListTransactionByCurrencyRequest) firefly.TransactionArray); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.TransactionArray)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiListTransactionByCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiListTransactionByCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// StoreCurrency provides a mock function with given fields: ctx
func (_m *CurrenciesApi) StoreCurrency(ctx context.Context) firefly.ApiStoreCurrencyRequest {
	ret := _m.Called(ctx)

	var r0 firefly.ApiStoreCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context) firefly.ApiStoreCurrencyRequest); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(firefly.ApiStoreCurrencyRequest)
	}

	return r0
}

// StoreCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) StoreCurrencyExecute(r firefly.ApiStoreCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiStoreCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiStoreCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiStoreCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateCurrency provides a mock function with given fields: ctx, code
func (_m *CurrenciesApi) UpdateCurrency(ctx context.Context, code string) firefly.ApiUpdateCurrencyRequest {
	ret := _m.Called(ctx, code)

	var r0 firefly.ApiUpdateCurrencyRequest
	if rf, ok := ret.Get(0).(func(context.Context, string) firefly.ApiUpdateCurrencyRequest); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(firefly.ApiUpdateCurrencyRequest)
	}

	return r0
}

// UpdateCurrencyExecute provides a mock function with given fields: r
func (_m *CurrenciesApi) UpdateCurrencyExecute(r firefly.ApiUpdateCurrencyRequest) (firefly.CurrencySingle, *http.Response, error) {
	ret := _m.Called(r)

	var r0 firefly.CurrencySingle
	if rf, ok := ret.Get(0).(func(firefly.ApiUpdateCurrencyRequest) firefly.CurrencySingle); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Get(0).(firefly.CurrencySingle)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(firefly.ApiUpdateCurrencyRequest) *http.Response); ok {
		r1 = rf(r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(firefly.ApiUpdateCurrencyRequest) error); ok {
		r2 = rf(r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
