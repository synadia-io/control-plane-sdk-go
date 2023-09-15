# JSKVBucketCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** |  | 
**Description** | **string** |  | 
**History** | **int32** |  | 
**MaxBytes** | **int64** |  | 
**MaxValSize** | **int32** |  | 
**Mirror** | Pointer to [**NullableJSKVBucketCreateRequestMirror**](JSKVBucketCreateRequestMirror.md) |  | [optional] 
**Placement** | Pointer to [**NullableJSCommonStreamConfigPlacement**](JSCommonStreamConfigPlacement.md) |  | [optional] 
**Replicas** | **int32** |  | 
**Republish** | Pointer to [**NullableJSCommonStreamConfigRepublish**](JSCommonStreamConfigRepublish.md) |  | [optional] 
**Sources** | Pointer to [**[]StreamSource**](StreamSource.md) |  | [optional] 
**Storage** | [**StorageType**](StorageType.md) |  | 
**Ttl** | **int64** |  | 

## Methods

### NewJSKVBucketCreateRequest

`func NewJSKVBucketCreateRequest(bucketName string, description string, history int32, maxBytes int64, maxValSize int32, replicas int32, storage StorageType, ttl int64, ) *JSKVBucketCreateRequest`

NewJSKVBucketCreateRequest instantiates a new JSKVBucketCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSKVBucketCreateRequestWithDefaults

`func NewJSKVBucketCreateRequestWithDefaults() *JSKVBucketCreateRequest`

NewJSKVBucketCreateRequestWithDefaults instantiates a new JSKVBucketCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *JSKVBucketCreateRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *JSKVBucketCreateRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *JSKVBucketCreateRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetDescription

`func (o *JSKVBucketCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSKVBucketCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSKVBucketCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHistory

`func (o *JSKVBucketCreateRequest) GetHistory() int32`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *JSKVBucketCreateRequest) GetHistoryOk() (*int32, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *JSKVBucketCreateRequest) SetHistory(v int32)`

SetHistory sets History field to given value.


### GetMaxBytes

`func (o *JSKVBucketCreateRequest) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *JSKVBucketCreateRequest) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *JSKVBucketCreateRequest) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.


### GetMaxValSize

`func (o *JSKVBucketCreateRequest) GetMaxValSize() int32`

GetMaxValSize returns the MaxValSize field if non-nil, zero value otherwise.

### GetMaxValSizeOk

`func (o *JSKVBucketCreateRequest) GetMaxValSizeOk() (*int32, bool)`

GetMaxValSizeOk returns a tuple with the MaxValSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValSize

`func (o *JSKVBucketCreateRequest) SetMaxValSize(v int32)`

SetMaxValSize sets MaxValSize field to given value.


### GetMirror

`func (o *JSKVBucketCreateRequest) GetMirror() JSKVBucketCreateRequestMirror`

GetMirror returns the Mirror field if non-nil, zero value otherwise.

### GetMirrorOk

`func (o *JSKVBucketCreateRequest) GetMirrorOk() (*JSKVBucketCreateRequestMirror, bool)`

GetMirrorOk returns a tuple with the Mirror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirror

`func (o *JSKVBucketCreateRequest) SetMirror(v JSKVBucketCreateRequestMirror)`

SetMirror sets Mirror field to given value.

### HasMirror

`func (o *JSKVBucketCreateRequest) HasMirror() bool`

HasMirror returns a boolean if a field has been set.

### SetMirrorNil

`func (o *JSKVBucketCreateRequest) SetMirrorNil(b bool)`

 SetMirrorNil sets the value for Mirror to be an explicit nil

### UnsetMirror
`func (o *JSKVBucketCreateRequest) UnsetMirror()`

UnsetMirror ensures that no value is present for Mirror, not even an explicit nil
### GetPlacement

`func (o *JSKVBucketCreateRequest) GetPlacement() JSCommonStreamConfigPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *JSKVBucketCreateRequest) GetPlacementOk() (*JSCommonStreamConfigPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *JSKVBucketCreateRequest) SetPlacement(v JSCommonStreamConfigPlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *JSKVBucketCreateRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *JSKVBucketCreateRequest) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *JSKVBucketCreateRequest) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetReplicas

`func (o *JSKVBucketCreateRequest) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *JSKVBucketCreateRequest) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *JSKVBucketCreateRequest) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetRepublish

`func (o *JSKVBucketCreateRequest) GetRepublish() JSCommonStreamConfigRepublish`

GetRepublish returns the Republish field if non-nil, zero value otherwise.

### GetRepublishOk

`func (o *JSKVBucketCreateRequest) GetRepublishOk() (*JSCommonStreamConfigRepublish, bool)`

GetRepublishOk returns a tuple with the Republish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepublish

`func (o *JSKVBucketCreateRequest) SetRepublish(v JSCommonStreamConfigRepublish)`

SetRepublish sets Republish field to given value.

### HasRepublish

`func (o *JSKVBucketCreateRequest) HasRepublish() bool`

HasRepublish returns a boolean if a field has been set.

### SetRepublishNil

`func (o *JSKVBucketCreateRequest) SetRepublishNil(b bool)`

 SetRepublishNil sets the value for Republish to be an explicit nil

### UnsetRepublish
`func (o *JSKVBucketCreateRequest) UnsetRepublish()`

UnsetRepublish ensures that no value is present for Republish, not even an explicit nil
### GetSources

`func (o *JSKVBucketCreateRequest) GetSources() []StreamSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *JSKVBucketCreateRequest) GetSourcesOk() (*[]StreamSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *JSKVBucketCreateRequest) SetSources(v []StreamSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *JSKVBucketCreateRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *JSKVBucketCreateRequest) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *JSKVBucketCreateRequest) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetStorage

`func (o *JSKVBucketCreateRequest) GetStorage() StorageType`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *JSKVBucketCreateRequest) GetStorageOk() (*StorageType, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *JSKVBucketCreateRequest) SetStorage(v StorageType)`

SetStorage sets Storage field to given value.


### GetTtl

`func (o *JSKVBucketCreateRequest) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *JSKVBucketCreateRequest) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *JSKVBucketCreateRequest) SetTtl(v int64)`

SetTtl sets Ttl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


