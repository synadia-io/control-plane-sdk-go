# Authorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAccounts** | Pointer to **[]string** | StringList is a wrapper for an array of strings | [optional] 
**AuthUsers** | **[]string** | StringList is a wrapper for an array of strings | 
**Xkey** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthorization

`func NewAuthorization(authUsers []string, ) *Authorization`

NewAuthorization instantiates a new Authorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationWithDefaults

`func NewAuthorizationWithDefaults() *Authorization`

NewAuthorizationWithDefaults instantiates a new Authorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAccounts

`func (o *Authorization) GetAllowedAccounts() []string`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *Authorization) GetAllowedAccountsOk() (*[]string, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccounts

`func (o *Authorization) SetAllowedAccounts(v []string)`

SetAllowedAccounts sets AllowedAccounts field to given value.

### HasAllowedAccounts

`func (o *Authorization) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### SetAllowedAccountsNil

`func (o *Authorization) SetAllowedAccountsNil(b bool)`

 SetAllowedAccountsNil sets the value for AllowedAccounts to be an explicit nil

### UnsetAllowedAccounts
`func (o *Authorization) UnsetAllowedAccounts()`

UnsetAllowedAccounts ensures that no value is present for AllowedAccounts, not even an explicit nil
### GetAuthUsers

`func (o *Authorization) GetAuthUsers() []string`

GetAuthUsers returns the AuthUsers field if non-nil, zero value otherwise.

### GetAuthUsersOk

`func (o *Authorization) GetAuthUsersOk() (*[]string, bool)`

GetAuthUsersOk returns a tuple with the AuthUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUsers

`func (o *Authorization) SetAuthUsers(v []string)`

SetAuthUsers sets AuthUsers field to given value.


### SetAuthUsersNil

`func (o *Authorization) SetAuthUsersNil(b bool)`

 SetAuthUsersNil sets the value for AuthUsers to be an explicit nil

### UnsetAuthUsers
`func (o *Authorization) UnsetAuthUsers()`

UnsetAuthUsers ensures that no value is present for AuthUsers, not even an explicit nil
### GetXkey

`func (o *Authorization) GetXkey() string`

GetXkey returns the Xkey field if non-nil, zero value otherwise.

### GetXkeyOk

`func (o *Authorization) GetXkeyOk() (*string, bool)`

GetXkeyOk returns a tuple with the Xkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXkey

`func (o *Authorization) SetXkey(v string)`

SetXkey sets Xkey field to given value.

### HasXkey

`func (o *Authorization) HasXkey() bool`

HasXkey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


