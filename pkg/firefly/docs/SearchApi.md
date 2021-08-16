# \SearchApi

All URIs are relative to *https://demo.firefly-iii.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAccounts**](SearchApi.md#SearchAccounts) | **Get** /api/v1/search/accounts | Search for accounts
[**SearchTransactions**](SearchApi.md#SearchTransactions) | **Get** /api/v1/search/transactions | Search for transactions



## SearchAccounts

> AccountArray SearchAccounts(ctx).Query(query).Field(field).Page(page).Type_(type_).Execute()

Search for accounts



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
    query := "checking" // string | The query you wish to search for.
    field := openapiclient.AccountSearchFieldFilter("all") // AccountSearchFieldFilter | The account field(s) you want to search in.
    page := int32(1) // int32 | Page number. The default pagination is 50 (optional)
    type_ := openapiclient.AccountTypeFilter("all") // AccountTypeFilter | The type of accounts you wish to limit the search to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchAccounts(context.Background()).Query(query).Field(field).Page(page).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAccounts`: AccountArray
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The query you wish to search for. | 
 **field** | [**AccountSearchFieldFilter**](AccountSearchFieldFilter.md) | The account field(s) you want to search in. | 
 **page** | **int32** | Page number. The default pagination is 50 | 
 **type_** | [**AccountTypeFilter**](AccountTypeFilter.md) | The type of accounts you wish to limit the search to. | 

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


## SearchTransactions

> TransactionArray SearchTransactions(ctx).Query(query).Page(page).Execute()

Search for transactions



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
    query := "groceries" // string | The query you wish to search for.
    page := int32(1) // int32 | Page number. The default pagination is 50 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SearchApi.SearchTransactions(context.Background()).Query(query).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.SearchTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTransactions`: TransactionArray
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.SearchTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The query you wish to search for. | 
 **page** | **int32** | Page number. The default pagination is 50 | 

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

