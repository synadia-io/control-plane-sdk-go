# SessionAccountListResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**AccountPublicKey** | **string** |  | 
**Claims** | [**AccountClaims**](AccountClaims.md) |  | 
**ClaimsInfo** | [**AccountClaimsInfo**](AccountClaimsInfo.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**IsSystemAccount** | **bool** |  | 
**Jwt** | **string** |  | 
**JwtSettings** | [**AccountJWTSettings**](AccountJWTSettings.md) |  | 
**Name** | **string** |  | 
**System** | [**SystemInfo**](SystemInfo.md) |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewSessionAccountListResponseItemsInner

`func NewSessionAccountListResponseItemsInner(accountPublicKey string, claims AccountClaims, claimsInfo AccountClaimsInfo, created time.Time, id string, isSystemAccount bool, jwt string, jwtSettings AccountJWTSettings, name string, system SystemInfo, userJwtExpiresInSecs int64, ) *SessionAccountListResponseItemsInner`

NewSessionAccountListResponseItemsInner instantiates a new SessionAccountListResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionAccountListResponseItemsInnerWithDefaults

`func NewSessionAccountListResponseItemsInnerWithDefaults() *SessionAccountListResponseItemsInner`

NewSessionAccountListResponseItemsInnerWithDefaults instantiates a new SessionAccountListResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SessionAccountListResponseItemsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SessionAccountListResponseItemsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SessionAccountListResponseItemsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SessionAccountListResponseItemsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAccountPublicKey

`func (o *SessionAccountListResponseItemsInner) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *SessionAccountListResponseItemsInner) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *SessionAccountListResponseItemsInner) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetClaims

`func (o *SessionAccountListResponseItemsInner) GetClaims() AccountClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SessionAccountListResponseItemsInner) GetClaimsOk() (*AccountClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SessionAccountListResponseItemsInner) SetClaims(v AccountClaims)`

SetClaims sets Claims field to given value.


### GetClaimsInfo

`func (o *SessionAccountListResponseItemsInner) GetClaimsInfo() AccountClaimsInfo`

GetClaimsInfo returns the ClaimsInfo field if non-nil, zero value otherwise.

### GetClaimsInfoOk

`func (o *SessionAccountListResponseItemsInner) GetClaimsInfoOk() (*AccountClaimsInfo, bool)`

GetClaimsInfoOk returns a tuple with the ClaimsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsInfo

`func (o *SessionAccountListResponseItemsInner) SetClaimsInfo(v AccountClaimsInfo)`

SetClaimsInfo sets ClaimsInfo field to given value.


### GetCreated

`func (o *SessionAccountListResponseItemsInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SessionAccountListResponseItemsInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SessionAccountListResponseItemsInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *SessionAccountListResponseItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionAccountListResponseItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionAccountListResponseItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemAccount

`func (o *SessionAccountListResponseItemsInner) GetIsSystemAccount() bool`

GetIsSystemAccount returns the IsSystemAccount field if non-nil, zero value otherwise.

### GetIsSystemAccountOk

`func (o *SessionAccountListResponseItemsInner) GetIsSystemAccountOk() (*bool, bool)`

GetIsSystemAccountOk returns a tuple with the IsSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAccount

`func (o *SessionAccountListResponseItemsInner) SetIsSystemAccount(v bool)`

SetIsSystemAccount sets IsSystemAccount field to given value.


### GetJwt

`func (o *SessionAccountListResponseItemsInner) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SessionAccountListResponseItemsInner) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SessionAccountListResponseItemsInner) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetJwtSettings

`func (o *SessionAccountListResponseItemsInner) GetJwtSettings() AccountJWTSettings`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *SessionAccountListResponseItemsInner) GetJwtSettingsOk() (*AccountJWTSettings, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *SessionAccountListResponseItemsInner) SetJwtSettings(v AccountJWTSettings)`

SetJwtSettings sets JwtSettings field to given value.


### GetName

`func (o *SessionAccountListResponseItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionAccountListResponseItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionAccountListResponseItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetSystem

`func (o *SessionAccountListResponseItemsInner) GetSystem() SystemInfo`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *SessionAccountListResponseItemsInner) GetSystemOk() (*SystemInfo, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *SessionAccountListResponseItemsInner) SetSystem(v SystemInfo)`

SetSystem sets System field to given value.


### GetUserJwtExpiresInSecs

`func (o *SessionAccountListResponseItemsInner) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *SessionAccountListResponseItemsInner) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *SessionAccountListResponseItemsInner) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


