# AccountSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountPublicKey** | **string** |  | 
**Id** | **string** |  | 
**IsSystemAccount** | **bool** |  | 
**Name** | **string** |  | 

## Methods

### NewAccountSearch

`func NewAccountSearch(accountPublicKey string, id string, isSystemAccount bool, name string, ) *AccountSearch`

NewAccountSearch instantiates a new AccountSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSearchWithDefaults

`func NewAccountSearchWithDefaults() *AccountSearch`

NewAccountSearchWithDefaults instantiates a new AccountSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountPublicKey

`func (o *AccountSearch) GetAccountPublicKey() string`

GetAccountPublicKey returns the AccountPublicKey field if non-nil, zero value otherwise.

### GetAccountPublicKeyOk

`func (o *AccountSearch) GetAccountPublicKeyOk() (*string, bool)`

GetAccountPublicKeyOk returns a tuple with the AccountPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPublicKey

`func (o *AccountSearch) SetAccountPublicKey(v string)`

SetAccountPublicKey sets AccountPublicKey field to given value.


### GetId

`func (o *AccountSearch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSearch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSearch) SetId(v string)`

SetId sets Id field to given value.


### GetIsSystemAccount

`func (o *AccountSearch) GetIsSystemAccount() bool`

GetIsSystemAccount returns the IsSystemAccount field if non-nil, zero value otherwise.

### GetIsSystemAccountOk

`func (o *AccountSearch) GetIsSystemAccountOk() (*bool, bool)`

GetIsSystemAccountOk returns a tuple with the IsSystemAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemAccount

`func (o *AccountSearch) SetIsSystemAccount(v bool)`

SetIsSystemAccount sets IsSystemAccount field to given value.


### GetName

`func (o *AccountSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountSearch) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


