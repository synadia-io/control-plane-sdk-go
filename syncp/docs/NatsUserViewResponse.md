# NatsUserViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountInfo**](AccountInfo.md) |  | 
**Claims** | [**UserClaims**](UserClaims.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**Jwt** | **string** |  | 
**JwtExpiresAtMax** | **int64** |  | 
**JwtExpiresInSecs** | **int64** |  | 
**JwtRevoked** | Pointer to **bool** | True when an account revocation covers the last JWT that the server signed for the user. | [optional] 
**JwtSettings** | [**NatsUserJwtSettings**](NatsUserJwtSettings.md) |  | 
**MissingSeed** | **bool** |  | 
**Name** | **string** |  | 
**SharedNkeyUsers** | Pointer to [][**NatsUserSharedNkeyUser**](NatsUserSharedNkeyUser.md) | Other users in the account with the same NKey. | [optional] 
**SkGroupId** | Pointer to **string** |  | [optional] 
**System** | [**SystemInfo**](SystemInfo.md) |  | 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 
**UserPublicKey** | **string** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


