# \SubjectImportAPI

All URIs are relative to *http://localhost/api/core/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubjectImport**](SubjectImportAPI.md#DeleteSubjectImport) | **Delete** /subject-imports/{subjectImportId} | Delete Subject Import
[**GetSubjectImport**](SubjectImportAPI.md#GetSubjectImport) | **Get** /subject-imports/{subjectImportId} | Get Subject Import
[**UpdateSubjectImport**](SubjectImportAPI.md#UpdateSubjectImport) | **Patch** /subject-imports/{subjectImportId} | Update Subject Import



## DeleteSubjectImport

> DeleteSubjectImport(ctx, subjectImportId).Execute()

Delete Subject Import



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
    subjectImportId := "subjectImportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubjectImportAPI.DeleteSubjectImport(context.Background(), subjectImportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectImportAPI.DeleteSubjectImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectImportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubjectImportRequest struct via the builder pattern


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


## GetSubjectImport

> SubjectImportViewResponse GetSubjectImport(ctx, subjectImportId).Execute()

Get Subject Import



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
    subjectImportId := "subjectImportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectImportAPI.GetSubjectImport(context.Background(), subjectImportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectImportAPI.GetSubjectImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubjectImport`: SubjectImportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectImportAPI.GetSubjectImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectImportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubjectImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubjectImportViewResponse**](SubjectImportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubjectImport

> SubjectImportViewResponse UpdateSubjectImport(ctx, subjectImportId).SubjectImportUpdateRequest(subjectImportUpdateRequest).Execute()

Update Subject Import



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
    subjectImportId := "subjectImportId_example" // string | 
    subjectImportUpdateRequest := *openapiclient.NewSubjectImportUpdateRequest() // SubjectImportUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubjectImportAPI.UpdateSubjectImport(context.Background(), subjectImportId).SubjectImportUpdateRequest(subjectImportUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubjectImportAPI.UpdateSubjectImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubjectImport`: SubjectImportViewResponse
    fmt.Fprintf(os.Stdout, "Response from `SubjectImportAPI.UpdateSubjectImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectImportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubjectImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subjectImportUpdateRequest** | [**SubjectImportUpdateRequest**](SubjectImportUpdateRequest.md) |  | 

### Return type

[**SubjectImportViewResponse**](SubjectImportViewResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

