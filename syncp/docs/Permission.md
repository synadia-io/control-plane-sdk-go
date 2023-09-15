# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allow** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**Deny** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllow

`func (o *Permission) GetAllow() []string`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *Permission) GetAllowOk() (*[]string, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *Permission) SetAllow(v []string)`

SetAllow sets Allow field to given value.

### HasAllow

`func (o *Permission) HasAllow() bool`

HasAllow returns a boolean if a field has been set.

### SetAllowNil

`func (o *Permission) SetAllowNil(b bool)`

 SetAllowNil sets the value for Allow to be an explicit nil

### UnsetAllow
`func (o *Permission) UnsetAllow()`

UnsetAllow ensures that no value is present for Allow, not even an explicit nil
### GetDeny

`func (o *Permission) GetDeny() []string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *Permission) GetDenyOk() (*[]string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *Permission) SetDeny(v []string)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *Permission) HasDeny() bool`

HasDeny returns a boolean if a field has been set.

### SetDenyNil

`func (o *Permission) SetDenyNil(b bool)`

 SetDenyNil sets the value for Deny to be an explicit nil

### UnsetDeny
`func (o *Permission) UnsetDeny()`

UnsetDeny ensures that no value is present for Deny, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


