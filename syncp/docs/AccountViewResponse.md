# AccountViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewAccountViewResponse

`func NewAccountViewResponse(accountPublicKey string, claims AccountClaims, claimsInfo AccountClaimsInfo, created time.Time, id string, isSystemAccount bool, jwt string, jwtSettings AccountJWTSettings, name string, system SystemInfo, userJwtExpiresInSecs int64, ) *AccountViewResponse`

NewAccountViewResponse instantiates a new AccountViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountViewResponseWithDefaults

`func NewAccountViewResponseWithDefaults() *AccountViewResponse`

NewAccountViewResponseWithDefaults instantiates a new AccountViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPublicKey

`func (o *AccountViewResponse) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *AccountViewResponse) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *AccountViewResponse) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetClaims

`func (o *AccountViewResponse) GetClaims() AccountClaims`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *AccountViewResponse) GetClaimsOk() (*AccountClaims, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *AccountViewResponse) SetClaims(v AccountClaims)`

SetClaims sets Claims field to given value.


### GetClaimsInfo

`func (o *AccountViewResponse) GetClaimsInfo() AccountClaimsInfo`

GetClaimsInfo returns the ClaimsInfo field if non-nil, zero value otherwise.

### GetClaimsInfoOk

`func (o *AccountViewResponse) GetClaimsInfoOk() (*AccountClaimsInfo, bool)`

GetClaimsInfoOk returns a tuple with the ClaimsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimsInfo

`func (o *AccountViewResponse) SetClaimsInfo(v AccountClaimsInfo)`

SetClaimsInfo sets ClaimsInfo field to given value.


### GetCreated

`func (o *AccountViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AccountViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AccountViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *AccountViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemAccount

`func (o *AccountViewResponse) GetIsSystemAccount() bool`

GetIsSystemAccount returns the IsSystemAccount field if non-nil, zero value otherwise.

### GetIsSystemAccountOk

`func (o *AccountViewResponse) GetIsSystemAccountOk() (*bool, bool)`

GetIsSystemAccountOk returns a tuple with the IsSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAccount

`func (o *AccountViewResponse) SetIsSystemAccount(v bool)`

SetIsSystemAccount sets IsSystemAccount field to given value.


### GetJwt

`func (o *AccountViewResponse) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *AccountViewResponse) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *AccountViewResponse) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetJwtSettings

`func (o *AccountViewResponse) GetJwtSettings() AccountJWTSettings`

GetJwtSettings returns the JwtSettings field if non-nil, zero value otherwise.

### GetJwtSettingsOk

`func (o *AccountViewResponse) GetJwtSettingsOk() (*AccountJWTSettings, bool)`

GetJwtSettingsOk returns a tuple with the JwtSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtSettings

`func (o *AccountViewResponse) SetJwtSettings(v AccountJWTSettings)`

SetJwtSettings sets JwtSettings field to given value.


### GetName

`func (o *AccountViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSystem

`func (o *AccountViewResponse) GetSystem() SystemInfo`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AccountViewResponse) GetSystemOk() (*SystemInfo, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AccountViewResponse) SetSystem(v SystemInfo)`

SetSystem sets System field to given value.


### GetUserJwtExpiresInSecs

`func (o *AccountViewResponse) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *AccountViewResponse) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *AccountViewResponse) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


