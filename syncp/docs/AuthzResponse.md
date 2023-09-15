# AuthzResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | **[]string** |  | 
**ResourceId** | **string** |  | 
**Service** | **string** |  | 

## Methods

### NewAuthzResponse

`func NewAuthzResponse(operations []string, resourceId string, service string, ) *AuthzResponse`

NewAuthzResponse instantiates a new AuthzResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzResponseWithDefaults

`func NewAuthzResponseWithDefaults() *AuthzResponse`

NewAuthzResponseWithDefaults instantiates a new AuthzResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *AuthzResponse) GetOperations() []string`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *AuthzResponse) GetOperationsOk() (*[]string, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *AuthzResponse) SetOperations(v []string)`

SetOperations sets Operations field to given value.


### GetResourceId

`func (o *AuthzResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AuthzResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AuthzResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetService

`func (o *AuthzResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AuthzResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AuthzResponse) SetService(v string)`

SetService sets Service field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


