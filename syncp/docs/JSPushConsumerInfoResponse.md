# JSPushConsumerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Config** | Pointer to [**JSPushConsumerConfigRequest**](JSPushConsumerConfigRequest.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NumAckPending** | **int64** |  | 
**NumPending** | **uint64** |  | 
**NumRedelivered** | **int64** |  | 
**PushBound** | Pointer to **bool** |  | [optional] 
**StreamName** | **string** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


