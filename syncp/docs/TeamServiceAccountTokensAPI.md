# \TeamServiceAccountTokensAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeamServiceAccountToken**](TeamServiceAccountTokensAPI.md#DeleteTeamServiceAccountToken) | **Delete** /service-account-tokens/team/{tokenId} | Delete Team Service Account Token
[**GetTeamServiceAccountToken**](TeamServiceAccountTokensAPI.md#GetTeamServiceAccountToken) | **Get** /service-account-tokens/team/{tokenId} | Get Team Service Account Token
[**UpdateTeamServiceAccountToken**](TeamServiceAccountTokensAPI.md#UpdateTeamServiceAccountToken) | **Patch** /service-account-tokens/team/{tokenId} | Update Team Service Account Token



## DeleteTeamServiceAccountToken

> DeleteTeamServiceAccountToken(ctx, tokenId).Execute()

Delete Team Service Account Token



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
    r, err := apiClient.TeamServiceAccountTokensAPI.DeleteTeamServiceAccountToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountTokensAPI.DeleteTeamServiceAccountToken``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTeamServiceAccountTokenRequest struct via the builder pattern


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


## GetTeamServiceAccountToken

> AppUserAccessTokenViewResponse GetTeamServiceAccountToken(ctx, tokenId).Execute()

Get Team Service Account Token



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
    resp, r, err := apiClient.TeamServiceAccountTokensAPI.GetTeamServiceAccountToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountTokensAPI.GetTeamServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamServiceAccountToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountTokensAPI.GetTeamServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamServiceAccountTokenRequest struct via the builder pattern


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


## UpdateTeamServiceAccountToken

> AppUserAccessTokenViewResponse UpdateTeamServiceAccountToken(ctx, tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()

Update Team Service Account Token



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
    resp, r, err := apiClient.TeamServiceAccountTokensAPI.UpdateTeamServiceAccountToken(context.Background(), tokenId).AppUserAccessTokenUpdateRequest(appUserAccessTokenUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountTokensAPI.UpdateTeamServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamServiceAccountToken`: AppUserAccessTokenViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountTokensAPI.UpdateTeamServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamServiceAccountTokenRequest struct via the builder pattern


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

