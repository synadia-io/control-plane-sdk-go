# \AdminAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAppAdmin**](AdminAPI.md#CheckAppAdmin) | **Get** /core/beta/admin/app-user | Check for Admin User
[**CreateAppAdmin**](AdminAPI.md#CreateAppAdmin) | **Post** /core/beta/admin/app-user | Create an App Admin User



## CheckAppAdmin

> CheckAppAdmin(ctx).Execute()

Check for Admin User



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CheckAppAdmin(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CheckAppAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAppAdminRequest struct via the builder pattern


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


## CreateAppAdmin

> AdminAppUserCreateResponse CreateAppAdmin(ctx).AdminAppUserCreateRequest(adminAppUserCreateRequest).Execute()

Create an App Admin User



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
    adminAppUserCreateRequest := *openapiclient.NewAdminAppUserCreateRequest("Password_example", "Username_example") // AdminAppUserCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.CreateAppAdmin(context.Background()).AdminAppUserCreateRequest(adminAppUserCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateAppAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAppAdmin`: AdminAppUserCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.CreateAppAdmin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminAppUserCreateRequest** | [**AdminAppUserCreateRequest**](AdminAppUserCreateRequest.md) |  | 

### Return type

[**AdminAppUserCreateResponse**](AdminAppUserCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

