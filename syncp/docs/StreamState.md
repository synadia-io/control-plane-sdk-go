# StreamState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** |  | 
**ConsumerCount** | **int32** |  | 
**Deleted** | Pointer to **[]int32** |  | [optional] 
**FirstSeq** | **int32** |  | 
**FirstTs** | **time.Time** |  | 
**LastSeq** | **int32** |  | 
**LastTs** | **time.Time** |  | 
**Lost** | Pointer to [**NullableStreamStateLost**](StreamStateLost.md) |  | [optional] 
**Messages** | **int32** |  | 
**NumDeleted** | Pointer to **int32** |  | [optional] 
**NumSubjects** | Pointer to **int32** |  | [optional] 
**Subjects** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewStreamState

`func NewStreamState(bytes int32, consumerCount int32, firstSeq int32, firstTs time.Time, lastSeq int32, lastTs time.Time, messages int32, ) *StreamState`

NewStreamState instantiates a new StreamState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamStateWithDefaults

`func NewStreamStateWithDefaults() *StreamState`

NewStreamStateWithDefaults instantiates a new StreamState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *StreamState) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *StreamState) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *StreamState) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetConsumerCount

`func (o *StreamState) GetConsumerCount() int32`

GetConsumerCount returns the ConsumerCount field if non-nil, zero value otherwise.

### GetConsumerCountOk

`func (o *StreamState) GetConsumerCountOk() (*int32, bool)`

GetConsumerCountOk returns a tuple with the ConsumerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerCount

`func (o *StreamState) SetConsumerCount(v int32)`

SetConsumerCount sets ConsumerCount field to given value.


### GetDeleted

`func (o *StreamState) GetDeleted() []int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *StreamState) GetDeletedOk() (*[]int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *StreamState) SetDeleted(v []int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *StreamState) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetFirstSeq

`func (o *StreamState) GetFirstSeq() int32`

GetFirstSeq returns the FirstSeq field if non-nil, zero value otherwise.

### GetFirstSeqOk

`func (o *StreamState) GetFirstSeqOk() (*int32, bool)`

GetFirstSeqOk returns a tuple with the FirstSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeq

`func (o *StreamState) SetFirstSeq(v int32)`

SetFirstSeq sets FirstSeq field to given value.


### GetFirstTs

`func (o *StreamState) GetFirstTs() time.Time`

GetFirstTs returns the FirstTs field if non-nil, zero value otherwise.

### GetFirstTsOk

`func (o *StreamState) GetFirstTsOk() (*time.Time, bool)`

GetFirstTsOk returns a tuple with the FirstTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTs

`func (o *StreamState) SetFirstTs(v time.Time)`

SetFirstTs sets FirstTs field to given value.


### GetLastSeq

`func (o *StreamState) GetLastSeq() int32`

GetLastSeq returns the LastSeq field if non-nil, zero value otherwise.

### GetLastSeqOk

`func (o *StreamState) GetLastSeqOk() (*int32, bool)`

GetLastSeqOk returns a tuple with the LastSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeq

`func (o *StreamState) SetLastSeq(v int32)`

SetLastSeq sets LastSeq field to given value.


### GetLastTs

`func (o *StreamState) GetLastTs() time.Time`

GetLastTs returns the LastTs field if non-nil, zero value otherwise.

### GetLastTsOk

`func (o *StreamState) GetLastTsOk() (*time.Time, bool)`

GetLastTsOk returns a tuple with the LastTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTs

`func (o *StreamState) SetLastTs(v time.Time)`

SetLastTs sets LastTs field to given value.


### GetLost

`func (o *StreamState) GetLost() StreamStateLost`

GetLost returns the Lost field if non-nil, zero value otherwise.

### GetLostOk

`func (o *StreamState) GetLostOk() (*StreamStateLost, bool)`

GetLostOk returns a tuple with the Lost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLost

`func (o *StreamState) SetLost(v StreamStateLost)`

SetLost sets Lost field to given value.

### HasLost

`func (o *StreamState) HasLost() bool`

HasLost returns a boolean if a field has been set.

### SetLostNil

`func (o *StreamState) SetLostNil(b bool)`

 SetLostNil sets the value for Lost to be an explicit nil

### UnsetLost
`func (o *StreamState) UnsetLost()`

UnsetLost ensures that no value is present for Lost, not even an explicit nil
### GetMessages

`func (o *StreamState) GetMessages() int32`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *StreamState) GetMessagesOk() (*int32, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *StreamState) SetMessages(v int32)`

SetMessages sets Messages field to given value.


### GetNumDeleted

`func (o *StreamState) GetNumDeleted() int32`

GetNumDeleted returns the NumDeleted field if non-nil, zero value otherwise.

### GetNumDeletedOk

`func (o *StreamState) GetNumDeletedOk() (*int32, bool)`

GetNumDeletedOk returns a tuple with the NumDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleted

`func (o *StreamState) SetNumDeleted(v int32)`

SetNumDeleted sets NumDeleted field to given value.

### HasNumDeleted

`func (o *StreamState) HasNumDeleted() bool`

HasNumDeleted returns a boolean if a field has been set.

### GetNumSubjects

`func (o *StreamState) GetNumSubjects() int32`

GetNumSubjects returns the NumSubjects field if non-nil, zero value otherwise.

### GetNumSubjectsOk

`func (o *StreamState) GetNumSubjectsOk() (*int32, bool)`

GetNumSubjectsOk returns a tuple with the NumSubjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSubjects

`func (o *StreamState) SetNumSubjects(v int32)`

SetNumSubjects sets NumSubjects field to given value.

### HasNumSubjects

`func (o *StreamState) HasNumSubjects() bool`

HasNumSubjects returns a boolean if a field has been set.

### GetSubjects

`func (o *StreamState) GetSubjects() map[string]int32`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *StreamState) GetSubjectsOk() (*map[string]int32, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *StreamState) SetSubjects(v map[string]int32)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *StreamState) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


