# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPublicKey** | **string** |  | 
**Id** | **string** |  | 
**IsScpAccount** | **bool** |  | 
**IsSystemAccount** | **bool** |  | 
**Name** | **string** |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewAccountInfo

`func NewAccountInfo(accountPublicKey string, id string, isScpAccount bool, isSystemAccount bool, name string, userJwtExpiresInSecs int64, ) *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPublicKey

`func (o *AccountInfo) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *AccountInfo) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *AccountInfo) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetId

`func (o *AccountInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountInfo) SetId(v string)`

SetId sets Id field to given value.


### GetIsScpAccount

`func (o *AccountInfo) GetIsScpAccount() bool`

GetIsScpAccount returns the IsScpAccount field if non-nil, zero value otherwise.

### GetIsScpAccountOk

`func (o *AccountInfo) GetIsScpAccountOk() (*bool, bool)`

GetIsScpAccountOk returns a tuple with the IsScpAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScpAccount

`func (o *AccountInfo) SetIsScpAccount(v bool)`

SetIsScpAccount sets IsScpAccount field to given value.


### GetIsSystemAccount

`func (o *AccountInfo) GetIsSystemAccount() bool`

GetIsSystemAccount returns the IsSystemAccount field if non-nil, zero value otherwise.

### GetIsSystemAccountOk

`func (o *AccountInfo) GetIsSystemAccountOk() (*bool, bool)`

GetIsSystemAccountOk returns a tuple with the IsSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAccount

`func (o *AccountInfo) SetIsSystemAccount(v bool)`

SetIsSystemAccount sets IsSystemAccount field to given value.


### GetName

`func (o *AccountInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountInfo) SetName(v string)`

SetName sets Name field to given value.


### GetUserJwtExpiresInSecs

`func (o *AccountInfo) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *AccountInfo) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *AccountInfo) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


