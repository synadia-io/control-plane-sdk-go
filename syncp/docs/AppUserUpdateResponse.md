# AppUserUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResetPasswordLink** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**Identifier** | **NullableString** |  | 
**Name** | **string** |  | 
**RoleId** | **string** |  | 
**RoleName** | **string** |  | 
**Type** | [**AppUserType**](AppUserType.md) |  | 

## Methods

### NewAppUserUpdateResponse

`func NewAppUserUpdateResponse(created time.Time, id string, identifier NullableString, name string, roleId string, roleName string, type_ AppUserType, ) *AppUserUpdateResponse`

NewAppUserUpdateResponse instantiates a new AppUserUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserUpdateResponseWithDefaults

`func NewAppUserUpdateResponseWithDefaults() *AppUserUpdateResponse`

NewAppUserUpdateResponseWithDefaults instantiates a new AppUserUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResetPasswordLink

`func (o *AppUserUpdateResponse) GetResetPasswordLink() string`

GetResetPasswordLink returns the ResetPasswordLink field if non-nil, zero value otherwise.

### GetResetPasswordLinkOk

`func (o *AppUserUpdateResponse) GetResetPasswordLinkOk() (*string, bool)`

GetResetPasswordLinkOk returns a tuple with the ResetPasswordLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordLink

`func (o *AppUserUpdateResponse) SetResetPasswordLink(v string)`

SetResetPasswordLink sets ResetPasswordLink field to given value.

### HasResetPasswordLink

`func (o *AppUserUpdateResponse) HasResetPasswordLink() bool`

HasResetPasswordLink returns a boolean if a field has been set.

### GetCreated

`func (o *AppUserUpdateResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUserUpdateResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUserUpdateResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AppUserUpdateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUserUpdateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUserUpdateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *AppUserUpdateResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AppUserUpdateResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AppUserUpdateResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### SetIdentifierNil

`func (o *AppUserUpdateResponse) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *AppUserUpdateResponse) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *AppUserUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserUpdateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRoleId

`func (o *AppUserUpdateResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AppUserUpdateResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AppUserUpdateResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetRoleName

`func (o *AppUserUpdateResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AppUserUpdateResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AppUserUpdateResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetType

`func (o *AppUserUpdateResponse) GetType() AppUserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppUserUpdateResponse) GetTypeOk() (*AppUserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppUserUpdateResponse) SetType(v AppUserType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


