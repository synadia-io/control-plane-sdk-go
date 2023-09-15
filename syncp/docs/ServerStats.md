# ServerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAccounts** | **int32** |  | 
**ActiveServers** | Pointer to **int32** |  | [optional] 
**Connections** | **int32** |  | 
**Cores** | **int32** |  | 
**Cpu** | **float64** |  | 
**Gateways** | Pointer to [**[]ServerStatsGatewaysInner**](ServerStatsGatewaysInner.md) |  | [optional] 
**Jetstream** | Pointer to [**NullableServerStatsJetstream**](ServerStatsJetstream.md) |  | [optional] 
**Mem** | **int64** |  | 
**Received** | [**DataStats**](DataStats.md) |  | 
**Routes** | Pointer to [**[]ServerStatsRoutesInner**](ServerStatsRoutesInner.md) |  | [optional] 
**Sent** | [**DataStats**](DataStats.md) |  | 
**SlowConsumers** | **int64** |  | 
**Start** | **time.Time** |  | 
**Subscriptions** | **int32** |  | 
**TotalConnections** | **int32** |  | 

## Methods

### NewServerStats

`func NewServerStats(activeAccounts int32, connections int32, cores int32, cpu float64, mem int64, received DataStats, sent DataStats, slowConsumers int64, start time.Time, subscriptions int32, totalConnections int32, ) *ServerStats`

NewServerStats instantiates a new ServerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatsWithDefaults

`func NewServerStatsWithDefaults() *ServerStats`

NewServerStatsWithDefaults instantiates a new ServerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAccounts

`func (o *ServerStats) GetActiveAccounts() int32`

GetActiveAccounts returns the ActiveAccounts field if non-nil, zero value otherwise.

### GetActiveAccountsOk

`func (o *ServerStats) GetActiveAccountsOk() (*int32, bool)`

GetActiveAccountsOk returns a tuple with the ActiveAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAccounts

`func (o *ServerStats) SetActiveAccounts(v int32)`

SetActiveAccounts sets ActiveAccounts field to given value.


### GetActiveServers

`func (o *ServerStats) GetActiveServers() int32`

GetActiveServers returns the ActiveServers field if non-nil, zero value otherwise.

### GetActiveServersOk

`func (o *ServerStats) GetActiveServersOk() (*int32, bool)`

GetActiveServersOk returns a tuple with the ActiveServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveServers

`func (o *ServerStats) SetActiveServers(v int32)`

SetActiveServers sets ActiveServers field to given value.

### HasActiveServers

`func (o *ServerStats) HasActiveServers() bool`

HasActiveServers returns a boolean if a field has been set.

### GetConnections

`func (o *ServerStats) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ServerStats) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ServerStats) SetConnections(v int32)`

SetConnections sets Connections field to given value.


### GetCores

`func (o *ServerStats) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ServerStats) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ServerStats) SetCores(v int32)`

SetCores sets Cores field to given value.


### GetCpu

`func (o *ServerStats) GetCpu() float64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ServerStats) GetCpuOk() (*float64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ServerStats) SetCpu(v float64)`

SetCpu sets Cpu field to given value.


### GetGateways

`func (o *ServerStats) GetGateways() []ServerStatsGatewaysInner`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *ServerStats) GetGatewaysOk() (*[]ServerStatsGatewaysInner, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *ServerStats) SetGateways(v []ServerStatsGatewaysInner)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *ServerStats) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetJetstream

`func (o *ServerStats) GetJetstream() ServerStatsJetstream`

GetJetstream returns the Jetstream field if non-nil, zero value otherwise.

### GetJetstreamOk

`func (o *ServerStats) GetJetstreamOk() (*ServerStatsJetstream, bool)`

GetJetstreamOk returns a tuple with the Jetstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstream

`func (o *ServerStats) SetJetstream(v ServerStatsJetstream)`

SetJetstream sets Jetstream field to given value.

### HasJetstream

`func (o *ServerStats) HasJetstream() bool`

HasJetstream returns a boolean if a field has been set.

### SetJetstreamNil

`func (o *ServerStats) SetJetstreamNil(b bool)`

 SetJetstreamNil sets the value for Jetstream to be an explicit nil

### UnsetJetstream
`func (o *ServerStats) UnsetJetstream()`

UnsetJetstream ensures that no value is present for Jetstream, not even an explicit nil
### GetMem

`func (o *ServerStats) GetMem() int64`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *ServerStats) GetMemOk() (*int64, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *ServerStats) SetMem(v int64)`

SetMem sets Mem field to given value.


### GetReceived

`func (o *ServerStats) GetReceived() DataStats`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ServerStats) GetReceivedOk() (*DataStats, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ServerStats) SetReceived(v DataStats)`

SetReceived sets Received field to given value.


### GetRoutes

`func (o *ServerStats) GetRoutes() []ServerStatsRoutesInner`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ServerStats) GetRoutesOk() (*[]ServerStatsRoutesInner, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ServerStats) SetRoutes(v []ServerStatsRoutesInner)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *ServerStats) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetSent

`func (o *ServerStats) GetSent() DataStats`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ServerStats) GetSentOk() (*DataStats, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ServerStats) SetSent(v DataStats)`

SetSent sets Sent field to given value.


### GetSlowConsumers

`func (o *ServerStats) GetSlowConsumers() int64`

GetSlowConsumers returns the SlowConsumers field if non-nil, zero value otherwise.

### GetSlowConsumersOk

`func (o *ServerStats) GetSlowConsumersOk() (*int64, bool)`

GetSlowConsumersOk returns a tuple with the SlowConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowConsumers

`func (o *ServerStats) SetSlowConsumers(v int64)`

SetSlowConsumers sets SlowConsumers field to given value.


### GetStart

`func (o *ServerStats) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ServerStats) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ServerStats) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetSubscriptions

`func (o *ServerStats) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ServerStats) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ServerStats) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.


### GetTotalConnections

`func (o *ServerStats) GetTotalConnections() int32`

GetTotalConnections returns the TotalConnections field if non-nil, zero value otherwise.

### GetTotalConnectionsOk

`func (o *ServerStats) GetTotalConnectionsOk() (*int32, bool)`

GetTotalConnectionsOk returns a tuple with the TotalConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConnections

`func (o *ServerStats) SetTotalConnections(v int32)`

SetTotalConnections sets TotalConnections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


