# AddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatPolicy** | Pointer to **string** | SchemaCompatPolicy represents the compatibility policy for schema evolution. Supported values are \&quot;none\&quot;, \&quot;backwards\&quot;, \&quot;forward\&quot;, and \&quot;full\&quot;. | [optional] 
**Definition** | **string** | TODO(jrm): we should validate the passed schema The schema definition (required). | 
**Description** | Pointer to **string** | Description of the schema (optional). | [optional] 
**Format** | **string** | SchemaFormat represents the format of a schema definition. Supported values are \&quot;jsonschema\&quot;. | 
**Metadata** | Pointer to **map[string]string** | Metadata is a map of key-value pairs (optional). | [optional] 
**Name** | Pointer to **string** | Unique name of the schema (set from subject, not from JSON payload). | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


