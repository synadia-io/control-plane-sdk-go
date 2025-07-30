# \PersonalAccessTokenAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePersonalAccessToken**](PersonalAccessTokenAPI.md#DeletePersonalAccessToken) | **Delete** /core/beta/personal-access-tokens/{tokenId} | Delete Personal Access Token
[**GetPersonalAccessToken**](PersonalAccessTokenAPI.md#GetPersonalAccessToken) | **Get** /core/beta/personal-access-tokens/{tokenId} | Get Personal Access Token
[**UpdatePersonalAccessToken**](PersonalAccessTokenAPI.md#UpdatePersonalAccessToken) | **Patch** /core/beta/personal-access-tokens/{tokenId} | Update Personal Access Token



## DeletePersonalAccessToken

> DeletePersonalAccessToken(ctx, tokenId).Execute()

Delete Personal Access Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
)

func main() {
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PersonalAccessTokenAPI.DeletePersonalAccessToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.DeletePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalAccessToken

> AppUserAccessTokenViewResponse GetPersonalAccessToken(ctx, tokenId).Execute()

Get Personal Access Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
)

func main() {
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokenAPI.GetPersonalAccessToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.GetPersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalAccessToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenAPI.GetPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUserAccessTokenViewResponse**](AppUserAccessTokenViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePersonalAccessToken

> AppUserAccessTokenViewResponse UpdatePersonalAccessToken(ctx, tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()

Update Personal Access Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
)

func main() {
    tokenId := "tokenId_example" // string | 
    appUserAccessTokenUpdateRequest := *openapiclient.NewAppUserAccessTokenUpdateRequest() // AppUserAccessTokenUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokenAPI.UpdatePersonalAccessToken(context.Background(), tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokenAPI.UpdatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePersonalAccessToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokenAPI.UpdatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserAccessTokenUpdateRequest** | [**AppUserAccessTokenUpdateRequest**](AppUserAccessTokenUpdateRequest.md) |  | 

### Return type

[**AppUserAccessTokenViewResponse**](AppUserAccessTokenViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

