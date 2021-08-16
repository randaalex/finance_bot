# \DataApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkAccountMoveTransactions**](DataApi.md#BulkAccountMoveTransactions) | **Post** /api/v1/data/bulk/accounts/transactions | Bulk move transactions from one account to another.
[**DestroyData**](DataApi.md#DestroyData) | **Delete** /api/v1/data/destroy | Endpoint to destroy user data
[**ExportAccounts**](DataApi.md#ExportAccounts) | **Get** /api/v1/data/export/accounts | Export account data from Firefly III
[**ExportBills**](DataApi.md#ExportBills) | **Get** /api/v1/data/export/bills | Export bills from Firefly III
[**ExportBudgets**](DataApi.md#ExportBudgets) | **Get** /api/v1/data/export/budgets | Export budgets and budget amount data from Firefly III
[**ExportCategories**](DataApi.md#ExportCategories) | **Get** /api/v1/data/export/categories | Export category data from Firefly III
[**ExportPiggies**](DataApi.md#ExportPiggies) | **Get** /api/v1/data/export/piggy-banks | Export piggy banks from Firefly III
[**ExportRecurring**](DataApi.md#ExportRecurring) | **Get** /api/v1/data/export/recurring | Export recurring transaction data from Firefly III
[**ExportRules**](DataApi.md#ExportRules) | **Get** /api/v1/data/export/rules | Export rule groups and rule data from Firefly III
[**ExportTags**](DataApi.md#ExportTags) | **Get** /api/v1/data/export/tags | Export tag data from Firefly III
[**ExportTransactions**](DataApi.md#ExportTransactions) | **Get** /api/v1/data/export/transactions | Export transaction data from Firefly III



## BulkAccountMoveTransactions

> BulkAccountMoveTransactions(ctx).BulkAccountTransactionObject(bulkAccountTransactionObject).Execute()

Bulk move transactions from one account to another.



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
    bulkAccountTransactionObject := *openapiclient.NewBulkAccountTransactionObject(float32(2), float32(17)) // BulkAccountTransactionObject | JSON array with the necessary information to facilitate the move.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.BulkAccountMoveTransactions(context.Background()).BulkAccountTransactionObject(bulkAccountTransactionObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.BulkAccountMoveTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkAccountMoveTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkAccountTransactionObject** | [**BulkAccountTransactionObject**](BulkAccountTransactionObject.md) | JSON array with the necessary information to facilitate the move. | 

### Return type

 (empty response body)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyData

> DestroyData(ctx).Objects(objects).Execute()

Endpoint to destroy user data



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
    objects := openapiclient.DataDestroyObject("budgets") // DataDestroyObject | The type of data that you wish to destroy.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.DestroyData(context.Background()).Objects(objects).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.DestroyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objects** | [**DataDestroyObject**](DataDestroyObject.md) | The type of data that you wish to destroy. | 

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


## ExportAccounts

> *os.File ExportAccounts(ctx).Type_(type_).Execute()

Export account data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportAccounts(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAccounts`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportBills

> *os.File ExportBills(ctx).Type_(type_).Execute()

Export bills from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportBills(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportBills``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBills`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportBills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportBillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportBudgets

> *os.File ExportBudgets(ctx).Type_(type_).Execute()

Export budgets and budget amount data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportBudgets(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBudgets`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCategories

> *os.File ExportCategories(ctx).Type_(type_).Execute()

Export category data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportCategories(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCategories`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPiggies

> *os.File ExportPiggies(ctx).Type_(type_).Execute()

Export piggy banks from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportPiggies(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportPiggies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPiggies`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportPiggies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportPiggiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRecurring

> *os.File ExportRecurring(ctx).Type_(type_).Execute()

Export recurring transaction data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportRecurring(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportRecurring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRecurring`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportRecurring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRecurringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRules

> *os.File ExportRules(ctx).Type_(type_).Execute()

Export rule groups and rule data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportRules(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRules`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTags

> *os.File ExportTags(ctx).Type_(type_).Execute()

Export tag data from Firefly III



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
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportTags(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTags`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportTransactions

> *os.File ExportTransactions(ctx).Start(start).End(end).Accounts(accounts).Type_(type_).Execute()

Export transaction data from Firefly III



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
    accounts := "1,2,3" // string | Limit the export of transactions to these accounts only. Only asset accounts will be accepted. Other types will be silently dropped.  (optional)
    type_ := openapiclient.ExportFileFilter("csv") // ExportFileFilter | The file type the export file (CSV is currently the only option). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataApi.ExportTransactions(context.Background()).Start(start).End(end).Accounts(accounts).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataApi.ExportTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportTransactions`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DataApi.ExportTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **accounts** | **string** | Limit the export of transactions to these accounts only. Only asset accounts will be accepted. Other types will be silently dropped.  | 
 **type_** | [**ExportFileFilter**](ExportFileFilter.md) | The file type the export file (CSV is currently the only option). | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

