# ActivationClaims

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
**Nats** | Pointer to [**Activation**](Activation.md) |  | [optional] 

## Methods

### NewActivationClaims

`func NewActivationClaims() *ActivationClaims`

NewActivationClaims instantiates a new ActivationClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationClaimsWithDefaults

`func NewActivationClaimsWithDefaults() *ActivationClaims`

NewActivationClaimsWithDefaults instantiates a new ActivationClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *ActivationClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *ActivationClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *ActivationClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *ActivationClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetExp

`func (o *ActivationClaims) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *ActivationClaims) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *ActivationClaims) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *ActivationClaims) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIat

`func (o *ActivationClaims) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *ActivationClaims) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *ActivationClaims) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *ActivationClaims) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *ActivationClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *ActivationClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *ActivationClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *ActivationClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *ActivationClaims) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *ActivationClaims) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *ActivationClaims) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *ActivationClaims) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetName

`func (o *ActivationClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivationClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivationClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivationClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNbf

`func (o *ActivationClaims) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *ActivationClaims) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *ActivationClaims) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *ActivationClaims) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetSub

`func (o *ActivationClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *ActivationClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *ActivationClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *ActivationClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetNats

`func (o *ActivationClaims) GetNats() Activation`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *ActivationClaims) GetNatsOk() (*Activation, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *ActivationClaims) SetNats(v Activation)`

SetNats sets Nats field to given value.

### HasNats

`func (o *ActivationClaims) HasNats() bool`

HasNats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


