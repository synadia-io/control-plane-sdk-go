# GenericFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Type** | Pointer to **string** | ClaimType is used to indicate the type of JWT being stored in a Claim | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewGenericFields

`func NewGenericFields() *GenericFields`

NewGenericFields instantiates a new GenericFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericFieldsWithDefaults

`func NewGenericFieldsWithDefaults() *GenericFields`

NewGenericFieldsWithDefaults instantiates a new GenericFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *GenericFields) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GenericFields) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GenericFields) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GenericFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GenericFields) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GenericFields) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetType

`func (o *GenericFields) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GenericFields) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GenericFields) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GenericFields) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *GenericFields) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GenericFields) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GenericFields) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GenericFields) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


