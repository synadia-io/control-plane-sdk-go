# ActivationClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimsData** | [**ClaimsData**](ClaimsData.md) |  | 
**Nats** | Pointer to [**Activation**](Activation.md) |  | [optional] 

## Methods

### NewActivationClaims

`func NewActivationClaims(claimsData ClaimsData, ) *ActivationClaims`

NewActivationClaims instantiates a new ActivationClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationClaimsWithDefaults

`func NewActivationClaimsWithDefaults() *ActivationClaims`

NewActivationClaimsWithDefaults instantiates a new ActivationClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimsData

`func (o *ActivationClaims) GetClaimsData() ClaimsData`

GetClaimsData returns the ClaimsData field if non-nil, zero value otherwise.

### GetClaimsDataOk

`func (o *ActivationClaims) GetClaimsDataOk() (*ClaimsData, bool)`

GetClaimsDataOk returns a tuple with the ClaimsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsData

`func (o *ActivationClaims) SetClaimsData(v ClaimsData)`

SetClaimsData sets ClaimsData field to given value.


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


