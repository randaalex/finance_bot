# \AboutApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAbout**](AboutApi.md#GetAbout) | **Get** /api/v1/about | System information end point.
[**GetCron**](AboutApi.md#GetCron) | **Get** /api/v1/cron/{cliToken} | Cron job endpoint
[**GetCurrentUser**](AboutApi.md#GetCurrentUser) | **Get** /api/v1/about/user | Currently authenticated user endpoint.



## GetAbout

> SystemInfo GetAbout(ctx).Execute()

System information end point.



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
    resp, r, err := api_client.AboutApi.GetAbout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutApi.GetAbout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAbout`: SystemInfo
    fmt.Fprintf(os.Stdout, "Response from `AboutApi.GetAbout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAboutRequest struct via the builder pattern


### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCron

> CronResult GetCron(ctx, cliToken).Date(date).Force(force).Execute()

Cron job endpoint



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
    cliToken := "d5ea6b5fb774618dd6ad6ba6e0a7f55c" // string | The CLI token of any user in Firefly III, required to run the cron job.
    date := time.Now() // string | A date formatted YYYY-MM-DD. This can be used to make the cron job pretend it's running on another day.  (optional)
    force := false // bool | Forces the cron job to fire, regardless of whether it has fired before. This may result in double transactions or weird budgets, so be careful.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AboutApi.GetCron(context.Background(), cliToken).Date(date).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutApi.GetCron``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCron`: CronResult
    fmt.Fprintf(os.Stdout, "Response from `AboutApi.GetCron`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cliToken** | **string** | The CLI token of any user in Firefly III, required to run the cron job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | A date formatted YYYY-MM-DD. This can be used to make the cron job pretend it&#39;s running on another day.  | 
 **force** | **bool** | Forces the cron job to fire, regardless of whether it has fired before. This may result in double transactions or weird budgets, so be careful.  | 

### Return type

[**CronResult**](CronResult.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> UserSingle GetCurrentUser(ctx).Execute()

Currently authenticated user endpoint.



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
    resp, r, err := api_client.AboutApi.GetCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AboutApi.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: UserSingle
    fmt.Fprintf(os.Stdout, "Response from `AboutApi.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**UserSingle**](UserSingle.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

