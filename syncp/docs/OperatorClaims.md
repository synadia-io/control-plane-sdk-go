# OperatorClaims

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
**Nats** | Pointer to [**Operator**](Operator.md) |  | [optional] 

## Methods

### NewOperatorClaims

`func NewOperatorClaims() *OperatorClaims`

NewOperatorClaims instantiates a new OperatorClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorClaimsWithDefaults

`func NewOperatorClaimsWithDefaults() *OperatorClaims`

NewOperatorClaimsWithDefaults instantiates a new OperatorClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *OperatorClaims) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *OperatorClaims) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *OperatorClaims) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *OperatorClaims) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetExp

`func (o *OperatorClaims) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *OperatorClaims) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *OperatorClaims) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *OperatorClaims) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIat

`func (o *OperatorClaims) GetIat() int64`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *OperatorClaims) GetIatOk() (*int64, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *OperatorClaims) SetIat(v int64)`

SetIat sets Iat field to given value.

### HasIat

`func (o *OperatorClaims) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *OperatorClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *OperatorClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *OperatorClaims) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *OperatorClaims) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *OperatorClaims) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *OperatorClaims) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *OperatorClaims) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *OperatorClaims) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetName

`func (o *OperatorClaims) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatorClaims) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatorClaims) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperatorClaims) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNbf

`func (o *OperatorClaims) GetNbf() int64`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *OperatorClaims) GetNbfOk() (*int64, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *OperatorClaims) SetNbf(v int64)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *OperatorClaims) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetSub

`func (o *OperatorClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *OperatorClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *OperatorClaims) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *OperatorClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetNats

`func (o *OperatorClaims) GetNats() Operator`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *OperatorClaims) GetNatsOk() (*Operator, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *OperatorClaims) SetNats(v Operator)`

SetNats sets Nats field to given value.

### HasNats

`func (o *OperatorClaims) HasNats() bool`

HasNats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


