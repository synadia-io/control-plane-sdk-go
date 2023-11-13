# SystemViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentType** | **string** |  | 
**HasManagedOperator** | **bool** |  | 
**HasManagedSystemAccount** | **bool** |  | 
**HostSystemId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**IsTenant** | **bool** |  | 
**JetstreamDomain** | Pointer to **NullableString** |  | [optional] 
**JetstreamEnabled** | **bool** |  | 
**JetstreamTiers** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | 
**OperatorClaims** | Pointer to [**OperatorClaims**](OperatorClaims.md) |  | [optional] 
**OperatorJwt** | Pointer to **string** |  | [optional] 
**State** | [**SystemState**](SystemState.md) |  | 
**SystemAccountJwt** | Pointer to **string** |  | [optional] 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewSystemViewResponse

`func NewSystemViewResponse(agentType string, hasManagedOperator bool, hasManagedSystemAccount bool, id string, isTenant bool, jetstreamEnabled bool, name string, state SystemState, team TeamInfo, userJwtExpiresInSecs int64, ) *SystemViewResponse`

NewSystemViewResponse instantiates a new SystemViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemViewResponseWithDefaults

`func NewSystemViewResponseWithDefaults() *SystemViewResponse`

NewSystemViewResponseWithDefaults instantiates a new SystemViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentType

