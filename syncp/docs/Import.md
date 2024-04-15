# Import

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**LocalSubject** | Pointer to **string** | Local subject used to subscribe (for streams) and publish (for services) to. This value only needs setting if you want to change the value of Subject. If the value of Subject ends in &gt; then LocalSubject needs to end in &gt; as well. LocalSubject can contain $&lt;number&gt; wildcard references where number references the nth wildcard in Subject. The sum of wildcard reference and * tokens needs to match the number of * token in Subject. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Share** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** | Subject field in an import is always from the perspective of the initial publisher - in the case of a stream it is the account owning the stream (the exporter), and in the case of a service it is the account making the request (the importer). | [optional] 
**To** | Pointer to **string** | Deprecated: use LocalSubject instead To field in an import is always from the perspective of the subscriber in the case of a stream it is the client of the stream (the importer), from the perspective of a service, it is the subscription waiting for requests (the exporter). If the field is empty, it will default to the value in the Subject field. | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ExportType**](ExportType.md) |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


