# \StreamExportAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStreamShares**](StreamExportAPI.md#CreateStreamShares) | **Post** /core/beta/stream-exports/{streamExportId}/shares | Create Stream Shares
[**DeleteStreamExport**](StreamExportAPI.md#DeleteStreamExport) | **Delete** /core/beta/stream-exports/{streamExportId} | Delete Stream Export
[**DeleteStreamShare**](StreamExportAPI.md#DeleteStreamShare) | **Delete** /core/beta/stream-exports/{streamExportId}/shares/{targetAccountNKeyPublic} | Delete Stream Share
[**GetStreamExport**](StreamExportAPI.md#GetStreamExport) | **Get** /core/beta/stream-exports/{streamExportId} | Get Stream Export
[**ListStreamShares**](StreamExportAPI.md#ListStreamShares) | **Get** /core/beta/stream-exports/{streamExportId}/shares | List Stream Shares



## CreateStreamShares

> StreamShareViewResponse CreateStreamShares(ctx, streamExportId).StreamShareCreateRequest(streamShareCreateRequest).Execute()

Create Stream Shares



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
    streamExportId := "streamExportId_example" // string | 
    streamShareCreateRequest := *openapiclient.NewStreamShareCreateRequest("TargetAccountNkeyPublic_example") // StreamShareCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamExportAPI.CreateStreamShares(context.Background(), streamExportId).StreamShareCreateRequest(streamShareCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamExportAPI.CreateStreamShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStreamShares`: StreamShareViewResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamExportAPI.CreateStreamShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStreamSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamShareCreateRequest** | [**StreamShareCreateRequest**](StreamShareCreateRequest.md) |  | 

### Return type

[**StreamShareViewResponse**](StreamShareViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStreamExport

> DeleteStreamExport(ctx, streamExportId).Execute()

Delete Stream Export



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
    streamExportId := "streamExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamExportAPI.DeleteStreamExport(context.Background(), streamExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamExportAPI.DeleteStreamExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamExportRequest struct via the builder pattern


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


## DeleteStreamShare

> DeleteStreamShare(ctx, streamExportId, targetAccountNKeyPublic).Execute()

Delete Stream Share



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
    streamExportId := "streamExportId_example" // string | 
    targetAccountNKeyPublic := "targetAccountNKeyPublic_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamExportAPI.DeleteStreamShare(context.Background(), streamExportId, targetAccountNKeyPublic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamExportAPI.DeleteStreamShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamExportId** | **string** |  | 
**targetAccountNKeyPublic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamShareRequest struct via the builder pattern


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


## GetStreamExport

> StreamExportViewResponse GetStreamExport(ctx, streamExportId).Execute()

Get Stream Export



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
    streamExportId := "streamExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamExportAPI.GetStreamExport(context.Background(), streamExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamExportAPI.GetStreamExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamExport`: StreamExportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamExportAPI.GetStreamExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamExportViewResponse**](StreamExportViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStreamShares

> StreamShareListResponse ListStreamShares(ctx, streamExportId).Execute()

List Stream Shares



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
    streamExportId := "streamExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamExportAPI.ListStreamShares(context.Background(), streamExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamExportAPI.ListStreamShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStreamShares`: StreamShareListResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamExportAPI.ListStreamShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStreamSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamShareListResponse**](StreamShareListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

