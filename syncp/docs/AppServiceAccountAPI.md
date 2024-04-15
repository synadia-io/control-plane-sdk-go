# \AppServiceAccountAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppServiceAccountToken**](AppServiceAccountAPI.md#CreateAppServiceAccountToken) | **Post** /service-accounts/app/{serviceAccountId}/tokens | Create Acess Token for App Service Account
[**DeleteAppServiceAccount**](AppServiceAccountAPI.md#DeleteAppServiceAccount) | **Delete** /service-accounts/app/{serviceAccountId} | Delete App Service Account
[**GetAppServiceAccount**](AppServiceAccountAPI.md#GetAppServiceAccount) | **Get** /service-accounts/app/{serviceAccountId} | Get App Service Account
[**ListAppServiceAccountTokens**](AppServiceAccountAPI.md#ListAppServiceAccountTokens) | **Get** /service-accounts/app/{serviceAccountId}/tokens | List Access Tokens for App Service Account
[**UpdateAppServiceAccount**](AppServiceAccountAPI.md#UpdateAppServiceAccount) | **Patch** /service-accounts/app/{serviceAccountId} | Update App Service Account



## CreateAppServiceAccountToken

> AppUserAccessTokenCreateResponse CreateAppServiceAccountToken(ctx, serviceAccountId).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()

Create Acess Token for App Service Account



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
    resp, r, err := apiClient.AppServiceAccountAPI.CreateAppServiceAccountToken(context.Background(), serviceAccountId).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountAPI.CreateAppServiceAccountToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppServiceAccountToken`: AppUserAccessTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountAPI.CreateAppServiceAccountToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceAccountTokenRequest struct via the builder pattern


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


## DeleteAppServiceAccount

> DeleteAppServiceAccount(ctx, serviceAccountId).Execute()

Delete App Service Account



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
    r, err := apiClient.AppServiceAccountAPI.DeleteAppServiceAccount(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountAPI.DeleteAppServiceAccount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAppServiceAccountRequest struct via the builder pattern


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


## GetAppServiceAccount

> ServiceAccountViewResponse GetAppServiceAccount(ctx, serviceAccountId).Execute()

Get App Service Account



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
    resp, r, err := apiClient.AppServiceAccountAPI.GetAppServiceAccount(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountAPI.GetAppServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountAPI.GetAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppServiceAccountRequest struct via the builder pattern


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


## ListAppServiceAccountTokens

> AppUserAccessTokenListResponse ListAppServiceAccountTokens(ctx, serviceAccountId).Execute()

List Access Tokens for App Service Account



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
    resp, r, err := apiClient.AppServiceAccountAPI.ListAppServiceAccountTokens(context.Background(), serviceAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountAPI.ListAppServiceAccountTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppServiceAccountTokens`: AppUserAccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountAPI.ListAppServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceAccountTokensRequest struct via the builder pattern


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


## UpdateAppServiceAccount

> ServiceAccountViewResponse UpdateAppServiceAccount(ctx, serviceAccountId).ServiceAccountUpdateRequest(serviceAccountUpdateRequest).Execute()

Update App Service Account



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
    resp, r, err := apiClient.AppServiceAccountAPI.UpdateAppServiceAccount(context.Background(), serviceAccountId).ServiceAccountUpdateRequest(serviceAccountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppServiceAccountAPI.UpdateAppServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppServiceAccountAPI.UpdateAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppServiceAccountRequest struct via the builder pattern


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

