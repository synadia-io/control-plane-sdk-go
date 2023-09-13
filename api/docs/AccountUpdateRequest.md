# AccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtSettings** | Pointer to [**AccountJWTSettings**](AccountJWTSettings.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserJwtExpiresInSecs** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountUpdateRequest

`func NewAccountUpdateRequest() *AccountUpdateRequest`

NewAccountUpdateRequest instantiates a new AccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUpdateRequestWithDefaults

`func NewAccountUpdateRequestWithDefaults() *AccountUpdateRequest`

NewAccountUpdateRequestWithDefaults instantiates a new AccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtSettings

`func (o *AccountUpdateRequest) GetJwtSettings() AccountJWTSettings`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *AccountUpdateRequest) GetJwtSettingsOk() (*AccountJWTSettings, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *AccountUpdateRequest) SetJwtSettings(v AccountJWTSettings)`

SetJwtSettings sets JwtSettings field to given value.

### HasJwtSettings

`func (o *AccountUpdateRequest) HasJwtSettings() bool`

HasJwtSettings returns a boolean if a field has been set.

### GetName

`func (o *AccountUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserJwtExpiresInSecs

`func (o *AccountUpdateRequest) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *AccountUpdateRequest) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *AccountUpdateRequest) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.

### HasUserJwtExpiresInSecs

`func (o *AccountUpdateRequest) HasUserJwtExpiresInSecs() bool`

HasUserJwtExpiresInSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


