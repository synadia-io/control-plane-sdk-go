# ServerStatsJetstream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**NullableJetStreamVarzConfig**](JetStreamVarzConfig.md) |  | [optional] 
**Meta** | Pointer to [**NullableJetStreamVarzMeta**](JetStreamVarzMeta.md) |  | [optional] 
**Stats** | Pointer to [**NullableJetStreamVarzStats**](JetStreamVarzStats.md) |  | [optional] 

## Methods

### NewServerStatsJetstream

`func NewServerStatsJetstream() *ServerStatsJetstream`

NewServerStatsJetstream instantiates a new ServerStatsJetstream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatsJetstreamWithDefaults

`func NewServerStatsJetstreamWithDefaults() *ServerStatsJetstream`

NewServerStatsJetstreamWithDefaults instantiates a new ServerStatsJetstream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ServerStatsJetstream) GetConfig() JetStreamVarzConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerStatsJetstream) GetConfigOk() (*JetStreamVarzConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerStatsJetstream) SetConfig(v JetStreamVarzConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerStatsJetstream) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ServerStatsJetstream) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ServerStatsJetstream) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetMeta

`func (o *ServerStatsJetstream) GetMeta() JetStreamVarzMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerStatsJetstream) GetMetaOk() (*JetStreamVarzMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerStatsJetstream) SetMeta(v JetStreamVarzMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServerStatsJetstream) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *ServerStatsJetstream) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *ServerStatsJetstream) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetStats

`func (o *ServerStatsJetstream) GetStats() JetStreamVarzStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ServerStatsJetstream) GetStatsOk() (*JetStreamVarzStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ServerStatsJetstream) SetStats(v JetStreamVarzStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ServerStatsJetstream) HasStats() bool`

HasStats returns a boolean if a field has been set.

### SetStatsNil

`func (o *ServerStatsJetstream) SetStatsNil(b bool)`

 SetStatsNil sets the value for Stats to be an explicit nil

### UnsetStats
`func (o *ServerStatsJetstream) UnsetStats()`

UnsetStats ensures that no value is present for Stats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


