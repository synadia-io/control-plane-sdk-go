# ImportsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**LocalSubject** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Share** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 
**To** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ExportType**](ExportType.md) |  | [optional] 

## Methods

### NewImportsInner

`func NewImportsInner() *ImportsInner`

NewImportsInner instantiates a new ImportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportsInnerWithDefaults

`func NewImportsInnerWithDefaults() *ImportsInner`

NewImportsInnerWithDefaults instantiates a new ImportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ImportsInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ImportsInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ImportsInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ImportsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLocalSubject

`func (o *ImportsInner) GetLocalSubject() string`

GetLocalSubject returns the LocalSubject field if non-nil, zero value otherwise.

### GetLocalSubjectOk

`func (o *ImportsInner) GetLocalSubjectOk() (*string, bool)`

GetLocalSubjectOk returns a tuple with the LocalSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubject

`func (o *ImportsInner) SetLocalSubject(v string)`

SetLocalSubject sets LocalSubject field to given value.

### HasLocalSubject

`func (o *ImportsInner) HasLocalSubject() bool`

HasLocalSubject returns a boolean if a field has been set.

### GetName

`func (o *ImportsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *ImportsInner) GetShare() bool`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *ImportsInner) GetShareOk() (*bool, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *ImportsInner) SetShare(v bool)`

SetShare sets Share field to given value.

### HasShare

`func (o *ImportsInner) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSubject

`func (o *ImportsInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ImportsInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ImportsInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ImportsInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *ImportsInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ImportsInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ImportsInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ImportsInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToken

`func (o *ImportsInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ImportsInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ImportsInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ImportsInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *ImportsInner) GetType() ExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportsInner) GetTypeOk() (*ExportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportsInner) SetType(v ExportType)`

SetType sets Type field to given value.

### HasType

`func (o *ImportsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


