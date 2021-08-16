# \SummaryApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBasicSummary**](SummaryApi.md#GetBasicSummary) | **Get** /api/v1/summary/basic | Returns basic sums of the users data.



## GetBasicSummary

> []BasicSummaryEntry GetBasicSummary(ctx).Start(start).End(end).CurrencyCode(currencyCode).Execute()

Returns basic sums of the users data.



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
    start := time.Now() // string | A date formatted YYYY-MM-DD. 
    end := time.Now() // string | A date formatted YYYY-MM-DD. 
    currencyCode := "currencyCode_example" // string | A currency code like EUR or USD, to filter the result.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SummaryApi.GetBasicSummary(context.Background()).Start(start).End(end).CurrencyCode(currencyCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SummaryApi.GetBasicSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBasicSummary`: []BasicSummaryEntry
    fmt.Fprintf(os.Stdout, "Response from `SummaryApi.GetBasicSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBasicSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | A date formatted YYYY-MM-DD.  | 
 **end** | **string** | A date formatted YYYY-MM-DD.  | 
 **currencyCode** | **string** | A currency code like EUR or USD, to filter the result.  | 

### Return type

[**[]BasicSummaryEntry**](BasicSummaryEntry.md)

### Authorization

[firefly_iii_auth](../README.md#firefly_iii_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

