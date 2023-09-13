# JSConsumerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerType** | Pointer to [**JSConsumerType**](JSConsumerType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
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

### NewJSConsumerInfoResponse

`func NewJSConsumerInfoResponse(ackFloor SequenceInfo, config ConsumerConfig, created time.Time, delivered SequenceInfo, name string, numAckPending int32, numPending int32, numRedelivered int32, numWaiting int32, streamName string, ) *JSConsumerInfoResponse`

NewJSConsumerInfoResponse instantiates a new JSConsumerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSConsumerInfoResponseWithDefaults

`func NewJSConsumerInfoResponseWithDefaults() *JSConsumerInfoResponse`

NewJSConsumerInfoResponseWithDefaults instantiates a new JSConsumerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerType

`func (o *JSConsumerInfoResponse) GetConsumerType() JSConsumerType`

GetConsumerType returns the ConsumerType field if non-nil, zero value otherwise.

### GetConsumerTypeOk

`func (o *JSConsumerInfoResponse) GetConsumerTypeOk() (*JSConsumerType, bool)`

GetConsumerTypeOk returns a tuple with the ConsumerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerType

`func (o *JSConsumerInfoResponse) SetConsumerType(v JSConsumerType)`

SetConsumerType sets ConsumerType field to given value.

### HasConsumerType

`func (o *JSConsumerInfoResponse) HasConsumerType() bool`

HasConsumerType returns a boolean if a field has been set.

### GetId

`func (o *JSConsumerInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSConsumerInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSConsumerInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JSConsumerInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAckFloor

`func (o *JSConsumerInfoResponse) GetAckFloor() SequenceInfo`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *JSConsumerInfoResponse) GetAckFloorOk() (*SequenceInfo, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *JSConsumerInfoResponse) SetAckFloor(v SequenceInfo)`

SetAckFloor sets AckFloor field to given value.


### GetCluster

`func (o *JSConsumerInfoResponse) GetCluster() ConsumerInfoCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSConsumerInfoResponse) GetClusterOk() (*ConsumerInfoCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSConsumerInfoResponse) SetCluster(v ConsumerInfoCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSConsumerInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *JSConsumerInfoResponse) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *JSConsumerInfoResponse) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetConfig

`func (o *JSConsumerInfoResponse) GetConfig() ConsumerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSConsumerInfoResponse) GetConfigOk() (*ConsumerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSConsumerInfoResponse) SetConfig(v ConsumerConfig)`

SetConfig sets Config field to given value.


### GetCreated

`func (o *JSConsumerInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSConsumerInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSConsumerInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDelivered

`func (o *JSConsumerInfoResponse) GetDelivered() SequenceInfo`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *JSConsumerInfoResponse) GetDeliveredOk() (*SequenceInfo, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *JSConsumerInfoResponse) SetDelivered(v SequenceInfo)`

SetDelivered sets Delivered field to given value.


### GetName

`func (o *JSConsumerInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSConsumerInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSConsumerInfoResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumAckPending

`func (o *JSConsumerInfoResponse) GetNumAckPending() int32`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *JSConsumerInfoResponse) GetNumAckPendingOk() (*int32, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *JSConsumerInfoResponse) SetNumAckPending(v int32)`

SetNumAckPending sets NumAckPending field to given value.


### GetNumPending

`func (o *JSConsumerInfoResponse) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *JSConsumerInfoResponse) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *JSConsumerInfoResponse) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.


### GetNumRedelivered

`func (o *JSConsumerInfoResponse) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *JSConsumerInfoResponse) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *JSConsumerInfoResponse) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.


### GetNumWaiting

`func (o *JSConsumerInfoResponse) GetNumWaiting() int32`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *JSConsumerInfoResponse) GetNumWaitingOk() (*int32, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *JSConsumerInfoResponse) SetNumWaiting(v int32)`

SetNumWaiting sets NumWaiting field to given value.


### GetPushBound

`func (o *JSConsumerInfoResponse) GetPushBound() bool`

GetPushBound returns the PushBound field if non-nil, zero value otherwise.

### GetPushBoundOk

`func (o *JSConsumerInfoResponse) GetPushBoundOk() (*bool, bool)`

GetPushBoundOk returns a tuple with the PushBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushBound

`func (o *JSConsumerInfoResponse) SetPushBound(v bool)`

SetPushBound sets PushBound field to given value.

### HasPushBound

`func (o *JSConsumerInfoResponse) HasPushBound() bool`

HasPushBound returns a boolean if a field has been set.

### GetStreamName

`func (o *JSConsumerInfoResponse) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *JSConsumerInfoResponse) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *JSConsumerInfoResponse) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


