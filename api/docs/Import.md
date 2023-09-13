# Import

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

### NewImport

`func NewImport() *Import`

NewImport instantiates a new Import object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportWithDefaults

`func NewImportWithDefaults() *Import`

NewImportWithDefaults instantiates a new Import object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Import) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Import) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Import) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Import) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetLocalSubject

`func (o *Import) GetLocalSubject() string`

GetLocalSubject returns the LocalSubject field if non-nil, zero value otherwise.

### GetLocalSubjectOk

`func (o *Import) GetLocalSubjectOk() (*string, bool)`

GetLocalSubjectOk returns a tuple with the LocalSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubject

`func (o *Import) SetLocalSubject(v string)`

SetLocalSubject sets LocalSubject field to given value.

### HasLocalSubject

`func (o *Import) HasLocalSubject() bool`

HasLocalSubject returns a boolean if a field has been set.

### GetName

`func (o *Import) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Import) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Import) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Import) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShare

`func (o *Import) GetShare() bool`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *Import) GetShareOk() (*bool, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *Import) SetShare(v bool)`

SetShare sets Share field to given value.

### HasShare

`func (o *Import) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetSubject

`func (o *Import) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Import) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Import) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Import) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTo

`func (o *Import) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Import) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Import) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *Import) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToken

`func (o *Import) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Import) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Import) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Import) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *Import) GetType() ExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Import) GetTypeOk() (*ExportType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Import) SetType(v ExportType)`

SetType sets Type field to given value.

### HasType

`func (o *Import) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


