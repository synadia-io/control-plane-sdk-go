# SigningKeyGroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to [**UserPermissionLimits**](UserPermissionLimits.md) |  | [optional] 

## Methods

### NewSigningKeyGroupUpdateRequest

`func NewSigningKeyGroupUpdateRequest() *SigningKeyGroupUpdateRequest`

NewSigningKeyGroupUpdateRequest instantiates a new SigningKeyGroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyGroupUpdateRequestWithDefaults

`func NewSigningKeyGroupUpdateRequestWithDefaults() *SigningKeyGroupUpdateRequest`

NewSigningKeyGroupUpdateRequestWithDefaults instantiates a new SigningKeyGroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SigningKeyGroupUpdateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SigningKeyGroupUpdateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SigningKeyGroupUpdateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SigningKeyGroupUpdateRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *SigningKeyGroupUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SigningKeyGroupUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SigningKeyGroupUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SigningKeyGroupUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *SigningKeyGroupUpdateRequest) GetScope() UserPermissionLimits`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SigningKeyGroupUpdateRequest) GetScopeOk() (*UserPermissionLimits, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SigningKeyGroupUpdateRequest) SetScope(v UserPermissionLimits)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SigningKeyGroupUpdateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


