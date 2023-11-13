# \MirrorAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMirrorPullConsumer**](MirrorAPI.md#CreateMirrorPullConsumer) | **Post** /jetstream/mirror/{streamId}/consumers/pull | Create Pull Consumer
[**CreateMirrorPushConsumer**](MirrorAPI.md#CreateMirrorPushConsumer) | **Post** /jetstream/mirror/{streamId}/consumers/push | Create Push consumer
[**DeleteMirror**](MirrorAPI.md#DeleteMirror) | **Delete** /jetstream/mirror/{streamId} | Delete Mirror
[**GetMirror**](MirrorAPI.md#GetMirror) | **Get** /jetstream/mirror/{streamId} | Get Mirror
[**ListMirrorConsumers**](MirrorAPI.md#ListMirrorConsumers) | **Get** /jetstream/mirror/{streamId}/consumers | List Consumers
[**UpdateMirror**](MirrorAPI.md#UpdateMirror) | **Patch** /jetstream/mirror/{streamId} | Update Mirror



## CreateMirrorPullConsumer

> JSPullConsumerInfoResponse CreateMirrorPullConsumer(ctx, streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()

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
    resp, r, err := apiClient.MirrorAPI.CreateMirrorPullConsumer(context.Background(), streamId).JSPullConsumerConfigRequest(jSPullConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.CreateMirrorPullConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMirrorPullConsumer`: JSPullConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MirrorAPI.CreateMirrorPullConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMirrorPullConsumerRequest struct via the builder pattern


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


## CreateMirrorPushConsumer

> JSPushConsumerInfoResponse CreateMirrorPushConsumer(ctx, streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()

Create Push consumer



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
    resp, r, err := apiClient.MirrorAPI.CreateMirrorPushConsumer(context.Background(), streamId).JSPushConsumerConfigRequest(jSPushConsumerConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.CreateMirrorPushConsumer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMirrorPushConsumer`: JSPushConsumerInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MirrorAPI.CreateMirrorPushConsumer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMirrorPushConsumerRequest struct via the builder pattern


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


## DeleteMirror

> DeleteMirror(ctx, streamId).Execute()

Delete Mirror



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
    r, err := apiClient.MirrorAPI.DeleteMirror(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.DeleteMirror``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMirrorRequest struct via the builder pattern


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


## GetMirror

> JSMirrorInfoResponse GetMirror(ctx, streamId).Execute()

Get Mirror



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
    resp, r, err := apiClient.MirrorAPI.GetMirror(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.GetMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMirror`: JSMirrorInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MirrorAPI.GetMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JSMirrorInfoResponse**](JSMirrorInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMirrorConsumers

> JSConsumerInfoListResponse ListMirrorConsumers(ctx, streamId).Execute()

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
    resp, r, err := apiClient.MirrorAPI.ListMirrorConsumers(context.Background(), streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.ListMirrorConsumers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMirrorConsumers`: JSConsumerInfoListResponse
    fmt.Fprintf(os.Stdout, "Response from `MirrorAPI.ListMirrorConsumers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMirrorConsumersRequest struct via the builder pattern


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


## UpdateMirror

> JSMirrorInfoResponse UpdateMirror(ctx, streamId).JSMirrorConfigRequest(jSMirrorConfigRequest).Execute()

Update Mirror



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
    jSMirrorConfigRequest := *openapiclient.NewJSMirrorConfigRequest(false, false, false, false, openapiclient.DiscardPolicy("old"), int64(123), int64(123), int32(123), int64(123), int64(123), "Name_example", int32(123), openapiclient.RetentionPolicy("limits"), false, openapiclient.StorageType("file"), *openapiclient.NewStreamSource("Name_example"), false) // JSMirrorConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MirrorAPI.UpdateMirror(context.Background(), streamId).JSMirrorConfigRequest(jSMirrorConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MirrorAPI.UpdateMirror``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMirror`: JSMirrorInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `MirrorAPI.UpdateMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jSMirrorConfigRequest** | [**JSMirrorConfigRequest**](JSMirrorConfigRequest.md) |  | 

### Return type

[**JSMirrorInfoResponse**](JSMirrorInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

