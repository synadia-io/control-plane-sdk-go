# SessionNatsUserListResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**Account** | [**AccountInfo**](AccountInfo.md) |  | 
**Claims** | [**UserClaims**](UserClaims.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**Jwt** | **string** |  | 
**JwtExpiresAtMax** | **int64** |  | 
**JwtExpiresInSecs** | **int64** |  | 
**Name** | **string** |  | 
**Revoked** | **bool** |  | 
**SkGroupId** | Pointer to **string** |  | [optional] 
**System** | [**SystemInfo**](SystemInfo.md) |  | 
**UserPublicKey** | **string** |  | 

## Methods

### NewSessionNatsUserListResponseItemsInner

`func NewSessionNatsUserListResponseItemsInner(account AccountInfo, claims UserClaims, created time.Time, id string, jwt string, jwtExpiresAtMax int64, jwtExpiresInSecs int64, name string, revoked bool, system SystemInfo, userPublicKey string, ) *SessionNatsUserListResponseItemsInner`

NewSessionNatsUserListResponseItemsInner instantiates a new SessionNatsUserListResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionNatsUserListResponseItemsInnerWithDefaults

`func NewSessionNatsUserListResponseItemsInnerWithDefaults() *SessionNatsUserListResponseItemsInner`

NewSessionNatsUserListResponseItemsInnerWithDefaults instantiates a new SessionNatsUserListResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SessionNatsUserListResponseItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SessionNatsUserListResponseItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SessionNatsUserListResponseItemsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SessionNatsUserListResponseItemsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAccount

`func (o *SessionNatsUserListResponseItemsInner) GetAccount() AccountInfo`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SessionNatsUserListResponseItemsInner) GetAccountOk() (*AccountInfo, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SessionNatsUserListResponseItemsInner) SetAccount(v AccountInfo)`

SetAccount sets Account field to given value.


### GetClaims

`func (o *SessionNatsUserListResponseItemsInner) GetClaims() UserClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SessionNatsUserListResponseItemsInner) GetClaimsOk() (*UserClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SessionNatsUserListResponseItemsInner) SetClaims(v UserClaims)`

SetClaims sets Claims field to given value.


### GetCreated

`func (o *SessionNatsUserListResponseItemsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SessionNatsUserListResponseItemsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SessionNatsUserListResponseItemsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *SessionNatsUserListResponseItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionNatsUserListResponseItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionNatsUserListResponseItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetJwt

`func (o *SessionNatsUserListResponseItemsInner) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SessionNatsUserListResponseItemsInner) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SessionNatsUserListResponseItemsInner) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetJwtExpiresAtMax

`func (o *SessionNatsUserListResponseItemsInner) GetJwtExpiresAtMax() int64`

GetJwtExpiresAtMax returns the JwtExpiresAtMax field if non-nil, zero value otherwise.

### GetJwtExpiresAtMaxOk

`func (o *SessionNatsUserListResponseItemsInner) GetJwtExpiresAtMaxOk() (*int64, bool)`

GetJwtExpiresAtMaxOk returns a tuple with the JwtExpiresAtMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresAtMax

`func (o *SessionNatsUserListResponseItemsInner) SetJwtExpiresAtMax(v int64)`

SetJwtExpiresAtMax sets JwtExpiresAtMax field to given value.


### GetJwtExpiresInSecs

`func (o *SessionNatsUserListResponseItemsInner) GetJwtExpiresInSecs() int64`

GetJwtExpiresInSecs returns the JwtExpiresInSecs field if non-nil, zero value otherwise.

### GetJwtExpiresInSecsOk

`func (o *SessionNatsUserListResponseItemsInner) GetJwtExpiresInSecsOk() (*int64, bool)`

GetJwtExpiresInSecsOk returns a tuple with the JwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresInSecs

`func (o *SessionNatsUserListResponseItemsInner) SetJwtExpiresInSecs(v int64)`

SetJwtExpiresInSecs sets JwtExpiresInSecs field to given value.


### GetName

`func (o *SessionNatsUserListResponseItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionNatsUserListResponseItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionNatsUserListResponseItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetRevoked

`func (o *SessionNatsUserListResponseItemsInner) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *SessionNatsUserListResponseItemsInner) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *SessionNatsUserListResponseItemsInner) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.


### GetSkGroupId

`func (o *SessionNatsUserListResponseItemsInner) GetSkGroupId() string`

GetSkGroupId returns the SkGroupId field if non-nil, zero value otherwise.

### GetSkGroupIdOk

`func (o *SessionNatsUserListResponseItemsInner) GetSkGroupIdOk() (*string, bool)`

GetSkGroupIdOk returns a tuple with the SkGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkGroupId

`func (o *SessionNatsUserListResponseItemsInner) SetSkGroupId(v string)`

SetSkGroupId sets SkGroupId field to given value.

### HasSkGroupId

`func (o *SessionNatsUserListResponseItemsInner) HasSkGroupId() bool`

HasSkGroupId returns a boolean if a field has been set.

### GetSystem

`func (o *SessionNatsUserListResponseItemsInner) GetSystem() SystemInfo`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SessionNatsUserListResponseItemsInner) GetSystemOk() (*SystemInfo, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SessionNatsUserListResponseItemsInner) SetSystem(v SystemInfo)`

SetSystem sets System field to given value.


### GetUserPublicKey

`func (o *SessionNatsUserListResponseItemsInner) GetUserPublicKey() string`

GetUserPublicKey returns the UserPublicKey field if non-nil, zero value otherwise.

### GetUserPublicKeyOk

`func (o *SessionNatsUserListResponseItemsInner) GetUserPublicKeyOk() (*string, bool)`

GetUserPublicKeyOk returns a tuple with the UserPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPublicKey

`func (o *SessionNatsUserListResponseItemsInner) SetUserPublicKey(v string)`

SetUserPublicKey sets UserPublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


