# JetStreamTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | **int32** |  | 
**Limits** | [**JetStreamAccountLimits**](JetStreamAccountLimits.md) |  | 
**Memory** | **int32** |  | 
**Storage** | **int32** |  | 
**Streams** | **int32** |  | 

## Methods

### NewJetStreamTier

`func NewJetStreamTier(consumers int32, limits JetStreamAccountLimits, memory int32, storage int32, streams int32, ) *JetStreamTier`

NewJetStreamTier instantiates a new JetStreamTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamTierWithDefaults

`func NewJetStreamTierWithDefaults() *JetStreamTier`

NewJetStreamTierWithDefaults instantiates a new JetStreamTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *JetStreamTier) GetConsumers() int32`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *JetStreamTier) GetConsumersOk() (*int32, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *JetStreamTier) SetConsumers(v int32)`

SetConsumers sets Consumers field to given value.


### GetLimits

`func (o *JetStreamTier) GetLimits() JetStreamAccountLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *JetStreamTier) GetLimitsOk() (*JetStreamAccountLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *JetStreamTier) SetLimits(v JetStreamAccountLimits)`

SetLimits sets Limits field to given value.


### GetMemory

`func (o *JetStreamTier) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JetStreamTier) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JetStreamTier) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetStorage

`func (o *JetStreamTier) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *JetStreamTier) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *JetStreamTier) SetStorage(v int32)`

SetStorage sets Storage field to given value.


### GetStreams

`func (o *JetStreamTier) GetStreams() int32`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *JetStreamTier) GetStreamsOk() (*int32, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *JetStreamTier) SetStreams(v int32)`

SetStreams sets Streams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


