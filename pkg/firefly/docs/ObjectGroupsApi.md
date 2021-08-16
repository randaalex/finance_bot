# \ObjectGroupsApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectGroup**](ObjectGroupsApi.md#DeleteObjectGroup) | **Delete** /api/v1/object_groups/{id} | Delete a object group.
[**GetObjectGroup**](ObjectGroupsApi.md#GetObjectGroup) | **Get** /api/v1/object_groups/{id} | Get a single object group.
[**ListBillByObjectGroup**](ObjectGroupsApi.md#ListBillByObjectGroup) | **Get** /api/v1/object_groups/{id}/bills | List all bills with this object group.
[**ListObjectGroups**](ObjectGroupsApi.md#ListObjectGroups) | **Get** /api/v1/object_groups | List all oject groups.
[**ListPiggyBankByObjectGroup**](ObjectGroupsApi.md#ListPiggyBankByObjectGroup) | **Get** /api/v1/object_groups/{id}/piggy_banks | List all piggy banks related to the object group.
[**UpdateObjectGroup**](ObjectGroupsApi.md#UpdateObjectGroup) | **Put** /api/v1/object_groups/{id} | Update existing object group.



## DeleteObjectGroup

> DeleteObjectGroup(ctx, id).Execute()

Delete a object group.



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
    id := int32(1) // int32 | The ID of the object group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectGroupsApi.DeleteObjectGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.DeleteObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the object group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectGroupRequest struct via the builder pattern


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


## GetObjectGroup

> ObjectGroupSingle GetObjectGroup(ctx, id).Execute()

Get a single object group.



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
    id := int32(1) // int32 | The ID of the object group.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectGroupsApi.GetObjectGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.GetObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectGroup`: ObjectGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.GetObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the object group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBillByObjectGroup

> BillArray ListBillByObjectGroup(ctx, id).Page(page).Execute()

List all bills with this object group.



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
    resp, r, err := api_client.ObjectGroupsApi.ListBillByObjectGroup(context.Background(), id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListBillByObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillByObjectGroup`: BillArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListBillByObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBillByObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number. The default pagination is per 50 items. | 

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


## ListObjectGroups

> ObjectGroupArray ListObjectGroups(ctx).Page(page).Execute()

List all oject groups.



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
    resp, r, err := api_client.ObjectGroupsApi.ListObjectGroups(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListObjectGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjectGroups`: ObjectGroupArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListObjectGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListObjectGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number. The default pagination is 50. | 

### Return type

[**ObjectGroupArray**](ObjectGroupArray.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPiggyBankByObjectGroup

> PiggyBankArray ListPiggyBankByObjectGroup(ctx, id).Page(page).Execute()

List all piggy banks related to the object group.



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
    resp, r, err := api_client.ObjectGroupsApi.ListPiggyBankByObjectGroup(context.Background(), id).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.ListPiggyBankByObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPiggyBankByObjectGroup`: PiggyBankArray
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.ListPiggyBankByObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPiggyBankByObjectGroupRequest struct via the builder pattern


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


## UpdateObjectGroup

> ObjectGroupSingle UpdateObjectGroup(ctx, id).ObjectGroupUpdate(objectGroupUpdate).Execute()

Update existing object group.



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
    id := int32(1) // int32 | The ID of the object group
    objectGroupUpdate := *openapiclient.NewObjectGroupUpdate() // ObjectGroupUpdate | JSON array with updated piggy bank information. See the model for the exact specifications.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ObjectGroupsApi.UpdateObjectGroup(context.Background(), id).ObjectGroupUpdate(objectGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectGroupsApi.UpdateObjectGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectGroup`: ObjectGroupSingle
    fmt.Fprintf(os.Stdout, "Response from `ObjectGroupsApi.UpdateObjectGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the object group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectGroupUpdate** | [**ObjectGroupUpdate**](ObjectGroupUpdate.md) | JSON array with updated piggy bank information. See the model for the exact specifications. | 

### Return type

[**ObjectGroupSingle**](ObjectGroupSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/vnd.api+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

