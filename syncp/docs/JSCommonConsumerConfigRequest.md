# JSCommonConsumerConfigRequest

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

## Methods

### NewJSCommonConsumerConfigRequest

`func NewJSCommonConsumerConfigRequest(ackPolicy AckPolicy, deliverPolicy DeliverPolicy, numReplicas int32, replayPolicy ReplayPolicy, ) *JSCommonConsumerConfigRequest`

NewJSCommonConsumerConfigRequest instantiates a new JSCommonConsumerConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSCommonConsumerConfigRequestWithDefaults

`func NewJSCommonConsumerConfigRequestWithDefaults() *JSCommonConsumerConfigRequest`

NewJSCommonConsumerConfigRequestWithDefaults instantiates a new JSCommonConsumerConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckPolicy

`func (o *JSCommonConsumerConfigRequest) GetAckPolicy() AckPolicy`

GetAckPolicy returns the AckPolicy field if non-nil, zero value otherwise.

### GetAckPolicyOk

`func (o *JSCommonConsumerConfigRequest) GetAckPolicyOk() (*AckPolicy, bool)`

GetAckPolicyOk returns a tuple with the AckPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckPolicy

`func (o *JSCommonConsumerConfigRequest) SetAckPolicy(v AckPolicy)`

SetAckPolicy sets AckPolicy field to given value.


### GetAckWait

`func (o *JSCommonConsumerConfigRequest) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *JSCommonConsumerConfigRequest) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *JSCommonConsumerConfigRequest) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *JSCommonConsumerConfigRequest) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetBackoff

`func (o *JSCommonConsumerConfigRequest) GetBackoff() []int64`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *JSCommonConsumerConfigRequest) GetBackoffOk() (*[]int64, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *JSCommonConsumerConfigRequest) SetBackoff(v []int64)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *JSCommonConsumerConfigRequest) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetDeliverPolicy

`func (o *JSCommonConsumerConfigRequest) GetDeliverPolicy() DeliverPolicy`

GetDeliverPolicy returns the DeliverPolicy field if non-nil, zero value otherwise.

### GetDeliverPolicyOk

`func (o *JSCommonConsumerConfigRequest) GetDeliverPolicyOk() (*DeliverPolicy, bool)`

GetDeliverPolicyOk returns a tuple with the DeliverPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverPolicy

`func (o *JSCommonConsumerConfigRequest) SetDeliverPolicy(v DeliverPolicy)`

SetDeliverPolicy sets DeliverPolicy field to given value.


### GetDescription

`func (o *JSCommonConsumerConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSCommonConsumerConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSCommonConsumerConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSCommonConsumerConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirect

`func (o *JSCommonConsumerConfigRequest) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *JSCommonConsumerConfigRequest) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *JSCommonConsumerConfigRequest) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *JSCommonConsumerConfigRequest) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetDurableName

`func (o *JSCommonConsumerConfigRequest) GetDurableName() string`

GetDurableName returns the DurableName field if non-nil, zero value otherwise.

### GetDurableNameOk

`func (o *JSCommonConsumerConfigRequest) GetDurableNameOk() (*string, bool)`

GetDurableNameOk returns a tuple with the DurableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurableName

`func (o *JSCommonConsumerConfigRequest) SetDurableName(v string)`

SetDurableName sets DurableName field to given value.

### HasDurableName

`func (o *JSCommonConsumerConfigRequest) HasDurableName() bool`

HasDurableName returns a boolean if a field has been set.

### GetFilterSubject

`func (o *JSCommonConsumerConfigRequest) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *JSCommonConsumerConfigRequest) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *JSCommonConsumerConfigRequest) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *JSCommonConsumerConfigRequest) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetInactiveThreshold

`func (o *JSCommonConsumerConfigRequest) GetInactiveThreshold() int64`

GetInactiveThreshold returns the InactiveThreshold field if non-nil, zero value otherwise.

### GetInactiveThresholdOk

`func (o *JSCommonConsumerConfigRequest) GetInactiveThresholdOk() (*int64, bool)`

GetInactiveThresholdOk returns a tuple with the InactiveThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveThreshold

`func (o *JSCommonConsumerConfigRequest) SetInactiveThreshold(v int64)`

SetInactiveThreshold sets InactiveThreshold field to given value.

### HasInactiveThreshold

`func (o *JSCommonConsumerConfigRequest) HasInactiveThreshold() bool`

HasInactiveThreshold returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *JSCommonConsumerConfigRequest) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *JSCommonConsumerConfigRequest) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *JSCommonConsumerConfigRequest) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *JSCommonConsumerConfigRequest) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *JSCommonConsumerConfigRequest) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *JSCommonConsumerConfigRequest) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *JSCommonConsumerConfigRequest) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *JSCommonConsumerConfigRequest) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetMemStorage

