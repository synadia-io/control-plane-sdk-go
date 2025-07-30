# \OidcProviderAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOidcProvider**](OidcProviderAPI.md#CreateOidcProvider) | **Post** /core/beta/oidc-providers | Create OIDC Provider
[**DeleteOidcProvider**](OidcProviderAPI.md#DeleteOidcProvider) | **Delete** /core/beta/oidc-providers/{oidcProviderId} | Delete OIDC Provider
[**GetOidcProvider**](OidcProviderAPI.md#GetOidcProvider) | **Get** /core/beta/oidc-providers/{oidcProviderId} | Get OIDC Provider
[**ListOidcDefaultProvider**](OidcProviderAPI.md#ListOidcDefaultProvider) | **Get** /core/beta/oidc-providers/defaults | List OIDC Provider Defaults
[**ListOidcProviders**](OidcProviderAPI.md#ListOidcProviders) | **Get** /core/beta/oidc-providers | List OIDC Providers
[**UpdateOidcProvider**](OidcProviderAPI.md#UpdateOidcProvider) | **Patch** /core/beta/oidc-providers/{oidcProviderId} | Update OIDC Provider



## CreateOidcProvider

> OidcProviderViewResponse CreateOidcProvider(ctx).OidcProviderAddRequest(oidcProviderAddRequest).Execute()

Create OIDC Provider



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
    oidcProviderAddRequest := *openapiclient.NewOidcProviderAddRequest("Domain_example", false) // OidcProviderAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcProviderAPI.CreateOidcProvider(context.Background()).OidcProviderAddRequest(oidcProviderAddRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.CreateOidcProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOidcProvider`: OidcProviderViewResponse
    fmt.Fprintf(os.Stdout, "Response from `OidcProviderAPI.CreateOidcProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oidcProviderAddRequest** | [**OidcProviderAddRequest**](OidcProviderAddRequest.md) |  | 

### Return type

[**OidcProviderViewResponse**](OidcProviderViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOidcProvider

> DeleteOidcProvider(ctx, oidcProviderId).Execute()

Delete OIDC Provider



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
    oidcProviderId := "oidcProviderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OidcProviderAPI.DeleteOidcProvider(context.Background(), oidcProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.DeleteOidcProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oidcProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOidcProviderRequest struct via the builder pattern


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


## GetOidcProvider

> OidcProviderViewResponse GetOidcProvider(ctx, oidcProviderId).Execute()

Get OIDC Provider



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
    oidcProviderId := "oidcProviderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcProviderAPI.GetOidcProvider(context.Background(), oidcProviderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.GetOidcProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOidcProvider`: OidcProviderViewResponse
    fmt.Fprintf(os.Stdout, "Response from `OidcProviderAPI.GetOidcProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oidcProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OidcProviderViewResponse**](OidcProviderViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOidcDefaultProvider

> OidcProviderDefaultsListResponse ListOidcDefaultProvider(ctx).Execute()

List OIDC Provider Defaults



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcProviderAPI.ListOidcDefaultProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.ListOidcDefaultProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOidcDefaultProvider`: OidcProviderDefaultsListResponse
    fmt.Fprintf(os.Stdout, "Response from `OidcProviderAPI.ListOidcDefaultProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOidcDefaultProviderRequest struct via the builder pattern


### Return type

[**OidcProviderDefaultsListResponse**](OidcProviderDefaultsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOidcProviders

> OidcProviderListResponse ListOidcProviders(ctx).Execute()

List OIDC Providers



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcProviderAPI.ListOidcProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.ListOidcProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOidcProviders`: OidcProviderListResponse
    fmt.Fprintf(os.Stdout, "Response from `OidcProviderAPI.ListOidcProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOidcProvidersRequest struct via the builder pattern


### Return type

[**OidcProviderListResponse**](OidcProviderListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOidcProvider

> OidcProviderViewResponse UpdateOidcProvider(ctx, oidcProviderId).OidcProviderUpdateRequest(oidcProviderUpdateRequest).Execute()

Update OIDC Provider



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
    oidcProviderId := "oidcProviderId_example" // string | 
    oidcProviderUpdateRequest := *openapiclient.NewOidcProviderUpdateRequest() // OidcProviderUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OidcProviderAPI.UpdateOidcProvider(context.Background(), oidcProviderId).OidcProviderUpdateRequest(oidcProviderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OidcProviderAPI.UpdateOidcProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOidcProvider`: OidcProviderViewResponse
    fmt.Fprintf(os.Stdout, "Response from `OidcProviderAPI.UpdateOidcProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oidcProviderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOidcProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oidcProviderUpdateRequest** | [**OidcProviderUpdateRequest**](OidcProviderUpdateRequest.md) |  | 

### Return type

[**OidcProviderViewResponse**](OidcProviderViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

