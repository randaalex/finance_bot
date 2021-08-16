# \RecurrencesApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecurrence**](RecurrencesApi.md#DeleteRecurrence) | **Delete** /api/v1/recurrences/{id} | Delete a recurring transaction.
[**GetRecurrence**](RecurrencesApi.md#GetRecurrence) | **Get** /api/v1/recurrences/{id} | Get a single recurring transaction.
[**ListRecurrence**](RecurrencesApi.md#ListRecurrence) | **Get** /api/v1/recurrences | List all recurring transactions.
[**ListTransactionByRecurrence**](RecurrencesApi.md#ListTransactionByRecurrence) | **Get** /api/v1/recurrences/{id}/transactions | List all transactions created by a recurring transaction.
[**StoreRecurrence**](RecurrencesApi.md#StoreRecurrence) | **Post** /api/v1/recurrences | Store a new recurring transaction
[**UpdateRecurrence**](RecurrencesApi.md#UpdateRecurrence) | **Put** /api/v1/recurrences/{id} | Update existing recurring transaction.



## DeleteRecurrence

> DeleteRecurrence(ctx, id).Execute()

Delete a recurring transaction.



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
    id := int32(1) // int32 | The ID of the recurring transaction.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrencesApi.DeleteRecurrence(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.DeleteRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the recurring transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurrenceRequest struct via the builder pattern


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


## GetRecurrence

> RecurrenceSingle GetRecurrence(ctx, id).Execute()

Get a single recurring transaction.



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
    id := int32(1) // int32 | The ID of the recurring transaction.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrencesApi.GetRecurrence(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.GetRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecurrence`: RecurrenceSingle
    fmt.Fprintf(os.Stdout, "Response from `RecurrencesApi.GetRecurrence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the recurring transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRecurrence

> RecurrenceArray ListRecurrence(ctx).Page(page).Execute()

List all recurring transactions.



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
    resp, r, err := api_client.RecurrencesApi.ListRecurrence(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.ListRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRecurrence`: RecurrenceArray
    fmt.Fprintf(os.Stdout, "Response from `RecurrencesApi.ListRecurrence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRecurrenceRequest struct via the builder pattern


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


## ListTransactionByRecurrence

> TransactionArray ListTransactionByRecurrence(ctx, id).Page(page).Start(start).End(end).Type_(type_).Execute()

List all transactions created by a recurring transaction.



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
    id := int32(1) // int32 | The ID of the recurring transaction.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD. Both the start and end date must be present.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD. Both the start and end date must be present.  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrencesApi.ListTransactionByRecurrence(context.Background(), id).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.ListTransactionByRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByRecurrence`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `RecurrencesApi.ListTransactionByRecurrence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the recurring transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD. Both the start and end date must be present.  | 
 **end** | **string** | A date formatted YYYY-MM-DD. Both the start and end date must be present.  | 
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


## StoreRecurrence

> RecurrenceSingle StoreRecurrence(ctx).RecurrenceStore(recurrenceStore).Execute()

Store a new recurring transaction



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
    recurrenceStore := *openapiclient.NewRecurrenceStore("withdrawal", "Rent", time.Now(), time.Now(), []openapiclient.RecurrenceRepetitionStore{*openapiclient.NewRecurrenceRepetitionStore("weekly", "3")}, []openapiclient.RecurrenceTransactionStore{*openapiclient.NewRecurrenceTransactionStore("Rent for the current month", "123.45", "913", "258")}) // RecurrenceStore | JSON array or key=value pairs with the necessary recurring transaction information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrencesApi.StoreRecurrence(context.Background()).RecurrenceStore(recurrenceStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.StoreRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreRecurrence`: RecurrenceSingle
    fmt.Fprintf(os.Stdout, "Response from `RecurrencesApi.StoreRecurrence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recurrenceStore** | [**RecurrenceStore**](RecurrenceStore.md) | JSON array or key&#x3D;value pairs with the necessary recurring transaction information. See the model for the exact specifications. | 

### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecurrence

> RecurrenceSingle UpdateRecurrence(ctx, id).RecurrenceUpdate(recurrenceUpdate).Execute()

Update existing recurring transaction.



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
    id := int32(1) // int32 | The ID of the recurring transaction.
    recurrenceUpdate := *openapiclient.NewRecurrenceUpdate() // RecurrenceUpdate | JSON array with updated recurring transaction information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RecurrencesApi.UpdateRecurrence(context.Background(), id).RecurrenceUpdate(recurrenceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecurrencesApi.UpdateRecurrence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecurrence`: RecurrenceSingle
    fmt.Fprintf(os.Stdout, "Response from `RecurrencesApi.UpdateRecurrence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the recurring transaction. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurrenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recurrenceUpdate** | [**RecurrenceUpdate**](RecurrenceUpdate.md) | JSON array with updated recurring transaction information. See the model for the exact specifications. | 

### Return type

[**RecurrenceSingle**](RecurrenceSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

