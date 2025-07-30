# \SessionAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptTerms**](SessionAPI.md#AcceptTerms) | **Post** /core/beta/terms/accept | Accept terms
[**CreateAppServiceAccount**](SessionAPI.md#CreateAppServiceAccount) | **Post** /core/beta/service-accounts/app | Create App Service Account
[**CreateAppUser**](SessionAPI.md#CreateAppUser) | **Post** /core/beta/app-users | Create App User
[**CreateAppUserAccessToken**](SessionAPI.md#CreateAppUserAccessToken) | **Post** /core/beta/personal-access-tokens | Create Personal Access Token
[**CreateTeam**](SessionAPI.md#CreateTeam) | **Post** /core/beta/teams | Create Team
[**DecideInvitation**](SessionAPI.md#DecideInvitation) | **Post** /core/beta/invitations/{teamId} | Accept or reject team invitation
[**GetVersion**](SessionAPI.md#GetVersion) | **Get** /core/beta/version | Get Version
[**ListAlerts**](SessionAPI.md#ListAlerts) | **Get** /core/beta/alerts | List Alerts
[**ListAppServiceAccounts**](SessionAPI.md#ListAppServiceAccounts) | **Get** /core/beta/service-accounts/app | List App Service Accounts
[**ListAppUserAccessTokens**](SessionAPI.md#ListAppUserAccessTokens) | **Get** /core/beta/personal-access-tokens | List Personal Access Tokens
[**ListAppUsers**](SessionAPI.md#ListAppUsers) | **Get** /core/beta/app-users | List App Users
[**ListInvitations**](SessionAPI.md#ListInvitations) | **Get** /core/beta/invitations | List of pending invitations
[**ListTeams**](SessionAPI.md#ListTeams) | **Get** /core/beta/teams | List Teams



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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppServiceAccount

> ServiceAccountViewResponse CreateAppServiceAccount(ctx).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()

Create App Service Account



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
    serviceAccountCreateRequest := *openapiclient.NewServiceAccountCreateRequest("Name_example", map[string]AppUserAssignRequest{"key": *openapiclient.NewAppUserAssignRequest("RoleId_example")}, "RoleId_example") // ServiceAccountCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateAppServiceAccount(context.Background()).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateAppServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateAppServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountCreateRequest** | [**ServiceAccountCreateRequest**](ServiceAccountCreateRequest.md) |  | 

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
    appUserCreateRequest := *openapiclient.NewAppUserCreateRequest("Identifier_example", "Name_example", "RoleId_example") // AppUserCreateRequest |  (optional)

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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppUserAccessToken

> AppUserAccessTokenCreateResponse CreateAppUserAccessToken(ctx).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()

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
    appUserAccessTokenCreateRequest := *openapiclient.NewAppUserAccessTokenCreateRequest("Name_example") // AppUserAccessTokenCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.CreateAppUserAccessToken(context.Background()).AppUserAccessTokenCreateRequest(appUserAccessTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.CreateAppUserAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppUserAccessToken`: AppUserAccessTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.CreateAppUserAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppUserAccessTokenRequest struct via the builder pattern


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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecideInvitation

> DecideInvitation(ctx, teamId).InvitationDecisionRequest(invitationDecisionRequest).Execute()

Accept or reject team invitation



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
    invitationDecisionRequest := *openapiclient.NewInvitationDecisionRequest(false) // InvitationDecisionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SessionAPI.DecideInvitation(context.Background(), teamId).InvitationDecisionRequest(invitationDecisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.DecideInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecideInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationDecisionRequest** | [**InvitationDecisionRequest**](InvitationDecisionRequest.md) |  | 

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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppServiceAccounts

> ServiceAccountListResponse ListAppServiceAccounts(ctx).Execute()

List App Service Accounts



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
    resp, r, err := apiClient.SessionAPI.ListAppServiceAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListAppServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppServiceAccounts`: ServiceAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListAppServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppServiceAccountsRequest struct via the builder pattern


### Return type

[**ServiceAccountListResponse**](ServiceAccountListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppUserAccessTokens

> AppUserAccessTokenListResponse ListAppUserAccessTokens(ctx).Execute()

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
    resp, r, err := apiClient.SessionAPI.ListAppUserAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListAppUserAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAppUserAccessTokens`: AppUserAccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListAppUserAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppUserAccessTokensRequest struct via the builder pattern


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

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvitations

> InvitationListResponse ListInvitations(ctx).Execute()

List of pending invitations



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
    resp, r, err := apiClient.SessionAPI.ListInvitations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvitations`: InvitationListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInvitationsRequest struct via the builder pattern


### Return type

[**InvitationListResponse**](InvitationListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeams

> TeamListResponse ListTeams(ctx).Assigned(assigned).Execute()

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
    assigned := "assigned_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionAPI.ListTeams(context.Background()).Assigned(assigned).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ListTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeams`: TeamListResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionAPI.ListTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assigned** | **string** |  | 

### Return type

[**TeamListResponse**](TeamListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

