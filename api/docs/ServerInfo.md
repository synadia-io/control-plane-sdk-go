# ServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Host** | **string** |  | 
**Id** | **string** |  | 
**Jetstream** | **bool** |  | 
**Name** | **string** |  | 
**Seq** | **int32** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Time** | **time.Time** |  | 
**Ver** | **string** |  | 

## Methods

### NewServerInfo

`func NewServerInfo(host string, id string, jetstream bool, name string, seq int32, time time.Time, ver string, ) *ServerInfo`

NewServerInfo instantiates a new ServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInfoWithDefaults

`func NewServerInfoWithDefaults() *ServerInfo`

NewServerInfoWithDefaults instantiates a new ServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ServerInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ServerInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ServerInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ServerInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDomain

`func (o *ServerInfo) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ServerInfo) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ServerInfo) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ServerInfo) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHost

`func (o *ServerInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ServerInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ServerInfo) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *ServerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInfo) SetId(v string)`

SetId sets Id field to given value.


### GetJetstream

`func (o *ServerInfo) GetJetstream() bool`

GetJetstream returns the Jetstream field if non-nil, zero value otherwise.

### GetJetstreamOk

`func (o *ServerInfo) GetJetstreamOk() (*bool, bool)`

GetJetstreamOk returns a tuple with the Jetstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstream

`func (o *ServerInfo) SetJetstream(v bool)`

SetJetstream sets Jetstream field to given value.


### GetName

`func (o *ServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerInfo) SetName(v string)`

SetName sets Name field to given value.


### GetSeq

`func (o *ServerInfo) GetSeq() int32`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *ServerInfo) GetSeqOk() (*int32, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *ServerInfo) SetSeq(v int32)`

SetSeq sets Seq field to given value.


### GetTags

`func (o *ServerInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTime

`func (o *ServerInfo) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ServerInfo) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ServerInfo) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetVer

`func (o *ServerInfo) GetVer() string`

GetVer returns the Ver field if non-nil, zero value otherwise.

### GetVerOk

`func (o *ServerInfo) GetVerOk() (*string, bool)`

GetVerOk returns a tuple with the Ver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVer

`func (o *ServerInfo) SetVer(v string)`

SetVer sets Ver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


