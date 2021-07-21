# \AutocompleteApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountsAC**](AutocompleteApi.md#GetAccountsAC) | **Get** /api/v1/autocomplete/accounts | All accounts of the user returned in a basic auto-complete array.
[**GetBillsAC**](AutocompleteApi.md#GetBillsAC) | **Get** /api/v1/autocomplete/bills | All bills of the user returned in a basic auto-complete array.
[**GetBudgetsAC**](AutocompleteApi.md#GetBudgetsAC) | **Get** /api/v1/autocomplete/budgets | All budgets of the user returned in a basic auto-complete array.
[**GetCategoriesAC**](AutocompleteApi.md#GetCategoriesAC) | **Get** /api/v1/autocomplete/categories | All categories of the user returned in a basic auto-complete array.
[**GetCurrenciesAC**](AutocompleteApi.md#GetCurrenciesAC) | **Get** /api/v1/autocomplete/currencies | All currencies of the user returned in a basic auto-complete array.
[**GetCurrenciesCodeAC**](AutocompleteApi.md#GetCurrenciesCodeAC) | **Get** /api/v1/autocomplete/currencies-with-code | All currencies of the user returned in a basic auto-complete array.
[**GetObjectGroupsAC**](AutocompleteApi.md#GetObjectGroupsAC) | **Get** /api/v1/autocomplete/object-groups | All object groups of the user returned in a basic auto-complete array.
[**GetPiggiesAC**](AutocompleteApi.md#GetPiggiesAC) | **Get** /api/v1/autocomplete/piggy-banks | All piggy banks of the user returned in a basic auto-complete array.
[**GetPiggiesBalanceAC**](AutocompleteApi.md#GetPiggiesBalanceAC) | **Get** /api/v1/autocomplete/piggy-banks-with-balance | All piggy banks of the user returned in a basic auto-complete array complemented with balance information.
[**GetRuleGroupsAC**](AutocompleteApi.md#GetRuleGroupsAC) | **Get** /api/v1/autocomplete/rule-groups | All rule groups of the user returned in a basic auto-complete array.
[**GetRulesAC**](AutocompleteApi.md#GetRulesAC) | **Get** /api/v1/autocomplete/rules | All rules of the user returned in a basic auto-complete array.
[**GetTagAC**](AutocompleteApi.md#GetTagAC) | **Get** /api/v1/autocomplete/tags | All tags of the user returned in a basic auto-complete array.
[**GetTransactionTypesAC**](AutocompleteApi.md#GetTransactionTypesAC) | **Get** /api/v1/autocomplete/transaction-types | All transaction types returned in a basic auto-complete array. English only.
[**GetTransactionsAC**](AutocompleteApi.md#GetTransactionsAC) | **Get** /api/v1/autocomplete/transactions | All transaction descriptions of the user returned in a basic auto-complete array.
[**GetTransactionsIDAC**](AutocompleteApi.md#GetTransactionsIDAC) | **Get** /api/v1/autocomplete/transactions-with-id | All transactions, complemented with their ID, of the user returned in a basic auto-complete array.



## GetAccountsAC

> []AutocompleteAccount GetAccountsAC(ctx).Query(query).Limit(limit).Date(date).Type_(type_).Execute()

All accounts of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)
    date := "2020-09-17" // string | For asset accounts, returns the balance on this date. (optional)
    type_ := openapiclient.AccountTypeFilter("all") // AccountTypeFilter | Optional filter on the account type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetAccountsAC(context.Background()).Query(query).Limit(limit).Date(date).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetAccountsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountsAC`: []AutocompleteAccount
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetAccountsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 
 **date** | **string** | For asset accounts, returns the balance on this date. | 
 **type_** | [**AccountTypeFilter**](AccountTypeFilter.md) | Optional filter on the account type(s) returned | 

### Return type

[**[]AutocompleteAccount**](AutocompleteAccount.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillsAC

> []AutocompleteBill GetBillsAC(ctx).Query(query).Limit(limit).Execute()

All bills of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetBillsAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetBillsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillsAC`: []AutocompleteBill
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetBillsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBillsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteBill**](AutocompleteBill.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBudgetsAC

> []AutocompleteBudget GetBudgetsAC(ctx).Query(query).Limit(limit).Execute()

All budgets of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetBudgetsAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetBudgetsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBudgetsAC`: []AutocompleteBudget
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetBudgetsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBudgetsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteBudget**](AutocompleteBudget.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategoriesAC

> []AutocompleteCategory GetCategoriesAC(ctx).Query(query).Limit(limit).Execute()

All categories of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetCategoriesAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetCategoriesAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoriesAC`: []AutocompleteCategory
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetCategoriesAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteCategory**](AutocompleteCategory.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesAC

> []AutocompleteCurrency GetCurrenciesAC(ctx).Query(query).Limit(limit).Execute()

All currencies of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetCurrenciesAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetCurrenciesAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesAC`: []AutocompleteCurrency
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetCurrenciesAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteCurrency**](AutocompleteCurrency.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesCodeAC

> []AutocompleteCurrencyCode GetCurrenciesCodeAC(ctx).Query(query).Limit(limit).Execute()

All currencies of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetCurrenciesCodeAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetCurrenciesCodeAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesCodeAC`: []AutocompleteCurrencyCode
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetCurrenciesCodeAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesCodeACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteCurrencyCode**](AutocompleteCurrencyCode.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectGroupsAC

> []AutocompleteObjectGroup GetObjectGroupsAC(ctx).Query(query).Limit(limit).Execute()

All object groups of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetObjectGroupsAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetObjectGroupsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectGroupsAC`: []AutocompleteObjectGroup
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetObjectGroupsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectGroupsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteObjectGroup**](AutocompleteObjectGroup.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPiggiesAC

> []AutocompletePiggy GetPiggiesAC(ctx).Query(query).Limit(limit).Execute()

All piggy banks of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetPiggiesAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetPiggiesAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPiggiesAC`: []AutocompletePiggy
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetPiggiesAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPiggiesACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompletePiggy**](AutocompletePiggy.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPiggiesBalanceAC

> []AutocompletePiggyBalance GetPiggiesBalanceAC(ctx).Query(query).Limit(limit).Execute()

All piggy banks of the user returned in a basic auto-complete array complemented with balance information.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetPiggiesBalanceAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetPiggiesBalanceAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPiggiesBalanceAC`: []AutocompletePiggyBalance
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetPiggiesBalanceAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPiggiesBalanceACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompletePiggyBalance**](AutocompletePiggyBalance.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleGroupsAC

> []AutocompleteRuleGroup GetRuleGroupsAC(ctx).Query(query).Limit(limit).Execute()

All rule groups of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetRuleGroupsAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetRuleGroupsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleGroupsAC`: []AutocompleteRuleGroup
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetRuleGroupsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleGroupsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteRuleGroup**](AutocompleteRuleGroup.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesAC

> []AutocompleteRule GetRulesAC(ctx).Query(query).Limit(limit).Execute()

All rules of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetRulesAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetRulesAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRulesAC`: []AutocompleteRule
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetRulesAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteRule**](AutocompleteRule.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagAC

> []AutocompleteTag GetTagAC(ctx).Query(query).Limit(limit).Execute()

All tags of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetTagAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetTagAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagAC`: []AutocompleteTag
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetTagAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteTag**](AutocompleteTag.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionTypesAC

> []AutocompleteTransactionType GetTransactionTypesAC(ctx).Query(query).Limit(limit).Execute()

All transaction types returned in a basic auto-complete array. English only.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetTransactionTypesAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetTransactionTypesAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionTypesAC`: []AutocompleteTransactionType
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetTransactionTypesAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionTypesACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteTransactionType**](AutocompleteTransactionType.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsAC

> []AutocompleteTransaction GetTransactionsAC(ctx).Query(query).Limit(limit).Execute()

All transaction descriptions of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetTransactionsAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetTransactionsAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsAC`: []AutocompleteTransaction
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetTransactionsAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteTransaction**](AutocompleteTransaction.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionsIDAC

> []AutocompleteTransactionID GetTransactionsIDAC(ctx).Query(query).Limit(limit).Execute()

All transactions, complemented with their ID, of the user returned in a basic auto-complete array.

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
    query := "str" // string | The autocomplete search query. (optional)
    limit := int32(10) // int32 | The autocomplete number of items returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AutocompleteApi.GetTransactionsIDAC(context.Background()).Query(query).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteApi.GetTransactionsIDAC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionsIDAC`: []AutocompleteTransactionID
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteApi.GetTransactionsIDAC`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsIDACRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The autocomplete search query. | 
 **limit** | **int32** | The autocomplete number of items returned | 

### Return type

[**[]AutocompleteTransactionID**](AutocompleteTransactionID.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

