# JSPullConsumerConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**MaxBatch** | Pointer to **int32** |  | [optional] 
**MaxBytes** | Pointer to **int32** |  | [optional] 
**MaxExpires** | Pointer to **int64** |  | [optional] 
**MaxWaiting** | Pointer to **int32** | Pull based options. | [optional] 

## Methods

### NewJSPullConsumerConfigRequest

`func NewJSPullConsumerConfigRequest(ackPolicy AckPolicy, deliverPolicy DeliverPolicy, numReplicas int32, replayPolicy ReplayPolicy, ) *JSPullConsumerConfigRequest`

NewJSPullConsumerConfigRequest instantiates a new JSPullConsumerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSPullConsumerConfigRequestWithDefaults

`func NewJSPullConsumerConfigRequestWithDefaults() *JSPullConsumerConfigRequest`

NewJSPullConsumerConfigRequestWithDefaults instantiates a new JSPullConsumerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckPolicy

`func (o *JSPullConsumerConfigRequest) GetAckPolicy() AckPolicy`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *JSPullConsumerConfigRequest) GetAckPolicyOk() (*AckPolicy, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *JSPullConsumerConfigRequest) SetAckPolicy(v AckPolicy)`

SetAckPolicy sets AckPolicy field to given value.


### GetAckWait

`func (o *JSPullConsumerConfigRequest) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *JSPullConsumerConfigRequest) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *JSPullConsumerConfigRequest) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *JSPullConsumerConfigRequest) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetBackoff

`func (o *JSPullConsumerConfigRequest) GetBackoff() []int64`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *JSPullConsumerConfigRequest) GetBackoffOk() (*[]int64, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *JSPullConsumerConfigRequest) SetBackoff(v []int64)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *JSPullConsumerConfigRequest) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *JSPullConsumerConfigRequest) GetDeliverPolicy() DeliverPolicy`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *JSPullConsumerConfigRequest) GetDeliverPolicyOk() (*DeliverPolicy, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *JSPullConsumerConfigRequest) SetDeliverPolicy(v DeliverPolicy)`

SetDeliverPolicy sets DeliverPolicy field to given value.


### GetDescription

`func (o *JSPullConsumerConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSPullConsumerConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSPullConsumerConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSPullConsumerConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirect

`func (o *JSPullConsumerConfigRequest) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *JSPullConsumerConfigRequest) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *JSPullConsumerConfigRequest) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *JSPullConsumerConfigRequest) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetDurableName

`func (o *JSPullConsumerConfigRequest) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *JSPullConsumerConfigRequest) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *JSPullConsumerConfigRequest) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *JSPullConsumerConfigRequest) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *JSPullConsumerConfigRequest) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *JSPullConsumerConfigRequest) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *JSPullConsumerConfigRequest) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *JSPullConsumerConfigRequest) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetInactiveThreshold

`func (o *JSPullConsumerConfigRequest) GetInactiveThreshold() int64`

GetInactiveThreshold returns the InactiveThreshold field if non-nil, zero value otherwise.

### GetInactiveThresholdOk

`func (o *JSPullConsumerConfigRequest) GetInactiveThresholdOk() (*int64, bool)`

GetInactiveThresholdOk returns a tuple with the InactiveThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveThreshold

`func (o *JSPullConsumerConfigRequest) SetInactiveThreshold(v int64)`

SetInactiveThreshold sets InactiveThreshold field to given value.

### HasInactiveThreshold

`func (o *JSPullConsumerConfigRequest) HasInactiveThreshold() bool`

HasInactiveThreshold returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *JSPullConsumerConfigRequest) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *JSPullConsumerConfigRequest) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *JSPullConsumerConfigRequest) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *JSPullConsumerConfigRequest) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *JSPullConsumerConfigRequest) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *JSPullConsumerConfigRequest) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *JSPullConsumerConfigRequest) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *JSPullConsumerConfigRequest) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetMemStorage

`func (o *JSPullConsumerConfigRequest) GetMemStorage() bool`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *JSPullConsumerConfigRequest) GetMemStorageOk() (*bool, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *JSPullConsumerConfigRequest) SetMemStorage(v bool)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *JSPullConsumerConfigRequest) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetName

`func (o *JSPullConsumerConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSPullConsumerConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSPullConsumerConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JSPullConsumerConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumReplicas

`func (o *JSPullConsumerConfigRequest) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *JSPullConsumerConfigRequest) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *JSPullConsumerConfigRequest) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetOptStartSeq

`func (o *JSPullConsumerConfigRequest) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *JSPullConsumerConfigRequest) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *JSPullConsumerConfigRequest) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *JSPullConsumerConfigRequest) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *JSPullConsumerConfigRequest) GetOptStartTime() string`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *JSPullConsumerConfigRequest) GetOptStartTimeOk() (*string, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *JSPullConsumerConfigRequest) SetOptStartTime(v string)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *JSPullConsumerConfigRequest) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### SetOptStartTimeNil

`func (o *JSPullConsumerConfigRequest) SetOptStartTimeNil(b bool)`

 SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil

### UnsetOptStartTime
`func (o *JSPullConsumerConfigRequest) UnsetOptStartTime()`

UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
### GetReplayPolicy

`func (o *JSPullConsumerConfigRequest) GetReplayPolicy() ReplayPolicy`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *JSPullConsumerConfigRequest) GetReplayPolicyOk() (*ReplayPolicy, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *JSPullConsumerConfigRequest) SetReplayPolicy(v ReplayPolicy)`

