# StreamSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **int64** |  | 
**Error** | Pointer to [**NullableStreamSourceInfoError**](StreamSourceInfoError.md) |  | [optional] 
**External** | Pointer to [**NullableStreamSourceExternal**](StreamSourceExternal.md) |  | [optional] 
**Lag** | **int32** |  | 
**Name** | **string** |  | 

## Methods

### NewStreamSourceInfo

`func NewStreamSourceInfo(active int64, lag int32, name string, ) *StreamSourceInfo`

NewStreamSourceInfo instantiates a new StreamSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSourceInfoWithDefaults

`func NewStreamSourceInfoWithDefaults() *StreamSourceInfo`

NewStreamSourceInfoWithDefaults instantiates a new StreamSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *StreamSourceInfo) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StreamSourceInfo) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StreamSourceInfo) SetActive(v int64)`

SetActive sets Active field to given value.


### GetError

`func (o *StreamSourceInfo) GetError() StreamSourceInfoError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamSourceInfo) GetErrorOk() (*StreamSourceInfoError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamSourceInfo) SetError(v StreamSourceInfoError)`

SetError sets Error field to given value.

### HasError

`func (o *StreamSourceInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *StreamSourceInfo) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *StreamSourceInfo) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetExternal

`func (o *StreamSourceInfo) GetExternal() StreamSourceExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *StreamSourceInfo) GetExternalOk() (*StreamSourceExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *StreamSourceInfo) SetExternal(v StreamSourceExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *StreamSourceInfo) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### SetExternalNil

`func (o *StreamSourceInfo) SetExternalNil(b bool)`

 SetExternalNil sets the value for External to be an explicit nil

### UnsetExternal
`func (o *StreamSourceInfo) UnsetExternal()`

UnsetExternal ensures that no value is present for External, not even an explicit nil
### GetLag

`func (o *StreamSourceInfo) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *StreamSourceInfo) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *StreamSourceInfo) SetLag(v int32)`

SetLag sets Lag field to given value.


### GetName

`func (o *StreamSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamSourceInfo) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


