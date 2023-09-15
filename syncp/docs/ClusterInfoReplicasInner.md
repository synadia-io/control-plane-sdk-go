# ClusterInfoReplicasInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int64** |  | 
**Current** | **bool** |  | 
**Lag** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Offline** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterInfoReplicasInner

`func NewClusterInfoReplicasInner(active int64, current bool, name string, ) *ClusterInfoReplicasInner`

NewClusterInfoReplicasInner instantiates a new ClusterInfoReplicasInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterInfoReplicasInnerWithDefaults

`func NewClusterInfoReplicasInnerWithDefaults() *ClusterInfoReplicasInner`

NewClusterInfoReplicasInnerWithDefaults instantiates a new ClusterInfoReplicasInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ClusterInfoReplicasInner) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ClusterInfoReplicasInner) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ClusterInfoReplicasInner) SetActive(v int64)`

SetActive sets Active field to given value.


### GetCurrent

`func (o *ClusterInfoReplicasInner) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ClusterInfoReplicasInner) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ClusterInfoReplicasInner) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetLag

`func (o *ClusterInfoReplicasInner) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ClusterInfoReplicasInner) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ClusterInfoReplicasInner) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *ClusterInfoReplicasInner) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetName

`func (o *ClusterInfoReplicasInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterInfoReplicasInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterInfoReplicasInner) SetName(v string)`

SetName sets Name field to given value.


### GetOffline

`func (o *ClusterInfoReplicasInner) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *ClusterInfoReplicasInner) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *ClusterInfoReplicasInner) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *ClusterInfoReplicasInner) HasOffline() bool`

HasOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


