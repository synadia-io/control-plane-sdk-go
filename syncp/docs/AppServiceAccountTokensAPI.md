# \AppServiceAccountTokensAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppServiceAccountToken**](AppServiceAccountTokensAPI.md#DeleteAppServiceAccountToken) | **Delete** /core/beta/service-account-tokens/app/{tokenId} | Delete App Service Account Token
[**GetAppServiceAccountToken**](AppServiceAccountTokensAPI.md#GetAppServiceAccountToken) | **Get** /core/beta/service-account-tokens/app/{tokenId} | Get App Service Account Token
[**UpdateAppServiceAccountToken**](AppServiceAccountTokensAPI.md#UpdateAppServiceAccountToken) | **Patch** /core/beta/service-account-tokens/app/{tokenId} | Update App Service Account Token



## DeleteAppServiceAccountToken

> DeleteAppServiceAccountToken(ctx, tokenId).Execute()

Delete App Service Account Token



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
    r, err := apiClient.AppServiceAccountTokensAPI.DeleteAppServiceAccountToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountTokensAPI.DeleteAppServiceAccountToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppServiceAccountTokenRequest struct via the builder pattern


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


## GetAppServiceAccountToken

> AppUserAccessTokenViewResponse GetAppServiceAccountToken(ctx, tokenId).Execute()

Get App Service Account Token



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
    resp, r, err := apiClient.AppServiceAccountTokensAPI.GetAppServiceAccountToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountTokensAPI.GetAppServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppServiceAccountToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountTokensAPI.GetAppServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceAccountTokenRequest struct via the builder pattern


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


## UpdateAppServiceAccountToken

> AppUserAccessTokenViewResponse UpdateAppServiceAccountToken(ctx, tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()

Update App Service Account Token



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
    resp, r, err := apiClient.AppServiceAccountTokensAPI.UpdateAppServiceAccountToken(context.Background(), tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountTokensAPI.UpdateAppServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppServiceAccountToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountTokensAPI.UpdateAppServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceAccountTokenRequest struct via the builder pattern


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

