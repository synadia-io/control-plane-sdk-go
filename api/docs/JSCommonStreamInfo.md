# JSCommonStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternates** | Pointer to [**[]StreamAlternate**](StreamAlternate.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Sources** | Pointer to [**[]StreamSourceInfo**](StreamSourceInfo.md) |  | [optional] 
**State** | [**StreamState**](StreamState.md) |  | 

## Methods

### NewJSCommonStreamInfo

`func NewJSCommonStreamInfo(created time.Time, state StreamState, ) *JSCommonStreamInfo`

NewJSCommonStreamInfo instantiates a new JSCommonStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSCommonStreamInfoWithDefaults

`func NewJSCommonStreamInfoWithDefaults() *JSCommonStreamInfo`

NewJSCommonStreamInfoWithDefaults instantiates a new JSCommonStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternates

`func (o *JSCommonStreamInfo) GetAlternates() []StreamAlternate`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *JSCommonStreamInfo) GetAlternatesOk() (*[]StreamAlternate, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *JSCommonStreamInfo) SetAlternates(v []StreamAlternate)`

SetAlternates sets Alternates field to given value.

### HasAlternates

`func (o *JSCommonStreamInfo) HasAlternates() bool`

HasAlternates returns a boolean if a field has been set.

### GetCluster

`func (o *JSCommonStreamInfo) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSCommonStreamInfo) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSCommonStreamInfo) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSCommonStreamInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreated

`func (o *JSCommonStreamInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSCommonStreamInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSCommonStreamInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSources

`func (o *JSCommonStreamInfo) GetSources() []StreamSourceInfo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSCommonStreamInfo) GetSourcesOk() (*[]StreamSourceInfo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSCommonStreamInfo) SetSources(v []StreamSourceInfo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSCommonStreamInfo) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *JSCommonStreamInfo) GetState() StreamState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JSCommonStreamInfo) GetStateOk() (*StreamState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JSCommonStreamInfo) SetState(v StreamState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


