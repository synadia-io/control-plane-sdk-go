# TeamAppUserViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUser** | [**AppUserInfo**](AppUserInfo.md) |  | 
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**PendingInvitation** | **bool** |  | 
**RoleId** | **string** |  | 
**RoleName** | **string** |  | 
**Team** | [**TeamInfo**](TeamInfo.md) |  | 

## Methods

### NewTeamAppUserViewResponse

`func NewTeamAppUserViewResponse(appUser AppUserInfo, created time.Time, id string, pendingInvitation bool, roleId string, roleName string, team TeamInfo, ) *TeamAppUserViewResponse`

NewTeamAppUserViewResponse instantiates a new TeamAppUserViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamAppUserViewResponseWithDefaults

`func NewTeamAppUserViewResponseWithDefaults() *TeamAppUserViewResponse`

NewTeamAppUserViewResponseWithDefaults instantiates a new TeamAppUserViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUser

`func (o *TeamAppUserViewResponse) GetAppUser() AppUserInfo`

GetAppUser returns the AppUser field if non-nil, zero value otherwise.

### GetAppUserOk

`func (o *TeamAppUserViewResponse) GetAppUserOk() (*AppUserInfo, bool)`

GetAppUserOk returns a tuple with the AppUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUser

`func (o *TeamAppUserViewResponse) SetAppUser(v AppUserInfo)`

SetAppUser sets AppUser field to given value.


### GetCreated

`func (o *TeamAppUserViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TeamAppUserViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TeamAppUserViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *TeamAppUserViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamAppUserViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamAppUserViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPendingInvitation

`func (o *TeamAppUserViewResponse) GetPendingInvitation() bool`

GetPendingInvitation returns the PendingInvitation field if non-nil, zero value otherwise.

### GetPendingInvitationOk

`func (o *TeamAppUserViewResponse) GetPendingInvitationOk() (*bool, bool)`

GetPendingInvitationOk returns a tuple with the PendingInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingInvitation

`func (o *TeamAppUserViewResponse) SetPendingInvitation(v bool)`

SetPendingInvitation sets PendingInvitation field to given value.


### GetRoleId

`func (o *TeamAppUserViewResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *TeamAppUserViewResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *TeamAppUserViewResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetRoleName

`func (o *TeamAppUserViewResponse) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *TeamAppUserViewResponse) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *TeamAppUserViewResponse) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.


### GetTeam

`func (o *TeamAppUserViewResponse) GetTeam() TeamInfo`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamAppUserViewResponse) GetTeamOk() (*TeamInfo, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamAppUserViewResponse) SetTeam(v TeamInfo)`

SetTeam sets Team field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


