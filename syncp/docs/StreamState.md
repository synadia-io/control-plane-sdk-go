# StreamState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **int32** |  | 
**ConsumerCount** | **int32** |  | 
**Deleted** | Pointer to []**int32** |  | [optional] 
**FirstSeq** | **int32** |  | 
**FirstTs** | **time.Time** |  | 
**LastSeq** | **int32** |  | 
**LastTs** | **time.Time** |  | 
**Lost** | Pointer to Nullable[[**LostStreamData**](LostStreamData.md)] |  | [optional] 
**Messages** | **int32** |  | 
**NumDeleted** | Pointer to **int32** |  | [optional] 
**NumSubjects** | Pointer to **int32** |  | [optional] 
**Subjects** | Pointer to **map[string]int32** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


