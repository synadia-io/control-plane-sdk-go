# AlertViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcknowledgedByAppUserID** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AcknowledgedAt** | Pointer to **time.Time** |  | [optional] 
**AlertRule** | [**AlertRuleViewResponse**](AlertRuleViewResponse.md) |  | 
**AlertRuleId** | **string** |  | 
**ClosedAt** | Pointer to **time.Time** |  | [optional] 
**ConsumerName** | Pointer to **string** |  | [optional] 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsAcknowledged** | **bool** |  | 
**IsClosed** | **bool** |  | 
**IsFiring** | **bool** |  | 
**Message** | **string** |  | 
**MetricValue** | **float64** |  | 
**SampleValues** | [][**PromSampleValue**](PromSampleValue.md) |  | 
**Severity** | [**AlertRuleSeverity**](AlertRuleSeverity.md) |  | 
**StreamName** | Pointer to **string** |  | [optional] 
**System** | Pointer to [**SystemInfo**](SystemInfo.md) |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**ThresholdValue** | **float64** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


