# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatPolicy** | **string** |  | 
**Definition** | **string** | The schema definition. | 
**Delete** | Pointer to **bool** | Deleted is a flag to indicate if the schema is deleted. | [optional] 
**Description** | **string** | Description of the schema. | 
**Format** | **string** |  | 
**Metadata** | **map[string]string** | Metadata is a map of key-value pairs. | 
**Name** | **string** | Unique name of the schema. | 
**Revision** | **int64** | Revision of the schema. | 
**Time** | **time.Time** | Timestamp when this schema was added. | 
**Version** | **int64** | TODO(jrm): We probably do not need the ID, as we can use subject + version as a compound unique identifier. Unique identifier of the schema. ID string &#x60;json:\&quot;id\&quot;&#x60; | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


