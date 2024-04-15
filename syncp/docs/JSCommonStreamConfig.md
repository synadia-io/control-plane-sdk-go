# JSCommonStreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDirect** | **bool** |  | 
**AllowRollupHdrs** | **bool** |  | 
**Compression** | Pointer to [**S2Compression**](S2Compression.md) |  | [optional] 
**DenyDelete** | **bool** |  | 
**DenyPurge** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Discard** | [**DiscardPolicy**](DiscardPolicy.md) |  | 
**DiscardNewPerSubject** | Pointer to **bool** |  | [optional] 
**DuplicateWindow** | Pointer to **int64** |  | [optional] 
**FirstSeq** | Pointer to **int32** |  | [optional] 
**MaxAge** | **int64** |  | 
**MaxBytes** | **int64** |  | 
**MaxConsumers** | **int32** |  | 
**MaxMsgSize** | Pointer to **int32** |  | [optional] 
**MaxMsgs** | **int64** |  | 
**MaxMsgsPerSubject** | **int64** |  | 
**Name** | **string** |  | 
**NoAck** | Pointer to **bool** |  | [optional] 
**NumReplicas** | **int32** |  | 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Republish** | Pointer to [**RePublish**](RePublish.md) |  | [optional] 
**Retention** | [**RetentionPolicy**](RetentionPolicy.md) |  | 
**Sealed** | **bool** |  | 
**Sources** | Pointer to [][**StreamSource**](StreamSource.md) |  | [optional] 
**Storage** | [**StorageType**](StorageType.md) |  | 
**SubjectTransform** | Pointer to [**SubjectTransformConfig**](SubjectTransformConfig.md) |  | [optional] 
**TemplateOwner** | Pointer to **string** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


