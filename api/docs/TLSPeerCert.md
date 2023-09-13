# TLSPeerCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertSha256** | Pointer to **string** |  | [optional] 
**SpkiSha256** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 

## Methods

### NewTLSPeerCert

`func NewTLSPeerCert() *TLSPeerCert`

NewTLSPeerCert instantiates a new TLSPeerCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSPeerCertWithDefaults

`func NewTLSPeerCertWithDefaults() *TLSPeerCert`

NewTLSPeerCertWithDefaults instantiates a new TLSPeerCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertSha256

`func (o *TLSPeerCert) GetCertSha256() string`

GetCertSha256 returns the CertSha256 field if non-nil, zero value otherwise.

### GetCertSha256Ok

`func (o *TLSPeerCert) GetCertSha256Ok() (*string, bool)`

GetCertSha256Ok returns a tuple with the CertSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSha256

`func (o *TLSPeerCert) SetCertSha256(v string)`

SetCertSha256 sets CertSha256 field to given value.

### HasCertSha256

`func (o *TLSPeerCert) HasCertSha256() bool`

HasCertSha256 returns a boolean if a field has been set.

### GetSpkiSha256

`func (o *TLSPeerCert) GetSpkiSha256() string`

GetSpkiSha256 returns the SpkiSha256 field if non-nil, zero value otherwise.

### GetSpkiSha256Ok

`func (o *TLSPeerCert) GetSpkiSha256Ok() (*string, bool)`

GetSpkiSha256Ok returns a tuple with the SpkiSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpkiSha256

`func (o *TLSPeerCert) SetSpkiSha256(v string)`

SetSpkiSha256 sets SpkiSha256 field to given value.

### HasSpkiSha256

`func (o *TLSPeerCert) HasSpkiSha256() bool`

HasSpkiSha256 returns a boolean if a field has been set.

### GetSubject

`func (o *TLSPeerCert) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TLSPeerCert) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TLSPeerCert) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *TLSPeerCert) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


