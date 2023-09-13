# \TeamAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignTeamAppUser**](TeamAPI.md#AssignTeamAppUser) | **Post** /teams/{teamId}/app-users/{appUserId} | Assign App User to Team
[**CreateSystem**](TeamAPI.md#CreateSystem) | **Post** /teams/{teamId}/systems | Create System
[**DeleteTeam**](TeamAPI.md#DeleteTeam) | **Delete** /teams/{teamId} | Delete Team
[**GetTeam**](TeamAPI.md#GetTeam) | **Get** /teams/{teamId} | Get Team
[**ImportSystem**](TeamAPI.md#ImportSystem) | **Post** /teams/{teamId}/import-system | Import a System
[**ListSystems**](TeamAPI.md#ListSystems) | **Get** /teams/{teamId}/systems | List Systems
[**ListTeamAccounts**](TeamAPI.md#ListTeamAccounts) | **Get** /teams/{teamId}/accounts | List Accounts
[**ListTeamAppUsers**](TeamAPI.md#ListTeamAppUsers) | **Get** /teams/{teamId}/app-users | List App Users
[**ListTeamNatsUsers**](TeamAPI.md#ListTeamNatsUsers) | **Get** /teams/{teamId}/nats-users | List NATS Users
[**UnAssignTeamAppUser**](TeamAPI.md#UnAssignTeamAppUser) | **Delete** /teams/{teamId}/app-users/{appUserId} | Unassign App User from Team
[**UpdateTeam**](TeamAPI.md#UpdateTeam) | **Patch** /teams/{teamId} | Update Team



## AssignTeamAppUser

> AppUserAssignResponse AssignTeamAppUser(ctx, teamId, appUserId).AppUserAssignRequest(appUserAssignRequest).Execute()

Assign App User to Team



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
    teamId := "teamId_example" // string | 
    appUserId := "appUserId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("Role_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.AssignTeamAppUser(context.Background(), teamId, appUserId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.AssignTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.AssignTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 
**appUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignTeamAppUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appUserAssignRequest** | [**AppUserAssignRequest**](AppUserAssignRequest.md) |  | 

### Return type

[**AppUserAssignResponse**](AppUserAssignResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    teamId := "teamId_example" // string | 
    systemCreateRequest := *openapiclient.NewSystemCreateRequest("Name_example", "TeamId_example") // SystemCreateRequest |  (optional)

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

No authorization required

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
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

No authorization required

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
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

No authorization required

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    teamId := "teamId_example" // string | 
    systemImportRequest := *openapiclient.NewSystemImportRequest("OperatorJwt_example", "OperatorKey_example", "SystemJwt_example", "SystemKey_example", "TeamId_example") // SystemImportRequest |  (optional)

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystems

> SystemListResponse ListSystems(ctx, teamId).Execute()

List Systems



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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ListSystems(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystems`: SystemListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListSystems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemListResponse**](SystemListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamAccounts

> AccountListResponse ListTeamAccounts(ctx, teamId).Execute()

List Accounts



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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ListTeamAccounts(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamAccounts`: AccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountListResponse**](AccountListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTeamNatsUsers

> NatsUserListResponse ListTeamNatsUsers(ctx, teamId).Execute()

List NATS Users



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
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamAPI.ListTeamNatsUsers(context.Background(), teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamAPI.ListTeamNatsUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTeamNatsUsers`: NatsUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamAPI.ListTeamNatsUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTeamNatsUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserListResponse**](NatsUserListResponse.md)

### Authorization

No authorization required

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
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

No authorization required

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
    openapiclient "github.com/synadia-io/control-plane-sdk-go/api"
)

func main() {
    teamId := "teamId_example" // string | 
    teamUpdateRequest := *openapiclient.NewTeamUpdateRequest("Name_example") // TeamUpdateRequest |  (optional)

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

