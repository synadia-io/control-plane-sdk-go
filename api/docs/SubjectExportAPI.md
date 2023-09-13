# \SubjectExportAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubjectShares**](SubjectExportAPI.md#CreateSubjectShares) | **Post** /subject-exports/{subjectExportId}/shares | Create Subject Shares
[**DeleteSubjectExport**](SubjectExportAPI.md#DeleteSubjectExport) | **Delete** /subject-exports/{subjectExportId} | Delete Subject Export
[**DeleteSubjectShare**](SubjectExportAPI.md#DeleteSubjectShare) | **Delete** /subject-exports/{subjectExportId}/shares/{targetAccountNKeyPublic} | Delete Subject Share
[**GetSubjectExport**](SubjectExportAPI.md#GetSubjectExport) | **Get** /subject-exports/{subjectExportId} | Get Subject Export
[**ListSubjectShares**](SubjectExportAPI.md#ListSubjectShares) | **Get** /subject-exports/{subjectExportId}/shares | List Subject Shares
[**UpdateSubjectExport**](SubjectExportAPI.md#UpdateSubjectExport) | **Patch** /subject-exports/{subjectExportId} | Update Subject Export



## CreateSubjectShares

> SubjectShareViewResponse CreateSubjectShares(ctx, subjectExportId).SubjectShareCreateRequest(subjectShareCreateRequest).Execute()

Create Subject Shares



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
    subjectExportId := "subjectExportId_example" // string | 
    subjectShareCreateRequest := *openapiclient.NewSubjectShareCreateRequest("TargetAccountNkeyPublic_example") // SubjectShareCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectExportAPI.CreateSubjectShares(context.Background(), subjectExportId).SubjectShareCreateRequest(subjectShareCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.CreateSubjectShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubjectShares`: SubjectShareViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectExportAPI.CreateSubjectShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubjectSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectShareCreateRequest** | [**SubjectShareCreateRequest**](SubjectShareCreateRequest.md) |  | 

### Return type

[**SubjectShareViewResponse**](SubjectShareViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubjectExport

> DeleteSubjectExport(ctx, subjectExportId).Execute()

Delete Subject Export



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
    subjectExportId := "subjectExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubjectExportAPI.DeleteSubjectExport(context.Background(), subjectExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.DeleteSubjectExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectExportRequest struct via the builder pattern


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


## DeleteSubjectShare

> DeleteSubjectShare(ctx, subjectExportId, targetAccountNKeyPublic).Execute()

Delete Subject Share



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
    subjectExportId := "subjectExportId_example" // string | 
    targetAccountNKeyPublic := "targetAccountNKeyPublic_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubjectExportAPI.DeleteSubjectShare(context.Background(), subjectExportId, targetAccountNKeyPublic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.DeleteSubjectShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 
**targetAccountNKeyPublic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectShareRequest struct via the builder pattern


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


## GetSubjectExport

> SubjectExportViewResponse GetSubjectExport(ctx, subjectExportId).Execute()

Get Subject Export



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
    subjectExportId := "subjectExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectExportAPI.GetSubjectExport(context.Background(), subjectExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.GetSubjectExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectExport`: SubjectExportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectExportAPI.GetSubjectExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectExportViewResponse**](SubjectExportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubjectShares

> SubjectShareListResponse ListSubjectShares(ctx, subjectExportId).Execute()

List Subject Shares



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
    subjectExportId := "subjectExportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectExportAPI.ListSubjectShares(context.Background(), subjectExportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.ListSubjectShares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubjectShares`: SubjectShareListResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectExportAPI.ListSubjectShares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubjectSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectShareListResponse**](SubjectShareListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubjectExport

> SubjectExportViewResponse UpdateSubjectExport(ctx, subjectExportId).SubjectExportUpdateRequest(subjectExportUpdateRequest).Execute()

Update Subject Export



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
    subjectExportId := "subjectExportId_example" // string | 
    subjectExportUpdateRequest := *openapiclient.NewSubjectExportUpdateRequest() // SubjectExportUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectExportAPI.UpdateSubjectExport(context.Background(), subjectExportId).SubjectExportUpdateRequest(subjectExportUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectExportAPI.UpdateSubjectExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubjectExport`: SubjectExportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectExportAPI.UpdateSubjectExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectExportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubjectExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectExportUpdateRequest** | [**SubjectExportUpdateRequest**](SubjectExportUpdateRequest.md) |  | 

### Return type

[**SubjectExportViewResponse**](SubjectExportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

