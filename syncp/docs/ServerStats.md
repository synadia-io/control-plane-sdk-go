# ServerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAccounts** | **int64** |  | 
**ActiveServers** | Pointer to **int64** |  | [optional] 
**Connections** | **int64** |  | 
**Cores** | **int64** |  | 
**Cpu** | **float64** |  | 
**Gateways** | Pointer to Nullable[[][**GatewayStat**](GatewayStat.md)] |  | [optional] 
**Jetstream** | Pointer to Nullable[[**JetStreamVarz**](JetStreamVarz.md)] |  | [optional] 
**Mem** | **int64** |  | 
**Received** | [**DataStats**](DataStats.md) |  | 
**Routes** | Pointer to Nullable[[][**RouteStat**](RouteStat.md)] |  | [optional] 
**Sent** | [**DataStats**](DataStats.md) |  | 
**SlowConsumers** | **int64** |  | 
**Start** | **time.Time** |  | 
**Subscriptions** | **uint32** |  | 
**TotalConnections** | **uint64** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


