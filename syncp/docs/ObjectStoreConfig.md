# ObjectStoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Compression** | Pointer to **bool** | Enable underlying stream compression. NOTE: Compression is supported for nats-server 2.10.0+ | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxAge** | Pointer to **int64** |  | [optional] 
**MaxBytes** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Bucket-specific metadata NOTE: Metadata requires nats-server v2.10.0+ | [optional] 
**NumReplicas** | Pointer to **int64** |  | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Storage** | Pointer to [**StorageType**](StorageType.md) |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


