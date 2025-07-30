# \WorkloadsAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkloadLimits**](WorkloadsAPI.md#GetWorkloadLimits) | **Get** /workloads/alpha/{accountId}/limits | Get Workloads limits



## GetWorkloadLimits

> WorkloadLimitsResponse GetWorkloadLimits(ctx, accountId).Execute()

Get Workloads limits



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
    accountId := "accountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkloadsAPI.GetWorkloadLimits(context.Background(), accountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsAPI.GetWorkloadLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkloadLimits`: WorkloadLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkloadsAPI.GetWorkloadLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkloadLimitsResponse**](WorkloadLimitsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

