# NatsUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtExpiresInSecs** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Revoked** | Pointer to **bool** |  | [optional] 
**UserClaims** | Pointer to [**NullableNatsUserUpdateRequestUserClaims**](NatsUserUpdateRequestUserClaims.md) |  | [optional] 

## Methods

### NewNatsUserUpdateRequest

`func NewNatsUserUpdateRequest() *NatsUserUpdateRequest`

NewNatsUserUpdateRequest instantiates a new NatsUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatsUserUpdateRequestWithDefaults

`func NewNatsUserUpdateRequestWithDefaults() *NatsUserUpdateRequest`

NewNatsUserUpdateRequestWithDefaults instantiates a new NatsUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtExpiresInSecs

`func (o *NatsUserUpdateRequest) GetJwtExpiresInSecs() int64`

GetJwtExpiresInSecs returns the JwtExpiresInSecs field if non-nil, zero value otherwise.

### GetJwtExpiresInSecsOk

`func (o *NatsUserUpdateRequest) GetJwtExpiresInSecsOk() (*int64, bool)`

GetJwtExpiresInSecsOk returns a tuple with the JwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresInSecs

`func (o *NatsUserUpdateRequest) SetJwtExpiresInSecs(v int64)`

SetJwtExpiresInSecs sets JwtExpiresInSecs field to given value.

### HasJwtExpiresInSecs

`func (o *NatsUserUpdateRequest) HasJwtExpiresInSecs() bool`

HasJwtExpiresInSecs returns a boolean if a field has been set.

### GetName

`func (o *NatsUserUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatsUserUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatsUserUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NatsUserUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevoked

`func (o *NatsUserUpdateRequest) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *NatsUserUpdateRequest) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *NatsUserUpdateRequest) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *NatsUserUpdateRequest) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetUserClaims

`func (o *NatsUserUpdateRequest) GetUserClaims() NatsUserUpdateRequestUserClaims`

GetUserClaims returns the UserClaims field if non-nil, zero value otherwise.

### GetUserClaimsOk

`func (o *NatsUserUpdateRequest) GetUserClaimsOk() (*NatsUserUpdateRequestUserClaims, bool)`

GetUserClaimsOk returns a tuple with the UserClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaims

`func (o *NatsUserUpdateRequest) SetUserClaims(v NatsUserUpdateRequestUserClaims)`

SetUserClaims sets UserClaims field to given value.

### HasUserClaims

`func (o *NatsUserUpdateRequest) HasUserClaims() bool`

HasUserClaims returns a boolean if a field has been set.

### SetUserClaimsNil

`func (o *NatsUserUpdateRequest) SetUserClaimsNil(b bool)`

 SetUserClaimsNil sets the value for UserClaims to be an explicit nil

### UnsetUserClaims
`func (o *NatsUserUpdateRequest) UnsetUserClaims()`

UnsetUserClaims ensures that no value is present for UserClaims, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


