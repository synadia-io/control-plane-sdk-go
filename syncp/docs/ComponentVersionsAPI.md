# \ComponentVersionsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableComponentVersions**](ComponentVersionsAPI.md#GetAvailableComponentVersions) | **Get** /core/beta/component-versions | Get Available Component Versions
[**GetAvailableComponentVersionsByType**](ComponentVersionsAPI.md#GetAvailableComponentVersionsByType) | **Post** /core/beta/component-versions | Get Available Component Versions By Type



## GetAvailableComponentVersions

> map[string][]AvailableChartVersion GetAvailableComponentVersions(ctx).Execute()

Get Available Component Versions



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
    resp, r, err := apiClient.ComponentVersionsAPI.GetAvailableComponentVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionsAPI.GetAvailableComponentVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableComponentVersions`: map[string][]AvailableChartVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionsAPI.GetAvailableComponentVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableComponentVersionsRequest struct via the builder pattern


### Return type

[**map[string][]AvailableChartVersion**](array.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableComponentVersionsByType

> []AvailableChartVersion GetAvailableComponentVersionsByType(ctx).AvailableComponentVersionsRequest(availableComponentVersionsRequest).Execute()

Get Available Component Versions By Type



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
    availableComponentVersionsRequest := *openapiclient.NewAvailableComponentVersionsRequest(int64(123), int64(123)) // AvailableComponentVersionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentVersionsAPI.GetAvailableComponentVersionsByType(context.Background()).AvailableComponentVersionsRequest(availableComponentVersionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentVersionsAPI.GetAvailableComponentVersionsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableComponentVersionsByType`: []AvailableChartVersion
    fmt.Fprintf(os.Stdout, "Response from `ComponentVersionsAPI.GetAvailableComponentVersionsByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableComponentVersionsByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **availableComponentVersionsRequest** | [**AvailableComponentVersionsRequest**](AvailableComponentVersionsRequest.md) |  | 

### Return type

[**[]AvailableChartVersion**](AvailableChartVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