`func (o *JSCommonConsumerConfigRequest) GetMemStorage() bool`

GetMemStorage returns the MemStorage field if non-nil, zero value otherwise.

### GetMemStorageOk

`func (o *JSCommonConsumerConfigRequest) GetMemStorageOk() (*bool, bool)`

GetMemStorageOk returns a tuple with the MemStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemStorage

`func (o *JSCommonConsumerConfigRequest) SetMemStorage(v bool)`

SetMemStorage sets MemStorage field to given value.

### HasMemStorage

`func (o *JSCommonConsumerConfigRequest) HasMemStorage() bool`

HasMemStorage returns a boolean if a field has been set.

### GetName

`func (o *JSCommonConsumerConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JSCommonConsumerConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JSCommonConsumerConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JSCommonConsumerConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumReplicas

`func (o *JSCommonConsumerConfigRequest) GetNumReplicas() int32`

GetNumReplicas returns the NumReplicas field if non-nil, zero value otherwise.

### GetNumReplicasOk

`func (o *JSCommonConsumerConfigRequest) GetNumReplicasOk() (*int32, bool)`

GetNumReplicasOk returns a tuple with the NumReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReplicas

`func (o *JSCommonConsumerConfigRequest) SetNumReplicas(v int32)`

SetNumReplicas sets NumReplicas field to given value.


### GetOptStartSeq

`func (o *JSCommonConsumerConfigRequest) GetOptStartSeq() int32`

GetOptStartSeq returns the OptStartSeq field if non-nil, zero value otherwise.

### GetOptStartSeqOk

`func (o *JSCommonConsumerConfigRequest) GetOptStartSeqOk() (*int32, bool)`

GetOptStartSeqOk returns a tuple with the OptStartSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartSeq

`func (o *JSCommonConsumerConfigRequest) SetOptStartSeq(v int32)`

SetOptStartSeq sets OptStartSeq field to given value.

### HasOptStartSeq

`func (o *JSCommonConsumerConfigRequest) HasOptStartSeq() bool`

HasOptStartSeq returns a boolean if a field has been set.

### GetOptStartTime

`func (o *JSCommonConsumerConfigRequest) GetOptStartTime() string`

GetOptStartTime returns the OptStartTime field if non-nil, zero value otherwise.

### GetOptStartTimeOk

`func (o *JSCommonConsumerConfigRequest) GetOptStartTimeOk() (*string, bool)`

GetOptStartTimeOk returns a tuple with the OptStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptStartTime

`func (o *JSCommonConsumerConfigRequest) SetOptStartTime(v string)`

SetOptStartTime sets OptStartTime field to given value.

### HasOptStartTime

`func (o *JSCommonConsumerConfigRequest) HasOptStartTime() bool`

HasOptStartTime returns a boolean if a field has been set.

### SetOptStartTimeNil

`func (o *JSCommonConsumerConfigRequest) SetOptStartTimeNil(b bool)`

 SetOptStartTimeNil sets the value for OptStartTime to be an explicit nil

### UnsetOptStartTime
`func (o *JSCommonConsumerConfigRequest) UnsetOptStartTime()`

UnsetOptStartTime ensures that no value is present for OptStartTime, not even an explicit nil
### GetReplayPolicy

`func (o *JSCommonConsumerConfigRequest) GetReplayPolicy() ReplayPolicy`

GetReplayPolicy returns the ReplayPolicy field if non-nil, zero value otherwise.

### GetReplayPolicyOk

`func (o *JSCommonConsumerConfigRequest) GetReplayPolicyOk() (*ReplayPolicy, bool)`

GetReplayPolicyOk returns a tuple with the ReplayPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayPolicy

`func (o *JSCommonConsumerConfigRequest) SetReplayPolicy(v ReplayPolicy)`

SetReplayPolicy sets ReplayPolicy field to given value.


### GetSampleFreq

`func (o *JSCommonConsumerConfigRequest) GetSampleFreq() string`

GetSampleFreq returns the SampleFreq field if non-nil, zero value otherwise.

### GetSampleFreqOk

`func (o *JSCommonConsumerConfigRequest) GetSampleFreqOk() (*string, bool)`

GetSampleFreqOk returns a tuple with the SampleFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFreq

`func (o *JSCommonConsumerConfigRequest) SetSampleFreq(v string)`

SetSampleFreq sets SampleFreq field to given value.

### HasSampleFreq

`func (o *JSCommonConsumerConfigRequest) HasSampleFreq() bool`

HasSampleFreq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


