# UserLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | Pointer to **[]string** |  | [optional] 
**Times** | Pointer to [**[]TimeRange**](TimeRange.md) |  | [optional] 
**TimesLocation** | Pointer to **string** |  | [optional] 

## Methods

### NewUserLimits

`func NewUserLimits() *UserLimits`

NewUserLimits instantiates a new UserLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLimitsWithDefaults

`func NewUserLimitsWithDefaults() *UserLimits`

NewUserLimitsWithDefaults instantiates a new UserLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *UserLimits) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *UserLimits) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *UserLimits) SetSrc(v []string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *UserLimits) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### SetSrcNil

`func (o *UserLimits) SetSrcNil(b bool)`

 SetSrcNil sets the value for Src to be an explicit nil

### UnsetSrc
`func (o *UserLimits) UnsetSrc()`

UnsetSrc ensures that no value is present for Src, not even an explicit nil
### GetTimes

`func (o *UserLimits) GetTimes() []TimeRange`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *UserLimits) GetTimesOk() (*[]TimeRange, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *UserLimits) SetTimes(v []TimeRange)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *UserLimits) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimesLocation

`func (o *UserLimits) GetTimesLocation() string`

GetTimesLocation returns the TimesLocation field if non-nil, zero value otherwise.

### GetTimesLocationOk

`func (o *UserLimits) GetTimesLocationOk() (*string, bool)`

GetTimesLocationOk returns a tuple with the TimesLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimesLocation

`func (o *UserLimits) SetTimesLocation(v string)`

SetTimesLocation sets TimesLocation field to given value.

### HasTimesLocation

`func (o *UserLimits) HasTimesLocation() bool`

HasTimesLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


