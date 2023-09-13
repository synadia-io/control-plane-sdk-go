# AppUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ResetPassword** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewAppUserUpdateRequest

`func NewAppUserUpdateRequest() *AppUserUpdateRequest`

NewAppUserUpdateRequest instantiates a new AppUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserUpdateRequestWithDefaults

`func NewAppUserUpdateRequestWithDefaults() *AppUserUpdateRequest`

NewAppUserUpdateRequestWithDefaults instantiates a new AppUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppUserUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppUserUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResetPassword

`func (o *AppUserUpdateRequest) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *AppUserUpdateRequest) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *AppUserUpdateRequest) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *AppUserUpdateRequest) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.

### GetRole

`func (o *AppUserUpdateRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AppUserUpdateRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AppUserUpdateRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AppUserUpdateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


