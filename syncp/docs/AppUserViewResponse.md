# AppUserViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**Identifier** | **NullableString** |  | 
**Name** | **string** |  | 
**OryId** | **NullableString** |  | 
**RoleId** | **string** |  | 
**RoleName** | **string** |  | 
**TermsAcceptedUpdated** | **NullableString** |  | 
**Type** | [**AppUserType**](AppUserType.md) |  | 

## Methods

### NewAppUserViewResponse

`func NewAppUserViewResponse(created time.Time, id string, identifier NullableString, name string, oryId NullableString, roleId string, roleName string, termsAcceptedUpdated NullableString, type_ AppUserType, ) *AppUserViewResponse`

NewAppUserViewResponse instantiates a new AppUserViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserViewResponseWithDefaults

`func NewAppUserViewResponseWithDefaults() *AppUserViewResponse`

NewAppUserViewResponseWithDefaults instantiates a new AppUserViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppUserViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUserViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUserViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AppUserViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUserViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUserViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *AppUserViewResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AppUserViewResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AppUserViewResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### SetIdentifierNil

`func (o *AppUserViewResponse) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *AppUserViewResponse) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetName

`func (o *AppUserViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOryId

`func (o *AppUserViewResponse) GetOryId() string`

GetOryId returns the OryId field if non-nil, zero value otherwise.

### GetOryIdOk

`func (o *AppUserViewResponse) GetOryIdOk() (*string, bool)`

GetOryIdOk returns a tuple with the OryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOryId

`func (o *AppUserViewResponse) SetOryId(v string)`

SetOryId sets OryId field to given value.


### SetOryIdNil

`func (o *AppUserViewResponse) SetOryIdNil(b bool)`

 SetOryIdNil sets the value for OryId to be an explicit nil

### UnsetOryId
`func (o *AppUserViewResponse) UnsetOryId()`

UnsetOryId ensures that no value is present for OryId, not even an explicit nil
### GetRoleId

`func (o *AppUserViewResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AppUserViewResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AppUserViewResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetRoleName

`func (o *AppUserViewResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AppUserViewResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AppUserViewResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetTermsAcceptedUpdated

`func (o *AppUserViewResponse) GetTermsAcceptedUpdated() string`

GetTermsAcceptedUpdated returns the TermsAcceptedUpdated field if non-nil, zero value otherwise.

### GetTermsAcceptedUpdatedOk

`func (o *AppUserViewResponse) GetTermsAcceptedUpdatedOk() (*string, bool)`

GetTermsAcceptedUpdatedOk returns a tuple with the TermsAcceptedUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAcceptedUpdated

`func (o *AppUserViewResponse) SetTermsAcceptedUpdated(v string)`

SetTermsAcceptedUpdated sets TermsAcceptedUpdated field to given value.


### SetTermsAcceptedUpdatedNil

`func (o *AppUserViewResponse) SetTermsAcceptedUpdatedNil(b bool)`

 SetTermsAcceptedUpdatedNil sets the value for TermsAcceptedUpdated to be an explicit nil

### UnsetTermsAcceptedUpdated
`func (o *AppUserViewResponse) UnsetTermsAcceptedUpdated()`

UnsetTermsAcceptedUpdated ensures that no value is present for TermsAcceptedUpdated, not even an explicit nil
### GetType

`func (o *AppUserViewResponse) GetType() AppUserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppUserViewResponse) GetTypeOk() (*AppUserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppUserViewResponse) SetType(v AppUserType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


