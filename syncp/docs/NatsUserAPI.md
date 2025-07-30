# \NatsUserAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNatsUserTeamAppUser**](NatsUserAPI.md#AssignNatsUserTeamAppUser) | **Post** /core/beta/nats-users/{userId}/app-users/{teamAppUserId} | Assign Team App User to NATS User
[**CopyNatsUser**](NatsUserAPI.md#CopyNatsUser) | **Post** /core/beta/nats-users/{userId}/copy | Copy nats user
[**DeleteNatsUser**](NatsUserAPI.md#DeleteNatsUser) | **Delete** /core/beta/nats-users/{userId} | Delete NATS User
[**DownloadNatsUserBearerJwt**](NatsUserAPI.md#DownloadNatsUserBearerJwt) | **Post** /core/beta/nats-users/{userId}/bearer-jwt | Get Bearer JWT
[**DownloadNatsUserCreds**](NatsUserAPI.md#DownloadNatsUserCreds) | **Post** /core/beta/nats-users/{userId}/creds | Get Creds
[**DownloadNatsUserHttpGwToken**](NatsUserAPI.md#DownloadNatsUserHttpGwToken) | **Post** /core/beta/nats-users/{userId}/http-gw-token | Get HTTP Gateway Token
[**GetNatsUser**](NatsUserAPI.md#GetNatsUser) | **Get** /core/beta/nats-users/{userId} | Get NATS User
[**ListNatsUserConnections**](NatsUserAPI.md#ListNatsUserConnections) | **Get** /core/beta/nats-users/{userId}/connections | List NATs User Connections
[**ListNatsUserIssuances**](NatsUserAPI.md#ListNatsUserIssuances) | **Get** /core/beta/nats-users/{userId}/issuances | List nats user issuances
[**ListNatsUserTeamAppUsers**](NatsUserAPI.md#ListNatsUserTeamAppUsers) | **Get** /core/beta/nats-users/{userId}/app-users | List Team App Users
[**RotateNatsUser**](NatsUserAPI.md#RotateNatsUser) | **Post** /core/beta/nats-users/{userId}/rotate | Rotate nats user nkey and seed
[**UnAssignNatsUserTeamAppUser**](NatsUserAPI.md#UnAssignNatsUserTeamAppUser) | **Delete** /core/beta/nats-users/{userId}/app-users/{teamAppUserId} | Unassign Team App User from NATS User
[**UpdateNatsUser**](NatsUserAPI.md#UpdateNatsUser) | **Patch** /core/beta/nats-users/{userId} | Update NATS User



## AssignNatsUserTeamAppUser

> AppUserAssignResponse AssignNatsUserTeamAppUser(ctx, userId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()

Assign Team App User to NATS User



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
    userId := "userId_example" // string | 
    teamAppUserId := "teamAppUserId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("RoleId_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.AssignNatsUserTeamAppUser(context.Background(), userId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.AssignNatsUserTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNatsUserTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.AssignNatsUserTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignNatsUserTeamAppUserRequest struct via the builder pattern


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


## CopyNatsUser

> NatsUserViewResponse CopyNatsUser(ctx, userId).NatsUserCopyRequest(natsUserCopyRequest).Execute()

Copy nats user



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
    userId := "userId_example" // string | 
    natsUserCopyRequest := *openapiclient.NewNatsUserCopyRequest("Name_example", "SkGroupId_example") // NatsUserCopyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.CopyNatsUser(context.Background(), userId).NatsUserCopyRequest(natsUserCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.CopyNatsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyNatsUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.CopyNatsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyNatsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **natsUserCopyRequest** | [**NatsUserCopyRequest**](NatsUserCopyRequest.md) |  | 

### Return type

[**NatsUserViewResponse**](NatsUserViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNatsUser

> DeleteNatsUser(ctx, userId).Execute()

Delete NATS User



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NatsUserAPI.DeleteNatsUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.DeleteNatsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNatsUserRequest struct via the builder pattern


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


## DownloadNatsUserBearerJwt

> string DownloadNatsUserBearerJwt(ctx, userId).Execute()

Get Bearer JWT



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.DownloadNatsUserBearerJwt(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.DownloadNatsUserBearerJwt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadNatsUserBearerJwt`: string
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.DownloadNatsUserBearerJwt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadNatsUserBearerJwtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadNatsUserCreds

> string DownloadNatsUserCreds(ctx, userId).Execute()

Get Creds



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.DownloadNatsUserCreds(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.DownloadNatsUserCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadNatsUserCreds`: string
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.DownloadNatsUserCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadNatsUserCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadNatsUserHttpGwToken

> NatsUserHTTPGWTokenCreateReply DownloadNatsUserHttpGwToken(ctx, userId).Execute()

Get HTTP Gateway Token



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.DownloadNatsUserHttpGwToken(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.DownloadNatsUserHttpGwToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadNatsUserHttpGwToken`: NatsUserHTTPGWTokenCreateReply
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.DownloadNatsUserHttpGwToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadNatsUserHttpGwTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserHTTPGWTokenCreateReply**](NatsUserHTTPGWTokenCreateReply.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNatsUser

> NatsUserViewResponse GetNatsUser(ctx, userId).Execute()

Get NATS User



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.GetNatsUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.GetNatsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNatsUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.GetNatsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNatsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserViewResponse**](NatsUserViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNatsUserConnections

> NatsUserConnectionsListResponse ListNatsUserConnections(ctx, userId).Sort(sort).Cid(cid).State(state).Subject(subject).Limit(limit).Execute()

List NATs User Connections



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
    userId := "userId_example" // string | 
    sort := "sort_example" // string |  (optional)
    cid := "cid_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    subject := "subject_example" // string |  (optional)
    limit := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.ListNatsUserConnections(context.Background(), userId).Sort(sort).Cid(cid).State(state).Subject(subject).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.ListNatsUserConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNatsUserConnections`: NatsUserConnectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.ListNatsUserConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNatsUserConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** |  | 
 **cid** | **string** |  | 
 **state** | **string** |  | 
 **subject** | **string** |  | 
 **limit** | **float32** |  | 

### Return type

[**NatsUserConnectionsListResponse**](NatsUserConnectionsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNatsUserIssuances

> NatsUserIssuancesListResponse ListNatsUserIssuances(ctx, userId).Execute()

List nats user issuances



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.ListNatsUserIssuances(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.ListNatsUserIssuances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNatsUserIssuances`: NatsUserIssuancesListResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.ListNatsUserIssuances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNatsUserIssuancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserIssuancesListResponse**](NatsUserIssuancesListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNatsUserTeamAppUsers

> AppUserAssignListResponse ListNatsUserTeamAppUsers(ctx, userId).Execute()

List Team App Users



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.ListNatsUserTeamAppUsers(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.ListNatsUserTeamAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNatsUserTeamAppUsers`: AppUserAssignListResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.ListNatsUserTeamAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNatsUserTeamAppUsersRequest struct via the builder pattern


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


## RotateNatsUser

> NatsUserViewResponse RotateNatsUser(ctx, userId).Execute()

Rotate nats user nkey and seed



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.RotateNatsUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.RotateNatsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateNatsUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.RotateNatsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateNatsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserViewResponse**](NatsUserViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnAssignNatsUserTeamAppUser

> UnAssignNatsUserTeamAppUser(ctx, userId, teamAppUserId).Execute()

Unassign Team App User from NATS User



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
    userId := "userId_example" // string | 
    teamAppUserId := "teamAppUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NatsUserAPI.UnAssignNatsUserTeamAppUser(context.Background(), userId, teamAppUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.UnAssignNatsUserTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnAssignNatsUserTeamAppUserRequest struct via the builder pattern


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


## UpdateNatsUser

> NatsUserViewResponse UpdateNatsUser(ctx, userId).NatsUserUpdateRequest(natsUserUpdateRequest).Execute()

Update NATS User



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
    userId := "userId_example" // string | 
    natsUserUpdateRequest := *openapiclient.NewNatsUserUpdateRequest() // NatsUserUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NatsUserAPI.UpdateNatsUser(context.Background(), userId).NatsUserUpdateRequest(natsUserUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NatsUserAPI.UpdateNatsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNatsUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `NatsUserAPI.UpdateNatsUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNatsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **natsUserUpdateRequest** | [**NatsUserUpdateRequest**](NatsUserUpdateRequest.md) |  | 

### Return type

[**NatsUserViewResponse**](NatsUserViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

