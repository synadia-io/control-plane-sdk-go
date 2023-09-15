# SigningKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** |  | 
**Seed** | Pointer to **string** |  | [optional] 

## Methods

### NewSigningKeyUpdateRequest

`func NewSigningKeyUpdateRequest(disabled bool, ) *SigningKeyUpdateRequest`

NewSigningKeyUpdateRequest instantiates a new SigningKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyUpdateRequestWithDefaults

`func NewSigningKeyUpdateRequestWithDefaults() *SigningKeyUpdateRequest`

NewSigningKeyUpdateRequestWithDefaults instantiates a new SigningKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SigningKeyUpdateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SigningKeyUpdateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SigningKeyUpdateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetSeed

`func (o *SigningKeyUpdateRequest) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *SigningKeyUpdateRequest) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *SigningKeyUpdateRequest) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *SigningKeyUpdateRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


