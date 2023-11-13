# JSAssetInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternates** | Pointer to [**[]StreamAlternate**](StreamAlternate.md) |  | [optional] 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Sources** | Pointer to [**[]StreamSourceInfo**](StreamSourceInfo.md) |  | [optional] 
**State** | [**StreamState**](StreamState.md) |  | 
**Config** | [**JSCommonStreamConfig**](JSCommonStreamConfig.md) |  | 
**Id** | **string** |  | 
**JetstreamType** | [**JSType**](JSType.md) |  | 

## Methods

### NewJSAssetInfoResponse

`func NewJSAssetInfoResponse(created time.Time, state StreamState, config JSCommonStreamConfig, id string, jetstreamType JSType, ) *JSAssetInfoResponse`

NewJSAssetInfoResponse instantiates a new JSAssetInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSAssetInfoResponseWithDefaults

`func NewJSAssetInfoResponseWithDefaults() *JSAssetInfoResponse`

NewJSAssetInfoResponseWithDefaults instantiates a new JSAssetInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternates

`func (o *JSAssetInfoResponse) GetAlternates() []StreamAlternate`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *JSAssetInfoResponse) GetAlternatesOk() (*[]StreamAlternate, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *JSAssetInfoResponse) SetAlternates(v []StreamAlternate)`

SetAlternates sets Alternates field to given value.

### HasAlternates

`func (o *JSAssetInfoResponse) HasAlternates() bool`

HasAlternates returns a boolean if a field has been set.

### GetCluster

`func (o *JSAssetInfoResponse) GetCluster() ClusterInfo`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *JSAssetInfoResponse) GetClusterOk() (*ClusterInfo, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *JSAssetInfoResponse) SetCluster(v ClusterInfo)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *JSAssetInfoResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreated

`func (o *JSAssetInfoResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JSAssetInfoResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JSAssetInfoResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetSources

`func (o *JSAssetInfoResponse) GetSources() []StreamSourceInfo`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSAssetInfoResponse) GetSourcesOk() (*[]StreamSourceInfo, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSAssetInfoResponse) SetSources(v []StreamSourceInfo)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSAssetInfoResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetState

`func (o *JSAssetInfoResponse) GetState() StreamState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *JSAssetInfoResponse) GetStateOk() (*StreamState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *JSAssetInfoResponse) SetState(v StreamState)`

SetState sets State field to given value.


### GetConfig

`func (o *JSAssetInfoResponse) GetConfig() JSCommonStreamConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *JSAssetInfoResponse) GetConfigOk() (*JSCommonStreamConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *JSAssetInfoResponse) SetConfig(v JSCommonStreamConfig)`

SetConfig sets Config field to given value.


### GetId

`func (o *JSAssetInfoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JSAssetInfoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JSAssetInfoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJetstreamType

`func (o *JSAssetInfoResponse) GetJetstreamType() JSType`

GetJetstreamType returns the JetstreamType field if non-nil, zero value otherwise.

### GetJetstreamTypeOk

`func (o *JSAssetInfoResponse) GetJetstreamTypeOk() (*JSType, bool)`

GetJetstreamTypeOk returns a tuple with the JetstreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamType

`func (o *JSAssetInfoResponse) SetJetstreamType(v JSType)`

SetJetstreamType sets JetstreamType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


