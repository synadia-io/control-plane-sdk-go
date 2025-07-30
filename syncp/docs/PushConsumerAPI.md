# \PushConsumerAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePushConsumer**](PushConsumerAPI.md#DeletePushConsumer) | **Delete** /core/beta/consumers/push/{consumerId} | Delete Push Consumer
[**GetPushConsumerInfo**](PushConsumerAPI.md#GetPushConsumerInfo) | **Get** /core/beta/consumers/push/{consumerId} | Get Push Consumer
[**UpdatePushConsumer**](PushConsumerAPI.md#UpdatePushConsumer) | **Patch** /core/beta/consumers/push/{consumerId} | Update Push Consumer



## DeletePushConsumer

> DeletePushConsumer(ctx, consumerId).Execute()

Delete Push Consumer



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
    consumerId := "consumerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PushConsumerAPI.DeletePushConsumer(context.Background(), consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushConsumerAPI.DeletePushConsumer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePushConsumerRequest struct via the builder pattern


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


## GetPushConsumerInfo

> JSPushConsumerInfoResponse GetPushConsumerInfo(ctx, consumerId).Execute()

Get Push Consumer



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
    consumerId := "consumerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushConsumerAPI.GetPushConsumerInfo(context.Background(), consumerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushConsumerAPI.GetPushConsumerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPushConsumerInfo`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PushConsumerAPI.GetPushConsumerInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPushConsumerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSPushConsumerInfoResponse**](JSPushConsumerInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePushConsumer

> JSPushConsumerInfoResponse UpdatePushConsumer(ctx, consumerId).JSPushConsumerUpdateRequest(jSPushConsumerUpdateRequest).Execute()

Update Push Consumer



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
    consumerId := "consumerId_example" // string | 
    jSPushConsumerUpdateRequest := *openapiclient.NewJSPushConsumerUpdateRequest() // JSPushConsumerUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushConsumerAPI.UpdatePushConsumer(context.Background(), consumerId).JSPushConsumerUpdateRequest(jSPushConsumerUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushConsumerAPI.UpdatePushConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePushConsumer`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PushConsumerAPI.UpdatePushConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**consumerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePushConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPushConsumerUpdateRequest** | [**JSPushConsumerUpdateRequest**](JSPushConsumerUpdateRequest.md) |  | 

### Return type

[**JSPushConsumerInfoResponse**](JSPushConsumerInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

