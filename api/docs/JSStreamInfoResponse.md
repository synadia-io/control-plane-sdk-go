# JSStreamInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**JSStreamConfigRequest**](JSStreamConfigRequest.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Alternates** | Pointer to [**[]StreamAlternate**](StreamAlternate.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Sources** | Pointer to [**[]StreamSourceInfo**](StreamSourceInfo.md) |  | [optional] 
**State** | [**StreamState**](StreamState.md) |  | 

## Methods

### NewJSStreamInfoResponse

`func NewJSStreamInfoResponse(created time.Time, state StreamState, ) *JSStreamInfoResponse`

NewJSStreamInfoResponse instantiates a new JSStreamInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSStreamInfoResponseWithDefaults

`func NewJSStreamInfoResponseWithDefaults() *JSStreamInfoResponse`

NewJSStreamInfoResponseWithDefaults instantiates a new JSStreamInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *JSStreamInfoResponse) GetConfig() JSStreamConfigRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSStreamInfoResponse) GetConfigOk() (*JSStreamConfigRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSStreamInfoResponse) SetConfig(v JSStreamConfigRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JSStreamInfoResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *JSStreamInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSStreamInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSStreamInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JSStreamInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlternates

`func (o *JSStreamInfoResponse) GetAlternates() []StreamAlternate`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *JSStreamInfoResponse) GetAlternatesOk() (*[]StreamAlternate, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *JSStreamInfoResponse) SetAlternates(v []StreamAlternate)`

SetAlternates sets Alternates field to given value.

### HasAlternates

`func (o *JSStreamInfoResponse) HasAlternates() bool`

HasAlternates returns a boolean if a field has been set.

### GetCluster

`func (o *JSStreamInfoResponse) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSStreamInfoResponse) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSStreamInfoResponse) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSStreamInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreated

`func (o *JSStreamInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSStreamInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSStreamInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSources

`func (o *JSStreamInfoResponse) GetSources() []StreamSourceInfo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSStreamInfoResponse) GetSourcesOk() (*[]StreamSourceInfo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSStreamInfoResponse) SetSources(v []StreamSourceInfo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSStreamInfoResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *JSStreamInfoResponse) GetState() StreamState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JSStreamInfoResponse) GetStateOk() (*StreamState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JSStreamInfoResponse) SetState(v StreamState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


