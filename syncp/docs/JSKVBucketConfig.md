# JSKVBucketConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | **string** |  | 
**Compression** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**History** | Pointer to **int64** |  | [optional] [default to 1]
**MaxAge** | Pointer to **int64** |  | [optional] 
**MaxBytes** | Pointer to **int64** |  | [optional] 
**MaxValueSize** | Pointer to **int64** |  | [optional] 
**Mirror** | Pointer to [**StreamSource**](StreamSource.md) |  | [optional] 
**NumReplicas** | Pointer to **int64** |  | [optional] [default to 1]
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Republish** | Pointer to [**RePublish**](RePublish.md) |  | [optional] 
**Sources** | Pointer to [][**StreamSource**](StreamSource.md) |  | [optional] 
**Storage** | [**StorageType**](StorageType.md) |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


