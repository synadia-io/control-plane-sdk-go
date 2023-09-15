# JSPushConsumerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliverGroup** | Pointer to **string** |  | [optional] 
**DeliverSubject** | Pointer to **string** | Push based options. | [optional] 
**FlowControl** | Pointer to **bool** |  | [optional] 
**HeadersOnly** | Pointer to **bool** |  | [optional] 
**IdleHeartbeat** | Pointer to **int64** |  | [optional] 
**RateLimitBps** | Pointer to **int32** |  | [optional] 
**AckPolicy** | [**AckPolicy**](AckPolicy.md) |  | 
**AckWait** | Pointer to **int64** |  | [optional] 
**Backoff** | Pointer to **[]int64** |  | [optional] 
**DeliverPolicy** | [**DeliverPolicy**](DeliverPolicy.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**Direct** | Pointer to **bool** |  | [optional] 
**DurableName** | Pointer to **string** |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**InactiveThreshold** | Pointer to **int64** |  | [optional] 
**MaxAckPending** | Pointer to **int32** |  | [optional] 
**MaxDeliver** | Pointer to **int32** |  | [optional] 
**MemStorage** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumReplicas** | **int32** |  | 
**OptStartSeq** | Pointer to **int32** |  | [optional] 
**OptStartTime** | Pointer to **NullableString** |  | [optional] 
**ReplayPolicy** | [**ReplayPolicy**](ReplayPolicy.md) |  | 
**SampleFreq** | Pointer to **string** |  | [optional] 

## Methods

### NewJSPushConsumerConfigRequest

`func NewJSPushConsumerConfigRequest(ackPolicy AckPolicy, deliverPolicy DeliverPolicy, numReplicas int32, replayPolicy ReplayPolicy, ) *JSPushConsumerConfigRequest`

NewJSPushConsumerConfigRequest instantiates a new JSPushConsumerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSPushConsumerConfigRequestWithDefaults

`func NewJSPushConsumerConfigRequestWithDefaults() *JSPushConsumerConfigRequest`

NewJSPushConsumerConfigRequestWithDefaults instantiates a new JSPushConsumerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliverGroup

`func (o *JSPushConsumerConfigRequest) GetDeliverGroup() string`

GetDeliverGroup returns the DeliverGroup field if non-nil, zero value otherwise.

### GetDeliverGroupOk

`func (o *JSPushConsumerConfigRequest) GetDeliverGroupOk() (*string, bool)`

GetDeliverGroupOk returns a tuple with the DeliverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverGroup

`func (o *JSPushConsumerConfigRequest) SetDeliverGroup(v string)`

SetDeliverGroup sets DeliverGroup field to given value.

### HasDeliverGroup

`func (o *JSPushConsumerConfigRequest) HasDeliverGroup() bool`

HasDeliverGroup returns a boolean if a field has been set.

### GetDeliverSubject

`func (o *JSPushConsumerConfigRequest) GetDeliverSubject() string`

GetDeliverSubject returns the DeliverSubject field if non-nil, zero value otherwise.

### GetDeliverSubjectOk

`func (o *JSPushConsumerConfigRequest) GetDeliverSubjectOk() (*string, bool)`

GetDeliverSubjectOk returns a tuple with the DeliverSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSubject

`func (o *JSPushConsumerConfigRequest) SetDeliverSubject(v string)`

SetDeliverSubject sets DeliverSubject field to given value.

### HasDeliverSubject

`func (o *JSPushConsumerConfigRequest) HasDeliverSubject() bool`

HasDeliverSubject returns a boolean if a field has been set.

### GetFlowControl

`func (o *JSPushConsumerConfigRequest) GetFlowControl() bool`

GetFlowControl returns the FlowControl field if non-nil, zero value otherwise.

### GetFlowControlOk

`func (o *JSPushConsumerConfigRequest) GetFlowControlOk() (*bool, bool)`

GetFlowControlOk returns a tuple with the FlowControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowControl

`func (o *JSPushConsumerConfigRequest) SetFlowControl(v bool)`

SetFlowControl sets FlowControl field to given value.

### HasFlowControl

`func (o *JSPushConsumerConfigRequest) HasFlowControl() bool`

HasFlowControl returns a boolean if a field has been set.

### GetHeadersOnly

`func (o *JSPushConsumerConfigRequest) GetHeadersOnly() bool`

GetHeadersOnly returns the HeadersOnly field if non-nil, zero value otherwise.

### GetHeadersOnlyOk

`func (o *JSPushConsumerConfigRequest) GetHeadersOnlyOk() (*bool, bool)`

GetHeadersOnlyOk returns a tuple with the HeadersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersOnly

`func (o *JSPushConsumerConfigRequest) SetHeadersOnly(v bool)`

SetHeadersOnly sets HeadersOnly field to given value.

### HasHeadersOnly

`func (o *JSPushConsumerConfigRequest) HasHeadersOnly() bool`

HasHeadersOnly returns a boolean if a field has been set.

### GetIdleHeartbeat

`func (o *JSPushConsumerConfigRequest) GetIdleHeartbeat() int64`

GetIdleHeartbeat returns the IdleHeartbeat field if non-nil, zero value otherwise.

### GetIdleHeartbeatOk

`func (o *JSPushConsumerConfigRequest) GetIdleHeartbeatOk() (*int64, bool)`

GetIdleHeartbeatOk returns a tuple with the IdleHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleHeartbeat

`func (o *JSPushConsumerConfigRequest) SetIdleHeartbeat(v int64)`

SetIdleHeartbeat sets IdleHeartbeat field to given value.

### HasIdleHeartbeat

`func (o *JSPushConsumerConfigRequest) HasIdleHeartbeat() bool`

HasIdleHeartbeat returns a boolean if a field has been set.

### GetRateLimitBps

`func (o *JSPushConsumerConfigRequest) GetRateLimitBps() int32`

GetRateLimitBps returns the RateLimitBps field if non-nil, zero value otherwise.

### GetRateLimitBpsOk

`func (o *JSPushConsumerConfigRequest) GetRateLimitBpsOk() (*int32, bool)`

GetRateLimitBpsOk returns a tuple with the RateLimitBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitBps

`func (o *JSPushConsumerConfigRequest) SetRateLimitBps(v int32)`

SetRateLimitBps sets RateLimitBps field to given value.

### HasRateLimitBps

`func (o *JSPushConsumerConfigRequest) HasRateLimitBps() bool`

HasRateLimitBps returns a boolean if a field has been set.

### GetAckPolicy

`func (o *JSPushConsumerConfigRequest) GetAckPolicy() AckPolicy`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *JSPushConsumerConfigRequest) GetAckPolicyOk() (*AckPolicy, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *JSPushConsumerConfigRequest) SetAckPolicy(v AckPolicy)`

SetAckPolicy sets AckPolicy field to given value.


### GetAckWait

`func (o *JSPushConsumerConfigRequest) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *JSPushConsumerConfigRequest) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *JSPushConsumerConfigRequest) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *JSPushConsumerConfigRequest) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetBackoff

