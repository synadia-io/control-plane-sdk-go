# AppPolicyStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **[]string** |  | 
**Effect** | [**[]AppRoleEffect**](AppRoleEffect.md) |  | 
**Name** | **string** |  | 
**OperationId** | Pointer to **string** |  | [optional] 
**Resource** | **[]string** |  | 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewAppPolicyStatement

`func NewAppPolicyStatement(action []string, effect []AppRoleEffect, name string, resource []string, ) *AppPolicyStatement`

NewAppPolicyStatement instantiates a new AppPolicyStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPolicyStatementWithDefaults

`func NewAppPolicyStatementWithDefaults() *AppPolicyStatement`

NewAppPolicyStatementWithDefaults instantiates a new AppPolicyStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AppPolicyStatement) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AppPolicyStatement) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AppPolicyStatement) SetAction(v []string)`

SetAction sets Action field to given value.


### GetEffect

`func (o *AppPolicyStatement) GetEffect() []AppRoleEffect`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *AppPolicyStatement) GetEffectOk() (*[]AppRoleEffect, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *AppPolicyStatement) SetEffect(v []AppRoleEffect)`

SetEffect sets Effect field to given value.


### GetName

`func (o *AppPolicyStatement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppPolicyStatement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppPolicyStatement) SetName(v string)`

SetName sets Name field to given value.


### GetOperationId

`func (o *AppPolicyStatement) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *AppPolicyStatement) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *AppPolicyStatement) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *AppPolicyStatement) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetResource

`func (o *AppPolicyStatement) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AppPolicyStatement) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AppPolicyStatement) SetResource(v []string)`

SetResource sets Resource field to given value.


### GetService

`func (o *AppPolicyStatement) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AppPolicyStatement) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AppPolicyStatement) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AppPolicyStatement) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


