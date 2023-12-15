# ServerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAccounts** | **int32** |  | 
**ActiveServers** | Pointer to **int32** |  | [optional] 
**Connections** | **int32** |  | 
**Cores** | **int32** |  | 
**Cpu** | **float64** |  | 
**Gateways** | Pointer to Nullable[[][**GatewayStat**](GatewayStat.md)] |  | [optional] 
**Jetstream** | Pointer to Nullable[[**JetStreamVarz**](JetStreamVarz.md)] |  | [optional] 
**Mem** | **int64** |  | 
**Received** | [**DataStats**](DataStats.md) |  | 
**Routes** | Pointer to Nullable[[][**RouteStat**](RouteStat.md)] |  | [optional] 
**Sent** | [**DataStats**](DataStats.md) |  | 
**SlowConsumers** | **int64** |  | 
**Start** | **time.Time** |  | 
**Subscriptions** | **int32** |  | 
**TotalConnections** | **int32** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


