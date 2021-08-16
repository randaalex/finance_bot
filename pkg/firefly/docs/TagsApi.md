# \TagsApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTag**](TagsApi.md#DeleteTag) | **Delete** /api/v1/tags/{tag} | Delete an tag.
[**GetTag**](TagsApi.md#GetTag) | **Get** /api/v1/tags/{tag} | Get a single tag.
[**ListAttachmentByTag**](TagsApi.md#ListAttachmentByTag) | **Get** /api/v1/tags/{tag}/attachments | Lists all attachments.
[**ListTag**](TagsApi.md#ListTag) | **Get** /api/v1/tags | List all tags.
[**ListTransactionByTag**](TagsApi.md#ListTransactionByTag) | **Get** /api/v1/tags/{tag}/transactions | List all transactions with this tag.
[**StoreTag**](TagsApi.md#StoreTag) | **Post** /api/v1/tags | Store a new tag
[**UpdateTag**](TagsApi.md#UpdateTag) | **Put** /api/v1/tags/{tag} | Update existing tag.



## DeleteTag

> DeleteTag(ctx, tag).Execute()

Delete an tag.



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
    tag := "groceries" // string | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.DeleteTag(context.Background(), tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetTag

> TagSingle GetTag(ctx, tag).Page(page).Execute()

Get a single tag.



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
    tag := "groceries" // string | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary.
    page := int32(56) // int32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.GetTag(context.Background(), tag).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: TagSingle
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttachmentByTag

> AttachmentArray ListAttachmentByTag(ctx, tag).Page(page).Execute()

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
    tag := "groceries" // string | Either the tag itself or the tag ID.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.ListAttachmentByTag(context.Background(), tag).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListAttachmentByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttachmentByTag`: AttachmentArray
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ListAttachmentByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | Either the tag itself or the tag ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttachmentByTagRequest struct via the builder pattern


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


## ListTag

> TagArray ListTag(ctx).Page(page).Execute()

List all tags.



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
    page := int32(56) // int32 | Page number (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.ListTag(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTag`: TagArray
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ListTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | 

### Return type

[**TagArray**](TagArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByTag

> TransactionArray ListTransactionByTag(ctx, tag).Page(page).Start(start).End(end).Type_(type_).Execute()

List all transactions with this tag.



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
    tag := "groceries" // string | Either the tag itself or the tag ID.
    page := int32(1) // int32 | Page number. The default pagination is 50. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD. This is the start date of the selected range (inclusive).  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD. This is the end date of the selected range (inclusive).  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.ListTransactionByTag(context.Background(), tag).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ListTransactionByTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByTag`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ListTransactionByTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | Either the tag itself or the tag ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByTagRequest struct via the builder pattern


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
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreTag

> TagSingle StoreTag(ctx).TagModelStore(tagModelStore).Execute()

Store a new tag



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
    tagModelStore := *openapiclient.NewTagModelStore("expensive") // TagModelStore | JSON array or key=value pairs with the necessary tag information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.StoreTag(context.Background()).TagModelStore(tagModelStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.StoreTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreTag`: TagSingle
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.StoreTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagModelStore** | [**TagModelStore**](TagModelStore.md) | JSON array or key&#x3D;value pairs with the necessary tag information. See the model for the exact specifications. | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> TagSingle UpdateTag(ctx, tag).TagModelUpdate(tagModelUpdate).Execute()

Update existing tag.



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
    tag := "groceries" // string | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary.
    tagModelUpdate := *openapiclient.NewTagModelUpdate() // TagModelUpdate | JSON array with updated tag information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.UpdateTag(context.Background(), tag).TagModelUpdate(tagModelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: TagSingle
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** | Either the tag itself or the tag ID. If you use the tag itself, and it contains international (non-ASCII) characters, your milage may vary. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagModelUpdate** | [**TagModelUpdate**](TagModelUpdate.md) | JSON array with updated tag information. See the model for the exact specifications. | 

### Return type

[**TagSingle**](TagSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

