# \PlatformComponentsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformComponentConnect**](PlatformComponentsAPI.md#PlatformComponentConnect) | **Post** /core/beta/platform-components/connect | Connect a Platform Component



## PlatformComponentConnect

> PlatformComponentsConnectViewResponse PlatformComponentConnect(ctx, id).PlatformComponentsConnectRequest(platformComponentsConnectRequest).Execute()

Connect a Platform Component



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
    id := "id_example" // string | 
    platformComponentsConnectRequest := *openapiclient.NewPlatformComponentsConnectRequest("NkeyPublic_example") // PlatformComponentsConnectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlatformComponentsAPI.PlatformComponentConnect(context.Background(), id).PlatformComponentsConnectRequest(platformComponentsConnectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlatformComponentsAPI.PlatformComponentConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlatformComponentConnect`: PlatformComponentsConnectViewResponse
    fmt.Fprintf(os.Stdout, "Response from `PlatformComponentsAPI.PlatformComponentConnect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComponentConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformComponentsConnectRequest** | [**PlatformComponentsConnectRequest**](PlatformComponentsConnectRequest.md) |  | 

### Return type

[**PlatformComponentsConnectViewResponse**](PlatformComponentsConnectViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

