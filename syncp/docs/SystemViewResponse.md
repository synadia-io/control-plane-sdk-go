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
**HttpGatewayConfig** | Pointer to [**HTTPGatewayConfig**](HTTPGatewayConfig.md) |  | [optional] 
**Id** | **string** |  | 
**IsTenant** | **bool** |  | 
**JetstreamDomain** | Pointer to **string** |  | [optional] 
**JetstreamEnabled** | **bool** |  | 
**JetstreamTiers** | Pointer to []**string** |  | [optional] 
**Name** | **string** |  | 
**OperatorClaims** | Pointer to [**OperatorClaims**](OperatorClaims.md) |  | [optional] 
**OperatorJwt** | Pointer to **string** |  | [optional] 
**State** | [**SystemState**](SystemState.md) |  | 
**SystemAccountJwt** | Pointer to **string** |  | [optional] 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


