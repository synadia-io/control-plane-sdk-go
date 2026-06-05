# AccountViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPublicKey** | Pointer to **string** |  | [optional] 
**AuthType** | [**AccountAuthType**](AccountAuthType.md) |  | 
**Claims** | Pointer to [**AccountClaims**](AccountClaims.md) |  | [optional] 
**ClaimsInfo** | Pointer to [**AccountClaimsInfo**](AccountClaimsInfo.md) |  | [optional] 
**Connectors** | **bool** |  | 
**ConnectorsLogStream** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsPlatformAccount** | **bool** |  | 
**IsScpAccount** | **bool** |  | 
**IsSystemAccount** | **bool** |  | 
**Jwt** | Pointer to **string** |  | [optional] 
**JwtSettings** | Pointer to [**AccountJWTSettings**](AccountJWTSettings.md) |  | [optional] 
**JwtSyncError** | **string** |  | 
**JwtSyncStatus** | [**JwtSyncStatus**](JwtSyncStatus.md) |  | 
**Name** | **string** |  | 
**NscManaged** | **bool** |  | 
**System** | [**SystemInfo**](SystemInfo.md) |  | 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 
**UserCred** | Pointer to **string** |  | [optional] 
**UserJwt** | Pointer to **string** |  | [optional] 
**UserJwtExpiresInSecs** | **int64** |  | 
**Workloads** | **bool** |  | 
**WorkloadsLogStream** | Pointer to **string** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


