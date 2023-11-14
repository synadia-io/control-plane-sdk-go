# \AccountAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignAccountTeamAppUser**](AccountAPI.md#AssignAccountTeamAppUser) | **Post** /accounts/{accountId}/app-users/{teamAppUserId} | Assign Team App User to Account
[**CreateAccountSkGroup**](AccountAPI.md#CreateAccountSkGroup) | **Post** /accounts/{accountId}/account-sk-groups | Create Account Signing Key Group
[**CreateAlertRule**](AccountAPI.md#CreateAlertRule) | **Post** /accounts/{accountId}/alert-rules | Create Account Alert Rule
[**CreateKvBucket**](AccountAPI.md#CreateKvBucket) | **Post** /accounts/{accountId}/jetstream/kv-buckets | Create KV Bucket
[**CreateMirror**](AccountAPI.md#CreateMirror) | **Post** /accounts/{accountId}/jetstream/mirrors | Create Mirror
[**CreateStream**](AccountAPI.md#CreateStream) | **Post** /accounts/{accountId}/jetstream/streams | Create Stream
[**CreateStreamExport**](AccountAPI.md#CreateStreamExport) | **Post** /accounts/{accountId}/stream-exports | Create Stream Export
[**CreateStreamImport**](AccountAPI.md#CreateStreamImport) | **Post** /accounts/{accountId}/stream-imports | Create Stream Import
[**CreateSubjectExport**](AccountAPI.md#CreateSubjectExport) | **Post** /accounts/{accountId}/subject-exports | Create Subject Export
[**CreateSubjectImport**](AccountAPI.md#CreateSubjectImport) | **Post** /accounts/{accountId}/subject-imports | Create Subject Import
[**CreateUser**](AccountAPI.md#CreateUser) | **Post** /accounts/{accountId}/nats-users | Create NATS User
[**DeleteAccount**](AccountAPI.md#DeleteAccount) | **Delete** /accounts/{accountId} | Delete Account
[**DeleteAlertRule**](AccountAPI.md#DeleteAlertRule) | **Delete** /accounts/{accountId}/alert-rules/{alertRuleId} | Delete Account Alert Rule
[**GetAccount**](AccountAPI.md#GetAccount) | **Get** /accounts/{accountId} | Get Account
[**GetAccountInfo**](AccountAPI.md#GetAccountInfo) | **Get** /accounts/{accountId}/info | Get Account Info
[**GetAccountMetrics**](AccountAPI.md#GetAccountMetrics) | **Get** /accounts/{accountId}/metrics | Get Account Metrics
[**GetAlertRule**](AccountAPI.md#GetAlertRule) | **Get** /accounts/{accountId}/alert-rules/{alertRuleId} | Get Account Alert Rule
[**GetJetStreamPlacementOptions**](AccountAPI.md#GetJetStreamPlacementOptions) | **Get** /accounts/{accountId}/jetstream/placement-options | Get JetStream Placement Options
[**ListAccountConnections**](AccountAPI.md#ListAccountConnections) | **Get** /accounts/{accountId}/connections | List Account Connections
[**ListAccountSkGroup**](AccountAPI.md#ListAccountSkGroup) | **Get** /accounts/{accountId}/account-sk-groups | List Account Signing Key Groups
[**ListAccountTeamAppUsers**](AccountAPI.md#ListAccountTeamAppUsers) | **Get** /accounts/{accountId}/app-users | List Account Team App Users
[**ListAlertRules**](AccountAPI.md#ListAlertRules) | **Get** /accounts/{accountId}/alert-rules | List Account Alert Rules
[**ListJetStreamAssets**](AccountAPI.md#ListJetStreamAssets) | **Get** /accounts/{accountId}/jetstream | List JetStream Assets
[**ListKvBuckets**](AccountAPI.md#ListKvBuckets) | **Get** /accounts/{accountId}/jetstream/kv-buckets | List KV buckets
[**ListMirrors**](AccountAPI.md#ListMirrors) | **Get** /accounts/{accountId}/jetstream/mirrors | List Mirrors
[**ListStreamExports**](AccountAPI.md#ListStreamExports) | **Get** /accounts/{accountId}/stream-exports | List Stream Exports
[**ListStreamExportsShared**](AccountAPI.md#ListStreamExportsShared) | **Get** /accounts/{accountId}/stream-imports/shared | List Shared Stream Exports
[**ListStreamImports**](AccountAPI.md#ListStreamImports) | **Get** /accounts/{accountId}/stream-imports | List Stream Imports
[**ListStreams**](AccountAPI.md#ListStreams) | **Get** /accounts/{accountId}/jetstream/streams | List Streams
[**ListSubjectExports**](AccountAPI.md#ListSubjectExports) | **Get** /accounts/{accountId}/subject-exports | List Subject Exports
[**ListSubjectExportsShared**](AccountAPI.md#ListSubjectExportsShared) | **Get** /accounts/{accountId}/subject-imports/shared | List Shared Subject Exports
[**ListSubjectImports**](AccountAPI.md#ListSubjectImports) | **Get** /accounts/{accountId}/subject-imports | List Subject Imports
[**ListUsers**](AccountAPI.md#ListUsers) | **Get** /accounts/{accountId}/nats-users | List NATS Users
[**RunAlertRule**](AccountAPI.md#RunAlertRule) | **Get** /accounts/{accountId}/alert-rules/{alertRuleId}/run | Run Account Alert Rule
[**UnAssignAccountTeamAppUser**](AccountAPI.md#UnAssignAccountTeamAppUser) | **Delete** /accounts/{accountId}/app-users/{teamAppUserId} | Unassign Team App User from Account
[**UnmanageAccount**](AccountAPI.md#UnmanageAccount) | **Delete** /accounts/{accountId}/unmanage | Unmanage Account
[**UpdateAccount**](AccountAPI.md#UpdateAccount) | **Patch** /accounts/{accountId} | Update Account
[**UpdateAlertRule**](AccountAPI.md#UpdateAlertRule) | **Patch** /accounts/{accountId}/alert-rules/{alertRuleId} | Update Account Alert Rule



## AssignAccountTeamAppUser

> AppUserAssignResponse AssignAccountTeamAppUser(ctx, accountId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()

Assign Team App User to Account



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
    accountId := "accountId_example" // string | 
    teamAppUserId := "teamAppUserId_example" // string | 
    appUserAssignRequest := *openapiclient.NewAppUserAssignRequest("Role_example") // AppUserAssignRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.AssignAccountTeamAppUser(context.Background(), accountId, teamAppUserId).AppUserAssignRequest(appUserAssignRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.AssignAccountTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignAccountTeamAppUser`: AppUserAssignResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.AssignAccountTeamAppUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignAccountTeamAppUserRequest struct via the builder pattern


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


## CreateAccountSkGroup

> SigningKeyGroupViewResponse CreateAccountSkGroup(ctx, accountId).SigningKeyGroupCreateRequest(signingKeyGroupCreateRequest).Execute()

Create Account Signing Key Group



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
    accountId := "accountId_example" // string | 
    signingKeyGroupCreateRequest := *openapiclient.NewSigningKeyGroupCreateRequest("Name_example") // SigningKeyGroupCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateAccountSkGroup(context.Background(), accountId).SigningKeyGroupCreateRequest(signingKeyGroupCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountSkGroup`: SigningKeyGroupViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAccountSkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signingKeyGroupCreateRequest** | [**SigningKeyGroupCreateRequest**](SigningKeyGroupCreateRequest.md) |  | 

### Return type

[**SigningKeyGroupViewResponse**](SigningKeyGroupViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlertRule

> AlertRuleViewResponse CreateAlertRule(ctx, accountId).AlertRuleAccountCreateRequest(alertRuleAccountCreateRequest).Execute()

Create Account Alert Rule



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
    accountId := "accountId_example" // string | 
    alertRuleAccountCreateRequest := *openapiclient.NewAlertRuleAccountCreateRequest(int32(123), "Message_example", "Metric_example", openapiclient.AlertRuleSeverity("Info"), openapiclient.AlertRuleOperator("Greater Than"), "RuleType_example") // AlertRuleAccountCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateAlertRule(context.Background(), accountId).AlertRuleAccountCreateRequest(alertRuleAccountCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertRuleAccountCreateRequest** | [**AlertRuleAccountCreateRequest**](AlertRuleAccountCreateRequest.md) |  | 

### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKvBucket

> JSKVBucketViewResponse CreateKvBucket(ctx, accountId).JSKVBucketCreateRequest(jSKVBucketCreateRequest).Execute()

Create KV Bucket



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
    accountId := "accountId_example" // string | 
    jSKVBucketCreateRequest := *openapiclient.NewJSKVBucketCreateRequest("BucketName_example", "Description_example", int32(123), int64(123), int32(123), int32(123), openapiclient.StorageType("file"), int64(123)) // JSKVBucketCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateKvBucket(context.Background(), accountId).JSKVBucketCreateRequest(jSKVBucketCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateKvBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKvBucket`: JSKVBucketViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateKvBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKvBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSKVBucketCreateRequest** | [**JSKVBucketCreateRequest**](JSKVBucketCreateRequest.md) |  | 

### Return type

[**JSKVBucketViewResponse**](JSKVBucketViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMirror

> JSMirrorInfoResponse CreateMirror(ctx, accountId).JSMirrorConfigRequest(jSMirrorConfigRequest).Execute()

Create Mirror



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
    accountId := "accountId_example" // string | 
    jSMirrorConfigRequest := *openapiclient.NewJSMirrorConfigRequest(false, false, false, false, openapiclient.DiscardPolicy("old"), int64(123), int64(123), int32(123), int64(123), int64(123), "Name_example", int32(123), openapiclient.RetentionPolicy("limits"), false, openapiclient.StorageType("file"), *openapiclient.NewStreamSource("Name_example"), false) // JSMirrorConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateMirror(context.Background(), accountId).JSMirrorConfigRequest(jSMirrorConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMirror`: JSMirrorInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSMirrorConfigRequest** | [**JSMirrorConfigRequest**](JSMirrorConfigRequest.md) |  | 

### Return type

[**JSMirrorInfoResponse**](JSMirrorInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStream

> JSStreamInfoResponse CreateStream(ctx, accountId).JSStreamConfigRequest(jSStreamConfigRequest).Execute()

Create Stream



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
    accountId := "accountId_example" // string | 
    jSStreamConfigRequest := *openapiclient.NewJSStreamConfigRequest(false, false, false, false, openapiclient.DiscardPolicy("old"), int64(123), int64(123), int32(123), int64(123), int64(123), "Name_example", int32(123), openapiclient.RetentionPolicy("limits"), false, openapiclient.StorageType("file")) // JSStreamConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateStream(context.Background(), accountId).JSStreamConfigRequest(jSStreamConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStream`: JSStreamInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSStreamConfigRequest** | [**JSStreamConfigRequest**](JSStreamConfigRequest.md) |  | 

### Return type

[**JSStreamInfoResponse**](JSStreamInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamExport

> StreamExportViewResponse CreateStreamExport(ctx, accountId).StreamExportCreateRequest(streamExportCreateRequest).Execute()

Create Stream Export



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
    accountId := "accountId_example" // string | 
    streamExportCreateRequest := *openapiclient.NewStreamExportCreateRequest(false, "StreamName_example") // StreamExportCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateStreamExport(context.Background(), accountId).StreamExportCreateRequest(streamExportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateStreamExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamExport`: StreamExportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateStreamExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamExportCreateRequest** | [**StreamExportCreateRequest**](StreamExportCreateRequest.md) |  | 

### Return type

[**StreamExportViewResponse**](StreamExportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStreamImport

> StreamImportViewResponse CreateStreamImport(ctx, accountId).StreamImportCreateRequest(streamImportCreateRequest).Execute()

Create Stream Import



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
    accountId := "accountId_example" // string | 
    streamImportCreateRequest := *openapiclient.NewStreamImportCreateRequest("DeliverSubjectPrefix_example", false, "JsSubjectPrefix_example", "RemoteAccountNkeyPublic_example", "StreamName_example") // StreamImportCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateStreamImport(context.Background(), accountId).StreamImportCreateRequest(streamImportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateStreamImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamImport`: StreamImportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateStreamImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamImportCreateRequest** | [**StreamImportCreateRequest**](StreamImportCreateRequest.md) |  | 

### Return type

[**StreamImportViewResponse**](StreamImportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubjectExport

> SubjectExportViewResponse CreateSubjectExport(ctx, accountId).SubjectExportCreateRequest(subjectExportCreateRequest).Execute()

Create Subject Export



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
    accountId := "accountId_example" // string | 
    subjectExportCreateRequest := *openapiclient.NewSubjectExportCreateRequest(*openapiclient.NewExport(), false, int64(123)) // SubjectExportCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateSubjectExport(context.Background(), accountId).SubjectExportCreateRequest(subjectExportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateSubjectExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubjectExport`: SubjectExportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateSubjectExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectExportCreateRequest** | [**SubjectExportCreateRequest**](SubjectExportCreateRequest.md) |  | 

### Return type

[**SubjectExportViewResponse**](SubjectExportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubjectImport

> SubjectImportViewResponse CreateSubjectImport(ctx, accountId).SubjectImportCreateRequest(subjectImportCreateRequest).Execute()

Create Subject Import



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
    accountId := "accountId_example" // string | 
    subjectImportCreateRequest := *openapiclient.NewSubjectImportCreateRequest(*openapiclient.NewImport()) // SubjectImportCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateSubjectImport(context.Background(), accountId).SubjectImportCreateRequest(subjectImportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateSubjectImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubjectImport`: SubjectImportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateSubjectImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectImportCreateRequest** | [**SubjectImportCreateRequest**](SubjectImportCreateRequest.md) |  | 

### Return type

[**SubjectImportViewResponse**](SubjectImportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> NatsUserViewResponse CreateUser(ctx, accountId).NatsUserCreateRequest(natsUserCreateRequest).Execute()

Create NATS User



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
    accountId := "accountId_example" // string | 
    natsUserCreateRequest := *openapiclient.NewNatsUserCreateRequest("Name_example", "SkGroupId_example") // NatsUserCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.CreateUser(context.Background(), accountId).NatsUserCreateRequest(natsUserCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: NatsUserViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **natsUserCreateRequest** | [**NatsUserCreateRequest**](NatsUserCreateRequest.md) |  | 

### Return type

[**NatsUserViewResponse**](NatsUserViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, accountId).Execute()

Delete Account



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.DeleteAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


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


## DeleteAlertRule

> DeleteAlertRule(ctx, accountId, alertRuleId).Execute()

Delete Account Alert Rule



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
    accountId := "accountId_example" // string | 
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.DeleteAlertRule(context.Background(), accountId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRuleRequest struct via the builder pattern


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


## GetAccount

> AccountViewResponse GetAccount(ctx, accountId).Execute()

Get Account



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: AccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountViewResponse**](AccountViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInfo

> JetStreamAccountStats GetAccountInfo(ctx, accountId).Execute()

Get Account Info



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetAccountInfo(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountInfo`: JetStreamAccountStats
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JetStreamAccountStats**](JetStreamAccountStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMetrics

> AccountMetrics GetAccountMetrics(ctx, accountId).Execute()

Get Account Metrics



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetAccountMetrics(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountMetrics`: AccountMetrics
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountMetrics**](AccountMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertRule

> AlertRuleViewResponse GetAlertRule(ctx, accountId, alertRuleId).Execute()

Get Account Alert Rule



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
    accountId := "accountId_example" // string | 
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetAlertRule(context.Background(), accountId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJetStreamPlacementOptions

> JSPlacementOptionsResponse GetJetStreamPlacementOptions(ctx, accountId).Execute()

Get JetStream Placement Options



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.GetJetStreamPlacementOptions(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetJetStreamPlacementOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJetStreamPlacementOptions`: JSPlacementOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetJetStreamPlacementOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJetStreamPlacementOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSPlacementOptionsResponse**](JSPlacementOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountConnections

> AccountConnectionsListResponse ListAccountConnections(ctx, accountId).Sort(sort).Cid(cid).State(state).Subject(subject).Limit(limit).User(user).Execute()

List Account Connections



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
    accountId := "accountId_example" // string | 
    sort := "sort_example" // string |  (optional)
    cid := "cid_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    subject := "subject_example" // string |  (optional)
    limit := float32(8.14) // float32 |  (optional)
    user := "user_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListAccountConnections(context.Background(), accountId).Sort(sort).Cid(cid).State(state).Subject(subject).Limit(limit).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListAccountConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountConnections`: AccountConnectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListAccountConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** |  | 
 **cid** | **string** |  | 
 **state** | **string** |  | 
 **subject** | **string** |  | 
 **limit** | **float32** |  | 
 **user** | **string** |  | 

### Return type

[**AccountConnectionsListResponse**](AccountConnectionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountSkGroup

> SigningKeyGroupListResponse ListAccountSkGroup(ctx, accountId).Execute()

List Account Signing Key Groups



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListAccountSkGroup(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountSkGroup`: SigningKeyGroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListAccountSkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountSkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyGroupListResponse**](SigningKeyGroupListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountTeamAppUsers

> AppUserAssignListResponse ListAccountTeamAppUsers(ctx, accountId).Execute()

List Account Team App Users



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListAccountTeamAppUsers(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListAccountTeamAppUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountTeamAppUsers`: AppUserAssignListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListAccountTeamAppUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountTeamAppUsersRequest struct via the builder pattern


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


## ListAlertRules

> AlertRuleListResponse ListAlertRules(ctx, accountId).Execute()

List Account Alert Rules



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListAlertRules(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListAlertRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlertRules`: AlertRuleListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListAlertRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAlertRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertRuleListResponse**](AlertRuleListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJetStreamAssets

> JSAssetInfoListResponse ListJetStreamAssets(ctx, accountId).Execute()

List JetStream Assets



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListJetStreamAssets(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListJetStreamAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJetStreamAssets`: JSAssetInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListJetStreamAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListJetStreamAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSAssetInfoListResponse**](JSAssetInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKvBuckets

> JSKVBucketListResponse ListKvBuckets(ctx, accountId).Execute()

List KV buckets



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListKvBuckets(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListKvBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKvBuckets`: JSKVBucketListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListKvBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKvBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSKVBucketListResponse**](JSKVBucketListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMirrors

> JSMirrorInfoListResponse ListMirrors(ctx, accountId).Execute()

List Mirrors



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListMirrors(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListMirrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMirrors`: JSMirrorInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSMirrorInfoListResponse**](JSMirrorInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamExports

> StreamExportListResponse ListStreamExports(ctx, accountId).Execute()

List Stream Exports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListStreamExports(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListStreamExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamExports`: StreamExportListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListStreamExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamExportListResponse**](StreamExportListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamExportsShared

> StreamExportSharedListResponse ListStreamExportsShared(ctx, accountId).Execute()

List Shared Stream Exports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListStreamExportsShared(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListStreamExportsShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamExportsShared`: StreamExportSharedListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListStreamExportsShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamExportsSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamExportSharedListResponse**](StreamExportSharedListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamImports

> StreamImportListResponse ListStreamImports(ctx, accountId).Execute()

List Stream Imports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListStreamImports(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListStreamImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamImports`: StreamImportListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListStreamImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamImportListResponse**](StreamImportListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreams

> JSStreamInfoListResponse ListStreams(ctx, accountId).Execute()

List Streams



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListStreams(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreams`: JSStreamInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListStreams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSStreamInfoListResponse**](JSStreamInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectExports

> SubjectExportListResponse ListSubjectExports(ctx, accountId).Execute()

List Subject Exports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListSubjectExports(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListSubjectExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubjectExports`: SubjectExportListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListSubjectExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectExportListResponse**](SubjectExportListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectExportsShared

> SubjectExportSharedListResponse ListSubjectExportsShared(ctx, accountId).Execute()

List Shared Subject Exports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListSubjectExportsShared(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListSubjectExportsShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubjectExportsShared`: SubjectExportSharedListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListSubjectExportsShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectExportsSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectExportSharedListResponse**](SubjectExportSharedListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectImports

> SubjectImportListResponse ListSubjectImports(ctx, accountId).Execute()

List Subject Imports



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListSubjectImports(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListSubjectImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubjectImports`: SubjectImportListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListSubjectImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectImportListResponse**](SubjectImportListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> NatsUserListResponse ListUsers(ctx, accountId).Execute()

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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.ListUsers(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: NatsUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


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


## RunAlertRule

> AlertViewResponse RunAlertRule(ctx, accountId, alertRuleId).Execute()

Run Account Alert Rule



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
    accountId := "accountId_example" // string | 
    alertRuleId := "alertRuleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.RunAlertRule(context.Background(), accountId, alertRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.RunAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunAlertRule`: AlertViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.RunAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertViewResponse**](AlertViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnAssignAccountTeamAppUser

> UnAssignAccountTeamAppUser(ctx, accountId, teamAppUserId).Execute()

Unassign Team App User from Account



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
    accountId := "accountId_example" // string | 
    teamAppUserId := "teamAppUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.UnAssignAccountTeamAppUser(context.Background(), accountId, teamAppUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UnAssignAccountTeamAppUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**teamAppUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnAssignAccountTeamAppUserRequest struct via the builder pattern


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


## UnmanageAccount

> UnmanageAccount(ctx, accountId).Execute()

Unmanage Account



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountAPI.UnmanageAccount(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UnmanageAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmanageAccountRequest struct via the builder pattern


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


## UpdateAccount

> AccountViewResponse UpdateAccount(ctx, accountId).AccountUpdateRequest(accountUpdateRequest).Execute()

Update Account



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
    accountId := "accountId_example" // string | 
    accountUpdateRequest := *openapiclient.NewAccountUpdateRequest() // AccountUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.UpdateAccount(context.Background(), accountId).AccountUpdateRequest(accountUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: AccountViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountUpdateRequest** | [**AccountUpdateRequest**](AccountUpdateRequest.md) |  | 

### Return type

[**AccountViewResponse**](AccountViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertRule

> AlertRuleViewResponse UpdateAlertRule(ctx, accountId, alertRuleId).AlertRuleUpdateRequest(alertRuleUpdateRequest).Execute()

Update Account Alert Rule



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
    accountId := "accountId_example" // string | 
    alertRuleId := "alertRuleId_example" // string | 
    alertRuleUpdateRequest := *openapiclient.NewAlertRuleUpdateRequest(int32(123), false, "Message_example", "Metric_example", openapiclient.AlertRuleSeverity("Info"), openapiclient.AlertRuleOperator("Greater Than")) // AlertRuleUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAPI.UpdateAlertRule(context.Background(), accountId, alertRuleId).AlertRuleUpdateRequest(alertRuleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.UpdateAlertRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertRule`: AlertRuleViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountAPI.UpdateAlertRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 
**alertRuleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertRuleUpdateRequest** | [**AlertRuleUpdateRequest**](AlertRuleUpdateRequest.md) |  | 

### Return type

[**AlertRuleViewResponse**](AlertRuleViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

