# AccountCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtSettings** | [**NullableAccountCreateRequestJwtSettings**](AccountCreateRequestJwtSettings.md) |  | 
**Name** | **string** |  | 
**UserJwtExpiresInSecs** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountCreateRequest

`func NewAccountCreateRequest(jwtSettings NullableAccountCreateRequestJwtSettings, name string, ) *AccountCreateRequest`

NewAccountCreateRequest instantiates a new AccountCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateRequestWithDefaults

`func NewAccountCreateRequestWithDefaults() *AccountCreateRequest`

NewAccountCreateRequestWithDefaults instantiates a new AccountCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtSettings

`func (o *AccountCreateRequest) GetJwtSettings() AccountCreateRequestJwtSettings`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *AccountCreateRequest) GetJwtSettingsOk() (*AccountCreateRequestJwtSettings, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *AccountCreateRequest) SetJwtSettings(v AccountCreateRequestJwtSettings)`

SetJwtSettings sets JwtSettings field to given value.


### SetJwtSettingsNil

`func (o *AccountCreateRequest) SetJwtSettingsNil(b bool)`

 SetJwtSettingsNil sets the value for JwtSettings to be an explicit nil

### UnsetJwtSettings
`func (o *AccountCreateRequest) UnsetJwtSettings()`

UnsetJwtSettings ensures that no value is present for JwtSettings, not even an explicit nil
### GetName

`func (o *AccountCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserJwtExpiresInSecs

`func (o *AccountCreateRequest) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *AccountCreateRequest) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *AccountCreateRequest) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.

### HasUserJwtExpiresInSecs

`func (o *AccountCreateRequest) HasUserJwtExpiresInSecs() bool`

HasUserJwtExpiresInSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


