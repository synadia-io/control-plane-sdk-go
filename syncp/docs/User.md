# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Type** | Pointer to **string** | ClaimType is used to indicate the type of JWT being stored in a Claim | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
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
**IssuerAccount** | Pointer to **string** | IssuerAccount stores the public key for the account the issuer represents. When set, the claim was issued by a signing key. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *User) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *User) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *User) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *User) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *User) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *User) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *User) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *User) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *User) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *User) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPub

`func (o *User) GetPub() Permission`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *User) GetPubOk() (*Permission, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *User) SetPub(v Permission)`

SetPub sets Pub field to given value.

### HasPub

`func (o *User) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetResp

`func (o *User) GetResp() PermissionsResp`

GetResp returns the Resp field if non-nil, zero value otherwise.

### GetRespOk

`func (o *User) GetRespOk() (*PermissionsResp, bool)`

GetRespOk returns a tuple with the Resp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResp

`func (o *User) SetResp(v PermissionsResp)`

SetResp sets Resp field to given value.

### HasResp

`func (o *User) HasResp() bool`

HasResp returns a boolean if a field has been set.

### SetRespNil

`func (o *User) SetRespNil(b bool)`

 SetRespNil sets the value for Resp to be an explicit nil

### UnsetResp
`func (o *User) UnsetResp()`

UnsetResp ensures that no value is present for Resp, not even an explicit nil
### GetSub

`func (o *User) GetSub() Permission`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *User) GetSubOk() (*Permission, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *User) SetSub(v Permission)`

SetSub sets Sub field to given value.

### HasSub

`func (o *User) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetSrc

`func (o *User) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *User) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *User) SetSrc(v []string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *User) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *User) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *User) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetTimes

`func (o *User) GetTimes() []TimeRange`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *User) GetTimesOk() (*[]TimeRange, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *User) SetTimes(v []TimeRange)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *User) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesLocation

`func (o *User) GetTimesLocation() string`

GetTimesLocation returns the TimesLocation field if non-nil, zero value otherwise.

### GetTimesLocationOk

`func (o *User) GetTimesLocationOk() (*string, bool)`

GetTimesLocationOk returns a tuple with the TimesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesLocation

`func (o *User) SetTimesLocation(v string)`

SetTimesLocation sets TimesLocation field to given value.

### HasTimesLocation

`func (o *User) HasTimesLocation() bool`

HasTimesLocation returns a boolean if a field has been set.

### GetData

`func (o *User) GetData() int64`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *User) GetDataOk() (*int64, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *User) SetData(v int64)`

SetData sets Data field to given value.

### HasData

`func (o *User) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPayload

`func (o *User) GetPayload() int64`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *User) GetPayloadOk() (*int64, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *User) SetPayload(v int64)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *User) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetSubs

`func (o *User) GetSubs() int64`

GetSubs returns the Subs field if non-nil, zero value otherwise.

### GetSubsOk

`func (o *User) GetSubsOk() (*int64, bool)`

GetSubsOk returns a tuple with the Subs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubs

`func (o *User) SetSubs(v int64)`

SetSubs sets Subs field to given value.

### HasSubs

`func (o *User) HasSubs() bool`

HasSubs returns a boolean if a field has been set.

### GetAllowedConnectionTypes

`func (o *User) GetAllowedConnectionTypes() []string`

GetAllowedConnectionTypes returns the AllowedConnectionTypes field if non-nil, zero value otherwise.

### GetAllowedConnectionTypesOk

`func (o *User) GetAllowedConnectionTypesOk() (*[]string, bool)`

GetAllowedConnectionTypesOk returns a tuple with the AllowedConnectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedConnectionTypes

`func (o *User) SetAllowedConnectionTypes(v []string)`

SetAllowedConnectionTypes sets AllowedConnectionTypes field to given value.

### HasAllowedConnectionTypes

`func (o *User) HasAllowedConnectionTypes() bool`

HasAllowedConnectionTypes returns a boolean if a field has been set.

### SetAllowedConnectionTypesNil

`func (o *User) SetAllowedConnectionTypesNil(b bool)`

 SetAllowedConnectionTypesNil sets the value for AllowedConnectionTypes to be an explicit nil

### UnsetAllowedConnectionTypes
`func (o *User) UnsetAllowedConnectionTypes()`

UnsetAllowedConnectionTypes ensures that no value is present for AllowedConnectionTypes, not even an explicit nil
### GetBearerToken

`func (o *User) GetBearerToken() bool`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *User) GetBearerTokenOk() (*bool, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *User) SetBearerToken(v bool)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *User) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetIssuerAccount

`func (o *User) GetIssuerAccount() string`

GetIssuerAccount returns the IssuerAccount field if non-nil, zero value otherwise.

### GetIssuerAccountOk

`func (o *User) GetIssuerAccountOk() (*string, bool)`

GetIssuerAccountOk returns a tuple with the IssuerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAccount

`func (o *User) SetIssuerAccount(v string)`

SetIssuerAccount sets IssuerAccount field to given value.

### HasIssuerAccount

`func (o *User) HasIssuerAccount() bool`

HasIssuerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


