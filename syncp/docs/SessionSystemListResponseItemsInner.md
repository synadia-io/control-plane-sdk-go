# SessionSystemListResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**HasManagedOperator** | **bool** |  | 
**HasManagedSystemAccount** | **bool** |  | 
**Id** | **string** |  | 
**JetstreamDomain** | Pointer to **NullableString** |  | [optional] 
**JetstreamEnabled** | **bool** |  | 
**Name** | **string** |  | 
**OperatorClaims** | Pointer to [**OperatorClaims**](OperatorClaims.md) |  | [optional] 
**OperatorJwt** | Pointer to **string** |  | [optional] 
**State** | [**SystemState**](SystemState.md) |  | 
**SystemAccountJwt** | Pointer to **string** |  | [optional] 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewSessionSystemListResponseItemsInner

`func NewSessionSystemListResponseItemsInner(hasManagedOperator bool, hasManagedSystemAccount bool, id string, jetstreamEnabled bool, name string, state SystemState, userJwtExpiresInSecs int64, ) *SessionSystemListResponseItemsInner`

NewSessionSystemListResponseItemsInner instantiates a new SessionSystemListResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSystemListResponseItemsInnerWithDefaults

`func NewSessionSystemListResponseItemsInnerWithDefaults() *SessionSystemListResponseItemsInner`

NewSessionSystemListResponseItemsInnerWithDefaults instantiates a new SessionSystemListResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SessionSystemListResponseItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SessionSystemListResponseItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SessionSystemListResponseItemsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SessionSystemListResponseItemsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetHasManagedOperator

`func (o *SessionSystemListResponseItemsInner) GetHasManagedOperator() bool`

GetHasManagedOperator returns the HasManagedOperator field if non-nil, zero value otherwise.

### GetHasManagedOperatorOk

`func (o *SessionSystemListResponseItemsInner) GetHasManagedOperatorOk() (*bool, bool)`

GetHasManagedOperatorOk returns a tuple with the HasManagedOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManagedOperator

`func (o *SessionSystemListResponseItemsInner) SetHasManagedOperator(v bool)`

SetHasManagedOperator sets HasManagedOperator field to given value.


### GetHasManagedSystemAccount

`func (o *SessionSystemListResponseItemsInner) GetHasManagedSystemAccount() bool`

GetHasManagedSystemAccount returns the HasManagedSystemAccount field if non-nil, zero value otherwise.

### GetHasManagedSystemAccountOk

`func (o *SessionSystemListResponseItemsInner) GetHasManagedSystemAccountOk() (*bool, bool)`

GetHasManagedSystemAccountOk returns a tuple with the HasManagedSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManagedSystemAccount

`func (o *SessionSystemListResponseItemsInner) SetHasManagedSystemAccount(v bool)`

SetHasManagedSystemAccount sets HasManagedSystemAccount field to given value.


### GetId

`func (o *SessionSystemListResponseItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionSystemListResponseItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionSystemListResponseItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetJetstreamDomain

`func (o *SessionSystemListResponseItemsInner) GetJetstreamDomain() string`

GetJetstreamDomain returns the JetstreamDomain field if non-nil, zero value otherwise.

### GetJetstreamDomainOk

`func (o *SessionSystemListResponseItemsInner) GetJetstreamDomainOk() (*string, bool)`

GetJetstreamDomainOk returns a tuple with the JetstreamDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamDomain

`func (o *SessionSystemListResponseItemsInner) SetJetstreamDomain(v string)`

SetJetstreamDomain sets JetstreamDomain field to given value.

### HasJetstreamDomain

`func (o *SessionSystemListResponseItemsInner) HasJetstreamDomain() bool`

HasJetstreamDomain returns a boolean if a field has been set.

### SetJetstreamDomainNil

`func (o *SessionSystemListResponseItemsInner) SetJetstreamDomainNil(b bool)`

 SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil

### UnsetJetstreamDomain
`func (o *SessionSystemListResponseItemsInner) UnsetJetstreamDomain()`

UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
### GetJetstreamEnabled

`func (o *SessionSystemListResponseItemsInner) GetJetstreamEnabled() bool`

GetJetstreamEnabled returns the JetstreamEnabled field if non-nil, zero value otherwise.

### GetJetstreamEnabledOk

`func (o *SessionSystemListResponseItemsInner) GetJetstreamEnabledOk() (*bool, bool)`

GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamEnabled

`func (o *SessionSystemListResponseItemsInner) SetJetstreamEnabled(v bool)`

SetJetstreamEnabled sets JetstreamEnabled field to given value.


### GetName

`func (o *SessionSystemListResponseItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionSystemListResponseItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionSystemListResponseItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorClaims

`func (o *SessionSystemListResponseItemsInner) GetOperatorClaims() OperatorClaims`

GetOperatorClaims returns the OperatorClaims field if non-nil, zero value otherwise.

### GetOperatorClaimsOk

`func (o *SessionSystemListResponseItemsInner) GetOperatorClaimsOk() (*OperatorClaims, bool)`

GetOperatorClaimsOk returns a tuple with the OperatorClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorClaims

`func (o *SessionSystemListResponseItemsInner) SetOperatorClaims(v OperatorClaims)`

SetOperatorClaims sets OperatorClaims field to given value.

### HasOperatorClaims

`func (o *SessionSystemListResponseItemsInner) HasOperatorClaims() bool`

HasOperatorClaims returns a boolean if a field has been set.

### GetOperatorJwt

`func (o *SessionSystemListResponseItemsInner) GetOperatorJwt() string`

GetOperatorJwt returns the OperatorJwt field if non-nil, zero value otherwise.

### GetOperatorJwtOk

`func (o *SessionSystemListResponseItemsInner) GetOperatorJwtOk() (*string, bool)`

GetOperatorJwtOk returns a tuple with the OperatorJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorJwt

`func (o *SessionSystemListResponseItemsInner) SetOperatorJwt(v string)`

SetOperatorJwt sets OperatorJwt field to given value.

### HasOperatorJwt

`func (o *SessionSystemListResponseItemsInner) HasOperatorJwt() bool`

HasOperatorJwt returns a boolean if a field has been set.

### GetState

`func (o *SessionSystemListResponseItemsInner) GetState() SystemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SessionSystemListResponseItemsInner) GetStateOk() (*SystemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SessionSystemListResponseItemsInner) SetState(v SystemState)`

SetState sets State field to given value.


### GetSystemAccountJwt

`func (o *SessionSystemListResponseItemsInner) GetSystemAccountJwt() string`

GetSystemAccountJwt returns the SystemAccountJwt field if non-nil, zero value otherwise.

### GetSystemAccountJwtOk

`func (o *SessionSystemListResponseItemsInner) GetSystemAccountJwtOk() (*string, bool)`

GetSystemAccountJwtOk returns a tuple with the SystemAccountJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAccountJwt

`func (o *SessionSystemListResponseItemsInner) SetSystemAccountJwt(v string)`

SetSystemAccountJwt sets SystemAccountJwt field to given value.

### HasSystemAccountJwt

`func (o *SessionSystemListResponseItemsInner) HasSystemAccountJwt() bool`

HasSystemAccountJwt returns a boolean if a field has been set.

### GetUserJwtExpiresInSecs

`func (o *SessionSystemListResponseItemsInner) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *SessionSystemListResponseItemsInner) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *SessionSystemListResponseItemsInner) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