`func (o *JSPushConsumerConfigRequest) GetBackoff() []int64`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *JSPushConsumerConfigRequest) GetBackoffOk() (*[]int64, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *JSPushConsumerConfigRequest) SetBackoff(v []int64)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *JSPushConsumerConfigRequest) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *JSPushConsumerConfigRequest) GetDeliverPolicy() DeliverPolicy`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *JSPushConsumerConfigRequest) GetDeliverPolicyOk() (*DeliverPolicy, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *JSPushConsumerConfigRequest) SetDeliverPolicy(v DeliverPolicy)`

SetDeliverPolicy sets DeliverPolicy field to given value.


### GetDescription

`func (o *JSPushConsumerConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSPushConsumerConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSPushConsumerConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSPushConsumerConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirect

`func (o *JSPushConsumerConfigRequest) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *JSPushConsumerConfigRequest) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *JSPushConsumerConfigRequest) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *JSPushConsumerConfigRequest) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetDurableName

`func (o *JSPushConsumerConfigRequest) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *JSPushConsumerConfigRequest) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *JSPushConsumerConfigRequest) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *JSPushConsumerConfigRequest) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *JSPushConsumerConfigRequest) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *JSPushConsumerConfigRequest) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *JSPushConsumerConfigRequest) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *JSPushConsumerConfigRequest) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetInactiveThreshold

`func (o *JSPushConsumerConfigRequest) GetInactiveThreshold() int64`

GetInactiveThreshold returns the InactiveThreshold field if non-nil, zero value otherwise.

### GetInactiveThresholdOk

`func (o *JSPushConsumerConfigRequest) GetInactiveThresholdOk() (*int64, bool)`

GetInactiveThresholdOk returns a tuple with the InactiveThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveThreshold

`func (o *JSPushConsumerConfigRequest) SetInactiveThreshold(v int64)`

SetInactiveThreshold sets InactiveThreshold field to given value.

### HasInactiveThreshold

`func (o *JSPushConsumerConfigRequest) HasInactiveThreshold() bool`

HasInactiveThreshold returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *JSPushConsumerConfigRequest) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *JSPushConsumerConfigRequest) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *JSPushConsumerConfigRequest) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *JSPushConsumerConfigRequest) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *JSPushConsumerConfigRequest) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *JSPushConsumerConfigRequest) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *JSPushConsumerConfigRequest) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *JSPushConsumerConfigRequest) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetMemStorage

`func (o *JSPushConsumerConfigRequest) GetMemStorage() bool`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *JSPushConsumerConfigRequest) GetMemStorageOk() (*bool, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *JSPushConsumerConfigRequest) SetMemStorage(v bool)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *JSPushConsumerConfigRequest) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetName

`func (o *JSPushConsumerConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSPushConsumerConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSPushConsumerConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JSPushConsumerConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumReplicas

`func (o *JSPushConsumerConfigRequest) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *JSPushConsumerConfigRequest) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *JSPushConsumerConfigRequest) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetOptStartSeq

`func (o *JSPushConsumerConfigRequest) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *JSPushConsumerConfigRequest) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *JSPushConsumerConfigRequest) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *JSPushConsumerConfigRequest) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *JSPushConsumerConfigRequest) GetOptStartTime() string`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *JSPushConsumerConfigRequest) GetOptStartTimeOk() (*string, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *JSPushConsumerConfigRequest) SetOptStartTime(v string)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *JSPushConsumerConfigRequest) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### SetOptStartTimeNil

`func (o *JSPushConsumerConfigRequest) SetOptStartTimeNil(b bool)`

 SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil

### UnsetOptStartTime
`func (o *JSPushConsumerConfigRequest) UnsetOptStartTime()`

UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
### GetReplayPolicy

`func (o *JSPushConsumerConfigRequest) GetReplayPolicy() ReplayPolicy`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *JSPushConsumerConfigRequest) GetReplayPolicyOk() (*ReplayPolicy, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *JSPushConsumerConfigRequest) SetReplayPolicy(v ReplayPolicy)`

SetReplayPolicy sets ReplayPolicy field to given value.


### GetSampleFreq

`func (o *JSPushConsumerConfigRequest) GetSampleFreq() string`

GetSampleFreq returns the SampleFreq field if non-nil, zero value otherwise.

### GetSampleFreqOk

`func (o *JSPushConsumerConfigRequest) GetSampleFreqOk() (*string, bool)`

GetSampleFreqOk returns a tuple with the SampleFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFreq

`func (o *JSPushConsumerConfigRequest) SetSampleFreq(v string)`

SetSampleFreq sets SampleFreq field to given value.

### HasSampleFreq

`func (o *JSPushConsumerConfigRequest) HasSampleFreq() bool`

HasSampleFreq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


