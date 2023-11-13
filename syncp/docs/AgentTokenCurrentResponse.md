# AgentTokenCurrentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsCurrent** | **bool** |  | 
**LastAccessedAt** | Pointer to **time.Time** |  | [optional] 
**NkeyPublic** | **string** |  | 
**RotatedAt** | Pointer to **time.Time** |  | [optional] 
**Token** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewAgentTokenCurrentResponse

`func NewAgentTokenCurrentResponse(created time.Time, id string, isCurrent bool, nkeyPublic string, token string, url string, ) *AgentTokenCurrentResponse`

NewAgentTokenCurrentResponse instantiates a new AgentTokenCurrentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentTokenCurrentResponseWithDefaults

`func NewAgentTokenCurrentResponseWithDefaults() *AgentTokenCurrentResponse`

NewAgentTokenCurrentResponseWithDefaults instantiates a new AgentTokenCurrentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AgentTokenCurrentResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AgentTokenCurrentResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AgentTokenCurrentResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AgentTokenCurrentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentTokenCurrentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentTokenCurrentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsCurrent

`func (o *AgentTokenCurrentResponse) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *AgentTokenCurrentResponse) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *AgentTokenCurrentResponse) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.


### GetLastAccessedAt

`func (o *AgentTokenCurrentResponse) GetLastAccessedAt() time.Time`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *AgentTokenCurrentResponse) GetLastAccessedAtOk() (*time.Time, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *AgentTokenCurrentResponse) SetLastAccessedAt(v time.Time)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *AgentTokenCurrentResponse) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### GetNkeyPublic

`func (o *AgentTokenCurrentResponse) GetNkeyPublic() string`

GetNkeyPublic returns the NkeyPublic field if non-nil, zero value otherwise.

### GetNkeyPublicOk

`func (o *AgentTokenCurrentResponse) GetNkeyPublicOk() (*string, bool)`

GetNkeyPublicOk returns a tuple with the NkeyPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNkeyPublic

`func (o *AgentTokenCurrentResponse) SetNkeyPublic(v string)`

SetNkeyPublic sets NkeyPublic field to given value.


### GetRotatedAt

`func (o *AgentTokenCurrentResponse) GetRotatedAt() time.Time`

GetRotatedAt returns the RotatedAt field if non-nil, zero value otherwise.

### GetRotatedAtOk

`func (o *AgentTokenCurrentResponse) GetRotatedAtOk() (*time.Time, bool)`

GetRotatedAtOk returns a tuple with the RotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedAt

`func (o *AgentTokenCurrentResponse) SetRotatedAt(v time.Time)`

SetRotatedAt sets RotatedAt field to given value.

### HasRotatedAt

`func (o *AgentTokenCurrentResponse) HasRotatedAt() bool`

HasRotatedAt returns a boolean if a field has been set.

### GetToken

`func (o *AgentTokenCurrentResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AgentTokenCurrentResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AgentTokenCurrentResponse) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *AgentTokenCurrentResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgentTokenCurrentResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgentTokenCurrentResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


