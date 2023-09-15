# JetStreamVarzConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompressOk** | Pointer to **bool** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**MaxMemory** | **int64** |  | 
**MaxStorage** | **int64** |  | 
**StoreDir** | Pointer to **string** |  | [optional] 

## Methods

### NewJetStreamVarzConfig

`func NewJetStreamVarzConfig(maxMemory int64, maxStorage int64, ) *JetStreamVarzConfig`

NewJetStreamVarzConfig instantiates a new JetStreamVarzConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamVarzConfigWithDefaults

`func NewJetStreamVarzConfigWithDefaults() *JetStreamVarzConfig`

NewJetStreamVarzConfigWithDefaults instantiates a new JetStreamVarzConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompressOk

`func (o *JetStreamVarzConfig) GetCompressOk() bool`

GetCompressOk returns the CompressOk field if non-nil, zero value otherwise.

### GetCompressOkOk

`func (o *JetStreamVarzConfig) GetCompressOkOk() (*bool, bool)`

GetCompressOkOk returns a tuple with the CompressOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOk

`func (o *JetStreamVarzConfig) SetCompressOk(v bool)`

SetCompressOk sets CompressOk field to given value.

### HasCompressOk

`func (o *JetStreamVarzConfig) HasCompressOk() bool`

HasCompressOk returns a boolean if a field has been set.

### GetDomain

`func (o *JetStreamVarzConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *JetStreamVarzConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *JetStreamVarzConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *JetStreamVarzConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaxMemory

`func (o *JetStreamVarzConfig) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *JetStreamVarzConfig) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *JetStreamVarzConfig) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.


### GetMaxStorage

`func (o *JetStreamVarzConfig) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *JetStreamVarzConfig) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *JetStreamVarzConfig) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.


### GetStoreDir

`func (o *JetStreamVarzConfig) GetStoreDir() string`

GetStoreDir returns the StoreDir field if non-nil, zero value otherwise.

### GetStoreDirOk

`func (o *JetStreamVarzConfig) GetStoreDirOk() (*string, bool)`

GetStoreDirOk returns a tuple with the StoreDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreDir

`func (o *JetStreamVarzConfig) SetStoreDir(v string)`

SetStoreDir sets StoreDir field to given value.

### HasStoreDir

`func (o *JetStreamVarzConfig) HasStoreDir() bool`

HasStoreDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


