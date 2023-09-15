# UserClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nats** | Pointer to [**User**](User.md) |  | [optional] 
**Aud** | Pointer to **string** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**Iat** | Pointer to **int64** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Jti** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 

## Methods

### NewUserClaims

`func NewUserClaims() *UserClaims`

NewUserClaims instantiates a new UserClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserClaimsWithDefaults

`func NewUserClaimsWithDefaults() *UserClaims`

NewUserClaimsWithDefaults instantiates a new UserClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNats

`func (o *UserClaims) GetNats() User`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *UserClaims) GetNatsOk() (*User, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *UserClaims) SetNats(v User)`

SetNats sets Nats field to given value.

### HasNats

`func (o *UserClaims) HasNats() bool`

HasNats returns a boolean if a field has been set.

### GetAud

`func (o *UserClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *UserClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *UserClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *UserClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetExp

`func (o *UserClaims) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *UserClaims) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *UserClaims) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *UserClaims) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIat

`func (o *UserClaims) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *UserClaims) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *UserClaims) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *UserClaims) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *UserClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *UserClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *UserClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *UserClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *UserClaims) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *UserClaims) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *UserClaims) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *UserClaims) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetName

`func (o *UserClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNbf

`func (o *UserClaims) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *UserClaims) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *UserClaims) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *UserClaims) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetSub

`func (o *UserClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UserClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UserClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *UserClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


