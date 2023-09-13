# SubDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Cid** | **int32** |  | 
**Max** | Pointer to **int64** |  | [optional] 
**Msgs** | **int64** |  | 
**Qgroup** | Pointer to **string** |  | [optional] 
**Sid** | **string** |  | 
**Subject** | **string** |  | 

## Methods

### NewSubDetail

`func NewSubDetail(cid int32, msgs int64, sid string, subject string, ) *SubDetail`

NewSubDetail instantiates a new SubDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubDetailWithDefaults

`func NewSubDetailWithDefaults() *SubDetail`

NewSubDetailWithDefaults instantiates a new SubDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SubDetail) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubDetail) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubDetail) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SubDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCid

`func (o *SubDetail) GetCid() int32`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SubDetail) GetCidOk() (*int32, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SubDetail) SetCid(v int32)`

SetCid sets Cid field to given value.


### GetMax

`func (o *SubDetail) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *SubDetail) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *SubDetail) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *SubDetail) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMsgs

`func (o *SubDetail) GetMsgs() int64`

GetMsgs returns the Msgs field if non-nil, zero value otherwise.

### GetMsgsOk

`func (o *SubDetail) GetMsgsOk() (*int64, bool)`

GetMsgsOk returns a tuple with the Msgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgs

`func (o *SubDetail) SetMsgs(v int64)`

SetMsgs sets Msgs field to given value.


### GetQgroup

`func (o *SubDetail) GetQgroup() string`

GetQgroup returns the Qgroup field if non-nil, zero value otherwise.

### GetQgroupOk

`func (o *SubDetail) GetQgroupOk() (*string, bool)`

GetQgroupOk returns a tuple with the Qgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQgroup

`func (o *SubDetail) SetQgroup(v string)`

SetQgroup sets Qgroup field to given value.

### HasQgroup

`func (o *SubDetail) HasQgroup() bool`

HasQgroup returns a boolean if a field has been set.

### GetSid

`func (o *SubDetail) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SubDetail) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SubDetail) SetSid(v string)`

SetSid sets Sid field to given value.


### GetSubject

`func (o *SubDetail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SubDetail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SubDetail) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


