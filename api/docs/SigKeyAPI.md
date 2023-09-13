# \SigKeyAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountSk**](SigKeyAPI.md#DeleteAccountSk) | **Delete** /account-sks/{keyId} | Delete Account Signing
[**GetAccountSk**](SigKeyAPI.md#GetAccountSk) | **Get** /account-sks/{keyId} | Get Account Signing
[**UpdateAccountSk**](SigKeyAPI.md#UpdateAccountSk) | **Patch** /account-sks/{keyId} | Update Account Signing



## DeleteAccountSk

> DeleteAccountSk(ctx, keyId).Execute()

Delete Account Signing



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
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SigKeyAPI.DeleteAccountSk(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyAPI.DeleteAccountSk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountSkRequest struct via the builder pattern


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


## GetAccountSk

> SigningKeyViewResponse GetAccountSk(ctx, keyId).Execute()

Get Account Signing



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
    keyId := "keyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyAPI.GetAccountSk(context.Background(), keyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyAPI.GetAccountSk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountSk`: SigningKeyViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyAPI.GetAccountSk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigningKeyViewResponse**](SigningKeyViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSk

> SigningKeyViewResponse UpdateAccountSk(ctx, keyId).SigningKeyUpdateRequest(signingKeyUpdateRequest).Execute()

Update Account Signing



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
    keyId := "keyId_example" // string | 
    signingKeyUpdateRequest := *openapiclient.NewSigningKeyUpdateRequest(false) // SigningKeyUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigKeyAPI.UpdateAccountSk(context.Background(), keyId).SigningKeyUpdateRequest(signingKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigKeyAPI.UpdateAccountSk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountSk`: SigningKeyViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SigKeyAPI.UpdateAccountSk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountSkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signingKeyUpdateRequest** | [**SigningKeyUpdateRequest**](SigningKeyUpdateRequest.md) |  | 

### Return type

[**SigningKeyViewResponse**](SigningKeyViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

