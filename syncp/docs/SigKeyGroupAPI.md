# \SigKeyGroupAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyAccountSkGroup**](SigKeyGroupAPI.md#CopyAccountSkGroup) | **Post** /core/beta/account-sk-groups/{groupId}/copy | Copy Account SK Group
[**DeleteAccountSkGroup**](SigKeyGroupAPI.md#DeleteAccountSkGroup) | **Delete** /core/beta/account-sk-groups/{groupId} | Delete Account Signing Key Group
[**GetAccountSkGroup**](SigKeyGroupAPI.md#GetAccountSkGroup) | **Get** /core/beta/account-sk-groups/{groupId} | Get Account Signing Key Group
[**ListAccountSKGroupUsers**](SigKeyGroupAPI.md#ListAccountSKGroupUsers) | **Get** /core/beta/account-sk-groups/{groupId}/nats-users | List NATS Users
[**ListAccountSkGroupKeys**](SigKeyGroupAPI.md#ListAccountSkGroupKeys) | **Get** /core/beta/account-sk-groups/{groupId}/account-sks | List Signing Keys
[**RotateAccountSk**](SigKeyGroupAPI.md#RotateAccountSk) | **Post** /core/beta/account-sk-groups/{groupId}/rotate-sk | Roate Active Signing Key
[**UpdateAccountSkGroup**](SigKeyGroupAPI.md#UpdateAccountSkGroup) | **Patch** /core/beta/account-sk-groups/{groupId} | Update Account Signing Key Group



## CopyAccountSkGroup

> SigningKeyGroupViewResponse CopyAccountSkGroup(ctx, groupId).SigningKeyGroupCopyRequest(signingKeyGroupCopyRequest).Execute()

Copy Account SK Group



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
    groupId := "groupId_example" // string | 
    signingKeyGroupCopyRequest := *openapiclient.NewSigningKeyGroupCopyRequest("Name_example") // SigningKeyGroupCopyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.CopyAccountSkGroup(context.Background(), groupId).SigningKeyGroupCopyRequest(signingKeyGroupCopyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.CopyAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyAccountSkGroup`: SigningKeyGroupViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.CopyAccountSkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyAccountSkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signingKeyGroupCopyRequest** | [**SigningKeyGroupCopyRequest**](SigningKeyGroupCopyRequest.md) |  | 

### Return type

[**SigningKeyGroupViewResponse**](SigningKeyGroupViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountSkGroup

> DeleteAccountSkGroup(ctx, groupId).Execute()

Delete Account Signing Key Group



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SigKeyGroupAPI.DeleteAccountSkGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.DeleteAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountSkGroupRequest struct via the builder pattern


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


## GetAccountSkGroup

> SigningKeyGroupViewResponse GetAccountSkGroup(ctx, groupId).Execute()

Get Account Signing Key Group



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.GetAccountSkGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.GetAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountSkGroup`: SigningKeyGroupViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.GetAccountSkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyGroupViewResponse**](SigningKeyGroupViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountSKGroupUsers

> NatsUserListResponse ListAccountSKGroupUsers(ctx, groupId).Execute()

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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.ListAccountSKGroupUsers(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.ListAccountSKGroupUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountSKGroupUsers`: NatsUserListResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.ListAccountSKGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountSKGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserListResponse**](NatsUserListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountSkGroupKeys

> SigningKeyListResponse ListAccountSkGroupKeys(ctx, groupId).Execute()

List Signing Keys



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.ListAccountSkGroupKeys(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.ListAccountSkGroupKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountSkGroupKeys`: SigningKeyListResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.ListAccountSkGroupKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountSkGroupKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyListResponse**](SigningKeyListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateAccountSk

> SigningKeyRotateResponse RotateAccountSk(ctx, groupId).Execute()

Roate Active Signing Key



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
    groupId := "groupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.RotateAccountSk(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.RotateAccountSk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RotateAccountSk`: SigningKeyRotateResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.RotateAccountSk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateAccountSkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyRotateResponse**](SigningKeyRotateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSkGroup

> SigningKeyGroupViewResponse UpdateAccountSkGroup(ctx, groupId).SigningKeyGroupUpdateRequest(signingKeyGroupUpdateRequest).Execute()

Update Account Signing Key Group



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
    groupId := "groupId_example" // string | 
    signingKeyGroupUpdateRequest := *openapiclient.NewSigningKeyGroupUpdateRequest() // SigningKeyGroupUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyGroupAPI.UpdateAccountSkGroup(context.Background(), groupId).SigningKeyGroupUpdateRequest(signingKeyGroupUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyGroupAPI.UpdateAccountSkGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountSkGroup`: SigningKeyGroupViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyGroupAPI.UpdateAccountSkGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountSkGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signingKeyGroupUpdateRequest** | [**SigningKeyGroupUpdateRequest**](SigningKeyGroupUpdateRequest.md) |  | 

### Return type

[**SigningKeyGroupViewResponse**](SigningKeyGroupViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

