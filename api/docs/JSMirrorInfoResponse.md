# JSMirrorInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**JSMirrorConfigRequest**](JSMirrorConfigRequest.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mirror** | Pointer to [**NullableStreamSourceInfo**](StreamSourceInfo.md) |  | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Alternates** | Pointer to [**[]StreamAlternate**](StreamAlternate.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Sources** | Pointer to [**[]StreamSourceInfo**](StreamSourceInfo.md) |  | [optional] 
**State** | [**StreamState**](StreamState.md) |  | 

## Methods

### NewJSMirrorInfoResponse

`func NewJSMirrorInfoResponse(created time.Time, state StreamState, ) *JSMirrorInfoResponse`

NewJSMirrorInfoResponse instantiates a new JSMirrorInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSMirrorInfoResponseWithDefaults

`func NewJSMirrorInfoResponseWithDefaults() *JSMirrorInfoResponse`

NewJSMirrorInfoResponseWithDefaults instantiates a new JSMirrorInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *JSMirrorInfoResponse) GetConfig() JSMirrorConfigRequest`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSMirrorInfoResponse) GetConfigOk() (*JSMirrorConfigRequest, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSMirrorInfoResponse) SetConfig(v JSMirrorConfigRequest)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *JSMirrorInfoResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetId

`func (o *JSMirrorInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSMirrorInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSMirrorInfoResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JSMirrorInfoResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMirror

`func (o *JSMirrorInfoResponse) GetMirror() StreamSourceInfo`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *JSMirrorInfoResponse) GetMirrorOk() (*StreamSourceInfo, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *JSMirrorInfoResponse) SetMirror(v StreamSourceInfo)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *JSMirrorInfoResponse) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### SetMirrorNil

`func (o *JSMirrorInfoResponse) SetMirrorNil(b bool)`

 SetMirrorNil sets the value for Mirror to be an explicit nil

### UnsetMirror
`func (o *JSMirrorInfoResponse) UnsetMirror()`

UnsetMirror ensures that no value is present for Mirror, not even an explicit nil
### GetShared

`func (o *JSMirrorInfoResponse) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *JSMirrorInfoResponse) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *JSMirrorInfoResponse) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *JSMirrorInfoResponse) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetAlternates

`func (o *JSMirrorInfoResponse) GetAlternates() []StreamAlternate`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *JSMirrorInfoResponse) GetAlternatesOk() (*[]StreamAlternate, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *JSMirrorInfoResponse) SetAlternates(v []StreamAlternate)`

SetAlternates sets Alternates field to given value.

### HasAlternates

`func (o *JSMirrorInfoResponse) HasAlternates() bool`

HasAlternates returns a boolean if a field has been set.

### GetCluster

`func (o *JSMirrorInfoResponse) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSMirrorInfoResponse) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSMirrorInfoResponse) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSMirrorInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreated

`func (o *JSMirrorInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSMirrorInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSMirrorInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSources

`func (o *JSMirrorInfoResponse) GetSources() []StreamSourceInfo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSMirrorInfoResponse) GetSourcesOk() (*[]StreamSourceInfo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSMirrorInfoResponse) SetSources(v []StreamSourceInfo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSMirrorInfoResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *JSMirrorInfoResponse) GetState() StreamState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JSMirrorInfoResponse) GetStateOk() (*StreamState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JSMirrorInfoResponse) SetState(v StreamState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


