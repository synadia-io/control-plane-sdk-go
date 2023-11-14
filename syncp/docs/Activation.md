# Activation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** | TagList is a unique array of lower case strings All tag list methods lower case the strings in the arguments | [optional] 
**Type** | Pointer to **string** | ClaimType is used to indicate the type of JWT being stored in a Claim | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**IssuerAccount** | Pointer to **string** | IssuerAccount stores the public key for the account the issuer represents. When set, the claim was issued by a signing key. | [optional] 
**Kind** | Pointer to [**ExportType**](ExportType.md) |  | [optional] 
**Subject** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 

## Methods

### NewActivation

`func NewActivation() *Activation`

NewActivation instantiates a new Activation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationWithDefaults

`func NewActivationWithDefaults() *Activation`

NewActivationWithDefaults instantiates a new Activation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *Activation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Activation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Activation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Activation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Activation) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Activation) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetType

`func (o *Activation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Activation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Activation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Activation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *Activation) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Activation) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Activation) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Activation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIssuerAccount

`func (o *Activation) GetIssuerAccount() string`

GetIssuerAccount returns the IssuerAccount field if non-nil, zero value otherwise.

### GetIssuerAccountOk

`func (o *Activation) GetIssuerAccountOk() (*string, bool)`

GetIssuerAccountOk returns a tuple with the IssuerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAccount

`func (o *Activation) SetIssuerAccount(v string)`

SetIssuerAccount sets IssuerAccount field to given value.

### HasIssuerAccount

`func (o *Activation) HasIssuerAccount() bool`

HasIssuerAccount returns a boolean if a field has been set.

### GetKind

`func (o *Activation) GetKind() ExportType`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Activation) GetKindOk() (*ExportType, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Activation) SetKind(v ExportType)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Activation) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSubject

`func (o *Activation) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *Activation) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *Activation) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *Activation) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


