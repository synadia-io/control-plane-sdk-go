# OperatorClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimsData** | [**ClaimsData**](ClaimsData.md) |  | 
**Nats** | Pointer to [**Operator**](Operator.md) |  | [optional] 

## Methods

### NewOperatorClaims

`func NewOperatorClaims(claimsData ClaimsData, ) *OperatorClaims`

NewOperatorClaims instantiates a new OperatorClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorClaimsWithDefaults

`func NewOperatorClaimsWithDefaults() *OperatorClaims`

NewOperatorClaimsWithDefaults instantiates a new OperatorClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimsData

`func (o *OperatorClaims) GetClaimsData() ClaimsData`

GetClaimsData returns the ClaimsData field if non-nil, zero value otherwise.

### GetClaimsDataOk

`func (o *OperatorClaims) GetClaimsDataOk() (*ClaimsData, bool)`

GetClaimsDataOk returns a tuple with the ClaimsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsData

`func (o *OperatorClaims) SetClaimsData(v ClaimsData)`

SetClaimsData sets ClaimsData field to given value.


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