`func (o *SystemViewResponse) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *SystemViewResponse) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *SystemViewResponse) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.


### GetHasManagedOperator

`func (o *SystemViewResponse) GetHasManagedOperator() bool`

GetHasManagedOperator returns the HasManagedOperator field if non-nil, zero value otherwise.

### GetHasManagedOperatorOk

`func (o *SystemViewResponse) GetHasManagedOperatorOk() (*bool, bool)`

GetHasManagedOperatorOk returns a tuple with the HasManagedOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManagedOperator

`func (o *SystemViewResponse) SetHasManagedOperator(v bool)`

SetHasManagedOperator sets HasManagedOperator field to given value.


### GetHasManagedSystemAccount

`func (o *SystemViewResponse) GetHasManagedSystemAccount() bool`

GetHasManagedSystemAccount returns the HasManagedSystemAccount field if non-nil, zero value otherwise.

### GetHasManagedSystemAccountOk

`func (o *SystemViewResponse) GetHasManagedSystemAccountOk() (*bool, bool)`

GetHasManagedSystemAccountOk returns a tuple with the HasManagedSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManagedSystemAccount

`func (o *SystemViewResponse) SetHasManagedSystemAccount(v bool)`

SetHasManagedSystemAccount sets HasManagedSystemAccount field to given value.


### GetHostSystemId

`func (o *SystemViewResponse) GetHostSystemId() string`

GetHostSystemId returns the HostSystemId field if non-nil, zero value otherwise.

### GetHostSystemIdOk

`func (o *SystemViewResponse) GetHostSystemIdOk() (*string, bool)`

GetHostSystemIdOk returns a tuple with the HostSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSystemId

`func (o *SystemViewResponse) SetHostSystemId(v string)`

SetHostSystemId sets HostSystemId field to given value.

### HasHostSystemId

`func (o *SystemViewResponse) HasHostSystemId() bool`

HasHostSystemId returns a boolean if a field has been set.

### SetHostSystemIdNil

`func (o *SystemViewResponse) SetHostSystemIdNil(b bool)`

 SetHostSystemIdNil sets the value for HostSystemId to be an explicit nil

### UnsetHostSystemId
`func (o *SystemViewResponse) UnsetHostSystemId()`

UnsetHostSystemId ensures that no value is present for HostSystemId, not even an explicit nil
### GetId

`func (o *SystemViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsTenant

`func (o *SystemViewResponse) GetIsTenant() bool`

GetIsTenant returns the IsTenant field if non-nil, zero value otherwise.

### GetIsTenantOk

`func (o *SystemViewResponse) GetIsTenantOk() (*bool, bool)`

GetIsTenantOk returns a tuple with the IsTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTenant

`func (o *SystemViewResponse) SetIsTenant(v bool)`

SetIsTenant sets IsTenant field to given value.


### GetJetstreamDomain

`func (o *SystemViewResponse) GetJetstreamDomain() string`

GetJetstreamDomain returns the JetstreamDomain field if non-nil, zero value otherwise.

### GetJetstreamDomainOk

`func (o *SystemViewResponse) GetJetstreamDomainOk() (*string, bool)`

GetJetstreamDomainOk returns a tuple with the JetstreamDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamDomain

`func (o *SystemViewResponse) SetJetstreamDomain(v string)`

SetJetstreamDomain sets JetstreamDomain field to given value.

### HasJetstreamDomain

`func (o *SystemViewResponse) HasJetstreamDomain() bool`

HasJetstreamDomain returns a boolean if a field has been set.

### SetJetstreamDomainNil

`func (o *SystemViewResponse) SetJetstreamDomainNil(b bool)`

 SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil

### UnsetJetstreamDomain
`func (o *SystemViewResponse) UnsetJetstreamDomain()`

UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
### GetJetstreamEnabled

`func (o *SystemViewResponse) GetJetstreamEnabled() bool`

GetJetstreamEnabled returns the JetstreamEnabled field if non-nil, zero value otherwise.

### GetJetstreamEnabledOk

`func (o *SystemViewResponse) GetJetstreamEnabledOk() (*bool, bool)`

GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamEnabled

`func (o *SystemViewResponse) SetJetstreamEnabled(v bool)`

SetJetstreamEnabled sets JetstreamEnabled field to given value.


### GetJetstreamTiers

`func (o *SystemViewResponse) GetJetstreamTiers() []string`

GetJetstreamTiers returns the JetstreamTiers field if non-nil, zero value otherwise.

### GetJetstreamTiersOk

`func (o *SystemViewResponse) GetJetstreamTiersOk() (*[]string, bool)`

GetJetstreamTiersOk returns a tuple with the JetstreamTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamTiers

`func (o *SystemViewResponse) SetJetstreamTiers(v []string)`

SetJetstreamTiers sets JetstreamTiers field to given value.

### HasJetstreamTiers

`func (o *SystemViewResponse) HasJetstreamTiers() bool`

HasJetstreamTiers returns a boolean if a field has been set.

### GetName

`func (o *SystemViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorClaims

`func (o *SystemViewResponse) GetOperatorClaims() OperatorClaims`

GetOperatorClaims returns the OperatorClaims field if non-nil, zero value otherwise.

### GetOperatorClaimsOk

`func (o *SystemViewResponse) GetOperatorClaimsOk() (*OperatorClaims, bool)`

GetOperatorClaimsOk returns a tuple with the OperatorClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorClaims

`func (o *SystemViewResponse) SetOperatorClaims(v OperatorClaims)`

SetOperatorClaims sets OperatorClaims field to given value.

### HasOperatorClaims

`func (o *SystemViewResponse) HasOperatorClaims() bool`

HasOperatorClaims returns a boolean if a field has been set.

### GetOperatorJwt

`func (o *SystemViewResponse) GetOperatorJwt() string`

GetOperatorJwt returns the OperatorJwt field if non-nil, zero value otherwise.

### GetOperatorJwtOk

`func (o *SystemViewResponse) GetOperatorJwtOk() (*string, bool)`

GetOperatorJwtOk returns a tuple with the OperatorJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorJwt

`func (o *SystemViewResponse) SetOperatorJwt(v string)`

SetOperatorJwt sets OperatorJwt field to given value.

### HasOperatorJwt

`func (o *SystemViewResponse) HasOperatorJwt() bool`

HasOperatorJwt returns a boolean if a field has been set.

### GetState

`func (o *SystemViewResponse) GetState() SystemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemViewResponse) GetStateOk() (*SystemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemViewResponse) SetState(v SystemState)`

SetState sets State field to given value.


### GetSystemAccountJwt

`func (o *SystemViewResponse) GetSystemAccountJwt() string`

GetSystemAccountJwt returns the SystemAccountJwt field if non-nil, zero value otherwise.

### GetSystemAccountJwtOk

`func (o *SystemViewResponse) GetSystemAccountJwtOk() (*string, bool)`

GetSystemAccountJwtOk returns a tuple with the SystemAccountJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccountJwt

`func (o *SystemViewResponse) SetSystemAccountJwt(v string)`

SetSystemAccountJwt sets SystemAccountJwt field to given value.

### HasSystemAccountJwt

`func (o *SystemViewResponse) HasSystemAccountJwt() bool`

HasSystemAccountJwt returns a boolean if a field has been set.

### GetTeam

`func (o *SystemViewResponse) GetTeam() TeamInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *SystemViewResponse) GetTeamOk() (*TeamInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *SystemViewResponse) SetTeam(v TeamInfo)`

SetTeam sets Team field to given value.


### GetUserJwtExpiresInSecs

`func (o *SystemViewResponse) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *SystemViewResponse) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *SystemViewResponse) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


