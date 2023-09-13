# AppPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Statements** | [**map[string]AppPolicyStatement**](AppPolicyStatement.md) |  | 

## Methods

### NewAppPolicy

`func NewAppPolicy(name string, statements map[string]AppPolicyStatement, ) *AppPolicy`

NewAppPolicy instantiates a new AppPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPolicyWithDefaults

`func NewAppPolicyWithDefaults() *AppPolicy`

NewAppPolicyWithDefaults instantiates a new AppPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetStatements

`func (o *AppPolicy) GetStatements() map[string]AppPolicyStatement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *AppPolicy) GetStatementsOk() (*map[string]AppPolicyStatement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *AppPolicy) SetStatements(v map[string]AppPolicyStatement)`

SetStatements sets Statements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


