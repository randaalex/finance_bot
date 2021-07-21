# \DataApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DestroyData**](DataApi.md#DestroyData) | **Delete** /api/v1/data/destroy | Endpoint to destroy user data



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

