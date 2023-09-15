# AppRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Enabled** | **bool** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Policies** | **[]string** |  | 
**Scope** | [**AppRoleScope**](AppRoleScope.md) |  | 
**SortOrder** | **float32** |  | 

## Methods

### NewAppRole

`func NewAppRole(description string, enabled bool, id string, name string, policies []string, scope AppRoleScope, sortOrder float32, ) *AppRole`

NewAppRole instantiates a new AppRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRoleWithDefaults

`func NewAppRoleWithDefaults() *AppRole`

NewAppRoleWithDefaults instantiates a new AppRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRole) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *AppRole) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppRole) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppRole) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *AppRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRole) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRole) SetName(v string)`

SetName sets Name field to given value.


### GetPolicies

`func (o *AppRole) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AppRole) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AppRole) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetScope

`func (o *AppRole) GetScope() AppRoleScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppRole) GetScopeOk() (*AppRoleScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppRole) SetScope(v AppRoleScope)`

SetScope sets Scope field to given value.


### GetSortOrder

`func (o *AppRole) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AppRole) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AppRole) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


