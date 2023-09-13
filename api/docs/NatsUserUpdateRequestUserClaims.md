# NatsUserUpdateRequestUserClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Pub** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**Resp** | Pointer to [**NullablePermissionsResp**](PermissionsResp.md) |  | [optional] 
**Sub** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**Src** | Pointer to **[]string** |  | [optional] 
**Times** | Pointer to [**[]TimeRange**](TimeRange.md) |  | [optional] 
**TimesLocation** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **int64** |  | [optional] 
**Subs** | Pointer to **int64** |  | [optional] 

## Methods

### NewNatsUserUpdateRequestUserClaims

`func NewNatsUserUpdateRequestUserClaims() *NatsUserUpdateRequestUserClaims`

NewNatsUserUpdateRequestUserClaims instantiates a new NatsUserUpdateRequestUserClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatsUserUpdateRequestUserClaimsWithDefaults

`func NewNatsUserUpdateRequestUserClaimsWithDefaults() *NatsUserUpdateRequestUserClaims`

NewNatsUserUpdateRequestUserClaimsWithDefaults instantiates a new NatsUserUpdateRequestUserClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *NatsUserUpdateRequestUserClaims) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NatsUserUpdateRequestUserClaims) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NatsUserUpdateRequestUserClaims) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NatsUserUpdateRequestUserClaims) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *NatsUserUpdateRequestUserClaims) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *NatsUserUpdateRequestUserClaims) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPub

`func (o *NatsUserUpdateRequestUserClaims) GetPub() Permission`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *NatsUserUpdateRequestUserClaims) GetPubOk() (*Permission, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *NatsUserUpdateRequestUserClaims) SetPub(v Permission)`

SetPub sets Pub field to given value.

### HasPub

`func (o *NatsUserUpdateRequestUserClaims) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetResp

`func (o *NatsUserUpdateRequestUserClaims) GetResp() PermissionsResp`

GetResp returns the Resp field if non-nil, zero value otherwise.

### GetRespOk

`func (o *NatsUserUpdateRequestUserClaims) GetRespOk() (*PermissionsResp, bool)`

GetRespOk returns a tuple with the Resp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResp

`func (o *NatsUserUpdateRequestUserClaims) SetResp(v PermissionsResp)`

SetResp sets Resp field to given value.

### HasResp

`func (o *NatsUserUpdateRequestUserClaims) HasResp() bool`

HasResp returns a boolean if a field has been set.

### SetRespNil

`func (o *NatsUserUpdateRequestUserClaims) SetRespNil(b bool)`

 SetRespNil sets the value for Resp to be an explicit nil

### UnsetResp
`func (o *NatsUserUpdateRequestUserClaims) UnsetResp()`

UnsetResp ensures that no value is present for Resp, not even an explicit nil
### GetSub

`func (o *NatsUserUpdateRequestUserClaims) GetSub() Permission`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *NatsUserUpdateRequestUserClaims) GetSubOk() (*Permission, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *NatsUserUpdateRequestUserClaims) SetSub(v Permission)`

SetSub sets Sub field to given value.

### HasSub

`func (o *NatsUserUpdateRequestUserClaims) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetSrc

`func (o *NatsUserUpdateRequestUserClaims) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *NatsUserUpdateRequestUserClaims) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *NatsUserUpdateRequestUserClaims) SetSrc(v []string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *NatsUserUpdateRequestUserClaims) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *NatsUserUpdateRequestUserClaims) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *NatsUserUpdateRequestUserClaims) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetTimes

`func (o *NatsUserUpdateRequestUserClaims) GetTimes() []TimeRange`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *NatsUserUpdateRequestUserClaims) GetTimesOk() (*[]TimeRange, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *NatsUserUpdateRequestUserClaims) SetTimes(v []TimeRange)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *NatsUserUpdateRequestUserClaims) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesLocation

`func (o *NatsUserUpdateRequestUserClaims) GetTimesLocation() string`

GetTimesLocation returns the TimesLocation field if non-nil, zero value otherwise.

### GetTimesLocationOk

`func (o *NatsUserUpdateRequestUserClaims) GetTimesLocationOk() (*string, bool)`

GetTimesLocationOk returns a tuple with the TimesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesLocation

`func (o *NatsUserUpdateRequestUserClaims) SetTimesLocation(v string)`

SetTimesLocation sets TimesLocation field to given value.

### HasTimesLocation

`func (o *NatsUserUpdateRequestUserClaims) HasTimesLocation() bool`

HasTimesLocation returns a boolean if a field has been set.

### GetData

`func (o *NatsUserUpdateRequestUserClaims) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NatsUserUpdateRequestUserClaims) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NatsUserUpdateRequestUserClaims) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *NatsUserUpdateRequestUserClaims) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPayload

`func (o *NatsUserUpdateRequestUserClaims) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NatsUserUpdateRequestUserClaims) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NatsUserUpdateRequestUserClaims) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NatsUserUpdateRequestUserClaims) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSubs

`func (o *NatsUserUpdateRequestUserClaims) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *NatsUserUpdateRequestUserClaims) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *NatsUserUpdateRequestUserClaims) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *NatsUserUpdateRequestUserClaims) HasSubs() bool`

HasSubs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


