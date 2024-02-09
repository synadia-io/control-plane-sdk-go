# NatsUserIssuanceViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | **string** | sha256 sum of (NatsUserId,Iss,Sub,Name,json_encode(Nats)) If a credential is downloaded and results in a unique checksum then a new issuance record is created. If a credential&#39;s checksum matches an existing record, a new event is appended to the existing record. | 
**Created** | **time.Time** |  | 
**Events** | [][**NatsUserIssuanceEvent**](NatsUserIssuanceEvent.md) | trail of download events | 
**ExpMax** | Pointer to **int64** | highest expiry time, undefined means Unlimited (exp not set) | [optional] 
**IatMax** | **int64** | most recent time this was issued | 
**IatMin** | **int64** | first time this was issued | 
**Id** | **string** |  | 
**Iss** | **string** | issuer account public nkey | 
**Name** | **string** |  | 
**Nats** | [**User**](User.md) |  | 
**Status** | [**NatsUserIssuanceStatus**](NatsUserIssuanceStatus.md) |  | 
**Sub** | **string** | nats user public nkey | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


