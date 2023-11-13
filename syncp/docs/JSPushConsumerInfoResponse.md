# JSPushConsumerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Config** | Pointer to [**JSPushConsumerConfigRequest**](JSPushConsumerConfigRequest.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NumAckPending** | **int32** |  | 
**NumPending** | **int32** |  | 
**NumRedelivered** | **int32** |  | 
**PushBound** | Pointer to **bool** |  | [optional] 
**StreamName** | **string** |  | 

## Methods

### NewJSPushConsumerInfoResponse

`func NewJSPushConsumerInfoResponse(ackFloor SequenceInfo, created time.Time, delivered SequenceInfo, id string, name string, numAckPending int32, numPending int32, numRedelivered int32, streamName string, ) *JSPushConsumerInfoResponse`

NewJSPushConsumerInfoResponse instantiates a new JSPushConsumerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSPushConsumerInfoResponseWithDefaults

`func NewJSPushConsumerInfoResponseWithDefaults() *JSPushConsumerInfoResponse`

NewJSPushConsumerInfoResponseWithDefaults instantiates a new JSPushConsumerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *JSPushConsumerInfoResponse) GetAckFloor() SequenceInfo`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *JSPushConsumerInfoResponse) GetAckFloorOk() (*SequenceInfo, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *JSPushConsumerInfoResponse) SetAckFloor(v SequenceInfo)`

SetAckFloor sets AckFloor field to given value.


### GetCluster

`func (o *JSPushConsumerInfoResponse) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSPushConsumerInfoResponse) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSPushConsumerInfoResponse) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSPushConsumerInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfig

`func (o *JSPushConsumerInfoResponse) GetConfig() JSPushConsumerConfigRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSPushConsumerInfoResponse) GetConfigOk() (*JSPushConsumerConfigRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSPushConsumerInfoResponse) SetConfig(v JSPushConsumerConfigRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JSPushConsumerInfoResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *JSPushConsumerInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSPushConsumerInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSPushConsumerInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDelivered

`func (o *JSPushConsumerInfoResponse) GetDelivered() SequenceInfo`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *JSPushConsumerInfoResponse) GetDeliveredOk() (*SequenceInfo, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *JSPushConsumerInfoResponse) SetDelivered(v SequenceInfo)`

SetDelivered sets Delivered field to given value.


### GetId

`func (o *JSPushConsumerInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSPushConsumerInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSPushConsumerInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JSPushConsumerInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSPushConsumerInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSPushConsumerInfoResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumAckPending

`func (o *JSPushConsumerInfoResponse) GetNumAckPending() int32`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *JSPushConsumerInfoResponse) GetNumAckPendingOk() (*int32, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *JSPushConsumerInfoResponse) SetNumAckPending(v int32)`

SetNumAckPending sets NumAckPending field to given value.


### GetNumPending

`func (o *JSPushConsumerInfoResponse) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *JSPushConsumerInfoResponse) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *JSPushConsumerInfoResponse) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.


### GetNumRedelivered

`func (o *JSPushConsumerInfoResponse) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *JSPushConsumerInfoResponse) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *JSPushConsumerInfoResponse) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.


### GetPushBound

`func (o *JSPushConsumerInfoResponse) GetPushBound() bool`

GetPushBound returns the PushBound field if non-nil, zero value otherwise.

### GetPushBoundOk

`func (o *JSPushConsumerInfoResponse) GetPushBoundOk() (*bool, bool)`

GetPushBoundOk returns a tuple with the PushBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushBound

`func (o *JSPushConsumerInfoResponse) SetPushBound(v bool)`

SetPushBound sets PushBound field to given value.

### HasPushBound

`func (o *JSPushConsumerInfoResponse) HasPushBound() bool`

HasPushBound returns a boolean if a field has been set.

### GetStreamName

`func (o *JSPushConsumerInfoResponse) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *JSPushConsumerInfoResponse) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *JSPushConsumerInfoResponse) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


