# JSPushConsumerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckWait** | Pointer to **int64** |  | [optional] 
**Backoff** | Pointer to **[]int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FilterSubject** | Pointer to **string** |  | [optional] 
**HeadersOnly** | Pointer to **bool** |  | [optional] 
**MaxAckPending** | Pointer to **int32** |  | [optional] 
**MaxDeliver** | Pointer to **int32** |  | [optional] 
**SampleFreq** | Pointer to **string** |  | [optional] 

## Methods

### NewJSPushConsumerUpdateRequest

`func NewJSPushConsumerUpdateRequest() *JSPushConsumerUpdateRequest`

NewJSPushConsumerUpdateRequest instantiates a new JSPushConsumerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSPushConsumerUpdateRequestWithDefaults

`func NewJSPushConsumerUpdateRequestWithDefaults() *JSPushConsumerUpdateRequest`

NewJSPushConsumerUpdateRequestWithDefaults instantiates a new JSPushConsumerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckWait

`func (o *JSPushConsumerUpdateRequest) GetAckWait() int64`

GetAckWait returns the AckWait field if non-nil, zero value otherwise.

### GetAckWaitOk

`func (o *JSPushConsumerUpdateRequest) GetAckWaitOk() (*int64, bool)`

GetAckWaitOk returns a tuple with the AckWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckWait

`func (o *JSPushConsumerUpdateRequest) SetAckWait(v int64)`

SetAckWait sets AckWait field to given value.

### HasAckWait

`func (o *JSPushConsumerUpdateRequest) HasAckWait() bool`

HasAckWait returns a boolean if a field has been set.

### GetBackoff

`func (o *JSPushConsumerUpdateRequest) GetBackoff() []int64`

GetBackoff returns the Backoff field if non-nil, zero value otherwise.

### GetBackoffOk

`func (o *JSPushConsumerUpdateRequest) GetBackoffOk() (*[]int64, bool)`

GetBackoffOk returns a tuple with the Backoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoff

`func (o *JSPushConsumerUpdateRequest) SetBackoff(v []int64)`

SetBackoff sets Backoff field to given value.

### HasBackoff

`func (o *JSPushConsumerUpdateRequest) HasBackoff() bool`

HasBackoff returns a boolean if a field has been set.

### GetDescription

`func (o *JSPushConsumerUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSPushConsumerUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSPushConsumerUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSPushConsumerUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilterSubject

`func (o *JSPushConsumerUpdateRequest) GetFilterSubject() string`

GetFilterSubject returns the FilterSubject field if non-nil, zero value otherwise.

### GetFilterSubjectOk

`func (o *JSPushConsumerUpdateRequest) GetFilterSubjectOk() (*string, bool)`

GetFilterSubjectOk returns a tuple with the FilterSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSubject

`func (o *JSPushConsumerUpdateRequest) SetFilterSubject(v string)`

SetFilterSubject sets FilterSubject field to given value.

### HasFilterSubject

`func (o *JSPushConsumerUpdateRequest) HasFilterSubject() bool`

HasFilterSubject returns a boolean if a field has been set.

### GetHeadersOnly

`func (o *JSPushConsumerUpdateRequest) GetHeadersOnly() bool`

GetHeadersOnly returns the HeadersOnly field if non-nil, zero value otherwise.

### GetHeadersOnlyOk

`func (o *JSPushConsumerUpdateRequest) GetHeadersOnlyOk() (*bool, bool)`

GetHeadersOnlyOk returns a tuple with the HeadersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersOnly

`func (o *JSPushConsumerUpdateRequest) SetHeadersOnly(v bool)`

SetHeadersOnly sets HeadersOnly field to given value.

### HasHeadersOnly

`func (o *JSPushConsumerUpdateRequest) HasHeadersOnly() bool`

HasHeadersOnly returns a boolean if a field has been set.

### GetMaxAckPending

`func (o *JSPushConsumerUpdateRequest) GetMaxAckPending() int32`

GetMaxAckPending returns the MaxAckPending field if non-nil, zero value otherwise.

### GetMaxAckPendingOk

`func (o *JSPushConsumerUpdateRequest) GetMaxAckPendingOk() (*int32, bool)`

GetMaxAckPendingOk returns a tuple with the MaxAckPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAckPending

`func (o *JSPushConsumerUpdateRequest) SetMaxAckPending(v int32)`

SetMaxAckPending sets MaxAckPending field to given value.

### HasMaxAckPending

`func (o *JSPushConsumerUpdateRequest) HasMaxAckPending() bool`

HasMaxAckPending returns a boolean if a field has been set.

### GetMaxDeliver

`func (o *JSPushConsumerUpdateRequest) GetMaxDeliver() int32`

GetMaxDeliver returns the MaxDeliver field if non-nil, zero value otherwise.

### GetMaxDeliverOk

`func (o *JSPushConsumerUpdateRequest) GetMaxDeliverOk() (*int32, bool)`

GetMaxDeliverOk returns a tuple with the MaxDeliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliver

`func (o *JSPushConsumerUpdateRequest) SetMaxDeliver(v int32)`

SetMaxDeliver sets MaxDeliver field to given value.

### HasMaxDeliver

`func (o *JSPushConsumerUpdateRequest) HasMaxDeliver() bool`

HasMaxDeliver returns a boolean if a field has been set.

### GetSampleFreq

`func (o *JSPushConsumerUpdateRequest) GetSampleFreq() string`

GetSampleFreq returns the SampleFreq field if non-nil, zero value otherwise.

### GetSampleFreqOk

`func (o *JSPushConsumerUpdateRequest) GetSampleFreqOk() (*string, bool)`

GetSampleFreqOk returns a tuple with the SampleFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleFreq

`func (o *JSPushConsumerUpdateRequest) SetSampleFreq(v string)`

SetSampleFreq sets SampleFreq field to given value.

### HasSampleFreq

`func (o *JSPushConsumerUpdateRequest) HasSampleFreq() bool`

HasSampleFreq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


