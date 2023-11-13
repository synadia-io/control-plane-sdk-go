# UserPermissionLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pub** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**Resp** | Pointer to [**NullablePermissionsResp**](PermissionsResp.md) |  | [optional] 
**Sub** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**Src** | Pointer to **[]string** |  | [optional] 
**Times** | Pointer to [**[]TimeRange**](TimeRange.md) |  | [optional] 
**TimesLocation** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to **int64** |  | [optional] 
**Subs** | Pointer to **int64** |  | [optional] 
**AllowedConnectionTypes** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**BearerToken** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserPermissionLimits

`func NewUserPermissionLimits() *UserPermissionLimits`

NewUserPermissionLimits instantiates a new UserPermissionLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionLimitsWithDefaults

`func NewUserPermissionLimitsWithDefaults() *UserPermissionLimits`

NewUserPermissionLimitsWithDefaults instantiates a new UserPermissionLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPub

`func (o *UserPermissionLimits) GetPub() Permission`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *UserPermissionLimits) GetPubOk() (*Permission, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *UserPermissionLimits) SetPub(v Permission)`

SetPub sets Pub field to given value.

### HasPub

`func (o *UserPermissionLimits) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetResp

`func (o *UserPermissionLimits) GetResp() PermissionsResp`

GetResp returns the Resp field if non-nil, zero value otherwise.

### GetRespOk

`func (o *UserPermissionLimits) GetRespOk() (*PermissionsResp, bool)`

GetRespOk returns a tuple with the Resp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResp

`func (o *UserPermissionLimits) SetResp(v PermissionsResp)`

SetResp sets Resp field to given value.

### HasResp

`func (o *UserPermissionLimits) HasResp() bool`

HasResp returns a boolean if a field has been set.

### SetRespNil

`func (o *UserPermissionLimits) SetRespNil(b bool)`

 SetRespNil sets the value for Resp to be an explicit nil

### UnsetResp
`func (o *UserPermissionLimits) UnsetResp()`

UnsetResp ensures that no value is present for Resp, not even an explicit nil
### GetSub

`func (o *UserPermissionLimits) GetSub() Permission`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *UserPermissionLimits) GetSubOk() (*Permission, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *UserPermissionLimits) SetSub(v Permission)`

SetSub sets Sub field to given value.

### HasSub

`func (o *UserPermissionLimits) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetSrc

`func (o *UserPermissionLimits) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *UserPermissionLimits) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *UserPermissionLimits) SetSrc(v []string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *UserPermissionLimits) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *UserPermissionLimits) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *UserPermissionLimits) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetTimes

`func (o *UserPermissionLimits) GetTimes() []TimeRange`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *UserPermissionLimits) GetTimesOk() (*[]TimeRange, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *UserPermissionLimits) SetTimes(v []TimeRange)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *UserPermissionLimits) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesLocation

`func (o *UserPermissionLimits) GetTimesLocation() string`

GetTimesLocation returns the TimesLocation field if non-nil, zero value otherwise.

### GetTimesLocationOk

`func (o *UserPermissionLimits) GetTimesLocationOk() (*string, bool)`

GetTimesLocationOk returns a tuple with the TimesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesLocation

`func (o *UserPermissionLimits) SetTimesLocation(v string)`

SetTimesLocation sets TimesLocation field to given value.

### HasTimesLocation

`func (o *UserPermissionLimits) HasTimesLocation() bool`

HasTimesLocation returns a boolean if a field has been set.

### GetData

`func (o *UserPermissionLimits) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserPermissionLimits) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserPermissionLimits) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *UserPermissionLimits) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPayload

`func (o *UserPermissionLimits) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UserPermissionLimits) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UserPermissionLimits) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *UserPermissionLimits) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSubs

`func (o *UserPermissionLimits) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *UserPermissionLimits) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *UserPermissionLimits) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *UserPermissionLimits) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetAllowedConnectionTypes

`func (o *UserPermissionLimits) GetAllowedConnectionTypes() []string`

GetAllowedConnectionTypes returns the AllowedConnectionTypes field if non-nil, zero value otherwise.

### GetAllowedConnectionTypesOk

`func (o *UserPermissionLimits) GetAllowedConnectionTypesOk() (*[]string, bool)`

GetAllowedConnectionTypesOk returns a tuple with the AllowedConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnectionTypes

`func (o *UserPermissionLimits) SetAllowedConnectionTypes(v []string)`

SetAllowedConnectionTypes sets AllowedConnectionTypes field to given value.

### HasAllowedConnectionTypes

`func (o *UserPermissionLimits) HasAllowedConnectionTypes() bool`

HasAllowedConnectionTypes returns a boolean if a field has been set.

### SetAllowedConnectionTypesNil

`func (o *UserPermissionLimits) SetAllowedConnectionTypesNil(b bool)`

 SetAllowedConnectionTypesNil sets the value for AllowedConnectionTypes to be an explicit nil

### UnsetAllowedConnectionTypes
`func (o *UserPermissionLimits) UnsetAllowedConnectionTypes()`

UnsetAllowedConnectionTypes ensures that no value is present for AllowedConnectionTypes, not even an explicit nil
### GetBearerToken

`func (o *UserPermissionLimits) GetBearerToken() bool`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *UserPermissionLimits) GetBearerTokenOk() (*bool, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *UserPermissionLimits) SetBearerToken(v bool)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *UserPermissionLimits) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


