# \IssuanceAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNatsUserIssuance**](IssuanceAPI.md#GetNatsUserIssuance) | **Get** /core/beta/nats-user-issuances/{issuanceId} | Get nats user issuance



## GetNatsUserIssuance

> NatsUserIssuanceViewResponse GetNatsUserIssuance(ctx, issuanceId).Execute()

Get nats user issuance



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
    issuanceId := "issuanceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IssuanceAPI.GetNatsUserIssuance(context.Background(), issuanceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IssuanceAPI.GetNatsUserIssuance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNatsUserIssuance`: NatsUserIssuanceViewResponse
    fmt.Fprintf(os.Stdout, "Response from `IssuanceAPI.GetNatsUserIssuance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issuanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNatsUserIssuanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NatsUserIssuanceViewResponse**](NatsUserIssuanceViewResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

