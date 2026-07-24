# JSCommonStreamConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAtomic** | Pointer to **bool** |  | [optional] 
**AllowBatched** | Pointer to **bool** |  | [optional] 
**AllowDirect** | **bool** |  | 
**AllowMsgCounter** | Pointer to **bool** |  | [optional] 
**AllowMsgSchedules** | Pointer to **bool** |  | [optional] 
**AllowMsgTtl** | Pointer to **bool** |  | [optional] 
**AllowRollupHdrs** | **bool** |  | 
**Compression** | Pointer to [**Compression**](Compression.md) |  | [optional] 
**ConsumerLimits** | Pointer to [**JSStreamConsumerLimits**](JSStreamConsumerLimits.md) |  | [optional] 
**DenyDelete** | **bool** |  | 
**DenyPurge** | **bool** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Discard** | [**DiscardPolicy**](DiscardPolicy.md) |  | 
**DiscardNewPerSubject** | Pointer to **bool** |  | [optional] 
**DuplicateWindow** | Pointer to **int64** |  | [optional] 
**FirstSeq** | Pointer to **uint64** |  | [optional] 
**MaxAge** | **int64** |  | 
**MaxBytes** | **int64** |  | 
**MaxConsumers** | **int64** |  | 
**MaxMsgSize** | Pointer to **int64** |  | [optional] 
**MaxMsgs** | **int64** |  | 
**MaxMsgsPerSubject** | **int64** |  | 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | **string** |  | 
**NoAck** | Pointer to **bool** |  | [optional] 
**NumReplicas** | **int64** |  | 
**PersistMode** | Pointer to [**JSPersistMode**](JSPersistMode.md) |  | [optional] 
**Placement** | Pointer to [**Placement**](Placement.md) |  | [optional] 
**Republish** | Pointer to [**RePublish**](RePublish.md) |  | [optional] 
**Retention** | [**RetentionPolicy**](RetentionPolicy.md) |  | 
**Sealed** | **bool** |  | 
**Sources** | Pointer to [][**StreamSource**](StreamSource.md) |  | [optional] 
**Storage** | [**StorageType**](StorageType.md) |  | 
**SubjectDeleteMarkerTtl** | Pointer to **int64** |  | [optional] 
**SubjectTransform** | Pointer to [**SubjectTransformConfig**](SubjectTransformConfig.md) |  | [optional] 
**TemplateOwner** | Pointer to **string** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


