# \SystemAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignSystemTeamAppUser**](SystemAPI.md#AssignSystemTeamAppUser) | **Post** /systems/{systemId}/app-users/{teamAppUserId} | Assign Team App User to System
[**CreateAccount**](SystemAPI.md#CreateAccount) | **Post** /systems/{systemId}/accounts | Create Account
[**CreateSystemAlertRule**](SystemAPI.md#CreateSystemAlertRule) | **Post** /systems/{systemId}/alert-rules | Create System Alert Rule
[**DeleteSystem**](SystemAPI.md#DeleteSystem) | **Delete** /systems/{systemId} | Delete System
[**DeleteSystemAlertRule**](SystemAPI.md#DeleteSystemAlertRule) | **Delete** /systems/{systemId}/alert-rules/{alertRuleId} | Delete System Alert Rule
[**DownloadSystemLogs**](SystemAPI.md#DownloadSystemLogs) | **Post** /systems/{systemId}/logs | Download Logs
[**GetCurrentAgentToken**](SystemAPI.md#GetCurrentAgentToken) | **Get** /systems/{systemId}/agent-tokens/current | Get Current Agent Token
[**GetSystem**](SystemAPI.md#GetSystem) | **Get** /systems/{systemId} | Get System
[**GetSystemAlertRule**](SystemAPI.md#GetSystemAlertRule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId} | Get System Alert Rule
[**GetSystemLimits**](SystemAPI.md#GetSystemLimits) | **Get** /systems/{systemId}/limits | Get System Limits
[**ImportAccount**](SystemAPI.md#ImportAccount) | **Post** /systems/{systemId}/import-account | Import Account
[**ImportUser**](SystemAPI.md#ImportUser) | **Post** /systems/{systemId}/import-user | Import User
[**ListAccounts**](SystemAPI.md#ListAccounts) | **Get** /systems/{systemId}/accounts | List Accounts
[**ListAccountsOverviewMetrics**](SystemAPI.md#ListAccountsOverviewMetrics) | **Get** /systems/{systemId}/accounts-overview-metrics | List Accounts overview metrics
[**ListAgentTokens**](SystemAPI.md#ListAgentTokens) | **Get** /systems/{systemId}/agent-tokens | List Agent Tokens
[**ListClusters**](SystemAPI.md#ListClusters) | **Get** /systems/{systemId}/nats-clusters | List Clusters
[**ListConnections**](SystemAPI.md#ListConnections) | **Get** /systems/{systemId}/connections | List Connections
[**ListServers**](SystemAPI.md#ListServers) | **Get** /systems/{systemId}/servers | List Servers
[**ListSystemAlertRules**](SystemAPI.md#ListSystemAlertRules) | **Get** /systems/{systemId}/alert-rules | List System Alert Rules
[**ListSystemInfoAccounts**](SystemAPI.md#ListSystemInfoAccounts) | **Get** /systems/{systemId}/info/accounts | List System Accounts Info
[**ListSystemInfoServers**](SystemAPI.md#ListSystemInfoServers) | **Get** /systems/{systemId}/info/servers | List System Servers info
[**ListSystemTeamAppUsers**](SystemAPI.md#ListSystemTeamAppUsers) | **Get** /systems/{systemId}/app-users | List System Team App Users
[**RotateAgentToken**](SystemAPI.md#RotateAgentToken) | **Post** /systems/{systemId}/agent-tokens | Rotate Agent Token
[**RunSystemAlertRule**](SystemAPI.md#RunSystemAlertRule) | **Get** /systems/{systemId}/alert-rules/{alertRuleId}/run | Run System Alert Rule
[**SystemJWTSync**](SystemAPI.md#SystemJWTSync) | **Post** /systems/{systemId}/jwt-sync | Re-sync JWTs of all accounts in this system
[**UnAssignSystemTeamAppUser**](SystemAPI.md#UnAssignSystemTeamAppUser) | **Delete** /systems/{systemId}/app-users/{teamAppUserId} | Unassign Team App User from System
[**UnmanageSystem**](SystemAPI.md#UnmanageSystem) | **Delete** /systems/{systemId}/unmanage | Unmanage System
[**UpdateSystem**](SystemAPI.md#UpdateSystem) | **Patch** /systems/{systemId} | Update System
[**UpdateSystemAlertRule**](SystemAPI.md#UpdateSystemAlertRule) | **Patch** /systems/{systemId}/alert-rules/{alertRuleId} | Update System Alert Rules



## AssignSystemTeamAppUser

> AppUserAssignResponse AssignSystemTeamAppUser(ctx, systemId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()

Assign Team App User to System



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
    teamAppUserId := "teamAppUserId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("RoleId_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.AssignSystemTeamAppUser(context.Background(), systemId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.AssignSystemTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignSystemTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.AssignSystemTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignSystemTeamAppUserRequest struct via the builder pattern


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


## CreateAccount

> AccountViewResponse CreateAccount(ctx, systemId).AccountCreateRequest(accountCreateRequest).Execute()

Create Account



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
    accountCreateRequest := *openapiclient.NewAccountCreateRequest("Name_example") // AccountCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.CreateAccount(context.Background(), systemId).AccountCreateRequest(accountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: AccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.CreateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountCreateRequest** | [**AccountCreateRequest**](AccountCreateRequest.md) |  | 

### Return type

[**AccountViewResponse**](AccountViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSystemAlertRule

> AlertRuleViewResponse CreateSystemAlertRule(ctx, systemId).Body(body).Execute()

Create System Alert Rule



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
    body := AlertRuleBaseCreateRequest(987) // AlertRuleBaseCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.CreateSystemAlertRule(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.CreateSystemAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystemAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.CreateSystemAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **AlertRuleBaseCreateRequest** |  | 

### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSystem

> DeleteSystem(ctx, systemId).Execute()

Delete System



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
    r, err := apiClient.SystemAPI.DeleteSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DeleteSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemRequest struct via the builder pattern


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


## DeleteSystemAlertRule

> DeleteSystemAlertRule(ctx, systemId, alertRuleId).Execute()

Delete System Alert Rule



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
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.DeleteSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DeleteSystemAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemAlertRuleRequest struct via the builder pattern


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


## DownloadSystemLogs

> *os.File DownloadSystemLogs(ctx, systemId).Execute()

Download Logs



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
    resp, r, err := apiClient.SystemAPI.DownloadSystemLogs(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.DownloadSystemLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadSystemLogs`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.DownloadSystemLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSystemLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentAgentToken

> AgentTokenCurrentResponse GetCurrentAgentToken(ctx, systemId).Execute()

Get Current Agent Token



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
    resp, r, err := apiClient.SystemAPI.GetCurrentAgentToken(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetCurrentAgentToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentAgentToken`: AgentTokenCurrentResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetCurrentAgentToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentAgentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentTokenCurrentResponse**](AgentTokenCurrentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystem

> SystemViewResponse GetSystem(ctx, systemId).Execute()

Get System



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
    resp, r, err := apiClient.SystemAPI.GetSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemViewResponse**](SystemViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemAlertRule

> AlertRuleViewResponse GetSystemAlertRule(ctx, systemId, alertRuleId).Execute()

Get System Alert Rule



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
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.GetSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemLimits

> SystemLimitsResponse GetSystemLimits(ctx, systemId).Execute()

Get System Limits



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
    resp, r, err := apiClient.SystemAPI.GetSystemLimits(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.GetSystemLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemLimits`: SystemLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.GetSystemLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemLimitsResponse**](SystemLimitsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccount

> ImportAccount(ctx, systemId).SystemAccountImportRequest(systemAccountImportRequest).Execute()

Import Account



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
    systemAccountImportRequest := *openapiclient.NewSystemAccountImportRequest("Jwt_example", "Seed_example") // SystemAccountImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.ImportAccount(context.Background(), systemId).SystemAccountImportRequest(systemAccountImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ImportAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemAccountImportRequest** | [**SystemAccountImportRequest**](SystemAccountImportRequest.md) |  | 

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


## ImportUser

> NatsUserViewResponse ImportUser(ctx, systemId).SystemUserImportRequest(systemUserImportRequest).Execute()

Import User



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
    systemUserImportRequest := *openapiclient.NewSystemUserImportRequest() // SystemUserImportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.ImportUser(context.Background(), systemId).SystemUserImportRequest(systemUserImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ImportUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ImportUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemUserImportRequest** | [**SystemUserImportRequest**](SystemUserImportRequest.md) |  | 

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


## ListAccounts

> AccountListResponse ListAccounts(ctx, systemId).Execute()

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
    systemId := "systemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.ListAccounts(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountListResponse**](AccountListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountsOverviewMetrics

> AccountsOverviewListResponse ListAccountsOverviewMetrics(ctx, systemId).Execute()

List Accounts overview metrics



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
    resp, r, err := apiClient.SystemAPI.ListAccountsOverviewMetrics(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListAccountsOverviewMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountsOverviewMetrics`: AccountsOverviewListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListAccountsOverviewMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsOverviewMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountsOverviewListResponse**](AccountsOverviewListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAgentTokens

> AgentTokenListResponse ListAgentTokens(ctx, systemId).Execute()

List Agent Tokens



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
    resp, r, err := apiClient.SystemAPI.ListAgentTokens(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListAgentTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAgentTokens`: AgentTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListAgentTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAgentTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentTokenListResponse**](AgentTokenListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> NatsClusterListResponse ListClusters(ctx, systemId).Execute()

List Clusters



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
    resp, r, err := apiClient.SystemAPI.ListClusters(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: NatsClusterListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsClusterListResponse**](NatsClusterListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> SystemConnectionsListResponse ListConnections(ctx, systemId).Sort(sort).Account(account).Cid(cid).State(state).Subject(subject).Limit(limit).User(user).Execute()

List Connections



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
    sort := "sort_example" // string |  (optional)
    account := "account_example" // string |  (optional)
    cid := "cid_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    subject := "subject_example" // string |  (optional)
    limit := float32(8.14) // float32 |  (optional)
    user := "user_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.ListConnections(context.Background(), systemId).Sort(sort).Account(account).Cid(cid).State(state).Subject(subject).Limit(limit).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: SystemConnectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** |  | 
 **account** | **string** |  | 
 **cid** | **string** |  | 
 **state** | **string** |  | 
 **subject** | **string** |  | 
 **limit** | **float32** |  | 
 **user** | **string** |  | 

### Return type

[**SystemConnectionsListResponse**](SystemConnectionsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServers

> NatsServerListResponse ListServers(ctx, systemId).Execute()

List Servers



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
    resp, r, err := apiClient.SystemAPI.ListServers(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServers`: NatsServerListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsServerListResponse**](NatsServerListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemAlertRules

> AlertRuleListResponse ListSystemAlertRules(ctx, systemId).Execute()

List System Alert Rules



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
    resp, r, err := apiClient.SystemAPI.ListSystemAlertRules(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemAlertRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemAlertRules`: AlertRuleListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleListResponse**](AlertRuleListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemInfoAccounts

> AccountSearchListResponse ListSystemInfoAccounts(ctx, systemId).Execute()

List System Accounts Info



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
    resp, r, err := apiClient.SystemAPI.ListSystemInfoAccounts(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemInfoAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemInfoAccounts`: AccountSearchListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemInfoAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemInfoAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountSearchListResponse**](AccountSearchListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemInfoServers

> NatsServerInfoListResponse ListSystemInfoServers(ctx, systemId).Execute()

List System Servers info



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
    resp, r, err := apiClient.SystemAPI.ListSystemInfoServers(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemInfoServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemInfoServers`: NatsServerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemInfoServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemInfoServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsServerInfoListResponse**](NatsServerInfoListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemTeamAppUsers

> AppUserAssignListResponse ListSystemTeamAppUsers(ctx, systemId).Execute()

List System Team App Users



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
    resp, r, err := apiClient.SystemAPI.ListSystemTeamAppUsers(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.ListSystemTeamAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemTeamAppUsers`: AppUserAssignListResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.ListSystemTeamAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemTeamAppUsersRequest struct via the builder pattern


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


## RotateAgentToken

> AgentTokenRotateResponse RotateAgentToken(ctx, systemId).Body(body).Execute()

Rotate Agent Token



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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.RotateAgentToken(context.Background(), systemId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RotateAgentToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAgentToken`: AgentTokenRotateResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.RotateAgentToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAgentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**AgentTokenRotateResponse**](AgentTokenRotateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSystemAlertRule

> AlertViewResponse RunSystemAlertRule(ctx, systemId, alertRuleId).Execute()

Run System Alert Rule



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
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.RunSystemAlertRule(context.Background(), systemId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.RunSystemAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunSystemAlertRule`: AlertViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.RunSystemAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunSystemAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertViewResponse**](AlertViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SystemJWTSync

> SystemJWTSync(ctx, systemId).Execute()

Re-sync JWTs of all accounts in this system



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
    r, err := apiClient.SystemAPI.SystemJWTSync(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.SystemJWTSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSystemJWTSyncRequest struct via the builder pattern


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


## UnAssignSystemTeamAppUser

> UnAssignSystemTeamAppUser(ctx, systemId, teamAppUserId).Execute()

Unassign Team App User from System



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
    teamAppUserId := "teamAppUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SystemAPI.UnAssignSystemTeamAppUser(context.Background(), systemId, teamAppUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UnAssignSystemTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnAssignSystemTeamAppUserRequest struct via the builder pattern


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


## UnmanageSystem

> UnmanageSystem(ctx, systemId).Execute()

Unmanage System



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
    r, err := apiClient.SystemAPI.UnmanageSystem(context.Background(), systemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UnmanageSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmanageSystemRequest struct via the builder pattern


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


## UpdateSystem

> SystemViewResponse UpdateSystem(ctx, systemId).TestConnection(testConnection).SystemUpdateRequest(systemUpdateRequest).Execute()

Update System



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
    testConnection := true // bool |  (optional)
    systemUpdateRequest := *openapiclient.NewSystemUpdateRequest() // SystemUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.UpdateSystem(context.Background(), systemId).TestConnection(testConnection).SystemUpdateRequest(systemUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UpdateSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSystem`: SystemViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.UpdateSystem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testConnection** | **bool** |  | 
 **systemUpdateRequest** | [**SystemUpdateRequest**](SystemUpdateRequest.md) |  | 

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


## UpdateSystemAlertRule

> AlertRuleViewResponse UpdateSystemAlertRule(ctx, systemId, alertRuleId).AlertRuleUpdateRequest(alertRuleUpdateRequest).Execute()

Update System Alert Rules



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
    alertRuleId := "alertRuleId_example" // string | 
    alertRuleUpdateRequest := *openapiclient.NewAlertRuleUpdateRequest() // AlertRuleUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAPI.UpdateSystemAlertRule(context.Background(), systemId, alertRuleId).AlertRuleUpdateRequest(alertRuleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAPI.UpdateSystemAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSystemAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAPI.UpdateSystemAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSystemAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertRuleUpdateRequest** | [**AlertRuleUpdateRequest**](AlertRuleUpdateRequest.md) |  | 

### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

