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
**NumAckPending** | **int64** |  | 
**NumPending** | **uint64** |  | 
**NumRedelivered** | **int64** |  | 
**NumWaiting** | **int64** |  | 
**PushBound** | Pointer to **bool** |  | [optional] 
**StreamName** | **string** |  | 
**Ts** | **time.Time** |  | 

## Methods


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


