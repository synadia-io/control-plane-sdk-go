# \TeamAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystem**](TeamAPI.md#CreateSystem) | **Post** /core/beta/teams/{teamId}/systems | Create System
[**CreateTeamServiceAccount**](TeamAPI.md#CreateTeamServiceAccount) | **Post** /core/beta/teams/{teamId}/service-accounts | Create Team Service Account
[**DeleteTeam**](TeamAPI.md#DeleteTeam) | **Delete** /core/beta/teams/{teamId} | Delete Team
[**GetTeam**](TeamAPI.md#GetTeam) | **Get** /core/beta/teams/{teamId} | Get Team
[**GetTeamLimits**](TeamAPI.md#GetTeamLimits) | **Get** /core/beta/teams/{teamId}/team-limits | Get Team Limits
[**ImportSystem**](TeamAPI.md#ImportSystem) | **Post** /core/beta/teams/{teamId}/import-system | Import a System
[**InviteAppUser**](TeamAPI.md#InviteAppUser) | **Post** /core/beta/teams/{teamId}/app-users/invitations | Invite App Users
[**LeaveTeam**](TeamAPI.md#LeaveTeam) | **Post** /core/beta/teams/{teamId}/app-users/leave | Leave Team
[**ListTeamAppUsers**](TeamAPI.md#ListTeamAppUsers) | **Get** /core/beta/teams/{teamId}/app-users | List App Users
[**ListTeamInfoAppUsers**](TeamAPI.md#ListTeamInfoAppUsers) | **Get** /core/beta/teams/{teamId}/info/app-users | List info of App Users in Team
[**ListTeamServiceAccounts**](TeamAPI.md#ListTeamServiceAccounts) | **Get** /core/beta/teams/{teamId}/service-accounts | List Team Service Accounts
[**ListTeamSystems**](TeamAPI.md#ListTeamSystems) | **Get** /core/beta/teams/{teamId}/systems | List Systems
[**UnAssignTeamAppUser**](TeamAPI.md#UnAssignTeamAppUser) | **Delete** /core/beta/teams/{teamId}/app-users/{appUserId} | Unassign App User from Team
[**UpdateTeam**](TeamAPI.md#UpdateTeam) | **Patch** /core/beta/teams/{teamId} | Update Team
[**UpdateTeamAppUser**](TeamAPI.md#UpdateTeamAppUser) | **Patch** /core/beta/teams/{teamId}/app-users/{appUserId} | Update App User Team Assignment



## CreateSystem

> SystemViewResponse CreateSystem(ctx, teamId).SystemCreateRequest(systemCreateRequest).Execute()

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
    teamId := "teamId_example" // string | 
    systemCreateRequest := *openapiclient.NewSystemCreateRequest("Name_example", "Url_example") // SystemCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.CreateSystem(context.Background(), teamId).SystemCreateRequest(systemCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CreateSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CreateSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemCreateRequest** | [**SystemCreateRequest**](SystemCreateRequest.md) |  | 

### Return type

[**SystemViewResponse**](SystemViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamServiceAccount

> ServiceAccountViewResponse CreateTeamServiceAccount(ctx, teamId).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()

Create Team Service Account



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
    serviceAccountCreateRequest := *openapiclient.NewServiceAccountCreateRequest("Name_example", map[string]AppUserAssignRequest{"key": *openapiclient.NewAppUserAssignRequest("RoleId_example")}, "RoleId_example") // ServiceAccountCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.CreateTeamServiceAccount(context.Background(), teamId).ServiceAccountCreateRequest(serviceAccountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.CreateTeamServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamServiceAccount`: ServiceAccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.CreateTeamServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamServiceAccountRequest struct via the builder pattern


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


## DeleteTeam

> DeleteTeam(ctx, teamId).Execute()

Delete Team



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
    r, err := apiClient.TeamAPI.DeleteTeam(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.DeleteTeam``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## GetTeam

> TeamViewResponse GetTeam(ctx, teamId).Execute()

Get Team



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
    resp, r, err := apiClient.TeamAPI.GetTeam(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeam`: TeamViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamViewResponse**](TeamViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamLimits

> TeamLimitsResponse GetTeamLimits(ctx, teamId).Execute()

Get Team Limits



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
    resp, r, err := apiClient.TeamAPI.GetTeamLimits(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.GetTeamLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeamLimits`: TeamLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.GetTeamLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamLimitsResponse**](TeamLimitsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSystem

> SystemViewResponse ImportSystem(ctx, teamId).SystemImportRequest(systemImportRequest).Execute()

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
    teamId := "teamId_example" // string | 
    systemImportRequest := *openapiclient.NewSystemImportRequest("OperatorJwt_example", "OperatorKey_example", "SystemJwt_example", "SystemKey_example") // SystemImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ImportSystem(context.Background(), teamId).SystemImportRequest(systemImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ImportSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ImportSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemImportRequest** | [**SystemImportRequest**](SystemImportRequest.md) |  | 

### Return type

[**SystemViewResponse**](SystemViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteAppUser

> AppUserInvitationResponse InviteAppUser(ctx, teamId).AppUserInvitationRequest(appUserInvitationRequest).Execute()

Invite App Users



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
    appUserInvitationRequest := *openapiclient.NewAppUserInvitationRequest("Identifier_example", "Name_example", "RoleId_example") // AppUserInvitationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.InviteAppUser(context.Background(), teamId).AppUserInvitationRequest(appUserInvitationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.InviteAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteAppUser`: AppUserInvitationResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.InviteAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appUserInvitationRequest** | [**AppUserInvitationRequest**](AppUserInvitationRequest.md) |  | 

### Return type

[**AppUserInvitationResponse**](AppUserInvitationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LeaveTeam

> LeaveTeam(ctx, teamId).Execute()

Leave Team



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
    r, err := apiClient.TeamAPI.LeaveTeam(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.LeaveTeam``: %v\n", err)
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

Other parameters are passed through a pointer to a apiLeaveTeamRequest struct via the builder pattern


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


## ListTeamAppUsers

> AppUserAssignListResponse ListTeamAppUsers(ctx, teamId).Execute()

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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ListTeamAppUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamAppUsers`: AppUserAssignListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamAppUsersRequest struct via the builder pattern


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


## ListTeamInfoAppUsers

> TeamAppUserListResponse ListTeamInfoAppUsers(ctx, teamId).Execute()

List info of App Users in Team



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
    resp, r, err := apiClient.TeamAPI.ListTeamInfoAppUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamInfoAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamInfoAppUsers`: TeamAppUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamInfoAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamInfoAppUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamAppUserListResponse**](TeamAppUserListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamServiceAccounts

> ServiceAccountListResponse ListTeamServiceAccounts(ctx, teamId).Execute()

List Team Service Accounts



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
    resp, r, err := apiClient.TeamAPI.ListTeamServiceAccounts(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamServiceAccounts`: ServiceAccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamServiceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListTeamSystems

> SystemListResponse ListTeamSystems(ctx, teamId).Execute()

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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ListTeamSystems(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamSystems`: SystemListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamSystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemListResponse**](SystemListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnAssignTeamAppUser

> UnAssignTeamAppUser(ctx, teamId, appUserId).Execute()

Unassign App User from Team



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
    appUserId := "appUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TeamAPI.UnAssignTeamAppUser(context.Background(), teamId, appUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UnAssignTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnAssignTeamAppUserRequest struct via the builder pattern


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


## UpdateTeam

> TeamViewResponse UpdateTeam(ctx, teamId).TeamUpdateRequest(teamUpdateRequest).Execute()

Update Team



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
    teamUpdateRequest := *openapiclient.NewTeamUpdateRequest() // TeamUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.UpdateTeam(context.Background(), teamId).TeamUpdateRequest(teamUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeam`: TeamViewResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamUpdateRequest** | [**TeamUpdateRequest**](TeamUpdateRequest.md) |  | 

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


## UpdateTeamAppUser

> AppUserAssignResponse UpdateTeamAppUser(ctx, teamId, appUserId).AppUserAssignRequest(appUserAssignRequest).Execute()

Update App User Team Assignment



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
    appUserId := "appUserId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("RoleId_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.UpdateTeamAppUser(context.Background(), teamId, appUserId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.UpdateTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.UpdateTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamAppUserRequest struct via the builder pattern


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

