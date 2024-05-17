# NatsUserIssuanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**EventsCount** | **int64** | total download events for this issuance | 
**ExpMax** | Pointer to **int64** | highest expiry time, undefined means Unlimited (exp not set) | [optional] 
**IatMax** | **int64** | most recent time this was issued | 
**IatMin** | **int64** | first time this was issued | 
**Id** | **string** |  | 
**Iss** | **string** | issuer account public nkey | 
**Name** | **string** |  | 
**Status** | [**NatsUserIssuanceStatus**](NatsUserIssuanceStatus.md) |  | 
**Sub** | **string** | nats user public nkey | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


