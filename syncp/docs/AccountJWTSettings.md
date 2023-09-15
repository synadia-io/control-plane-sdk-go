# AccountJWTSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**Authorization**](Authorization.md) |  | [optional] 
**Limits** | Pointer to [**OperatorLimits**](OperatorLimits.md) |  | [optional] 
**Mappings** | Pointer to [**map[string][]WeightedMapping**](array.md) |  | [optional] 
**Revocations** | Pointer to **map[string]int64** | RevocationList is used to store a mapping of public keys to unix timestamps | [optional] 
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InfoUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountJWTSettings

`func NewAccountJWTSettings() *AccountJWTSettings`

NewAccountJWTSettings instantiates a new AccountJWTSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountJWTSettingsWithDefaults

`func NewAccountJWTSettingsWithDefaults() *AccountJWTSettings`

NewAccountJWTSettingsWithDefaults instantiates a new AccountJWTSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *AccountJWTSettings) GetAuthorization() Authorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *AccountJWTSettings) GetAuthorizationOk() (*Authorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *AccountJWTSettings) SetAuthorization(v Authorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *AccountJWTSettings) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetLimits

`func (o *AccountJWTSettings) GetLimits() OperatorLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *AccountJWTSettings) GetLimitsOk() (*OperatorLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *AccountJWTSettings) SetLimits(v OperatorLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *AccountJWTSettings) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetMappings

`func (o *AccountJWTSettings) GetMappings() map[string][]WeightedMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *AccountJWTSettings) GetMappingsOk() (*map[string][]WeightedMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *AccountJWTSettings) SetMappings(v map[string][]WeightedMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *AccountJWTSettings) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetRevocations

`func (o *AccountJWTSettings) GetRevocations() map[string]int64`

GetRevocations returns the Revocations field if non-nil, zero value otherwise.

### GetRevocationsOk

`func (o *AccountJWTSettings) GetRevocationsOk() (*map[string]int64, bool)`

GetRevocationsOk returns a tuple with the Revocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocations

`func (o *AccountJWTSettings) SetRevocations(v map[string]int64)`

SetRevocations sets Revocations field to given value.

### HasRevocations

`func (o *AccountJWTSettings) HasRevocations() bool`

HasRevocations returns a boolean if a field has been set.

### GetTags

`func (o *AccountJWTSettings) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountJWTSettings) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountJWTSettings) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountJWTSettings) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AccountJWTSettings) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AccountJWTSettings) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDescription

`func (o *AccountJWTSettings) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountJWTSettings) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountJWTSettings) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountJWTSettings) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfoUrl

`func (o *AccountJWTSettings) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *AccountJWTSettings) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *AccountJWTSettings) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *AccountJWTSettings) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


