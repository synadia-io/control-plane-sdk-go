# ConsumerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to [**NullableConsumerInfoCluster**](ConsumerInfoCluster.md) |  | [optional] 
**Config** | [**ConsumerConfig**](ConsumerConfig.md) |  | 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Name** | **string** |  | 
**NumAckPending** | **int32** |  | 
**NumPending** | **int32** |  | 
**NumRedelivered** | **int32** |  | 
**NumWaiting** | **int32** |  | 
**PushBound** | Pointer to **bool** |  | [optional] 
**StreamName** | **string** |  | 

## Methods

### NewConsumerInfo

`func NewConsumerInfo(ackFloor SequenceInfo, config ConsumerConfig, created time.Time, delivered SequenceInfo, name string, numAckPending int32, numPending int32, numRedelivered int32, numWaiting int32, streamName string, ) *ConsumerInfo`

NewConsumerInfo instantiates a new ConsumerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerInfoWithDefaults

`func NewConsumerInfoWithDefaults() *ConsumerInfo`

NewConsumerInfoWithDefaults instantiates a new ConsumerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *ConsumerInfo) GetAckFloor() SequenceInfo`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *ConsumerInfo) GetAckFloorOk() (*SequenceInfo, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *ConsumerInfo) SetAckFloor(v SequenceInfo)`

SetAckFloor sets AckFloor field to given value.


### GetCluster

`func (o *ConsumerInfo) GetCluster() ConsumerInfoCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ConsumerInfo) GetClusterOk() (*ConsumerInfoCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ConsumerInfo) SetCluster(v ConsumerInfoCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ConsumerInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *ConsumerInfo) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *ConsumerInfo) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetConfig

`func (o *ConsumerInfo) GetConfig() ConsumerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConsumerInfo) GetConfigOk() (*ConsumerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConsumerInfo) SetConfig(v ConsumerConfig)`

SetConfig sets Config field to given value.


### GetCreated

`func (o *ConsumerInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsumerInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsumerInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDelivered

`func (o *ConsumerInfo) GetDelivered() SequenceInfo`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *ConsumerInfo) GetDeliveredOk() (*SequenceInfo, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *ConsumerInfo) SetDelivered(v SequenceInfo)`

SetDelivered sets Delivered field to given value.


### GetName

`func (o *ConsumerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNumAckPending

`func (o *ConsumerInfo) GetNumAckPending() int32`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *ConsumerInfo) GetNumAckPendingOk() (*int32, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *ConsumerInfo) SetNumAckPending(v int32)`

SetNumAckPending sets NumAckPending field to given value.


### GetNumPending

`func (o *ConsumerInfo) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *ConsumerInfo) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *ConsumerInfo) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.


### GetNumRedelivered

`func (o *ConsumerInfo) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *ConsumerInfo) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *ConsumerInfo) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.


### GetNumWaiting

`func (o *ConsumerInfo) GetNumWaiting() int32`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *ConsumerInfo) GetNumWaitingOk() (*int32, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *ConsumerInfo) SetNumWaiting(v int32)`

SetNumWaiting sets NumWaiting field to given value.


### GetPushBound

`func (o *ConsumerInfo) GetPushBound() bool`

GetPushBound returns the PushBound field if non-nil, zero value otherwise.

### GetPushBoundOk

`func (o *ConsumerInfo) GetPushBoundOk() (*bool, bool)`

GetPushBoundOk returns a tuple with the PushBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushBound

`func (o *ConsumerInfo) SetPushBound(v bool)`

SetPushBound sets PushBound field to given value.

### HasPushBound

`func (o *ConsumerInfo) HasPushBound() bool`

HasPushBound returns a boolean if a field has been set.

### GetStreamName

`func (o *ConsumerInfo) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *ConsumerInfo) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *ConsumerInfo) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


