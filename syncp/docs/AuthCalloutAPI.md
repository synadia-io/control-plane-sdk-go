# \AuthCalloutAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAuthCalloutTargetAccount**](AuthCalloutAPI.md#AddAuthCalloutTargetAccount) | **Post** /core/beta/auth-callout/{authCalloutId}/target-accounts | Configure Target Account
[**AddAuthCalloutUser**](AuthCalloutAPI.md#AddAuthCalloutUser) | **Post** /core/beta/auth-callout/{authCalloutId}/users | Create Auth Callout User
[**DeleteAuthCallout**](AuthCalloutAPI.md#DeleteAuthCallout) | **Delete** /core/beta/auth-callout/{authCalloutId} | Delete Auth Callout Config
[**DeleteAuthCalloutTargetAccount**](AuthCalloutAPI.md#DeleteAuthCalloutTargetAccount) | **Delete** /core/beta/auth-callout/{authCalloutId}/target-accounts/{targetAccountId} | Delete Target Account
[**DeleteAuthCalloutUser**](AuthCalloutAPI.md#DeleteAuthCalloutUser) | **Delete** /core/beta/auth-callout/{authCalloutId}/users/{acUserId} | Delete Control Account User
[**GetAuthCallout**](AuthCalloutAPI.md#GetAuthCallout) | **Get** /core/beta/auth-callout/{authCalloutId} | Auth Callout Config
[**ListAuthCalloutTargetAccounts**](AuthCalloutAPI.md#ListAuthCalloutTargetAccounts) | **Get** /core/beta/auth-callout/{authCalloutId}/target-accounts | Get Target Account List
[**ListAuthCalloutUsers**](AuthCalloutAPI.md#ListAuthCalloutUsers) | **Get** /core/beta/auth-callout/{authCalloutId}/users | Get Target Account List



## AddAuthCalloutTargetAccount

> AddAuthCalloutTargetAccount(ctx, authCalloutId).AuthCalloutAddTargetAccountRequest(authCalloutAddTargetAccountRequest).Execute()

Configure Target Account



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
    authCalloutId := "authCalloutId_example" // string | 
    authCalloutAddTargetAccountRequest := *openapiclient.NewAuthCalloutAddTargetAccountRequest("AccountId_example", "ControlAccountSkGroupId_example", "SkGroupId_example") // AuthCalloutAddTargetAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthCalloutAPI.AddAuthCalloutTargetAccount(context.Background(), authCalloutId).AuthCalloutAddTargetAccountRequest(authCalloutAddTargetAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.AddAuthCalloutTargetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthCalloutTargetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authCalloutAddTargetAccountRequest** | [**AuthCalloutAddTargetAccountRequest**](AuthCalloutAddTargetAccountRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAuthCalloutUser

> AddAuthCalloutUser(ctx, authCalloutId).AuthCalloutAddUserRequest(authCalloutAddUserRequest).Execute()

Create Auth Callout User



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
    authCalloutId := "authCalloutId_example" // string | 
    authCalloutAddUserRequest := *openapiclient.NewAuthCalloutAddUserRequest("NatsUserId_example") // AuthCalloutAddUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthCalloutAPI.AddAuthCalloutUser(context.Background(), authCalloutId).AuthCalloutAddUserRequest(authCalloutAddUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.AddAuthCalloutUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAuthCalloutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authCalloutAddUserRequest** | [**AuthCalloutAddUserRequest**](AuthCalloutAddUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAuthCallout

> DeleteAuthCallout(ctx, authCalloutId).Execute()

Delete Auth Callout Config



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
    authCalloutId := "authCalloutId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthCalloutAPI.DeleteAuthCallout(context.Background(), authCalloutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.DeleteAuthCallout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCalloutRequest struct via the builder pattern


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


## DeleteAuthCalloutTargetAccount

> DeleteAuthCalloutTargetAccount(ctx, authCalloutId, targetAccountId).Execute()

Delete Target Account



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
    authCalloutId := "authCalloutId_example" // string | 
    targetAccountId := "targetAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthCalloutAPI.DeleteAuthCalloutTargetAccount(context.Background(), authCalloutId, targetAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.DeleteAuthCalloutTargetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 
**targetAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCalloutTargetAccountRequest struct via the builder pattern


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


## DeleteAuthCalloutUser

> DeleteAuthCalloutUser(ctx, authCalloutId, acUserId).Execute()

Delete Control Account User



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
    authCalloutId := "authCalloutId_example" // string | 
    acUserId := "acUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthCalloutAPI.DeleteAuthCalloutUser(context.Background(), authCalloutId, acUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.DeleteAuthCalloutUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 
**acUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthCalloutUserRequest struct via the builder pattern


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


## GetAuthCallout

> AuthCalloutViewResponse GetAuthCallout(ctx, authCalloutId).Execute()

Auth Callout Config



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
    authCalloutId := "authCalloutId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCalloutAPI.GetAuthCallout(context.Background(), authCalloutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.GetAuthCallout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthCallout`: AuthCalloutViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthCalloutAPI.GetAuthCallout`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthCalloutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthCalloutViewResponse**](AuthCalloutViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthCalloutTargetAccounts

> AuthCalloutTargetAccountListResponse ListAuthCalloutTargetAccounts(ctx, authCalloutId).Execute()

Get Target Account List



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
    authCalloutId := "authCalloutId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCalloutAPI.ListAuthCalloutTargetAccounts(context.Background(), authCalloutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.ListAuthCalloutTargetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthCalloutTargetAccounts`: AuthCalloutTargetAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthCalloutAPI.ListAuthCalloutTargetAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthCalloutTargetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthCalloutTargetAccountListResponse**](AuthCalloutTargetAccountListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthCalloutUsers

> AuthCalloutUsersListResponse ListAuthCalloutUsers(ctx, authCalloutId).Execute()

Get Target Account List



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
    authCalloutId := "authCalloutId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthCalloutAPI.ListAuthCalloutUsers(context.Background(), authCalloutId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthCalloutAPI.ListAuthCalloutUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthCalloutUsers`: AuthCalloutUsersListResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthCalloutAPI.ListAuthCalloutUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCalloutId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthCalloutUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthCalloutUsersListResponse**](AuthCalloutUsersListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

