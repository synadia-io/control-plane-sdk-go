# SystemViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | [**SystemConnectionType**](SystemConnectionType.md) |  | [default to SYSTEMCONNECTIONTYPE_AGENT]
**Created** | **time.Time** |  | 
**DirectConnectionOpts** | Pointer to [**SystemDirectConnectionOpts**](SystemDirectConnectionOpts.md) |  | [optional] 
**HasManagedOperator** | **bool** |  | 
**HasManagedSystemAccount** | **bool** |  | 
**HostSystemId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IsTenant** | **bool** |  | 
**JetstreamDomain** | Pointer to **string** |  | [optional] 
**JetstreamEnabled** | **bool** |  | 
**JetstreamTiers** | Pointer to []**string** |  | [optional] 
**Limited** | **bool** |  | 
**LogForwardingOpts** | Pointer to [**SystemLogForwardingOpts**](SystemLogForwardingOpts.md) |  | [optional] 
**ManagedBy** | Pointer to [**SystemManagedBy**](SystemManagedBy.md) |  | [optional] 
**MaxNodes** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**NatsAlertingConfig** | Pointer to [**NatsAlertingConfig**](NatsAlertingConfig.md) |  | [optional] 
**OperatorClaims** | Pointer to [**OperatorClaims**](OperatorClaims.md) |  | [optional] 
**OperatorJwt** | Pointer to **string** |  | [optional] 
**PlatformComponents** | Pointer to [**PlatformComponentsViewResponse**](PlatformComponentsViewResponse.md) |  | [optional] 
**State** | [**SystemState**](SystemState.md) |  | 
**SystemAccountJwt** | Pointer to **string** |  | [optional] 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 
**UpdatesAvailable** | [**map[string]Update**](Update.md) |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


