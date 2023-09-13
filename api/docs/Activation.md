# Activation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenericFields** | [**GenericFields**](GenericFields.md) |  | 
**IssuerAccount** | Pointer to **string** | IssuerAccount stores the public key for the account the issuer represents. When set, the claim was issued by a signing key. | [optional] 
**Kind** | Pointer to [**ExportType**](ExportType.md) |  | [optional] 
**Subject** | Pointer to **string** | Subject is a string that represents a NATS subject | [optional] 

## Methods

### NewActivation

`func NewActivation(genericFields GenericFields, ) *Activation`

NewActivation instantiates a new Activation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationWithDefaults

`func NewActivationWithDefaults() *Activation`

NewActivationWithDefaults instantiates a new Activation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenericFields

`func (o *Activation) GetGenericFields() GenericFields`

GetGenericFields returns the GenericFields field if non-nil, zero value otherwise.

### GetGenericFieldsOk

`func (o *Activation) GetGenericFieldsOk() (*GenericFields, bool)`

GetGenericFieldsOk returns a tuple with the GenericFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericFields

`func (o *Activation) SetGenericFields(v GenericFields)`

SetGenericFields sets GenericFields field to given value.


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


