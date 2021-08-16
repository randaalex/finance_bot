# \LinksApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLinkType**](LinksApi.md#DeleteLinkType) | **Delete** /api/v1/link_types/{id} | Permanently delete link type.
[**DeleteTransactionLink**](LinksApi.md#DeleteTransactionLink) | **Delete** /api/v1/transaction_links/{id} | Permanently delete link between transactions.
[**GetLinkType**](LinksApi.md#GetLinkType) | **Get** /api/v1/link_types/{id} | Get single a link type.
[**GetTransactionLink**](LinksApi.md#GetTransactionLink) | **Get** /api/v1/transaction_links/{id} | Get a single link.
[**ListLinkType**](LinksApi.md#ListLinkType) | **Get** /api/v1/link_types | List all types of links.
[**ListTransactionByLinkType**](LinksApi.md#ListTransactionByLinkType) | **Get** /api/v1/link_types/{id}/transactions | List all transactions under this link type.
[**ListTransactionLink**](LinksApi.md#ListTransactionLink) | **Get** /api/v1/transaction_links | List all transaction links.
[**StoreLinkType**](LinksApi.md#StoreLinkType) | **Post** /api/v1/link_types | Create a new link type
[**StoreTransactionLink**](LinksApi.md#StoreTransactionLink) | **Post** /api/v1/transaction_links | Create a new link between transactions
[**UpdateLinkType**](LinksApi.md#UpdateLinkType) | **Put** /api/v1/link_types/{id} | Update existing link type.
[**UpdateTransactionLink**](LinksApi.md#UpdateTransactionLink) | **Put** /api/v1/transaction_links/{id} | Update an existing link between transactions.



## DeleteLinkType

> DeleteLinkType(ctx, id).Execute()

Permanently delete link type.



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
    id := int32(1) // int32 | The ID of the link type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.DeleteLinkType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.DeleteLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkTypeRequest struct via the builder pattern


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


## DeleteTransactionLink

> DeleteTransactionLink(ctx, id).Execute()

Permanently delete link between transactions.



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
    id := int32(1) // int32 | The ID of the transaction link.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.DeleteTransactionLink(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.DeleteTransactionLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the transaction link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionLinkRequest struct via the builder pattern


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


## GetLinkType

> LinkTypeSingle GetLinkType(ctx, id).Execute()

Get single a link type.



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
    id := int32(1) // int32 | The ID of the link type.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.GetLinkType(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GetLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinkType`: LinkTypeSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.GetLinkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionLink

> TransactionLinkSingle GetTransactionLink(ctx, id).Execute()

Get a single link.



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
    id := int32(1) // int32 | The ID of the transaction link.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.GetTransactionLink(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GetTransactionLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionLink`: TransactionLinkSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.GetTransactionLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the transaction link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLinkType

> LinkTypeArray ListLinkType(ctx).Page(page).Execute()

List all types of links.



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
    page := int32(1) // int32 | Page number. The default pagination is 50 items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.ListLinkType(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.ListLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLinkType`: LinkTypeArray
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.ListLinkType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is 50 items. | 

### Return type

[**LinkTypeArray**](LinkTypeArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionByLinkType

> TransactionArray ListTransactionByLinkType(ctx, id).Page(page).Start(start).End(end).Type_(type_).Execute()

List all transactions under this link type.



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
    id := int32(1) // int32 | The ID of the link type.
    page := int32(1) // int32 | Page number. The default pagination is per 50 items. (optional)
    start := time.Now() // string | A date formatted YYYY-MM-DD, to limit the results.  (optional)
    end := time.Now() // string | A date formatted YYYY-MM-DD, to limit the results.  (optional)
    type_ := openapiclient.TransactionTypeFilter("all") // TransactionTypeFilter | Optional filter on the transaction type(s) returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.ListTransactionByLinkType(context.Background(), id).Page(page).Start(start).End(end).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.ListTransactionByLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionByLinkType`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.ListTransactionByLinkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionByLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is per 50 items. | 
 **start** | **string** | A date formatted YYYY-MM-DD, to limit the results.  | 
 **end** | **string** | A date formatted YYYY-MM-DD, to limit the results.  | 
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


## ListTransactionLink

> TransactionLinkArray ListTransactionLink(ctx).Page(page).Execute()

List all transaction links.



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
    resp, r, err := api_client.LinksApi.ListTransactionLink(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.ListTransactionLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactionLink`: TransactionLinkArray
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.ListTransactionLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is per 50 items. | 

### Return type

[**TransactionLinkArray**](TransactionLinkArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreLinkType

> LinkTypeSingle StoreLinkType(ctx).LinkType(linkType).Execute()

Create a new link type



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
    linkType := *openapiclient.NewLinkType("Paid", "is (partially) paid for by", "(partially) pays for") // LinkType | JSON array with the necessary link type information or key=value pairs. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.StoreLinkType(context.Background()).LinkType(linkType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.StoreLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreLinkType`: LinkTypeSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.StoreLinkType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkType** | [**LinkType**](LinkType.md) | JSON array with the necessary link type information or key&#x3D;value pairs. See the model for the exact specifications. | 

### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreTransactionLink

> TransactionLinkSingle StoreTransactionLink(ctx).TransactionLinkStore(transactionLinkStore).Execute()

Create a new link between transactions



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
    transactionLinkStore := *openapiclient.NewTransactionLinkStore("5", "131", "131") // TransactionLinkStore | JSON array with the necessary link type information or key=value pairs. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.StoreTransactionLink(context.Background()).TransactionLinkStore(transactionLinkStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.StoreTransactionLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreTransactionLink`: TransactionLinkSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.StoreTransactionLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreTransactionLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionLinkStore** | [**TransactionLinkStore**](TransactionLinkStore.md) | JSON array with the necessary link type information or key&#x3D;value pairs. See the model for the exact specifications. | 

### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkType

> LinkTypeSingle UpdateLinkType(ctx, id).LinkTypeUpdate(linkTypeUpdate).Execute()

Update existing link type.



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
    id := int32(1) // int32 | The ID of the link type.
    linkTypeUpdate := *openapiclient.NewLinkTypeUpdate() // LinkTypeUpdate | JSON array or formdata with updated link type information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.UpdateLinkType(context.Background(), id).LinkTypeUpdate(linkTypeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.UpdateLinkType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLinkType`: LinkTypeSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.UpdateLinkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkTypeUpdate** | [**LinkTypeUpdate**](LinkTypeUpdate.md) | JSON array or formdata with updated link type information. See the model for the exact specifications. | 

### Return type

[**LinkTypeSingle**](LinkTypeSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransactionLink

> TransactionLinkSingle UpdateTransactionLink(ctx, id).TransactionLinkUpdate(transactionLinkUpdate).Execute()

Update an existing link between transactions.



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
    id := int32(1) // int32 | The ID of the transaction link.
    transactionLinkUpdate := *openapiclient.NewTransactionLinkUpdate() // TransactionLinkUpdate | JSON array or formdata with updated link type information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LinksApi.UpdateTransactionLink(context.Background(), id).TransactionLinkUpdate(transactionLinkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.UpdateTransactionLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransactionLink`: TransactionLinkSingle
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.UpdateTransactionLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the transaction link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransactionLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionLinkUpdate** | [**TransactionLinkUpdate**](TransactionLinkUpdate.md) | JSON array or formdata with updated link type information. See the model for the exact specifications. | 

### Return type

[**TransactionLinkSingle**](TransactionLinkSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

