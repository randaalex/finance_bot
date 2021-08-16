# \CurrenciesApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DefaultCurrency**](CurrenciesApi.md#DefaultCurrency) | **Post** /api/v1/currencies/{code}/default | Make currency default currency.
[**DeleteCurrency**](CurrenciesApi.md#DeleteCurrency) | **Delete** /api/v1/currencies/{code} | Delete a currency.
[**DisableCurrency**](CurrenciesApi.md#DisableCurrency) | **Post** /api/v1/currencies/{code}/disable | Disable a currency.
[**EnableCurrency**](CurrenciesApi.md#EnableCurrency) | **Post** /api/v1/currencies/{code}/enable | Enable a single currency.
[**GetCurrency**](CurrenciesApi.md#GetCurrency) | **Get** /api/v1/currencies/{code} | Get a single currency.
[**GetDefaultCurrency**](CurrenciesApi.md#GetDefaultCurrency) | **Get** /api/v1/currencies/default | Get the user&#39;s default currency.
[**ListAccountByCurrency**](CurrenciesApi.md#ListAccountByCurrency) | **Get** /api/v1/currencies/{code}/accounts | List all accounts with this currency.
[**ListAvailableBudgetByCurrency**](CurrenciesApi.md#ListAvailableBudgetByCurrency) | **Get** /api/v1/currencies/{code}/available_budgets | List all available budgets with this currency.
[**ListBillByCurrency**](CurrenciesApi.md#ListBillByCurrency) | **Get** /api/v1/currencies/{code}/bills | List all bills with this currency.
[**ListBudgetLimitByCurrency**](CurrenciesApi.md#ListBudgetLimitByCurrency) | **Get** /api/v1/currencies/{code}/budget_limits | List all budget limits with this currency
[**ListCurrency**](CurrenciesApi.md#ListCurrency) | **Get** /api/v1/currencies | List all currencies.
[**ListRecurrenceByCurrency**](CurrenciesApi.md#ListRecurrenceByCurrency) | **Get** /api/v1/currencies/{code}/recurrences | List all recurring transactions with this currency.
[**ListRuleByCurrency**](CurrenciesApi.md#ListRuleByCurrency) | **Get** /api/v1/currencies/{code}/rules | List all rules with this currency.
[**ListTransactionByCurrency**](CurrenciesApi.md#ListTransactionByCurrency) | **Get** /api/v1/currencies/{code}/transactions | List all transactions with this currency.
[**StoreCurrency**](CurrenciesApi.md#StoreCurrency) | **Post** /api/v1/currencies | Store a new currency
[**UpdateCurrency**](CurrenciesApi.md#UpdateCurrency) | **Put** /api/v1/currencies/{code} | Update existing currency.



## DefaultCurrency

> CurrencySingle DefaultCurrency(ctx, code).Execute()

Make currency default currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.DefaultCurrency(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.DefaultCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DefaultCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.DefaultCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDefaultCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCurrency

> DeleteCurrency(ctx, code).Execute()

Delete a currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "GBP" // string | The currency code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.DeleteCurrency(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.DeleteCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableCurrency

> CurrencySingle DisableCurrency(ctx, code).Execute()

Disable a currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := int32(56) // int32 | The currency code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.DisableCurrency(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.DisableCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.DisableCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **int32** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCurrency

> CurrencySingle EnableCurrency(ctx, code).Execute()

Enable a single currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.EnableCurrency(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.EnableCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.EnableCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrency

> CurrencySingle GetCurrency(ctx, code).Execute()

Get a single currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.GetCurrency(context.Background(), code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultCurrency

> CurrencySingle GetDefaultCurrency(ctx).Execute()

Get the user's default currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.GetDefaultCurrency(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.GetDefaultCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.GetDefaultCurrency`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultCurrencyRequest struct via the builder pattern


### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountByCurrency

> AccountArray ListAccountByCurrency(ctx, code).Page(page).Date(date).Type_(type_).Execute()

List all accounts with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    date := time.Now() // string | A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account's balance on that day.  (optional)
    type_ := openapiclient.AccountTypeFilter("all") // AccountTypeFilter | Optional filter on the account type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListAccountByCurrency(context.Background(), code).Page(page).Date(date).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListAccountByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountByCurrency`: AccountArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListAccountByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 
 **date** | **string** | A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account&#39;s balance on that day.  | 
 **type_** | [**AccountTypeFilter**](AccountTypeFilter.md) | Optional filter on the account type(s) returned | 

### Return type

[**AccountArray**](AccountArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableBudgetByCurrency

> AvailableBudgetArray ListAvailableBudgetByCurrency(ctx, code).Page(page).Execute()

List all available budgets with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "EUR" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is 50 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListAvailableBudgetByCurrency(context.Background(), code).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListAvailableBudgetByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableBudgetByCurrency`: AvailableBudgetArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListAvailableBudgetByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableBudgetByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50 | 

### Return type

[**AvailableBudgetArray**](AvailableBudgetArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillByCurrency

> BillArray ListBillByCurrency(ctx, code).Page(page).Execute()

List all bills with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListBillByCurrency(context.Background(), code).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListBillByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillByCurrency`: BillArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListBillByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBillByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**BillArray**](BillArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBudgetLimitByCurrency

> BudgetLimitArray ListBudgetLimitByCurrency(ctx, code).Page(page).Start(start).End(end).Execute()

List all budget limits with this currency



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | Start date for the budget limit list. (optional)
    end := time.Now() // string | End date for the budget limit list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListBudgetLimitByCurrency(context.Background(), code).Page(page).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListBudgetLimitByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBudgetLimitByCurrency`: BudgetLimitArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListBudgetLimitByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBudgetLimitByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | Start date for the budget limit list. | 
 **end** | **string** | End date for the budget limit list. | 

### Return type

[**BudgetLimitArray**](BudgetLimitArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCurrency

> CurrencyArray ListCurrency(ctx).Page(page).Execute()

List all currencies.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListCurrency(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCurrency`: CurrencyArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListCurrency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**CurrencyArray**](CurrencyArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecurrenceByCurrency

> RecurrenceArray ListRecurrenceByCurrency(ctx, code).Page(page).Execute()

List all recurring transactions with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "EUR" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListRecurrenceByCurrency(context.Background(), code).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListRecurrenceByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecurrenceByCurrency`: RecurrenceArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListRecurrenceByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRecurrenceByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**RecurrenceArray**](RecurrenceArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuleByCurrency

> RuleArray ListRuleByCurrency(ctx, code).Page(page).Execute()

List all rules with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination per 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListRuleByCurrency(context.Background(), code).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListRuleByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRuleByCurrency`: RuleArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListRuleByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRuleByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination per 50. | 

### Return type

[**RuleArray**](RuleArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByCurrency

> TransactionArray ListTransactionByCurrency(ctx, code).Page(page).Start(start).End(end).Type_(type_).Execute()

List all transactions with this currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    code := "USD" // string | The currency code.
    page := int32(1) // int32 | Page number. The default pagination is per 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the list of transactions.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the list of transactions.  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.ListTransactionByCurrency(context.Background(), code).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.ListTransactionByCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByCurrency`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.ListTransactionByCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is per 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the list of transactions.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the list of transactions.  | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreCurrency

> CurrencySingle StoreCurrency(ctx).CurrencyStore(currencyStore).Execute()

Store a new currency



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    currencyStore := *openapiclient.NewCurrencyStore("AMS", "Ankh-Morpork dollar", "AM$") // CurrencyStore | JSON array or key=value pairs with the necessary currency information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.StoreCurrency(context.Background()).CurrencyStore(currencyStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.StoreCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.StoreCurrency`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencyStore** | [**CurrencyStore**](CurrencyStore.md) | JSON array or key&#x3D;value pairs with the necessary currency information. See the model for the exact specifications. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrency

> CurrencySingle UpdateCurrency(ctx, code).CurrencyUpdate(currencyUpdate).Execute()

Update existing currency.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    code := "EUR" // string | The currency code.
    currencyUpdate := *openapiclient.NewCurrencyUpdate() // CurrencyUpdate | JSON array with updated currency information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CurrenciesApi.UpdateCurrency(context.Background(), code).CurrencyUpdate(currencyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CurrenciesApi.UpdateCurrency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCurrency`: CurrencySingle
    fmt.Fprintf(os.Stdout, "Response from `CurrenciesApi.UpdateCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | The currency code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencyUpdate** | [**CurrencyUpdate**](CurrencyUpdate.md) | JSON array with updated currency information. See the model for the exact specifications. | 

### Return type

[**CurrencySingle**](CurrencySingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

