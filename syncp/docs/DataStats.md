# DataStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int64** |  | 
**Msgs** | **int64** |  | 

## Methods

### NewDataStats

`func NewDataStats(bytes int64, msgs int64, ) *DataStats`

NewDataStats instantiates a new DataStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStatsWithDefaults

`func NewDataStatsWithDefaults() *DataStats`

NewDataStatsWithDefaults instantiates a new DataStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *DataStats) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *DataStats) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *DataStats) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### GetMsgs

`func (o *DataStats) GetMsgs() int64`

GetMsgs returns the Msgs field if non-nil, zero value otherwise.

### GetMsgsOk

`func (o *DataStats) GetMsgsOk() (*int64, bool)`

GetMsgsOk returns a tuple with the Msgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgs

`func (o *DataStats) SetMsgs(v int64)`

SetMsgs sets Msgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


