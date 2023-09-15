# JetStreamAccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAckPending** | **int32** |  | 
**MaxBytesRequired** | **bool** |  | 
**MaxConsumers** | **int32** |  | 
**MaxMemory** | **int64** |  | 
**MaxStorage** | **int64** |  | 
**MaxStreams** | **int32** |  | 
**MemoryMaxStreamBytes** | **int64** |  | 
**StorageMaxStreamBytes** | **int64** |  | 

## Methods

### NewJetStreamAccountLimits

`func NewJetStreamAccountLimits(maxAckPending int32, maxBytesRequired bool, maxConsumers int32, maxMemory int64, maxStorage int64, maxStreams int32, memoryMaxStreamBytes int64, storageMaxStreamBytes int64, ) *JetStreamAccountLimits`

NewJetStreamAccountLimits instantiates a new JetStreamAccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamAccountLimitsWithDefaults

`func NewJetStreamAccountLimitsWithDefaults() *JetStreamAccountLimits`

NewJetStreamAccountLimitsWithDefaults instantiates a new JetStreamAccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAckPending

`func (o *JetStreamAccountLimits) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *JetStreamAccountLimits) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *JetStreamAccountLimits) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.


### GetMaxBytesRequired

`func (o *JetStreamAccountLimits) GetMaxBytesRequired() bool`

GetMaxBytesRequired returns the MaxBytesRequired field if non-nil, zero value otherwise.

### GetMaxBytesRequiredOk

`func (o *JetStreamAccountLimits) GetMaxBytesRequiredOk() (*bool, bool)`

GetMaxBytesRequiredOk returns a tuple with the MaxBytesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytesRequired

`func (o *JetStreamAccountLimits) SetMaxBytesRequired(v bool)`

SetMaxBytesRequired sets MaxBytesRequired field to given value.


### GetMaxConsumers

`func (o *JetStreamAccountLimits) GetMaxConsumers() int32`

GetMaxConsumers returns the MaxConsumers field if non-nil, zero value otherwise.

### GetMaxConsumersOk

`func (o *JetStreamAccountLimits) GetMaxConsumersOk() (*int32, bool)`

GetMaxConsumersOk returns a tuple with the MaxConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsumers

`func (o *JetStreamAccountLimits) SetMaxConsumers(v int32)`

SetMaxConsumers sets MaxConsumers field to given value.


### GetMaxMemory

`func (o *JetStreamAccountLimits) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *JetStreamAccountLimits) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *JetStreamAccountLimits) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxStorage

`func (o *JetStreamAccountLimits) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *JetStreamAccountLimits) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *JetStreamAccountLimits) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.


### GetMaxStreams

`func (o *JetStreamAccountLimits) GetMaxStreams() int32`

GetMaxStreams returns the MaxStreams field if non-nil, zero value otherwise.

### GetMaxStreamsOk

`func (o *JetStreamAccountLimits) GetMaxStreamsOk() (*int32, bool)`

GetMaxStreamsOk returns a tuple with the MaxStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreams

`func (o *JetStreamAccountLimits) SetMaxStreams(v int32)`

SetMaxStreams sets MaxStreams field to given value.


### GetMemoryMaxStreamBytes

`func (o *JetStreamAccountLimits) GetMemoryMaxStreamBytes() int64`

GetMemoryMaxStreamBytes returns the MemoryMaxStreamBytes field if non-nil, zero value otherwise.

### GetMemoryMaxStreamBytesOk

`func (o *JetStreamAccountLimits) GetMemoryMaxStreamBytesOk() (*int64, bool)`

GetMemoryMaxStreamBytesOk returns a tuple with the MemoryMaxStreamBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMaxStreamBytes

`func (o *JetStreamAccountLimits) SetMemoryMaxStreamBytes(v int64)`

SetMemoryMaxStreamBytes sets MemoryMaxStreamBytes field to given value.


### GetStorageMaxStreamBytes

`func (o *JetStreamAccountLimits) GetStorageMaxStreamBytes() int64`

GetStorageMaxStreamBytes returns the StorageMaxStreamBytes field if non-nil, zero value otherwise.

### GetStorageMaxStreamBytesOk

`func (o *JetStreamAccountLimits) GetStorageMaxStreamBytesOk() (*int64, bool)`

GetStorageMaxStreamBytesOk returns a tuple with the StorageMaxStreamBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMaxStreamBytes

`func (o *JetStreamAccountLimits) SetStorageMaxStreamBytes(v int64)`

SetStorageMaxStreamBytes sets StorageMaxStreamBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


