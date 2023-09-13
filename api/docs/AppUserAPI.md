# \AppUserAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppUser**](AppUserAPI.md#DeleteAppUser) | **Delete** /app-users/{appUserId} | Delete App User
[**GetAppUser**](AppUserAPI.md#GetAppUser) | **Get** /app-users/{appUserId} | Get App User
[**ListAppUserRoles**](AppUserAPI.md#ListAppUserRoles) | **Get** /app-users/{appUserId}/roles | List Roles
[**UpdateAppUser**](AppUserAPI.md#UpdateAppUser) | **Patch** /app-users/{appUserId} | Update App User



## DeleteAppUser

> DeleteAppUser(ctx, appUserId).Execute()

Delete App User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    appUserId := "appUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppUserAPI.DeleteAppUser(context.Background(), appUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserAPI.DeleteAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppUser

> AppUserViewResponse GetAppUser(ctx, appUserId).Execute()

Get App User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    appUserId := "appUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUserAPI.GetAppUser(context.Background(), appUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserAPI.GetAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppUser`: AppUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AppUserAPI.GetAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUserViewResponse**](AppUserViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppUserRoles

> AppUserAssignListResponse ListAppUserRoles(ctx, appUserId).Execute()

List Roles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    appUserId := "appUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUserAPI.ListAppUserRoles(context.Background(), appUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserAPI.ListAppUserRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppUserRoles`: AppUserAssignListResponse
    fmt.Fprintf(os.Stdout, "Response from `AppUserAPI.ListAppUserRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAppUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppUserAssignListResponse**](AppUserAssignListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppUser

> AppUserUpdateResponse UpdateAppUser(ctx, appUserId).AppUserUpdateRequest(appUserUpdateRequest).Execute()

Update App User



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    appUserId := "appUserId_example" // string | 
    appUserUpdateRequest := *openapiclient.NewAppUserUpdateRequest() // AppUserUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUserAPI.UpdateAppUser(context.Background(), appUserId).AppUserUpdateRequest(appUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserAPI.UpdateAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppUser`: AppUserUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `AppUserAPI.UpdateAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserUpdateRequest** | [**AppUserUpdateRequest**](AppUserUpdateRequest.md) |  | 

### Return type

[**AppUserUpdateResponse**](AppUserUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

