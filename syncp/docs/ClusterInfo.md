# ClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Leader** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**[]ClusterInfoReplicasInner**](ClusterInfoReplicasInner.md) |  | [optional] 

## Methods

### NewClusterInfo

`func NewClusterInfo() *ClusterInfo`

NewClusterInfo instantiates a new ClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoWithDefaults

`func NewClusterInfoWithDefaults() *ClusterInfo`

NewClusterInfoWithDefaults instantiates a new ClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeader

`func (o *ClusterInfo) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ClusterInfo) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ClusterInfo) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *ClusterInfo) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetName

`func (o *ClusterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicas

`func (o *ClusterInfo) GetReplicas() []ClusterInfoReplicasInner`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ClusterInfo) GetReplicasOk() (*[]ClusterInfoReplicasInner, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ClusterInfo) SetReplicas(v []ClusterInfoReplicasInner)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *ClusterInfo) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


