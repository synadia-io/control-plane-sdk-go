# SystemLimitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | [**TenantLimits**](TenantLimits.md) |  | 
**Available** | [**TenantLimits**](TenantLimits.md) |  | 
**Total** | [**TenantLimits**](TenantLimits.md) |  | 

## Methods

### NewSystemLimitsResponse

`func NewSystemLimitsResponse(allocated TenantLimits, available TenantLimits, total TenantLimits, ) *SystemLimitsResponse`

NewSystemLimitsResponse instantiates a new SystemLimitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemLimitsResponseWithDefaults

`func NewSystemLimitsResponseWithDefaults() *SystemLimitsResponse`

NewSystemLimitsResponseWithDefaults instantiates a new SystemLimitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *SystemLimitsResponse) GetAllocated() TenantLimits`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *SystemLimitsResponse) GetAllocatedOk() (*TenantLimits, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *SystemLimitsResponse) SetAllocated(v TenantLimits)`

SetAllocated sets Allocated field to given value.


### GetAvailable

`func (o *SystemLimitsResponse) GetAvailable() TenantLimits`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *SystemLimitsResponse) GetAvailableOk() (*TenantLimits, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *SystemLimitsResponse) SetAvailable(v TenantLimits)`

SetAvailable sets Available field to given value.


### GetTotal

`func (o *SystemLimitsResponse) GetTotal() TenantLimits`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SystemLimitsResponse) GetTotalOk() (*TenantLimits, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SystemLimitsResponse) SetTotal(v TenantLimits)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


