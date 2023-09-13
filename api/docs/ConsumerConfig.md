# ConsumerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckPolicy** | [**AckPolicy**](AckPolicy.md) |  | 
**AckWait** | Pointer to **int64** |  | [optional] 
**Backoff** | Pointer to **[]int64** |  | [optional] 
**DeliverGroup** | Pointer to **string** |  | [optional] 
**DeliverPolicy** | [**DeliverPolicy**](DeliverPolicy.md) |  | 
**DeliverSubject** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Direct** | Pointer to **bool** | Don&#39;t add to general clients. | [optional] 
**DurableName** | Pointer to **string** |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**FlowControl** | Pointer to **bool** |  | [optional] 
**HeadersOnly** | Pointer to **bool** |  | [optional] 
**IdleHeartbeat** | Pointer to **int64** |  | [optional] 
**InactiveThreshold** | Pointer to **int64** |  | [optional] 
**MaxAckPending** | Pointer to **int32** |  | [optional] 
**MaxBatch** | Pointer to **int32** |  | [optional] 
**MaxBytes** | Pointer to **int32** |  | [optional] 
**MaxDeliver** | Pointer to **int32** |  | [optional] 
**MaxExpires** | Pointer to **int64** |  | [optional] 
**MaxWaiting** | Pointer to **int32** |  | [optional] 
**MemStorage** | Pointer to **bool** |  | [optional] 
**NumReplicas** | **int32** |  | 
**OptStartSeq** | Pointer to **int32** |  | [optional] 
**OptStartTime** | Pointer to **NullableString** |  | [optional] 
**RateLimitBps** | Pointer to **int32** |  | [optional] 
**ReplayPolicy** | [**ReplayPolicy**](ReplayPolicy.md) |  | 
**SampleFreq** | Pointer to **string** |  | [optional] 

## Methods

### NewConsumerConfig

`func NewConsumerConfig(ackPolicy AckPolicy, deliverPolicy DeliverPolicy, numReplicas int32, replayPolicy ReplayPolicy, ) *ConsumerConfig`

NewConsumerConfig instantiates a new ConsumerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerConfigWithDefaults

`func NewConsumerConfigWithDefaults() *ConsumerConfig`

NewConsumerConfigWithDefaults instantiates a new ConsumerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckPolicy

`func (o *ConsumerConfig) GetAckPolicy() AckPolicy`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *ConsumerConfig) GetAckPolicyOk() (*AckPolicy, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *ConsumerConfig) SetAckPolicy(v AckPolicy)`

SetAckPolicy sets AckPolicy field to given value.


### GetAckWait

`func (o *ConsumerConfig) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *ConsumerConfig) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *ConsumerConfig) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *ConsumerConfig) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetBackoff

`func (o *ConsumerConfig) GetBackoff() []int64`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *ConsumerConfig) GetBackoffOk() (*[]int64, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *ConsumerConfig) SetBackoff(v []int64)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *ConsumerConfig) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetDeliverGroup

`func (o *ConsumerConfig) GetDeliverGroup() string`

GetDeliverGroup returns the DeliverGroup field if non-nil, zero value otherwise.

### GetDeliverGroupOk

`func (o *ConsumerConfig) GetDeliverGroupOk() (*string, bool)`

GetDeliverGroupOk returns a tuple with the DeliverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverGroup

`func (o *ConsumerConfig) SetDeliverGroup(v string)`

SetDeliverGroup sets DeliverGroup field to given value.

### HasDeliverGroup

`func (o *ConsumerConfig) HasDeliverGroup() bool`

HasDeliverGroup returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *ConsumerConfig) GetDeliverPolicy() DeliverPolicy`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *ConsumerConfig) GetDeliverPolicyOk() (*DeliverPolicy, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *ConsumerConfig) SetDeliverPolicy(v DeliverPolicy)`

SetDeliverPolicy sets DeliverPolicy field to given value.


### GetDeliverSubject

`func (o *ConsumerConfig) GetDeliverSubject() string`

GetDeliverSubject returns the DeliverSubject field if non-nil, zero value otherwise.

