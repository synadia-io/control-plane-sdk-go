# JSPullConsumerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to [**ClusterInfo**](ClusterInfo.md) |  | [optional] 
**Config** | Pointer to [**JSPullConsumerConfigRequest**](JSPullConsumerConfigRequest.md) |  | [optional] 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NumAckPending** | **int32** |  | 
**NumPending** | **int32** |  | 
**NumRedelivered** | **int32** |  | 
**NumWaiting** | **int32** |  | 
**StreamName** | **string** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


