# \LicenseAPI

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLicense**](LicenseAPI.md#DeleteLicense) | **Delete** /core/beta/license | Delete License
[**GetLicenseStatus**](LicenseAPI.md#GetLicenseStatus) | **Get** /core/beta/license | Get License Status
[**UpdateLicense**](LicenseAPI.md#UpdateLicense) | **Put** /core/beta/license | Update License



## DeleteLicense

> LicenseStatusResponse DeleteLicense(ctx).Execute()

Delete License



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
    resp, r, err := apiClient.LicenseAPI.DeleteLicense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.DeleteLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLicense`: LicenseStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.DeleteLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLicenseRequest struct via the builder pattern


### Return type

[**LicenseStatusResponse**](LicenseStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseStatus

> LicenseStatusResponse GetLicenseStatus(ctx).Execute()

Get License Status



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
    resp, r, err := apiClient.LicenseAPI.GetLicenseStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.GetLicenseStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseStatus`: LicenseStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.GetLicenseStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseStatusRequest struct via the builder pattern


### Return type

[**LicenseStatusResponse**](LicenseStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLicense

> LicenseStatusResponse UpdateLicense(ctx).LicenseUpdateRequest(licenseUpdateRequest).Execute()

Update License



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
    licenseUpdateRequest := *openapiclient.NewLicenseUpdateRequest("LicenseKey_example") // LicenseUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseAPI.UpdateLicense(context.Background()).LicenseUpdateRequest(licenseUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseAPI.UpdateLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLicense`: LicenseStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `LicenseAPI.UpdateLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licenseUpdateRequest** | [**LicenseUpdateRequest**](LicenseUpdateRequest.md) |  | 

### Return type

[**LicenseStatusResponse**](LicenseStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [sessionAuth](../README.md#sessionAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

