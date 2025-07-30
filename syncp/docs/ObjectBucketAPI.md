# \ObjectBucketAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjPullConsumer**](ObjectBucketAPI.md#CreateObjPullConsumer) | **Post** /core/beta/jetstream/object-bucket/{streamId}/consumers/pull | Create Pull Consumer
[**CreateObjPushConsumer**](ObjectBucketAPI.md#CreateObjPushConsumer) | **Post** /core/beta/jetstream/object-bucket/{streamId}/consumers/push | Create Push Consumer
[**DeleteObjectBucket**](ObjectBucketAPI.md#DeleteObjectBucket) | **Delete** /core/beta/jetstream/object-bucket/{streamId} | Delete Object Bucket
[**GetObjectBucket**](ObjectBucketAPI.md#GetObjectBucket) | **Get** /core/beta/jetstream/object-bucket/{streamId} | Get Object Bucket
[**ListObjConsumers**](ObjectBucketAPI.md#ListObjConsumers) | **Get** /core/beta/jetstream/object-bucket/{streamId}/consumers | List Consumers
[**PurgeObjBucket**](ObjectBucketAPI.md#PurgeObjBucket) | **Delete** /core/beta/jetstream/object-bucket/{streamId}/purge | Purge Object Bucket
[**UpdateObjectBucket**](ObjectBucketAPI.md#UpdateObjectBucket) | **Patch** /core/beta/jetstream/object-bucket/{streamId} | Update Object Bucket



## CreateObjPullConsumer

> JSPullConsumerInfoResponse CreateObjPullConsumer(ctx, streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()

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
    jSPullConsumerConfigRequest := *openapiclient.NewJSPullConsumerConfigRequest(openapiclient.AckPolicy("none"), openapiclient.DeliverPolicy("all"), int64(123), openapiclient.ReplayPolicy("instant")) // JSPullConsumerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectBucketAPI.CreateObjPullConsumer(context.Background(), streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.CreateObjPullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjPullConsumer`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectBucketAPI.CreateObjPullConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjPullConsumerRequest struct via the builder pattern


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


## CreateObjPushConsumer

> JSPushConsumerInfoResponse CreateObjPushConsumer(ctx, streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()

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
    jSPushConsumerConfigRequest := *openapiclient.NewJSPushConsumerConfigRequest(openapiclient.AckPolicy("none"), openapiclient.DeliverPolicy("all"), int64(123), openapiclient.ReplayPolicy("instant")) // JSPushConsumerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectBucketAPI.CreateObjPushConsumer(context.Background(), streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.CreateObjPushConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjPushConsumer`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectBucketAPI.CreateObjPushConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjPushConsumerRequest struct via the builder pattern


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


## DeleteObjectBucket

> DeleteObjectBucket(ctx, streamId).Execute()

Delete Object Bucket



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
    r, err := apiClient.ObjectBucketAPI.DeleteObjectBucket(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.DeleteObjectBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteObjectBucketRequest struct via the builder pattern


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


## GetObjectBucket

> JSObjectBucketViewResponse GetObjectBucket(ctx, streamId).Execute()

Get Object Bucket



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
    resp, r, err := apiClient.ObjectBucketAPI.GetObjectBucket(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.GetObjectBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectBucket`: JSObjectBucketViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectBucketAPI.GetObjectBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSObjectBucketViewResponse**](JSObjectBucketViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjConsumers

> JSConsumerInfoListResponse ListObjConsumers(ctx, streamId).Execute()

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
    resp, r, err := apiClient.ObjectBucketAPI.ListObjConsumers(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.ListObjConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjConsumers`: JSConsumerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectBucketAPI.ListObjConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjConsumersRequest struct via the builder pattern


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


## PurgeObjBucket

> PurgeObjBucket(ctx, streamId).Execute()

Purge Object Bucket



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
    r, err := apiClient.ObjectBucketAPI.PurgeObjBucket(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.PurgeObjBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPurgeObjBucketRequest struct via the builder pattern


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


## UpdateObjectBucket

> JSObjectBucketViewResponse UpdateObjectBucket(ctx, streamId).JSObjectBucketUpdateRequest(jSObjectBucketUpdateRequest).Execute()

Update Object Bucket



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
    jSObjectBucketUpdateRequest := *openapiclient.NewJSObjectBucketUpdateRequest(*openapiclient.NewUpdatableObjectBucketConfig()) // JSObjectBucketUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectBucketAPI.UpdateObjectBucket(context.Background(), streamId).JSObjectBucketUpdateRequest(jSObjectBucketUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectBucketAPI.UpdateObjectBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectBucket`: JSObjectBucketViewResponse
    fmt.Fprintf(os.Stdout, "Response from `ObjectBucketAPI.UpdateObjectBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSObjectBucketUpdateRequest** | [**JSObjectBucketUpdateRequest**](JSObjectBucketUpdateRequest.md) |  | 

### Return type

[**JSObjectBucketViewResponse**](JSObjectBucketViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

