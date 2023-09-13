# SequenceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerSeq** | **int32** |  | 
**LastActive** | Pointer to **NullableString** |  | [optional] 
**StreamSeq** | **int32** |  | 

## Methods

### NewSequenceInfo

`func NewSequenceInfo(consumerSeq int32, streamSeq int32, ) *SequenceInfo`

NewSequenceInfo instantiates a new SequenceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSequenceInfoWithDefaults

`func NewSequenceInfoWithDefaults() *SequenceInfo`

NewSequenceInfoWithDefaults instantiates a new SequenceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerSeq

`func (o *SequenceInfo) GetConsumerSeq() int32`

GetConsumerSeq returns the ConsumerSeq field if non-nil, zero value otherwise.

### GetConsumerSeqOk

`func (o *SequenceInfo) GetConsumerSeqOk() (*int32, bool)`

GetConsumerSeqOk returns a tuple with the ConsumerSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSeq

`func (o *SequenceInfo) SetConsumerSeq(v int32)`

SetConsumerSeq sets ConsumerSeq field to given value.


### GetLastActive

`func (o *SequenceInfo) GetLastActive() string`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *SequenceInfo) GetLastActiveOk() (*string, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *SequenceInfo) SetLastActive(v string)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *SequenceInfo) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### SetLastActiveNil

`func (o *SequenceInfo) SetLastActiveNil(b bool)`

 SetLastActiveNil sets the value for LastActive to be an explicit nil

### UnsetLastActive
`func (o *SequenceInfo) UnsetLastActive()`

UnsetLastActive ensures that no value is present for LastActive, not even an explicit nil
### GetStreamSeq

`func (o *SequenceInfo) GetStreamSeq() int32`

GetStreamSeq returns the StreamSeq field if non-nil, zero value otherwise.

### GetStreamSeqOk

`func (o *SequenceInfo) GetStreamSeqOk() (*int32, bool)`

GetStreamSeqOk returns a tuple with the StreamSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSeq

`func (o *SequenceInfo) SetStreamSeq(v int32)`

SetStreamSeq sets StreamSeq field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


