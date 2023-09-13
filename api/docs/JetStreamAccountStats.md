# JetStreamAccountStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JetStreamTier** | [**JetStreamTier**](JetStreamTier.md) |  | 
**Api** | [**JetStreamAPIStats**](JetStreamAPIStats.md) |  | 
**Domain** | Pointer to **string** |  | [optional] 
**Tiers** | Pointer to [**map[string]JetStreamTier**](JetStreamTier.md) |  | [optional] 

## Methods

### NewJetStreamAccountStats

`func NewJetStreamAccountStats(jetStreamTier JetStreamTier, api JetStreamAPIStats, ) *JetStreamAccountStats`

NewJetStreamAccountStats instantiates a new JetStreamAccountStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJetStreamAccountStatsWithDefaults

`func NewJetStreamAccountStatsWithDefaults() *JetStreamAccountStats`

NewJetStreamAccountStatsWithDefaults instantiates a new JetStreamAccountStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJetStreamTier

`func (o *JetStreamAccountStats) GetJetStreamTier() JetStreamTier`

GetJetStreamTier returns the JetStreamTier field if non-nil, zero value otherwise.

### GetJetStreamTierOk

`func (o *JetStreamAccountStats) GetJetStreamTierOk() (*JetStreamTier, bool)`

GetJetStreamTierOk returns a tuple with the JetStreamTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetStreamTier

`func (o *JetStreamAccountStats) SetJetStreamTier(v JetStreamTier)`

SetJetStreamTier sets JetStreamTier field to given value.


### GetApi

`func (o *JetStreamAccountStats) GetApi() JetStreamAPIStats`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *JetStreamAccountStats) GetApiOk() (*JetStreamAPIStats, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *JetStreamAccountStats) SetApi(v JetStreamAPIStats)`

SetApi sets Api field to given value.


### GetDomain

`func (o *JetStreamAccountStats) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *JetStreamAccountStats) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *JetStreamAccountStats) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *JetStreamAccountStats) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTiers

`func (o *JetStreamAccountStats) GetTiers() map[string]JetStreamTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *JetStreamAccountStats) GetTiersOk() (*map[string]JetStreamTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *JetStreamAccountStats) SetTiers(v map[string]JetStreamTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *JetStreamAccountStats) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


