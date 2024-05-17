# ConsumerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckPolicy** | [**AckPolicy**](AckPolicy.md) |  | 
**AckWait** | Pointer to **int64** |  | [optional] 
**Backoff** | Pointer to []**int64** |  | [optional] 
**DeliverGroup** | Pointer to **string** |  | [optional] 
**DeliverPolicy** | [**DeliverPolicy**](DeliverPolicy.md) |  | 
**DeliverSubject** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Direct** | Pointer to **bool** | Don&#39;t add to general clients. | [optional] 
**DurableName** | Pointer to **string** |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**FilterSubjects** | Pointer to []**string** |  | [optional] 
**FlowControl** | Pointer to **bool** |  | [optional] 
**HeadersOnly** | Pointer to **bool** |  | [optional] 
**IdleHeartbeat** | Pointer to **int64** |  | [optional] 
**InactiveThreshold** | Pointer to **int64** |  | [optional] 
**MaxAckPending** | Pointer to **int64** |  | [optional] 
**MaxBatch** | Pointer to **int64** |  | [optional] 
**MaxBytes** | Pointer to **int64** |  | [optional] 
**MaxDeliver** | Pointer to **int64** |  | [optional] 
**MaxExpires** | Pointer to **int64** |  | [optional] 
**MaxWaiting** | Pointer to **int64** |  | [optional] 
**MemStorage** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata is additional metadata for the Consumer. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumReplicas** | **int64** |  | 
**OptStartSeq** | Pointer to **uint64** |  | [optional] 
**OptStartTime** | Pointer to Nullable[**string**] |  | [optional] 
**RateLimitBps** | Pointer to **uint64** |  | [optional] 
**ReplayPolicy** | [**ReplayPolicy**](ReplayPolicy.md) |  | 
**SampleFreq** | Pointer to **string** |  | [optional] 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


