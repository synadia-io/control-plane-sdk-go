# AgentTokenViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsCurrent** | **bool** |  | 
**LastAccessedAt** | Pointer to **time.Time** |  | [optional] 
**NkeyPublic** | **string** |  | 
**RotatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAgentTokenViewResponse

`func NewAgentTokenViewResponse(created time.Time, id string, isCurrent bool, nkeyPublic string, ) *AgentTokenViewResponse`

NewAgentTokenViewResponse instantiates a new AgentTokenViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTokenViewResponseWithDefaults

`func NewAgentTokenViewResponseWithDefaults() *AgentTokenViewResponse`

NewAgentTokenViewResponseWithDefaults instantiates a new AgentTokenViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AgentTokenViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentTokenViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentTokenViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AgentTokenViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentTokenViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentTokenViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsCurrent

`func (o *AgentTokenViewResponse) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *AgentTokenViewResponse) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *AgentTokenViewResponse) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.


### GetLastAccessedAt

`func (o *AgentTokenViewResponse) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *AgentTokenViewResponse) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *AgentTokenViewResponse) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *AgentTokenViewResponse) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### GetNkeyPublic

`func (o *AgentTokenViewResponse) GetNkeyPublic() string`

GetNkeyPublic returns the NkeyPublic field if non-nil, zero value otherwise.

### GetNkeyPublicOk

`func (o *AgentTokenViewResponse) GetNkeyPublicOk() (*string, bool)`

GetNkeyPublicOk returns a tuple with the NkeyPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNkeyPublic

`func (o *AgentTokenViewResponse) SetNkeyPublic(v string)`

SetNkeyPublic sets NkeyPublic field to given value.


### GetRotatedAt

`func (o *AgentTokenViewResponse) GetRotatedAt() time.Time`

GetRotatedAt returns the RotatedAt field if non-nil, zero value otherwise.

### GetRotatedAtOk

`func (o *AgentTokenViewResponse) GetRotatedAtOk() (*time.Time, bool)`

GetRotatedAtOk returns a tuple with the RotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedAt

`func (o *AgentTokenViewResponse) SetRotatedAt(v time.Time)`

SetRotatedAt sets RotatedAt field to given value.

### HasRotatedAt

`func (o *AgentTokenViewResponse) HasRotatedAt() bool`

HasRotatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


