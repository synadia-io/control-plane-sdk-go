# JetStreamAPIStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewJetStreamAPIStats

`func NewJetStreamAPIStats(errors int32, total int32, ) *JetStreamAPIStats`

NewJetStreamAPIStats instantiates a new JetStreamAPIStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamAPIStatsWithDefaults

`func NewJetStreamAPIStatsWithDefaults() *JetStreamAPIStats`

NewJetStreamAPIStatsWithDefaults instantiates a new JetStreamAPIStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *JetStreamAPIStats) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *JetStreamAPIStats) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *JetStreamAPIStats) SetErrors(v int32)`

SetErrors sets Errors field to given value.


### GetTotal

`func (o *JetStreamAPIStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *JetStreamAPIStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *JetStreamAPIStats) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


