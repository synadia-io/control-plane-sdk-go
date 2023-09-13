# ConsumerInfoCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leader** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**[]ClusterInfoReplicasInner**](ClusterInfoReplicasInner.md) |  | [optional] 

## Methods

### NewConsumerInfoCluster

`func NewConsumerInfoCluster() *ConsumerInfoCluster`

NewConsumerInfoCluster instantiates a new ConsumerInfoCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerInfoClusterWithDefaults

`func NewConsumerInfoClusterWithDefaults() *ConsumerInfoCluster`

NewConsumerInfoClusterWithDefaults instantiates a new ConsumerInfoCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeader

`func (o *ConsumerInfoCluster) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ConsumerInfoCluster) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ConsumerInfoCluster) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *ConsumerInfoCluster) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetName

`func (o *ConsumerInfoCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsumerInfoCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsumerInfoCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsumerInfoCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicas

`func (o *ConsumerInfoCluster) GetReplicas() []ClusterInfoReplicasInner`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ConsumerInfoCluster) GetReplicasOk() (*[]ClusterInfoReplicasInner, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ConsumerInfoCluster) SetReplicas(v []ClusterInfoReplicasInner)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *ConsumerInfoCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


