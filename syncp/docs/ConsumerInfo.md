# ConsumerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckFloor** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Cluster** | Pointer to Nullable[[**ClusterInfo**](ClusterInfo.md)] |  | [optional] 
**Config** | [**ConsumerConfig**](ConsumerConfig.md) |  | 
**Created** | **time.Time** |  | 
**Delivered** | [**SequenceInfo**](SequenceInfo.md) |  | 
**Name** | **string** |  | 
**NumAckPending** | **int32** |  | 
**NumPending** | **int32** |  | 
**NumRedelivered** | **int32** |  | 
**NumWaiting** | **int32** |  | 
**PushBound** | Pointer to **bool** |  | [optional] 
**StreamName** | **string** |  | 
**Ts** | **time.Time** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


