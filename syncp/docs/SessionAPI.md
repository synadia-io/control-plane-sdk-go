# \SessionAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptTerms**](SessionAPI.md#AcceptTerms) | **Post** /terms/accept | Accept terms
[**CreateAppUser**](SessionAPI.md#CreateAppUser) | **Post** /app-users | Create App User
[**CreatePersonalAccessToken**](SessionAPI.md#CreatePersonalAccessToken) | **Post** /personal-access-tokens | Create Personal Access Token
[**CreateTeam**](SessionAPI.md#CreateTeam) | **Post** /teams | Create Team
[**GetVersion**](SessionAPI.md#GetVersion) | **Get** /version | Get Version
[**ListAlerts**](SessionAPI.md#ListAlerts) | **Get** /alerts | List Alerts
[**ListAppUsers**](SessionAPI.md#ListAppUsers) | **Get** /app-users | List App Users
[**ListPersonalAccessTokens**](SessionAPI.md#ListPersonalAccessTokens) | **Get** /personal-access-tokens | List Personal Access Tokens
[**ListTeams**](SessionAPI.md#ListTeams) | **Get** /teams | List Teams
[**SearchSystemAccounts**](SessionAPI.md#SearchSystemAccounts) | **Get** /search/systems/{systemId}/accounts | Search System Accounts
[**SearchSystemServers**](SessionAPI.md#SearchSystemServers) | **Get** /search/systems/{systemId}/servers | Search System Servers
[**SearchTeamAppUsers**](SessionAPI.md#SearchTeamAppUsers) | **Get** /search/teams/{teamId}/app-users | Search App Users in Team



## AcceptTerms

> AcceptTermsResponse AcceptTerms(ctx).Execute()

Accept terms



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
    resp, r, err := apiClient.SessionAPI.AcceptTerms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.AcceptTerms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptTerms`: AcceptTermsResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.AcceptTerms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptTermsRequest struct via the builder pattern


### Return type

[**AcceptTermsResponse**](AcceptTermsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## CreateTeam

> TeamViewResponse CreateTeam(ctx).TeamCreateRequest(teamCreateRequest).Execute()

Create Team



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
    teamCreateRequest := *openapiclient.NewTeamCreateRequest("Name_example") // TeamCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateTeam(context.Background()).TeamCreateRequest(teamCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeam`: TeamViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamCreateRequest** | [**TeamCreateRequest**](TeamCreateRequest.md) |  | 

### Return type

[**TeamViewResponse**](TeamViewResponse.md)

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


## ListTeams

> TeamListResponse ListTeams(ctx).Execute()

List Teams



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
    resp, r, err := apiClient.SessionAPI.ListTeams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeams`: TeamListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsRequest struct via the builder pattern


### Return type

[**TeamListResponse**](TeamListResponse.md)

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


## SearchTeamAppUsers

> TeamAppUserListResponse SearchTeamAppUsers(ctx, teamId).Execute()

Search App Users in Team



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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.SearchTeamAppUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SearchTeamAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTeamAppUsers`: TeamAppUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.SearchTeamAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTeamAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamAppUserListResponse**](TeamAppUserListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

