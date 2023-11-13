# AppUserInviteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppUser** | [**AppUserInfo**](AppUserInfo.md) |  | 
**AppUserCreated** | **bool** |  | 
**AutoAccepted** | **bool** |  | 
**TeamAppUser** | [**TeamAppUserInfo**](TeamAppUserInfo.md) |  | 
**TeamAppUserCreated** | **bool** |  | 

## Methods

### NewAppUserInviteResponse

`func NewAppUserInviteResponse(appUser AppUserInfo, appUserCreated bool, autoAccepted bool, teamAppUser TeamAppUserInfo, teamAppUserCreated bool, ) *AppUserInviteResponse`

NewAppUserInviteResponse instantiates a new AppUserInviteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserInviteResponseWithDefaults

`func NewAppUserInviteResponseWithDefaults() *AppUserInviteResponse`

NewAppUserInviteResponseWithDefaults instantiates a new AppUserInviteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppUser

`func (o *AppUserInviteResponse) GetAppUser() AppUserInfo`

GetAppUser returns the AppUser field if non-nil, zero value otherwise.

### GetAppUserOk

`func (o *AppUserInviteResponse) GetAppUserOk() (*AppUserInfo, bool)`

GetAppUserOk returns a tuple with the AppUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUser

`func (o *AppUserInviteResponse) SetAppUser(v AppUserInfo)`

SetAppUser sets AppUser field to given value.


### GetAppUserCreated

`func (o *AppUserInviteResponse) GetAppUserCreated() bool`

GetAppUserCreated returns the AppUserCreated field if non-nil, zero value otherwise.

### GetAppUserCreatedOk

`func (o *AppUserInviteResponse) GetAppUserCreatedOk() (*bool, bool)`

GetAppUserCreatedOk returns a tuple with the AppUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUserCreated

`func (o *AppUserInviteResponse) SetAppUserCreated(v bool)`

SetAppUserCreated sets AppUserCreated field to given value.


### GetAutoAccepted

`func (o *AppUserInviteResponse) GetAutoAccepted() bool`

GetAutoAccepted returns the AutoAccepted field if non-nil, zero value otherwise.

### GetAutoAcceptedOk

`func (o *AppUserInviteResponse) GetAutoAcceptedOk() (*bool, bool)`

GetAutoAcceptedOk returns a tuple with the AutoAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAccepted

`func (o *AppUserInviteResponse) SetAutoAccepted(v bool)`

SetAutoAccepted sets AutoAccepted field to given value.


### GetTeamAppUser

`func (o *AppUserInviteResponse) GetTeamAppUser() TeamAppUserInfo`

GetTeamAppUser returns the TeamAppUser field if non-nil, zero value otherwise.

### GetTeamAppUserOk

`func (o *AppUserInviteResponse) GetTeamAppUserOk() (*TeamAppUserInfo, bool)`

GetTeamAppUserOk returns a tuple with the TeamAppUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAppUser

`func (o *AppUserInviteResponse) SetTeamAppUser(v TeamAppUserInfo)`

SetTeamAppUser sets TeamAppUser field to given value.


### GetTeamAppUserCreated

`func (o *AppUserInviteResponse) GetTeamAppUserCreated() bool`

GetTeamAppUserCreated returns the TeamAppUserCreated field if non-nil, zero value otherwise.

### GetTeamAppUserCreatedOk

`func (o *AppUserInviteResponse) GetTeamAppUserCreatedOk() (*bool, bool)`

GetTeamAppUserCreatedOk returns a tuple with the TeamAppUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamAppUserCreated

`func (o *AppUserInviteResponse) SetTeamAppUserCreated(v bool)`

SetTeamAppUserCreated sets TeamAppUserCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


