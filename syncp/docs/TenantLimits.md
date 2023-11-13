# TenantLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumAccounts** | **int32** |  | 
**ResourceLimits** | [**OperatorLimits**](OperatorLimits.md) |  | 

## Methods

### NewTenantLimits

`func NewTenantLimits(numAccounts int32, resourceLimits OperatorLimits, ) *TenantLimits`

NewTenantLimits instantiates a new TenantLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantLimitsWithDefaults

`func NewTenantLimitsWithDefaults() *TenantLimits`

NewTenantLimitsWithDefaults instantiates a new TenantLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumAccounts

`func (o *TenantLimits) GetNumAccounts() int32`

GetNumAccounts returns the NumAccounts field if non-nil, zero value otherwise.

### GetNumAccountsOk

`func (o *TenantLimits) GetNumAccountsOk() (*int32, bool)`

GetNumAccountsOk returns a tuple with the NumAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAccounts

`func (o *TenantLimits) SetNumAccounts(v int32)`

SetNumAccounts sets NumAccounts field to given value.


### GetResourceLimits

`func (o *TenantLimits) GetResourceLimits() OperatorLimits`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *TenantLimits) GetResourceLimitsOk() (*OperatorLimits, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *TenantLimits) SetResourceLimits(v OperatorLimits)`

SetResourceLimits sets ResourceLimits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


