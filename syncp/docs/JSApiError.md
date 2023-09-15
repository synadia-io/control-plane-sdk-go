# JSApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ErrCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewJSApiError

`func NewJSApiError(code int32, ) *JSApiError`

NewJSApiError instantiates a new JSApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJSApiErrorWithDefaults

`func NewJSApiErrorWithDefaults() *JSApiError`

NewJSApiErrorWithDefaults instantiates a new JSApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JSApiError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JSApiError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JSApiError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetDescription

`func (o *JSApiError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JSApiError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JSApiError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JSApiError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrCode

`func (o *JSApiError) GetErrCode() int32`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *JSApiError) GetErrCodeOk() (*int32, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *JSApiError) SetErrCode(v int32)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *JSApiError) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


