# AlertViewResponseAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPublicKey** | **string** |  | 
**Id** | **string** |  | 
**IsSystemAccount** | **bool** |  | 
**Name** | **string** |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewAlertViewResponseAccount

`func NewAlertViewResponseAccount(accountPublicKey string, id string, isSystemAccount bool, name string, userJwtExpiresInSecs int64, ) *AlertViewResponseAccount`

NewAlertViewResponseAccount instantiates a new AlertViewResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertViewResponseAccountWithDefaults

`func NewAlertViewResponseAccountWithDefaults() *AlertViewResponseAccount`

NewAlertViewResponseAccountWithDefaults instantiates a new AlertViewResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPublicKey

`func (o *AlertViewResponseAccount) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *AlertViewResponseAccount) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *AlertViewResponseAccount) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetId

`func (o *AlertViewResponseAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertViewResponseAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertViewResponseAccount) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemAccount

`func (o *AlertViewResponseAccount) GetIsSystemAccount() bool`

GetIsSystemAccount returns the IsSystemAccount field if non-nil, zero value otherwise.

### GetIsSystemAccountOk

`func (o *AlertViewResponseAccount) GetIsSystemAccountOk() (*bool, bool)`

GetIsSystemAccountOk returns a tuple with the IsSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAccount

`func (o *AlertViewResponseAccount) SetIsSystemAccount(v bool)`

SetIsSystemAccount sets IsSystemAccount field to given value.


### GetName

`func (o *AlertViewResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertViewResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertViewResponseAccount) SetName(v string)`

SetName sets Name field to given value.


### GetUserJwtExpiresInSecs

`func (o *AlertViewResponseAccount) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *AlertViewResponseAccount) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *AlertViewResponseAccount) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


