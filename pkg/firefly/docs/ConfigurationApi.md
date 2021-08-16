# \ConfigurationApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationApi.md#GetConfiguration) | **Get** /api/v1/configuration | Get Firefly III system configuration values.
[**GetSingleConfiguration**](ConfigurationApi.md#GetSingleConfiguration) | **Get** /api/v1/configuration/{name} | Get a single Firefly III system configuration value
[**SetConfiguration**](ConfigurationApi.md#SetConfiguration) | **Put** /api/v1/configuration/{name} | Update configuration value



## GetConfiguration

> []Configuration GetConfiguration(ctx).Execute()

Get Firefly III system configuration values.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: []Configuration
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

[**[]Configuration**](Configuration.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleConfiguration

> ConfigurationSingle GetSingleConfiguration(ctx, name).Execute()

Get a single Firefly III system configuration value



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
    name := openapiclient.ConfigValueFilter("configuration.is_demo_site") // ConfigValueFilter | The name of the configuration value you want to know.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.GetSingleConfiguration(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSingleConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleConfiguration`: ConfigurationSingle
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSingleConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**ConfigValueFilter**](.md) | The name of the configuration value you want to know. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration

> ConfigurationSingle SetConfiguration(ctx, name).Value(value).Execute()

Update configuration value



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
    name := openapiclient.ConfigValueUpdateFilter("configuration.is_demo_site") // ConfigValueUpdateFilter | The name of the configuration value you want to update.
    value := TODO // string | Can be a number, a string, boolean or object. This depends on the actual configuration value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigurationApi.SetConfiguration(context.Background(), name).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.SetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetConfiguration`: ConfigurationSingle
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.SetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | [**ConfigValueUpdateFilter**](.md) | The name of the configuration value you want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | [**string**](oneOf&lt;boolean,string,object,array&gt;.md) | Can be a number, a string, boolean or object. This depends on the actual configuration value. | 

### Return type

[**ConfigurationSingle**](ConfigurationSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

