# AlertViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedByAppUserID** | Pointer to Nullable[**string**] |  | [optional] 
**Account** | Pointer to Nullable[[**AccountInfo**](AccountInfo.md)] |  | [optional] 
**AccountId** | Pointer to Nullable[**string**] |  | [optional] 
**AcknowledgedAt** | Pointer to Nullable[**string**] |  | [optional] 
**AlertRule** | [**AlertRuleViewResponse**](AlertRuleViewResponse.md) |  | 
**AlertRuleId** | **string** |  | 
**ClosedAt** | Pointer to Nullable[**string**] |  | [optional] 
**ConsumerName** | Pointer to Nullable[**string**] |  | [optional] 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsAcknowledged** | **bool** |  | 
**IsClosed** | **bool** |  | 
**IsFiring** | **bool** |  | 
**Message** | **string** |  | 
**MetricValue** | **float64** |  | 
**SampleValues** | [][**PromSampleValue**](PromSampleValue.md) |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**StreamName** | Pointer to Nullable[**string**] |  | [optional] 
**System** | Pointer to Nullable[[**SystemInfo**](SystemInfo.md)] |  | [optional] 
**SystemId** | Pointer to Nullable[**string**] |  | [optional] 
**ThresholdValue** | **float64** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


