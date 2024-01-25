# \StreamAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePullConsumer**](StreamAPI.md#CreatePullConsumer) | **Post** /jetstream/stream/{streamId}/consumers/pull | Create Pull Consumer
[**CreatePushConsumer**](StreamAPI.md#CreatePushConsumer) | **Post** /jetstream/stream/{streamId}/consumers/push | Create Push Consumer
[**DeleteStream**](StreamAPI.md#DeleteStream) | **Delete** /jetstream/stream/{streamId} | Delete Stream
[**GetStreamInfo**](StreamAPI.md#GetStreamInfo) | **Get** /jetstream/stream/{streamId} | Get Stream
[**ListConsumers**](StreamAPI.md#ListConsumers) | **Get** /jetstream/stream/{streamId}/consumers | List Consumers
[**UpdateStream**](StreamAPI.md#UpdateStream) | **Patch** /jetstream/stream/{streamId} | Update Stream



## CreatePullConsumer

> JSPullConsumerInfoResponse CreatePullConsumer(ctx, streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()

Create Pull Consumer



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
    streamId := "streamId_example" // string | 
    jSPullConsumerConfigRequest := *openapiclient.NewJSPullConsumerConfigRequest(openapiclient.AckPolicy("none"), openapiclient.DeliverPolicy("all"), int32(123), openapiclient.ReplayPolicy("instant")) // JSPullConsumerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamAPI.CreatePullConsumer(context.Background(), streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.CreatePullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePullConsumer`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamAPI.CreatePullConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePullConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPullConsumerConfigRequest** | [**JSPullConsumerConfigRequest**](JSPullConsumerConfigRequest.md) |  | 

### Return type

[**JSPullConsumerInfoResponse**](JSPullConsumerInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePushConsumer

> JSPushConsumerInfoResponse CreatePushConsumer(ctx, streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()

Create Push Consumer



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
    streamId := "streamId_example" // string | 
    jSPushConsumerConfigRequest := *openapiclient.NewJSPushConsumerConfigRequest(openapiclient.AckPolicy("none"), openapiclient.DeliverPolicy("all"), int32(123), openapiclient.ReplayPolicy("instant")) // JSPushConsumerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamAPI.CreatePushConsumer(context.Background(), streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.CreatePushConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePushConsumer`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamAPI.CreatePushConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePushConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPushConsumerConfigRequest** | [**JSPushConsumerConfigRequest**](JSPushConsumerConfigRequest.md) |  | 

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


## DeleteStream

> DeleteStream(ctx, streamId).Execute()

Delete Stream



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
    streamId := "streamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.DeleteStream(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.DeleteStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStreamRequest struct via the builder pattern


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


## GetStreamInfo

> JSStreamInfoResponse GetStreamInfo(ctx, streamId).Execute()

Get Stream



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
    streamId := "streamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamAPI.GetStreamInfo(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.GetStreamInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStreamInfo`: JSStreamInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamAPI.GetStreamInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStreamInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSStreamInfoResponse**](JSStreamInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConsumers

> JSConsumerInfoListResponse ListConsumers(ctx, streamId).Execute()

List Consumers



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
    streamId := "streamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamAPI.ListConsumers(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.ListConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConsumers`: JSConsumerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamAPI.ListConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSConsumerInfoListResponse**](JSConsumerInfoListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStream

> JSStreamInfoResponse UpdateStream(ctx, streamId).JSStreamConfigRequest(jSStreamConfigRequest).Execute()

Update Stream



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
    streamId := "streamId_example" // string | 
    jSStreamConfigRequest := *openapiclient.NewJSStreamConfigRequest(false, false, false, false, openapiclient.DiscardPolicy("old"), int64(123), int64(123), int32(123), int64(123), int64(123), "Name_example", int32(123), openapiclient.RetentionPolicy("limits"), false, openapiclient.StorageType("file")) // JSStreamConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StreamAPI.UpdateStream(context.Background(), streamId).JSStreamConfigRequest(jSStreamConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.UpdateStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStream`: JSStreamInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `StreamAPI.UpdateStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSStreamConfigRequest** | [**JSStreamConfigRequest**](JSStreamConfigRequest.md) |  | 

### Return type

[**JSStreamInfoResponse**](JSStreamInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

