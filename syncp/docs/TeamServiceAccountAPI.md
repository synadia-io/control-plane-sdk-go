# \TeamServiceAccountAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeamServiceAccountToken**](TeamServiceAccountAPI.md#CreateTeamServiceAccountToken) | **Post** /core/beta/service-accounts/team/{serviceAccountId}/tokens | Create Acess Token for Team Service Account
[**DeleteTeamServiceAccount**](TeamServiceAccountAPI.md#DeleteTeamServiceAccount) | **Delete** /core/beta/service-accounts/team/{serviceAccountId} | Delete Team Service Account
[**GetTeamServiceAccount**](TeamServiceAccountAPI.md#GetTeamServiceAccount) | **Get** /core/beta/service-accounts/team/{serviceAccountId} | Get Team Service Account
[**ListTeamServiceAccountTokens**](TeamServiceAccountAPI.md#ListTeamServiceAccountTokens) | **Get** /core/beta/service-accounts/team/{serviceAccountId}/tokens | List Access Tokens for Team Service Account
[**UpdateTeamServiceAccount**](TeamServiceAccountAPI.md#UpdateTeamServiceAccount) | **Patch** /core/beta/service-accounts/team/{serviceAccountId} | Update Team Service Account



## CreateTeamServiceAccountToken

> AppUserAccessTokenCreateResponse CreateTeamServiceAccountToken(ctx, serviceAccountId).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()

Create Acess Token for Team Service Account



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
    serviceAccountId := "serviceAccountId_example" // string | 
    appUserAccessTokenCreateRequest := *openapiclient.NewAppUserAccessTokenCreateRequest("Name_example") // AppUserAccessTokenCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamServiceAccountAPI.CreateTeamServiceAccountToken(context.Background(), serviceAccountId).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountAPI.CreateTeamServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamServiceAccountToken`: AppUserAccessTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountAPI.CreateTeamServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamServiceAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserAccessTokenCreateRequest** | [**AppUserAccessTokenCreateRequest**](AppUserAccessTokenCreateRequest.md) |  | 

### Return type

[**AppUserAccessTokenCreateResponse**](AppUserAccessTokenCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamServiceAccount

> DeleteTeamServiceAccount(ctx, serviceAccountId).Execute()

Delete Team Service Account



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
    serviceAccountId := "serviceAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamServiceAccountAPI.DeleteTeamServiceAccount(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountAPI.DeleteTeamServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamServiceAccountRequest struct via the builder pattern


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


## GetTeamServiceAccount

> ServiceAccountViewResponse GetTeamServiceAccount(ctx, serviceAccountId).Execute()

Get Team Service Account



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
    serviceAccountId := "serviceAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamServiceAccountAPI.GetTeamServiceAccount(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountAPI.GetTeamServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountAPI.GetTeamServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccountViewResponse**](ServiceAccountViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamServiceAccountTokens

> AppUserAccessTokenListResponse ListTeamServiceAccountTokens(ctx, serviceAccountId).Execute()

List Access Tokens for Team Service Account



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
    serviceAccountId := "serviceAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamServiceAccountAPI.ListTeamServiceAccountTokens(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountAPI.ListTeamServiceAccountTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamServiceAccountTokens`: AppUserAccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountAPI.ListTeamServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamServiceAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUserAccessTokenListResponse**](AppUserAccessTokenListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamServiceAccount

> ServiceAccountViewResponse UpdateTeamServiceAccount(ctx, serviceAccountId).ServiceAccountUpdateRequest(serviceAccountUpdateRequest).Execute()

Update Team Service Account



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
    serviceAccountId := "serviceAccountId_example" // string | 
    serviceAccountUpdateRequest := *openapiclient.NewServiceAccountUpdateRequest() // ServiceAccountUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamServiceAccountAPI.UpdateTeamServiceAccount(context.Background(), serviceAccountId).ServiceAccountUpdateRequest(serviceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamServiceAccountAPI.UpdateTeamServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamServiceAccountAPI.UpdateTeamServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAccountUpdateRequest** | [**ServiceAccountUpdateRequest**](ServiceAccountUpdateRequest.md) |  | 

### Return type

[**ServiceAccountViewResponse**](ServiceAccountViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

