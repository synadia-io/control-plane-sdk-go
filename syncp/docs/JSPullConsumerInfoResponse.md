# JSPullConsumerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Config** | Pointer to [**JSPullConsumerConfigRequest**](JSPullConsumerConfigRequest.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NumAckPending** | **int32** |  | 
**NumPending** | **int32** |  | 
**NumRedelivered** | **int32** |  | 
**NumWaiting** | **int32** |  | 
**StreamName** | **string** |  | 

## Methods

### NewJSPullConsumerInfoResponse

`func NewJSPullConsumerInfoResponse(ackFloor SequenceInfo, created time.Time, delivered SequenceInfo, id string, name string, numAckPending int32, numPending int32, numRedelivered int32, numWaiting int32, streamName string, ) *JSPullConsumerInfoResponse`

NewJSPullConsumerInfoResponse instantiates a new JSPullConsumerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSPullConsumerInfoResponseWithDefaults

`func NewJSPullConsumerInfoResponseWithDefaults() *JSPullConsumerInfoResponse`

NewJSPullConsumerInfoResponseWithDefaults instantiates a new JSPullConsumerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckFloor

`func (o *JSPullConsumerInfoResponse) GetAckFloor() SequenceInfo`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *JSPullConsumerInfoResponse) GetAckFloorOk() (*SequenceInfo, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *JSPullConsumerInfoResponse) SetAckFloor(v SequenceInfo)`

SetAckFloor sets AckFloor field to given value.


### GetCluster

`func (o *JSPullConsumerInfoResponse) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSPullConsumerInfoResponse) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSPullConsumerInfoResponse) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSPullConsumerInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetConfig

`func (o *JSPullConsumerInfoResponse) GetConfig() JSPullConsumerConfigRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSPullConsumerInfoResponse) GetConfigOk() (*JSPullConsumerConfigRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSPullConsumerInfoResponse) SetConfig(v JSPullConsumerConfigRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JSPullConsumerInfoResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *JSPullConsumerInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSPullConsumerInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSPullConsumerInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDelivered

`func (o *JSPullConsumerInfoResponse) GetDelivered() SequenceInfo`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *JSPullConsumerInfoResponse) GetDeliveredOk() (*SequenceInfo, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *JSPullConsumerInfoResponse) SetDelivered(v SequenceInfo)`

SetDelivered sets Delivered field to given value.


### GetId

`func (o *JSPullConsumerInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSPullConsumerInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSPullConsumerInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JSPullConsumerInfoResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSPullConsumerInfoResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSPullConsumerInfoResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNumAckPending

`func (o *JSPullConsumerInfoResponse) GetNumAckPending() int32`

GetNumAckPending returns the NumAckPending field if non-nil, zero value otherwise.

### GetNumAckPendingOk

`func (o *JSPullConsumerInfoResponse) GetNumAckPendingOk() (*int32, bool)`

GetNumAckPendingOk returns a tuple with the NumAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAckPending

`func (o *JSPullConsumerInfoResponse) SetNumAckPending(v int32)`

SetNumAckPending sets NumAckPending field to given value.


### GetNumPending

`func (o *JSPullConsumerInfoResponse) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *JSPullConsumerInfoResponse) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *JSPullConsumerInfoResponse) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.


### GetNumRedelivered

`func (o *JSPullConsumerInfoResponse) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *JSPullConsumerInfoResponse) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *JSPullConsumerInfoResponse) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.


### GetNumWaiting

`func (o *JSPullConsumerInfoResponse) GetNumWaiting() int32`

GetNumWaiting returns the NumWaiting field if non-nil, zero value otherwise.

### GetNumWaitingOk

`func (o *JSPullConsumerInfoResponse) GetNumWaitingOk() (*int32, bool)`

GetNumWaitingOk returns a tuple with the NumWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWaiting

`func (o *JSPullConsumerInfoResponse) SetNumWaiting(v int32)`

SetNumWaiting sets NumWaiting field to given value.


### GetStreamName

`func (o *JSPullConsumerInfoResponse) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *JSPullConsumerInfoResponse) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *JSPullConsumerInfoResponse) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


