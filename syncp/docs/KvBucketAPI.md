# \KvBucketAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKvPullConsumer**](KvBucketAPI.md#CreateKvPullConsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/pull | Create Pull Consumer
[**CreateKvPushConsumer**](KvBucketAPI.md#CreateKvPushConsumer) | **Post** /jetstream/kv-bucket/{streamId}/consumers/push | Create Push Consumer
[**DeleteKvBucket**](KvBucketAPI.md#DeleteKvBucket) | **Delete** /jetstream/kv-bucket/{streamId} | Delete KV Bucket
[**GetKvBucket**](KvBucketAPI.md#GetKvBucket) | **Get** /jetstream/kv-bucket/{streamId} | Get KV Bucket
[**ListKvConsumers**](KvBucketAPI.md#ListKvConsumers) | **Get** /jetstream/kv-bucket/{streamId}/consumers | List Consumers
[**UpdateKvBucket**](KvBucketAPI.md#UpdateKvBucket) | **Patch** /jetstream/kv-bucket/{streamId} | Update KV Bucket



## CreateKvPullConsumer

> JSPullConsumerInfoResponse CreateKvPullConsumer(ctx, streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()

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
    resp, r, err := apiClient.KvBucketAPI.CreateKvPullConsumer(context.Background(), streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.CreateKvPullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKvPullConsumer`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `KvBucketAPI.CreateKvPullConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKvPullConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPullConsumerConfigRequest** | [**JSPullConsumerConfigRequest**](JSPullConsumerConfigRequest.md) |  | 

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


## CreateKvPushConsumer

> JSPushConsumerInfoResponse CreateKvPushConsumer(ctx, streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()

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
    resp, r, err := apiClient.KvBucketAPI.CreateKvPushConsumer(context.Background(), streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.CreateKvPushConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKvPushConsumer`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `KvBucketAPI.CreateKvPushConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKvPushConsumerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSPushConsumerConfigRequest** | [**JSPushConsumerConfigRequest**](JSPushConsumerConfigRequest.md) |  | 

### Return type

[**JSPushConsumerInfoResponse**](JSPushConsumerInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKvBucket

> DeleteKvBucket(ctx, streamId).Execute()

Delete KV Bucket



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
    r, err := apiClient.KvBucketAPI.DeleteKvBucket(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.DeleteKvBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKvBucketRequest struct via the builder pattern


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


## GetKvBucket

> JSKVBucketViewResponse GetKvBucket(ctx, streamId).Execute()

Get KV Bucket



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
    resp, r, err := apiClient.KvBucketAPI.GetKvBucket(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.GetKvBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKvBucket`: JSKVBucketViewResponse
    fmt.Fprintf(os.Stdout, "Response from `KvBucketAPI.GetKvBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKvBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSKVBucketViewResponse**](JSKVBucketViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKvConsumers

> JSConsumerInfoListResponse ListKvConsumers(ctx, streamId).Execute()

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
    resp, r, err := apiClient.KvBucketAPI.ListKvConsumers(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.ListKvConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKvConsumers`: JSConsumerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `KvBucketAPI.ListKvConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKvConsumersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSConsumerInfoListResponse**](JSConsumerInfoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKvBucket

> JSKVBucketViewResponse UpdateKvBucket(ctx, streamId).JSKVBucketUpdateRequest(jSKVBucketUpdateRequest).Execute()

Update KV Bucket



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
    jSKVBucketUpdateRequest := *openapiclient.NewJSKVBucketUpdateRequest(*openapiclient.NewUpdatableKVBucketConfig()) // JSKVBucketUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KvBucketAPI.UpdateKvBucket(context.Background(), streamId).JSKVBucketUpdateRequest(jSKVBucketUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KvBucketAPI.UpdateKvBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKvBucket`: JSKVBucketViewResponse
    fmt.Fprintf(os.Stdout, "Response from `KvBucketAPI.UpdateKvBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKvBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSKVBucketUpdateRequest** | [**JSKVBucketUpdateRequest**](JSKVBucketUpdateRequest.md) |  | 

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

