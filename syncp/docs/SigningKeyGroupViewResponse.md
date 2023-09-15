# SigningKeyGroupViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Disabled** | **bool** |  | 
**DisabledAt** | Pointer to **time.Time** |  | [optional] 
**Id** | **string** |  | 
**IsScoped** | **bool** |  | 
**Name** | **string** |  | 
**Scope** | Pointer to [**UserPermissionLimits**](UserPermissionLimits.md) |  | [optional] 

## Methods

### NewSigningKeyGroupViewResponse

`func NewSigningKeyGroupViewResponse(created time.Time, disabled bool, id string, isScoped bool, name string, ) *SigningKeyGroupViewResponse`

NewSigningKeyGroupViewResponse instantiates a new SigningKeyGroupViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyGroupViewResponseWithDefaults

`func NewSigningKeyGroupViewResponseWithDefaults() *SigningKeyGroupViewResponse`

NewSigningKeyGroupViewResponseWithDefaults instantiates a new SigningKeyGroupViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SigningKeyGroupViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SigningKeyGroupViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SigningKeyGroupViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDisabled

`func (o *SigningKeyGroupViewResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SigningKeyGroupViewResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SigningKeyGroupViewResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetDisabledAt

`func (o *SigningKeyGroupViewResponse) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *SigningKeyGroupViewResponse) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *SigningKeyGroupViewResponse) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *SigningKeyGroupViewResponse) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetId

`func (o *SigningKeyGroupViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningKeyGroupViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningKeyGroupViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsScoped

`func (o *SigningKeyGroupViewResponse) GetIsScoped() bool`

GetIsScoped returns the IsScoped field if non-nil, zero value otherwise.

### GetIsScopedOk

`func (o *SigningKeyGroupViewResponse) GetIsScopedOk() (*bool, bool)`

GetIsScopedOk returns a tuple with the IsScoped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScoped

`func (o *SigningKeyGroupViewResponse) SetIsScoped(v bool)`

SetIsScoped sets IsScoped field to given value.


### GetName

`func (o *SigningKeyGroupViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SigningKeyGroupViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SigningKeyGroupViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *SigningKeyGroupViewResponse) GetScope() UserPermissionLimits`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SigningKeyGroupViewResponse) GetScopeOk() (*UserPermissionLimits, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SigningKeyGroupViewResponse) SetScope(v UserPermissionLimits)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SigningKeyGroupViewResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


