# \SessionAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppUser**](SessionAPI.md#CreateAppUser) | **Post** /app-users | Create App User
[**CreatePersonalAccessToken**](SessionAPI.md#CreatePersonalAccessToken) | **Post** /personal-access-tokens | Create Personal Access Token
[**CreateSystem**](SessionAPI.md#CreateSystem) | **Post** /systems | Create System
[**GetVersion**](SessionAPI.md#GetVersion) | **Get** /version | Get Version
[**ImportSystem**](SessionAPI.md#ImportSystem) | **Post** /import-system | Import a System
[**ListAlerts**](SessionAPI.md#ListAlerts) | **Get** /alerts | List Alerts
[**ListAppUsers**](SessionAPI.md#ListAppUsers) | **Get** /app-users | List App Users
[**ListPersonalAccessTokens**](SessionAPI.md#ListPersonalAccessTokens) | **Get** /personal-access-tokens | List Personal Access Tokens
[**ListSessionAccounts**](SessionAPI.md#ListSessionAccounts) | **Get** /accounts | List Accounts
[**ListSessionNatsUsers**](SessionAPI.md#ListSessionNatsUsers) | **Get** /nats-users | List NATS Users
[**ListSystems**](SessionAPI.md#ListSystems) | **Get** /systems | List Systems
[**SearchAppUsers**](SessionAPI.md#SearchAppUsers) | **Get** /search/app-users | Search App Users
[**SearchSystemAccounts**](SessionAPI.md#SearchSystemAccounts) | **Get** /search/systems/{systemId}/accounts | Search System Accounts
[**SearchSystemServers**](SessionAPI.md#SearchSystemServers) | **Get** /search/systems/{systemId}/servers | Search System Servers



## CreateAppUser

> AppUserCreateResponse CreateAppUser(ctx).AppUserCreateRequest(appUserCreateRequest).Execute()

Create App User



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
    appUserCreateRequest := *openapiclient.NewAppUserCreateRequest("Name_example", "Role_example", "Username_example") // AppUserCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateAppUser(context.Background()).AppUserCreateRequest(appUserCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppUser`: AppUserCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateAppUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appUserCreateRequest** | [**AppUserCreateRequest**](AppUserCreateRequest.md) |  | 

### Return type

[**AppUserCreateResponse**](AppUserCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePersonalAccessToken

> AppUserAccessTokenCreateResponse CreatePersonalAccessToken(ctx).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()

Create Personal Access Token



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
    appUserAccessTokenCreateRequest := *openapiclient.NewAppUserAccessTokenCreateRequest("Expires_example", "Name_example") // AppUserAccessTokenCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreatePersonalAccessToken(context.Background()).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersonalAccessToken`: AppUserAccessTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appUserAccessTokenCreateRequest** | [**AppUserAccessTokenCreateRequest**](AppUserAccessTokenCreateRequest.md) |  | 

### Return type

[**AppUserAccessTokenCreateResponse**](AppUserAccessTokenCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSystem

> SystemViewResponse CreateSystem(ctx).SystemCreateRequest(systemCreateRequest).Execute()

Create System



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
    systemCreateRequest := *openapiclient.NewSystemCreateRequest("Name_example") // SystemCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateSystem(context.Background()).SystemCreateRequest(systemCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemCreateRequest** | [**SystemCreateRequest**](SystemCreateRequest.md) |  | 

### Return type

[**SystemViewResponse**](SystemViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> VersionResponse GetVersion(ctx).Execute()

Get Version



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
    resp, r, err := apiClient.SessionAPI.GetVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: VersionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


### Return type

[**VersionResponse**](VersionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSystem

> SystemViewResponse ImportSystem(ctx).SystemImportRequest(systemImportRequest).Execute()

Import a System



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
    systemImportRequest := *openapiclient.NewSystemImportRequest("OperatorJwt_example", "OperatorKey_example", "SystemJwt_example", "SystemKey_example") // SystemImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.ImportSystem(context.Background()).SystemImportRequest(systemImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ImportSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ImportSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemImportRequest** | [**SystemImportRequest**](SystemImportRequest.md) |  | 

### Return type

[**SystemViewResponse**](SystemViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlerts

> AlertListResponse ListAlerts(ctx).Status(status).Execute()

List Alerts



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
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.ListAlerts(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlerts`: AlertListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 

### Return type

[**AlertListResponse**](AlertListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppUsers

> AppUserListResponse ListAppUsers(ctx).Execute()

List App Users



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
    resp, r, err := apiClient.SessionAPI.ListAppUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppUsers`: AppUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListAppUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppUsersRequest struct via the builder pattern


### Return type

[**AppUserListResponse**](AppUserListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersonalAccessTokens

> AppUserAccessTokenListResponse ListPersonalAccessTokens(ctx).Execute()

List Personal Access Tokens



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
    resp, r, err := apiClient.SessionAPI.ListPersonalAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListPersonalAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersonalAccessTokens`: AppUserAccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListPersonalAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalAccessTokensRequest struct via the builder pattern


### Return type

[**AppUserAccessTokenListResponse**](AppUserAccessTokenListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionAccounts

> SessionAccountListResponse ListSessionAccounts(ctx).Execute()

List Accounts



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
    resp, r, err := apiClient.SessionAPI.ListSessionAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListSessionAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionAccounts`: SessionAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListSessionAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionAccountsRequest struct via the builder pattern


### Return type

[**SessionAccountListResponse**](SessionAccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionNatsUsers

> SessionNatsUserListResponse ListSessionNatsUsers(ctx).Execute()

List NATS Users



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
    resp, r, err := apiClient.SessionAPI.ListSessionNatsUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListSessionNatsUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionNatsUsers`: SessionNatsUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListSessionNatsUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionNatsUsersRequest struct via the builder pattern


### Return type

[**SessionNatsUserListResponse**](SessionNatsUserListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystems

> SessionSystemListResponse ListSystems(ctx).Execute()

List Systems



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
    resp, r, err := apiClient.SessionAPI.ListSystems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystems`: SessionSystemListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListSystems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemsRequest struct via the builder pattern


### Return type

[**SessionSystemListResponse**](SessionSystemListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAppUsers

> AppUserInfoListResponse SearchAppUsers(ctx).Execute()

Search App Users



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
    resp, r, err := apiClient.SessionAPI.SearchAppUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SearchAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAppUsers`: AppUserInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.SearchAppUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAppUsersRequest struct via the builder pattern


### Return type

[**AppUserInfoListResponse**](AppUserInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSystemAccounts

> AccountSearchListResponse SearchSystemAccounts(ctx, systemId).Execute()

Search System Accounts



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
    systemId := "systemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.SearchSystemAccounts(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SearchSystemAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSystemAccounts`: AccountSearchListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.SearchSystemAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSystemAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountSearchListResponse**](AccountSearchListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSystemServers

> NatsServerInfoListResponse SearchSystemServers(ctx, systemId).Execute()

Search System Servers



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
    systemId := "systemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.SearchSystemServers(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SearchSystemServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSystemServers`: NatsServerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.SearchSystemServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSystemServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsServerInfoListResponse**](NatsServerInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

