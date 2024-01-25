# \AppUserAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignTeamAppUser**](AppUserAPI.md#AssignTeamAppUser) | **Post** /app-users/{appUserId}/teams/{teamId} | Assign App User to Team
[**DeleteAppUser**](AppUserAPI.md#DeleteAppUser) | **Delete** /app-users/{appUserId} | Delete App User
[**GetAppUser**](AppUserAPI.md#GetAppUser) | **Get** /app-users/{appUserId} | Get App User
[**ListAppUserRoles**](AppUserAPI.md#ListAppUserRoles) | **Get** /app-users/{appUserId}/roles | List Roles
[**UpdateAppUser**](AppUserAPI.md#UpdateAppUser) | **Patch** /app-users/{appUserId} | Update App User



## AssignTeamAppUser

> AppUserAssignResponse AssignTeamAppUser(ctx, appUserId, teamId).AppUserAssignRequest(appUserAssignRequest).Execute()

Assign App User to Team



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
    appUserId := "appUserId_example" // string | 
    teamId := "teamId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("RoleId_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppUserAPI.AssignTeamAppUser(context.Background(), appUserId, teamId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppUserAPI.AssignTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `AppUserAPI.AssignTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUserId** | **string** |  | 
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignTeamAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appUserAssignRequest** | [**AppUserAssignRequest**](AppUserAssignRequest.md) |  | 

### Return type

[**AppUserAssignResponse**](AppUserAssignResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/syncp"
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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

