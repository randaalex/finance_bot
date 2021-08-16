# \AvailableBudgetsApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAvailableBudget**](AvailableBudgetsApi.md#DeleteAvailableBudget) | **Delete** /api/v1/available_budgets/{id} | Delete an available budget.
[**GetAvailableBudget**](AvailableBudgetsApi.md#GetAvailableBudget) | **Get** /api/v1/available_budgets/{id} | Get a single available budget.
[**ListAvailableBudget**](AvailableBudgetsApi.md#ListAvailableBudget) | **Get** /api/v1/available_budgets | List all available budget amounts.
[**StoreAvailableBudget**](AvailableBudgetsApi.md#StoreAvailableBudget) | **Post** /api/v1/available_budgets | Store a new available budget
[**UpdateAvailableBudget**](AvailableBudgetsApi.md#UpdateAvailableBudget) | **Put** /api/v1/available_budgets/{id} | Update existing available budget, to change for example the date range of the amount or the amount itself.



## DeleteAvailableBudget

> DeleteAvailableBudget(ctx, id).Execute()

Delete an available budget.



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
    id := int32(1) // int32 | The ID of the available budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailableBudgetsApi.DeleteAvailableBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailableBudgetsApi.DeleteAvailableBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the available budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvailableBudgetRequest struct via the builder pattern


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


## GetAvailableBudget

> AvailableBudgetSingle GetAvailableBudget(ctx, id).Execute()

Get a single available budget.



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
    id := int32(1) // int32 | The ID of the available budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailableBudgetsApi.GetAvailableBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailableBudgetsApi.GetAvailableBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableBudget`: AvailableBudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `AvailableBudgetsApi.GetAvailableBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the available budget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableBudget

> AvailableBudgetArray ListAvailableBudget(ctx).Page(page).Start(start).End(end).Execute()

List all available budget amounts.



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
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailableBudgetsApi.ListAvailableBudget(context.Background()).Page(page).Start(start).End(end).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailableBudgetsApi.ListAvailableBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableBudget`: AvailableBudgetArray
    fmt.Fprintf(os.Stdout, "Response from `AvailableBudgetsApi.ListAvailableBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is 50. | 
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 

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


## StoreAvailableBudget

> AvailableBudgetSingle StoreAvailableBudget(ctx).AvailableBudgetStore(availableBudgetStore).Execute()

Store a new available budget



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
    availableBudgetStore := *openapiclient.NewAvailableBudgetStore("123.45", time.Now(), time.Now()) // AvailableBudgetStore | JSON array or key=value pairs with the necessary available budget information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailableBudgetsApi.StoreAvailableBudget(context.Background()).AvailableBudgetStore(availableBudgetStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailableBudgetsApi.StoreAvailableBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreAvailableBudget`: AvailableBudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `AvailableBudgetsApi.StoreAvailableBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreAvailableBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **availableBudgetStore** | [**AvailableBudgetStore**](AvailableBudgetStore.md) | JSON array or key&#x3D;value pairs with the necessary available budget information. See the model for the exact specifications. | 

### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAvailableBudget

> AvailableBudgetSingle UpdateAvailableBudget(ctx, id).AvailableBudgetUpdate(availableBudgetUpdate).Execute()

Update existing available budget, to change for example the date range of the amount or the amount itself.



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
    id := int32(1) // int32 | The ID of the object.X
    availableBudgetUpdate := *openapiclient.NewAvailableBudgetUpdate() // AvailableBudgetUpdate | JSON array or form value with updated available budget information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AvailableBudgetsApi.UpdateAvailableBudget(context.Background(), id).AvailableBudgetUpdate(availableBudgetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvailableBudgetsApi.UpdateAvailableBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAvailableBudget`: AvailableBudgetSingle
    fmt.Fprintf(os.Stdout, "Response from `AvailableBudgetsApi.UpdateAvailableBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the object.X | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAvailableBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **availableBudgetUpdate** | [**AvailableBudgetUpdate**](AvailableBudgetUpdate.md) | JSON array or form value with updated available budget information. See the model for the exact specifications. | 

### Return type

[**AvailableBudgetSingle**](AvailableBudgetSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

