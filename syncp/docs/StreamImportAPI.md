# \StreamImportAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStreamImport**](StreamImportAPI.md#DeleteStreamImport) | **Delete** /stream-imports/{streamImportId} | Delete Stream Import
[**GetStreamImport**](StreamImportAPI.md#GetStreamImport) | **Get** /stream-imports/{streamImportId} | Get Stream Import



## DeleteStreamImport

> DeleteStreamImport(ctx, streamImportId).Execute()

Delete Stream Import



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
    streamImportId := "streamImportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamImportAPI.DeleteStreamImport(context.Background(), streamImportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamImportAPI.DeleteStreamImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamImportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamImportRequest struct via the builder pattern


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


## GetStreamImport

> StreamImportViewResponse GetStreamImport(ctx, streamImportId).Execute()

Get Stream Import



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
    streamImportId := "streamImportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamImportAPI.GetStreamImport(context.Background(), streamImportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamImportAPI.GetStreamImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamImport`: StreamImportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamImportAPI.GetStreamImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamImportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamImportViewResponse**](StreamImportViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

