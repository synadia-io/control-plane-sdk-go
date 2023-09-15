# JetStreamVarzStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | **int32** |  | 
**Api** | [**JetStreamAPIStats**](JetStreamAPIStats.md) |  | 
**HaAssets** | **int32** |  | 
**Memory** | **int32** |  | 
**ReservedMemory** | **int32** |  | 
**ReservedStorage** | **int32** |  | 
**Storage** | **int32** |  | 

## Methods

### NewJetStreamVarzStats

`func NewJetStreamVarzStats(accounts int32, api JetStreamAPIStats, haAssets int32, memory int32, reservedMemory int32, reservedStorage int32, storage int32, ) *JetStreamVarzStats`

NewJetStreamVarzStats instantiates a new JetStreamVarzStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamVarzStatsWithDefaults

`func NewJetStreamVarzStatsWithDefaults() *JetStreamVarzStats`

NewJetStreamVarzStatsWithDefaults instantiates a new JetStreamVarzStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *JetStreamVarzStats) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *JetStreamVarzStats) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *JetStreamVarzStats) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.


### GetApi

`func (o *JetStreamVarzStats) GetApi() JetStreamAPIStats`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *JetStreamVarzStats) GetApiOk() (*JetStreamAPIStats, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *JetStreamVarzStats) SetApi(v JetStreamAPIStats)`

SetApi sets Api field to given value.


### GetHaAssets

`func (o *JetStreamVarzStats) GetHaAssets() int32`

GetHaAssets returns the HaAssets field if non-nil, zero value otherwise.

### GetHaAssetsOk

`func (o *JetStreamVarzStats) GetHaAssetsOk() (*int32, bool)`

GetHaAssetsOk returns a tuple with the HaAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaAssets

`func (o *JetStreamVarzStats) SetHaAssets(v int32)`

SetHaAssets sets HaAssets field to given value.


### GetMemory

`func (o *JetStreamVarzStats) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *JetStreamVarzStats) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *JetStreamVarzStats) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetReservedMemory

`func (o *JetStreamVarzStats) GetReservedMemory() int32`

GetReservedMemory returns the ReservedMemory field if non-nil, zero value otherwise.

### GetReservedMemoryOk

`func (o *JetStreamVarzStats) GetReservedMemoryOk() (*int32, bool)`

GetReservedMemoryOk returns a tuple with the ReservedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedMemory

`func (o *JetStreamVarzStats) SetReservedMemory(v int32)`

SetReservedMemory sets ReservedMemory field to given value.


### GetReservedStorage

`func (o *JetStreamVarzStats) GetReservedStorage() int32`

GetReservedStorage returns the ReservedStorage field if non-nil, zero value otherwise.

### GetReservedStorageOk

`func (o *JetStreamVarzStats) GetReservedStorageOk() (*int32, bool)`

GetReservedStorageOk returns a tuple with the ReservedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStorage

`func (o *JetStreamVarzStats) SetReservedStorage(v int32)`

SetReservedStorage sets ReservedStorage field to given value.


### GetStorage

`func (o *JetStreamVarzStats) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *JetStreamVarzStats) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *JetStreamVarzStats) SetStorage(v int32)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