### GetDeliverSubjectOk

`func (o *ConsumerConfig) GetDeliverSubjectOk() (*string, bool)`

GetDeliverSubjectOk returns a tuple with the DeliverSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubject

`func (o *ConsumerConfig) SetDeliverSubject(v string)`

SetDeliverSubject sets DeliverSubject field to given value.

### HasDeliverSubject

`func (o *ConsumerConfig) HasDeliverSubject() bool`

HasDeliverSubject returns a boolean if a field has been set.

### GetDescription

`func (o *ConsumerConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsumerConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsumerConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsumerConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirect

`func (o *ConsumerConfig) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *ConsumerConfig) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *ConsumerConfig) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *ConsumerConfig) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetDurableName

`func (o *ConsumerConfig) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *ConsumerConfig) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *ConsumerConfig) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *ConsumerConfig) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *ConsumerConfig) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *ConsumerConfig) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *ConsumerConfig) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *ConsumerConfig) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetFlowControl

`func (o *ConsumerConfig) GetFlowControl() bool`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *ConsumerConfig) GetFlowControlOk() (*bool, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *ConsumerConfig) SetFlowControl(v bool)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *ConsumerConfig) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetHeadersOnly

`func (o *ConsumerConfig) GetHeadersOnly() bool`

GetHeadersOnly returns the HeadersOnly field if non-nil, zero value otherwise.

### GetHeadersOnlyOk

`func (o *ConsumerConfig) GetHeadersOnlyOk() (*bool, bool)`

GetHeadersOnlyOk returns a tuple with the HeadersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersOnly

`func (o *ConsumerConfig) SetHeadersOnly(v bool)`

SetHeadersOnly sets HeadersOnly field to given value.

### HasHeadersOnly

`func (o *ConsumerConfig) HasHeadersOnly() bool`

HasHeadersOnly returns a boolean if a field has been set.

### GetIdleHeartbeat

`func (o *ConsumerConfig) GetIdleHeartbeat() int64`

GetIdleHeartbeat returns the IdleHeartbeat field if non-nil, zero value otherwise.

### GetIdleHeartbeatOk

`func (o *ConsumerConfig) GetIdleHeartbeatOk() (*int64, bool)`

GetIdleHeartbeatOk returns a tuple with the IdleHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleHeartbeat

`func (o *ConsumerConfig) SetIdleHeartbeat(v int64)`

SetIdleHeartbeat sets IdleHeartbeat field to given value.

### HasIdleHeartbeat

`func (o *ConsumerConfig) HasIdleHeartbeat() bool`

HasIdleHeartbeat returns a boolean if a field has been set.

### GetInactiveThreshold

`func (o *ConsumerConfig) GetInactiveThreshold() int64`

GetInactiveThreshold returns the InactiveThreshold field if non-nil, zero value otherwise.

### GetInactiveThresholdOk

`func (o *ConsumerConfig) GetInactiveThresholdOk() (*int64, bool)`

GetInactiveThresholdOk returns a tuple with the InactiveThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveThreshold

`func (o *ConsumerConfig) SetInactiveThreshold(v int64)`

SetInactiveThreshold sets InactiveThreshold field to given value.

### HasInactiveThreshold

`func (o *ConsumerConfig) HasInactiveThreshold() bool`

HasInactiveThreshold returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *ConsumerConfig) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *ConsumerConfig) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *ConsumerConfig) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *ConsumerConfig) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxBatch

`func (o *ConsumerConfig) GetMaxBatch() int32`

GetMaxBatch returns the MaxBatch field if non-nil, zero value otherwise.

### GetMaxBatchOk

`func (o *ConsumerConfig) GetMaxBatchOk() (*int32, bool)`

GetMaxBatchOk returns a tuple with the MaxBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBatch

`func (o *ConsumerConfig) SetMaxBatch(v int32)`

SetMaxBatch sets MaxBatch field to given value.

### HasMaxBatch

`func (o *ConsumerConfig) HasMaxBatch() bool`

HasMaxBatch returns a boolean if a field has been set.

### GetMaxBytes

`func (o *ConsumerConfig) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *ConsumerConfig) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *ConsumerConfig) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *ConsumerConfig) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *ConsumerConfig) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *ConsumerConfig) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *ConsumerConfig) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *ConsumerConfig) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetMaxExpires

