# \AlertAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcknowledgeAlert**](AlertAPI.md#AcknowledgeAlert) | **Patch** /alerts/{alertId}/acknowledge | Acknowledge Alert
[**GetAlert**](AlertAPI.md#GetAlert) | **Get** /alerts/{alertId} | Get Alert



## AcknowledgeAlert

> AcknowledgeAlert(ctx, alertId).Execute()

Acknowledge Alert



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
    alertId := "alertId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertAPI.AcknowledgeAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.AcknowledgeAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcknowledgeAlertRequest struct via the builder pattern


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


## GetAlert

> AlertViewResponse GetAlert(ctx, alertId).Execute()

Get Alert



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
    alertId := "alertId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertAPI.GetAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlert`: AlertViewResponse
    fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertViewResponse**](AlertViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

