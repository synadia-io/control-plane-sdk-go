# AccountJWTSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
 | [**Info**](Info.md) |   | Embedded Struct
**Authorization** | Pointer to [**ExternalAuthorization**](ExternalAuthorization.md) |  | [optional] 
**Limits** | Pointer to [**OperatorLimits**](OperatorLimits.md) |  | [optional] 
**Mappings** | Pointer to [**map[string][]WeightedMapping**](WeightedMapping.md) |  | [optional] 
**Revocations** | Pointer to **map[string]int64** | RevocationList is used to store a mapping of public keys to unix timestamps | [optional] 
**Tags** | Pointer to []**string** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