`func (o *ConsumerConfig) GetMaxExpires() int64`

GetMaxExpires returns the MaxExpires field if non-nil, zero value otherwise.

### GetMaxExpiresOk

`func (o *ConsumerConfig) GetMaxExpiresOk() (*int64, bool)`

GetMaxExpiresOk returns a tuple with the MaxExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExpires

`func (o *ConsumerConfig) SetMaxExpires(v int64)`

SetMaxExpires sets MaxExpires field to given value.

### HasMaxExpires

`func (o *ConsumerConfig) HasMaxExpires() bool`

HasMaxExpires returns a boolean if a field has been set.

### GetMaxWaiting

`func (o *ConsumerConfig) GetMaxWaiting() int32`

GetMaxWaiting returns the MaxWaiting field if non-nil, zero value otherwise.

### GetMaxWaitingOk

`func (o *ConsumerConfig) GetMaxWaitingOk() (*int32, bool)`

GetMaxWaitingOk returns a tuple with the MaxWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaiting

`func (o *ConsumerConfig) SetMaxWaiting(v int32)`

SetMaxWaiting sets MaxWaiting field to given value.

### HasMaxWaiting

`func (o *ConsumerConfig) HasMaxWaiting() bool`

HasMaxWaiting returns a boolean if a field has been set.

### GetMemStorage

`func (o *ConsumerConfig) GetMemStorage() bool`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *ConsumerConfig) GetMemStorageOk() (*bool, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *ConsumerConfig) SetMemStorage(v bool)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *ConsumerConfig) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetNumReplicas

`func (o *ConsumerConfig) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *ConsumerConfig) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *ConsumerConfig) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetOptStartSeq

`func (o *ConsumerConfig) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *ConsumerConfig) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *ConsumerConfig) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *ConsumerConfig) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *ConsumerConfig) GetOptStartTime() string`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *ConsumerConfig) GetOptStartTimeOk() (*string, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *ConsumerConfig) SetOptStartTime(v string)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *ConsumerConfig) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### SetOptStartTimeNil

`func (o *ConsumerConfig) SetOptStartTimeNil(b bool)`

 SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil

### UnsetOptStartTime
`func (o *ConsumerConfig) UnsetOptStartTime()`

UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
### GetRateLimitBps

`func (o *ConsumerConfig) GetRateLimitBps() int32`

GetRateLimitBps returns the RateLimitBps field if non-nil, zero value otherwise.

### GetRateLimitBpsOk

`func (o *ConsumerConfig) GetRateLimitBpsOk() (*int32, bool)`

GetRateLimitBpsOk returns a tuple with the RateLimitBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitBps

`func (o *ConsumerConfig) SetRateLimitBps(v int32)`

SetRateLimitBps sets RateLimitBps field to given value.

### HasRateLimitBps

`func (o *ConsumerConfig) HasRateLimitBps() bool`

HasRateLimitBps returns a boolean if a field has been set.

### GetReplayPolicy

`func (o *ConsumerConfig) GetReplayPolicy() ReplayPolicy`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *ConsumerConfig) GetReplayPolicyOk() (*ReplayPolicy, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *ConsumerConfig) SetReplayPolicy(v ReplayPolicy)`

SetReplayPolicy sets ReplayPolicy field to given value.


### GetSampleFreq

`func (o *ConsumerConfig) GetSampleFreq() string`

GetSampleFreq returns the SampleFreq field if non-nil, zero value otherwise.

### GetSampleFreqOk

`func (o *ConsumerConfig) GetSampleFreqOk() (*string, bool)`

GetSampleFreqOk returns a tuple with the SampleFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFreq

`func (o *ConsumerConfig) SetSampleFreq(v string)`

SetSampleFreq sets SampleFreq field to given value.

### HasSampleFreq

`func (o *ConsumerConfig) HasSampleFreq() bool`

HasSampleFreq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


