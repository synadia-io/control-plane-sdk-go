# LostStreamData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** | How many bytes were lost | 
**Msgs** | **[]int32** | Message IDs of lost messages | 

## Methods

### NewLostStreamData

`func NewLostStreamData(bytes int32, msgs []int32, ) *LostStreamData`

NewLostStreamData instantiates a new LostStreamData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLostStreamDataWithDefaults

`func NewLostStreamDataWithDefaults() *LostStreamData`

NewLostStreamDataWithDefaults instantiates a new LostStreamData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *LostStreamData) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *LostStreamData) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *LostStreamData) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetMsgs

`func (o *LostStreamData) GetMsgs() []int32`

GetMsgs returns the Msgs field if non-nil, zero value otherwise.

### GetMsgsOk

`func (o *LostStreamData) GetMsgsOk() (*[]int32, bool)`

GetMsgsOk returns a tuple with the Msgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgs

`func (o *LostStreamData) SetMsgs(v []int32)`

SetMsgs sets Msgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


