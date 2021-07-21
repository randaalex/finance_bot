# \ImportApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImport**](ImportApi.md#GetImport) | **Get** /api/v1/import/{key} | Show info on a single import
[**ListImport**](ImportApi.md#ListImport) | **Get** /api/v1/import/list | List al imports
[**ListTransactionByImport**](ImportApi.md#ListTransactionByImport) | **Get** /api/v1/import/{key}/transactions | List all transactions related to the import job. The correlation is made through the tag.



## GetImport

> ImportJobSingle GetImport(ctx, key).Execute()

Show info on a single import



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
    key := "x2Akaijm2" // string | The job key of an import job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.GetImport(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.GetImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImport`: ImportJobSingle
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.GetImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The job key of an import job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImportJobSingle**](ImportJobSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImport

> ImportJobArray ListImport(ctx).Page(page).Execute()

List al imports



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
    page := int32(1) // int32 | Page number. The default pagination is per 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.ListImport(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ListImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListImport`: ImportJobArray
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ListImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is per 50 items. | 

### Return type

[**ImportJobArray**](ImportJobArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByImport

> TransactionArray ListTransactionByImport(ctx, key).Page(page).Start(start).End(end).Type_(type_).Execute()

List all transactions related to the import job. The correlation is made through the tag.



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
    key := "abcde" // string | The key of the import job
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ImportApi.ListTransactionByImport(context.Background(), key).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImportApi.ListTransactionByImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByImport`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `ImportApi.ListTransactionByImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The key of the import job | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  | 
 **end** | **string** | A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned. | 

### Return type

[**TransactionArray**](TransactionArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

