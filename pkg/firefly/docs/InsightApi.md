# \InsightApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightExpenseAsset**](InsightApi.md#InsightExpenseAsset) | **Get** /api/v1/insight/expense/asset | Insight into expenses, grouped by asset account.
[**InsightExpenseBill**](InsightApi.md#InsightExpenseBill) | **Get** /api/v1/insight/expense/bill | Insight into expenses, grouped by bill.
[**InsightExpenseBudget**](InsightApi.md#InsightExpenseBudget) | **Get** /api/v1/insight/expense/budget | Insight into expenses, grouped by budget.
[**InsightExpenseCategory**](InsightApi.md#InsightExpenseCategory) | **Get** /api/v1/insight/expense/category | Insight into expenses, grouped by category.
[**InsightExpenseExpense**](InsightApi.md#InsightExpenseExpense) | **Get** /api/v1/insight/expense/expense | Insight into expenses, grouped by expense account.
[**InsightExpenseNoBill**](InsightApi.md#InsightExpenseNoBill) | **Get** /api/v1/insight/expense/no-bill | Insight into expenses, without bill.
[**InsightExpenseNoBudget**](InsightApi.md#InsightExpenseNoBudget) | **Get** /api/v1/insight/expense/no-budget | Insight into expenses, without budget.
[**InsightExpenseNoCategory**](InsightApi.md#InsightExpenseNoCategory) | **Get** /api/v1/insight/expense/no-category | Insight into expenses, without category.
[**InsightExpenseNoTag**](InsightApi.md#InsightExpenseNoTag) | **Get** /api/v1/insight/expense/no-tag | Insight into expenses, without tag.
[**InsightExpenseTag**](InsightApi.md#InsightExpenseTag) | **Get** /api/v1/insight/expense/tag | Insight into expenses, grouped by tag.
[**InsightExpenseTotal**](InsightApi.md#InsightExpenseTotal) | **Get** /api/v1/insight/expense/total | Insight into total expenses.
[**InsightIncomeAsset**](InsightApi.md#InsightIncomeAsset) | **Get** /api/v1/insight/income/asset | Insight into income, grouped by asset account.
[**InsightIncomeCategory**](InsightApi.md#InsightIncomeCategory) | **Get** /api/v1/insight/income/category | Insight into income, grouped by category.
[**InsightIncomeNoCategory**](InsightApi.md#InsightIncomeNoCategory) | **Get** /api/v1/insight/income/no-category | Insight into income, without category.
[**InsightIncomeNoTag**](InsightApi.md#InsightIncomeNoTag) | **Get** /api/v1/insight/income/no-tag | Insight into income, without tag.
[**InsightIncomeRevenue**](InsightApi.md#InsightIncomeRevenue) | **Get** /api/v1/insight/income/revenue | Insight into income, grouped by revenue account.
[**InsightIncomeTag**](InsightApi.md#InsightIncomeTag) | **Get** /api/v1/insight/income/tag | Insight into income, grouped by tag.
[**InsightIncomeTotal**](InsightApi.md#InsightIncomeTotal) | **Get** /api/v1/insight/income/total | Insight into total income.
[**InsightTransferCategory**](InsightApi.md#InsightTransferCategory) | **Get** /api/v1/insight/transfer/category | Insight into transfers, grouped by category.
[**InsightTransferNoCategory**](InsightApi.md#InsightTransferNoCategory) | **Get** /api/v1/insight/transfer/no-category | Insight into transfers, without category.
[**InsightTransferNoTag**](InsightApi.md#InsightTransferNoTag) | **Get** /api/v1/insight/transfer/no-tag | Insight into expenses, without tag.
[**InsightTransferTag**](InsightApi.md#InsightTransferTag) | **Get** /api/v1/insight/transfer/tag | Insight into transfers, grouped by tag.
[**InsightTransferTotal**](InsightApi.md#InsightTransferTotal) | **Get** /api/v1/insight/transfer/total | Insight into total transfers.
[**InsightTransfers**](InsightApi.md#InsightTransfers) | **Get** /api/v1/insight/transfer/asset | Insight into transfers, grouped by account.



## InsightExpenseAsset

> []InsightGroupEntry InsightExpenseAsset(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, grouped by asset account.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseAsset(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseAsset`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseBill

> []InsightGroupEntry InsightExpenseBill(ctx).Start(start).End(end).Bills(bills).Accounts(accounts).Execute()

Insight into expenses, grouped by bill.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    bills := []int64{int64(123)} // []int64 | The bills to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseBill(context.Background()).Start(start).End(end).Bills(bills).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseBill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseBill`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseBill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseBillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **bills** | **[]int64** | The bills to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseBudget

> []InsightGroupEntry InsightExpenseBudget(ctx).Start(start).End(end).Budgets(budgets).Accounts(accounts).Execute()

Insight into expenses, grouped by budget.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    budgets := []int64{int64(123)} // []int64 | The budgets to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseBudget(context.Background()).Start(start).End(end).Budgets(budgets).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseBudget`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **budgets** | **[]int64** | The budgets to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseCategory

> []InsightGroupEntry InsightExpenseCategory(ctx).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()

Insight into expenses, grouped by category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    categories := []int64{int64(123)} // []int64 | The categories to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseCategory(context.Background()).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseCategory`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **categories** | **[]int64** | The categories to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseExpense

> []InsightGroupEntry InsightExpenseExpense(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, grouped by expense account.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you add the accounts ID's of expense accounts, only those accounts are included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. You can combine both asset / liability and expense account ID's. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseExpense(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseExpense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseExpense`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseExpense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseExpenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you add the accounts ID&#39;s of expense accounts, only those accounts are included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. You can combine both asset / liability and expense account ID&#39;s. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseNoBill

> []InsightTotalEntry InsightExpenseNoBill(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, without bill.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseNoBill(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseNoBill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseNoBill`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseNoBill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseNoBillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseNoBudget

> []InsightTotalEntry InsightExpenseNoBudget(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, without budget.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseNoBudget(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseNoBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseNoBudget`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseNoBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseNoBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseNoCategory

> []InsightTotalEntry InsightExpenseNoCategory(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, without category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseNoCategory(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseNoCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseNoCategory`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseNoCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseNoCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseNoTag

> []InsightTotalEntry InsightExpenseNoTag(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, without tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseNoTag(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseNoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseNoTag`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseNoTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseNoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseTag

> []InsightGroupEntry InsightExpenseTag(ctx).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()

Insight into expenses, grouped by tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    tags := []int64{int64(123)} // []int64 | The tags to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseTag(context.Background()).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseTag`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **tags** | **[]int64** | The tags to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightExpenseTotal

> []InsightTotalEntry InsightExpenseTotal(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into total expenses.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightExpenseTotal(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightExpenseTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightExpenseTotal`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightExpenseTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightExpenseTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only withdrawals from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeAsset

> []InsightGroupEntry InsightIncomeAsset(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into income, grouped by asset account.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeAsset(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeAsset`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeCategory

> []InsightGroupEntry InsightIncomeCategory(ctx).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()

Insight into income, grouped by category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    categories := []int64{int64(123)} // []int64 | The categories to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeCategory(context.Background()).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeCategory`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **categories** | **[]int64** | The categories to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeNoCategory

> []InsightTotalEntry InsightIncomeNoCategory(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into income, without category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeNoCategory(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeNoCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeNoCategory`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeNoCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeNoCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeNoTag

> []InsightTotalEntry InsightIncomeNoTag(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into income, without tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeNoTag(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeNoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeNoTag`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeNoTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeNoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeRevenue

> []InsightGroupEntry InsightIncomeRevenue(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into income, grouped by revenue account.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you add the accounts ID's of revenue accounts, only those accounts are included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. You can combine both asset / liability and deposit account ID's. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeRevenue(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeRevenue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeRevenue`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeRevenue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeRevenueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you add the accounts ID&#39;s of revenue accounts, only those accounts are included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. You can combine both asset / liability and deposit account ID&#39;s. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeTag

> []InsightGroupEntry InsightIncomeTag(ctx).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()

Insight into income, grouped by tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    tags := []int64{int64(123)} // []int64 | The tags to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeTag(context.Background()).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeTag`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **tags** | **[]int64** | The tags to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightIncomeTotal

> []InsightTotalEntry InsightIncomeTotal(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into total income.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightIncomeTotal(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightIncomeTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightIncomeTotal`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightIncomeTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightIncomeTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only deposits to those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransferCategory

> []InsightGroupEntry InsightTransferCategory(ctx).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()

Insight into transfers, grouped by category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    categories := []int64{int64(123)} // []int64 | The categories to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransferCategory(context.Background()).Start(start).End(end).Categories(categories).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransferCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransferCategory`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransferCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransferCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **categories** | **[]int64** | The categories to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransferNoCategory

> []InsightTotalEntry InsightTransferNoCategory(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into transfers, without category.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransferNoCategory(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransferNoCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransferNoCategory`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransferNoCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransferNoCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransferNoTag

> []InsightTotalEntry InsightTransferNoTag(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into expenses, without tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers from those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransferNoTag(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransferNoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransferNoTag`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransferNoTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransferNoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers from those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransferTag

> []InsightGroupEntry InsightTransferTag(ctx).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()

Insight into transfers, grouped by tag.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    tags := []int64{int64(123)} // []int64 | The tags to be included in the results.  (optional)
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransferTag(context.Background()).Start(start).End(end).Tags(tags).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransferTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransferTag`: []InsightGroupEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransferTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransferTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **tags** | **[]int64** | The tags to be included in the results.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightGroupEntry**](InsightGroupEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransferTotal

> []InsightTotalEntry InsightTransferTotal(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into total transfers.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransferTotal(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransferTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransferTotal`: []InsightTotalEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransferTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransferTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTotalEntry**](InsightTotalEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsightTransfers

> []InsightTransferEntry InsightTransfers(ctx).Start(start).End(end).Accounts(accounts).Execute()

Insight into transfers, grouped by account.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    accounts := []int64{int64(123)} // []int64 | The accounts to be included in the results. If you include ID's of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID's will be ignored.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InsightApi.InsightTransfers(context.Background()).Start(start).End(end).Accounts(accounts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.InsightTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsightTransfers`: []InsightTransferEntry
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.InsightTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsightTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **[]int64** | The accounts to be included in the results. If you include ID&#39;s of asset accounts or liabilities, only transfers between those asset accounts / liabilities will be included. Other account ID&#39;s will be ignored.  | 

### Return type

[**[]InsightTransferEntry**](InsightTransferEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

