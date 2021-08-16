# \AccountsApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /api/v1/accounts/{id} | Permanently delete account.
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /api/v1/accounts/{id} | Get single account.
[**ListAccount**](AccountsApi.md#ListAccount) | **Get** /api/v1/accounts | List all accounts.
[**ListAttachmentByAccount**](AccountsApi.md#ListAttachmentByAccount) | **Get** /api/v1/accounts/{id}/attachments | Lists all attachments.
[**ListPiggyBankByAccount**](AccountsApi.md#ListPiggyBankByAccount) | **Get** /api/v1/accounts/{id}/piggy_banks | List all piggy banks related to the account.
[**ListTransactionByAccount**](AccountsApi.md#ListTransactionByAccount) | **Get** /api/v1/accounts/{id}/transactions | List all transactions related to the account.
[**StoreAccount**](AccountsApi.md#StoreAccount) | **Post** /api/v1/accounts | Create new account.
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Put** /api/v1/accounts/{id} | Update existing account.



## DeleteAccount

> DeleteAccount(ctx, id).Execute()

Permanently delete account.



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
    id := int32(1) // int32 | The ID of the account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## GetAccount

> AccountSingle GetAccount(ctx, id).Date(date).Execute()

Get single account.



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
    id := int32(1) // int32 | The ID of the account.
    date := time.Now() // string | A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account's balance on that day.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), id).Date(date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AccountSingle
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account&#39;s balance on that day.  | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccount

> AccountArray ListAccount(ctx).Page(page).Date(date).Type_(type_).Execute()

List all accounts.



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
    page := int32(1) // int32 | Page number. The default pagination is per 50 items. (optional)
    date := time.Now() // string | A date formatted YYYY-MM-DD. When added to the request, Firefly III will show the account's balance on that day.  (optional)
    type_ := openapiclient.AccountTypeFilter("all") // AccountTypeFilter | Optional filter on the account type(s) returned (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAccount(context.Background()).Page(page).Date(date).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccount`: AccountArray
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is per 50 items. | 
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


## ListAttachmentByAccount

> AttachmentArray ListAttachmentByAccount(ctx, id).Page(page).Execute()

Lists all attachments.



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
    id := int32(1) // int32 | The ID of the account.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListAttachmentByAccount(context.Background(), id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAttachmentByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachmentByAccount`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAttachmentByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**AttachmentArray**](AttachmentArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPiggyBankByAccount

> PiggyBankArray ListPiggyBankByAccount(ctx, id).Page(page).Execute()

List all piggy banks related to the account.



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
    id := int32(1) // int32 | The ID of the account.
    page := int32(56) // int32 | Page number. The default pagination is per 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListPiggyBankByAccount(context.Background(), id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListPiggyBankByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPiggyBankByAccount`: PiggyBankArray
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListPiggyBankByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPiggyBankByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is per 50 items. | 

### Return type

[**PiggyBankArray**](PiggyBankArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByAccount

> TransactionArray ListTransactionByAccount(ctx, id).Page(page).Limit(limit).Start(start).End(end).Type_(type_).Execute()

List all transactions related to the account.



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
    id := int32(1) // int32 | The ID of the account.
    page := int32(1) // int32 | Page number. The default pagination is per 50 items. (optional)
    limit := int32(5) // int32 | Limits the number of results on one page. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD.  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ListTransactionByAccount(context.Background(), id).Page(page).Limit(limit).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListTransactionByAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByAccount`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListTransactionByAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is per 50 items. | 
 **limit** | **int32** | Limits the number of results on one page. | 
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **type_** | [**TransactionTypeFilter**](TransactionTypeFilter.md) | Optional filter on the transaction type(s) returned. | 

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


## StoreAccount

> AccountSingle StoreAccount(ctx).AccountStore(accountStore).Execute()

Create new account.



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
    accountStore := *openapiclient.NewAccountStore("My checking account", "asset") // AccountStore | JSON array with the necessary account information or key=value pairs. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.StoreAccount(context.Background()).AccountStore(accountStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.StoreAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreAccount`: AccountSingle
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.StoreAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountStore** | [**AccountStore**](AccountStore.md) | JSON array with the necessary account information or key&#x3D;value pairs. See the model for the exact specifications. | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> AccountSingle UpdateAccount(ctx, id).AccountUpdate(accountUpdate).Execute()

Update existing account.



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
    id := int32(1) // int32 | The ID of the account.
    accountUpdate := *openapiclient.NewAccountUpdate("My checking account") // AccountUpdate | JSON array or formdata with updated account information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.UpdateAccount(context.Background(), id).AccountUpdate(accountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: AccountSingle
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUpdate** | [**AccountUpdate**](AccountUpdate.md) | JSON array or formdata with updated account information. See the model for the exact specifications. | 

### Return type

[**AccountSingle**](AccountSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

