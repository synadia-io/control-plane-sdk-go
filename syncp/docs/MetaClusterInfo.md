# MetaClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterSize** | **int32** |  | 
**Leader** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Replicas** | Pointer to [**[]ClusterInfoReplicasInner**](ClusterInfoReplicasInner.md) |  | [optional] 

## Methods

### NewMetaClusterInfo

`func NewMetaClusterInfo(clusterSize int32, ) *MetaClusterInfo`

NewMetaClusterInfo instantiates a new MetaClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaClusterInfoWithDefaults

`func NewMetaClusterInfoWithDefaults() *MetaClusterInfo`

NewMetaClusterInfoWithDefaults instantiates a new MetaClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterSize

`func (o *MetaClusterInfo) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *MetaClusterInfo) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *MetaClusterInfo) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.


### GetLeader

`func (o *MetaClusterInfo) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *MetaClusterInfo) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *MetaClusterInfo) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *MetaClusterInfo) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetName

`func (o *MetaClusterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaClusterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaClusterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaClusterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplicas

`func (o *MetaClusterInfo) GetReplicas() []ClusterInfoReplicasInner`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MetaClusterInfo) GetReplicasOk() (*[]ClusterInfoReplicasInner, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MetaClusterInfo) SetReplicas(v []ClusterInfoReplicasInner)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *MetaClusterInfo) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


