# AppUserAssignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountInfo**](AccountInfo.md) |  | [optional] 
**AppUser** | [**AppUserInfo**](AppUserInfo.md) |  | 
**Created** | **time.Time** |  | 
**NatsUser** | Pointer to [**NatsUserInfo**](NatsUserInfo.md) |  | [optional] 
**ResourceId** | **string** |  | 
**RoleId** | **string** |  | 
**RoleName** | **string** |  | 
**Scope** | [**AppRoleScope**](AppRoleScope.md) |  | 
**System** | Pointer to [**SystemInfo**](SystemInfo.md) |  | [optional] 
**Team** | Pointer to [**TeamInfo**](TeamInfo.md) |  | [optional] 

## Methods

### NewAppUserAssignResponse

`func NewAppUserAssignResponse(appUser AppUserInfo, created time.Time, resourceId string, roleId string, roleName string, scope AppRoleScope, ) *AppUserAssignResponse`

NewAppUserAssignResponse instantiates a new AppUserAssignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserAssignResponseWithDefaults

`func NewAppUserAssignResponseWithDefaults() *AppUserAssignResponse`

NewAppUserAssignResponseWithDefaults instantiates a new AppUserAssignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AppUserAssignResponse) GetAccount() AccountInfo`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AppUserAssignResponse) GetAccountOk() (*AccountInfo, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AppUserAssignResponse) SetAccount(v AccountInfo)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AppUserAssignResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAppUser

`func (o *AppUserAssignResponse) GetAppUser() AppUserInfo`

GetAppUser returns the AppUser field if non-nil, zero value otherwise.

### GetAppUserOk

`func (o *AppUserAssignResponse) GetAppUserOk() (*AppUserInfo, bool)`

GetAppUserOk returns a tuple with the AppUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUser

`func (o *AppUserAssignResponse) SetAppUser(v AppUserInfo)`

SetAppUser sets AppUser field to given value.


### GetCreated

`func (o *AppUserAssignResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUserAssignResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUserAssignResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetNatsUser

`func (o *AppUserAssignResponse) GetNatsUser() NatsUserInfo`

GetNatsUser returns the NatsUser field if non-nil, zero value otherwise.

### GetNatsUserOk

`func (o *AppUserAssignResponse) GetNatsUserOk() (*NatsUserInfo, bool)`

GetNatsUserOk returns a tuple with the NatsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatsUser

`func (o *AppUserAssignResponse) SetNatsUser(v NatsUserInfo)`

SetNatsUser sets NatsUser field to given value.

### HasNatsUser

`func (o *AppUserAssignResponse) HasNatsUser() bool`

HasNatsUser returns a boolean if a field has been set.

### GetResourceId

`func (o *AppUserAssignResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AppUserAssignResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AppUserAssignResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetRoleId

`func (o *AppUserAssignResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *AppUserAssignResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *AppUserAssignResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetRoleName

`func (o *AppUserAssignResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AppUserAssignResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AppUserAssignResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetScope

`func (o *AppUserAssignResponse) GetScope() AppRoleScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppUserAssignResponse) GetScopeOk() (*AppRoleScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppUserAssignResponse) SetScope(v AppRoleScope)`

SetScope sets Scope field to given value.


### GetSystem

`func (o *AppUserAssignResponse) GetSystem() SystemInfo`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AppUserAssignResponse) GetSystemOk() (*SystemInfo, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AppUserAssignResponse) SetSystem(v SystemInfo)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AppUserAssignResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetTeam

`func (o *AppUserAssignResponse) GetTeam() TeamInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *AppUserAssignResponse) GetTeamOk() (*TeamInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *AppUserAssignResponse) SetTeam(v TeamInfo)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *AppUserAssignResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


