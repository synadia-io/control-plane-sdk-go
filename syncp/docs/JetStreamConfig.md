# JetStreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressOk** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**MaxMemory** | **int64** |  | 
**MaxStorage** | **int64** |  | 
**StoreDir** | Pointer to **string** |  | [optional] 

## Methods

### NewJetStreamConfig

`func NewJetStreamConfig(maxMemory int64, maxStorage int64, ) *JetStreamConfig`

NewJetStreamConfig instantiates a new JetStreamConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamConfigWithDefaults

`func NewJetStreamConfigWithDefaults() *JetStreamConfig`

NewJetStreamConfigWithDefaults instantiates a new JetStreamConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressOk

`func (o *JetStreamConfig) GetCompressOk() bool`

GetCompressOk returns the CompressOk field if non-nil, zero value otherwise.

### GetCompressOkOk

`func (o *JetStreamConfig) GetCompressOkOk() (*bool, bool)`

GetCompressOkOk returns a tuple with the CompressOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOk

`func (o *JetStreamConfig) SetCompressOk(v bool)`

SetCompressOk sets CompressOk field to given value.

### HasCompressOk

`func (o *JetStreamConfig) HasCompressOk() bool`

HasCompressOk returns a boolean if a field has been set.

### GetDomain

`func (o *JetStreamConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *JetStreamConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *JetStreamConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *JetStreamConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaxMemory

`func (o *JetStreamConfig) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *JetStreamConfig) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *JetStreamConfig) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxStorage

`func (o *JetStreamConfig) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *JetStreamConfig) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *JetStreamConfig) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.


### GetStoreDir

`func (o *JetStreamConfig) GetStoreDir() string`

GetStoreDir returns the StoreDir field if non-nil, zero value otherwise.

### GetStoreDirOk

`func (o *JetStreamConfig) GetStoreDirOk() (*string, bool)`

GetStoreDirOk returns a tuple with the StoreDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreDir

`func (o *JetStreamConfig) SetStoreDir(v string)`

SetStoreDir sets StoreDir field to given value.

### HasStoreDir

`func (o *JetStreamConfig) HasStoreDir() bool`

HasStoreDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


