# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
 | [**Info**](Info.md) |   | Embedded Struct
 | [**GenericFields**](GenericFields.md) |   | Embedded Struct
**Authorization** | Pointer to [**ExternalAuthorization**](ExternalAuthorization.md) |  | [optional] 
**DefaultPermissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**Exports** | Pointer to Nullable[[][**Export**](Export.md)] | Exports is a slice of exports | [optional] 
**Imports** | Pointer to Nullable[[][**Import**](Import.md)] | Imports is a list of import structs | [optional] 
**Limits** | Pointer to [**OperatorLimits**](OperatorLimits.md) |  | [optional] 
**Mappings** | Pointer to [**map[string][]WeightedMapping**](WeightedMapping.md) |  | [optional] 
**Revocations** | Pointer to **map[string]int64** | RevocationList is used to store a mapping of public keys to unix timestamps | [optional] 
**SigningKeys** | Pointer to [**SigningKeys**](SigningKeys.md) |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


