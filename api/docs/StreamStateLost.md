# StreamStateLost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** | How many bytes were lost | 
**Msgs** | **[]int32** | Message IDs of lost messages | 

## Methods

### NewStreamStateLost

`func NewStreamStateLost(bytes int32, msgs []int32, ) *StreamStateLost`

NewStreamStateLost instantiates a new StreamStateLost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamStateLostWithDefaults

`func NewStreamStateLostWithDefaults() *StreamStateLost`

NewStreamStateLostWithDefaults instantiates a new StreamStateLost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *StreamStateLost) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *StreamStateLost) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *StreamStateLost) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetMsgs

`func (o *StreamStateLost) GetMsgs() []int32`

GetMsgs returns the Msgs field if non-nil, zero value otherwise.

### GetMsgsOk

`func (o *StreamStateLost) GetMsgsOk() (*[]int32, bool)`

GetMsgsOk returns a tuple with the Msgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgs

`func (o *StreamStateLost) SetMsgs(v []int32)`

SetMsgs sets Msgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


