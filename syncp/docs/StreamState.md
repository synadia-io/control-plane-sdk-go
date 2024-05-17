# StreamState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | **uint64** |  | 
**ConsumerCount** | **int64** |  | 
**Deleted** | Pointer to []**uint64** |  | [optional] 
**FirstSeq** | **uint64** |  | 
**FirstTs** | **time.Time** |  | 
**LastSeq** | **uint64** |  | 
**LastTs** | **time.Time** |  | 
**Lost** | Pointer to Nullable[[**LostStreamData**](LostStreamData.md)] |  | [optional] 
**Messages** | **uint64** |  | 
**NumDeleted** | Pointer to **int64** |  | [optional] 
**NumSubjects** | Pointer to **int64** |  | [optional] 
**Subjects** | Pointer to **map[string]uint64** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


