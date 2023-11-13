# AccountClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **string** |  | [optional] 
**Exp** | Pointer to **int64** |  | [optional] 
**Iat** | Pointer to **int64** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Jti** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Nbf** | Pointer to **int64** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Nats** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewAccountClaims

`func NewAccountClaims() *AccountClaims`

NewAccountClaims instantiates a new AccountClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountClaimsWithDefaults

`func NewAccountClaimsWithDefaults() *AccountClaims`

NewAccountClaimsWithDefaults instantiates a new AccountClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *AccountClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *AccountClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *AccountClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *AccountClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetExp

`func (o *AccountClaims) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AccountClaims) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AccountClaims) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *AccountClaims) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIat

`func (o *AccountClaims) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *AccountClaims) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *AccountClaims) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *AccountClaims) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *AccountClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AccountClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AccountClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *AccountClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *AccountClaims) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *AccountClaims) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *AccountClaims) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *AccountClaims) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetName

`func (o *AccountClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNbf

`func (o *AccountClaims) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *AccountClaims) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *AccountClaims) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *AccountClaims) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetSub

`func (o *AccountClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AccountClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AccountClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *AccountClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetNats

`func (o *AccountClaims) GetNats() Account`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *AccountClaims) GetNatsOk() (*Account, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *AccountClaims) SetNats(v Account)`

SetNats sets Nats field to given value.

### HasNats

`func (o *AccountClaims) HasNats() bool`

HasNats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