SetReplayPolicy sets ReplayPolicy field to given value.


### GetSampleFreq

`func (o *JSPullConsumerConfigRequest) GetSampleFreq() string`

GetSampleFreq returns the SampleFreq field if non-nil, zero value otherwise.

### GetSampleFreqOk

`func (o *JSPullConsumerConfigRequest) GetSampleFreqOk() (*string, bool)`

GetSampleFreqOk returns a tuple with the SampleFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFreq

`func (o *JSPullConsumerConfigRequest) SetSampleFreq(v string)`

SetSampleFreq sets SampleFreq field to given value.

### HasSampleFreq

`func (o *JSPullConsumerConfigRequest) HasSampleFreq() bool`

HasSampleFreq returns a boolean if a field has been set.

### GetMaxBatch

`func (o *JSPullConsumerConfigRequest) GetMaxBatch() int32`

GetMaxBatch returns the MaxBatch field if non-nil, zero value otherwise.

### GetMaxBatchOk

`func (o *JSPullConsumerConfigRequest) GetMaxBatchOk() (*int32, bool)`

GetMaxBatchOk returns a tuple with the MaxBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBatch

`func (o *JSPullConsumerConfigRequest) SetMaxBatch(v int32)`

SetMaxBatch sets MaxBatch field to given value.

### HasMaxBatch

`func (o *JSPullConsumerConfigRequest) HasMaxBatch() bool`

HasMaxBatch returns a boolean if a field has been set.

### GetMaxBytes

`func (o *JSPullConsumerConfigRequest) GetMaxBytes() int32`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *JSPullConsumerConfigRequest) GetMaxBytesOk() (*int32, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *JSPullConsumerConfigRequest) SetMaxBytes(v int32)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *JSPullConsumerConfigRequest) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.

### GetMaxExpires

`func (o *JSPullConsumerConfigRequest) GetMaxExpires() int64`

GetMaxExpires returns the MaxExpires field if non-nil, zero value otherwise.

### GetMaxExpiresOk

`func (o *JSPullConsumerConfigRequest) GetMaxExpiresOk() (*int64, bool)`

GetMaxExpiresOk returns a tuple with the MaxExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExpires

`func (o *JSPullConsumerConfigRequest) SetMaxExpires(v int64)`

SetMaxExpires sets MaxExpires field to given value.

### HasMaxExpires

`func (o *JSPullConsumerConfigRequest) HasMaxExpires() bool`

HasMaxExpires returns a boolean if a field has been set.

### GetMaxWaiting

`func (o *JSPullConsumerConfigRequest) GetMaxWaiting() int32`

GetMaxWaiting returns the MaxWaiting field if non-nil, zero value otherwise.

### GetMaxWaitingOk

`func (o *JSPullConsumerConfigRequest) GetMaxWaitingOk() (*int32, bool)`

GetMaxWaitingOk returns a tuple with the MaxWaiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaiting

`func (o *JSPullConsumerConfigRequest) SetMaxWaiting(v int32)`

SetMaxWaiting sets MaxWaiting field to given value.

### HasMaxWaiting

`func (o *JSPullConsumerConfigRequest) HasMaxWaiting() bool`

HasMaxWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


