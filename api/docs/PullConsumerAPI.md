# \PullConsumerAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePullConsumer**](PullConsumerAPI.md#DeletePullConsumer) | **Delete** /consumers/pull/{consumerId} | Delete Pull Consumer
[**GetPullConsumerInfo**](PullConsumerAPI.md#GetPullConsumerInfo) | **Get** /consumers/pull/{consumerId} | Get Pull Consumer
[**UpdatePullConsumer**](PullConsumerAPI.md#UpdatePullConsumer) | **Patch** /consumers/pull/{consumerId} | Update Pull Consumer



## DeletePullConsumer

> DeletePullConsumer(ctx, consumerId).Execute()

Delete Pull Consumer



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
    consumerId := "consumerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PullConsumerAPI.DeletePullConsumer(context.Background(), consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullConsumerAPI.DeletePullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePullConsumerRequest struct via the builder pattern


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


## GetPullConsumerInfo

> JSPullConsumerInfoResponse GetPullConsumerInfo(ctx, consumerId).Execute()

Get Pull Consumer



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
    consumerId := "consumerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullConsumerAPI.GetPullConsumerInfo(context.Background(), consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullConsumerAPI.GetPullConsumerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPullConsumerInfo`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PullConsumerAPI.GetPullConsumerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPullConsumerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSPullConsumerInfoResponse**](JSPullConsumerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePullConsumer

> JSPullConsumerInfoResponse UpdatePullConsumer(ctx, consumerId).JSPullConsumerUpdateRequest(jSPullConsumerUpdateRequest).Execute()

Update Pull Consumer



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
    consumerId := "consumerId_example" // string | 
    jSPullConsumerUpdateRequest := *openapiclient.NewJSPullConsumerUpdateRequest() // JSPullConsumerUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PullConsumerAPI.UpdatePullConsumer(context.Background(), consumerId).JSPullConsumerUpdateRequest(jSPullConsumerUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PullConsumerAPI.UpdatePullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePullConsumer`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PullConsumerAPI.UpdatePullConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePullConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPullConsumerUpdateRequest** | [**JSPullConsumerUpdateRequest**](JSPullConsumerUpdateRequest.md) |  | 

### Return type

[**JSPullConsumerInfoResponse**](JSPullConsumerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

