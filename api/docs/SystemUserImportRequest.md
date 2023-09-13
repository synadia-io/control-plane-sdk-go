# SystemUserImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to **string** |  | [optional] 
**Jwt** | Pointer to **string** |  | [optional] 
**Seed** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemUserImportRequest

`func NewSystemUserImportRequest() *SystemUserImportRequest`

NewSystemUserImportRequest instantiates a new SystemUserImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemUserImportRequestWithDefaults

`func NewSystemUserImportRequestWithDefaults() *SystemUserImportRequest`

NewSystemUserImportRequestWithDefaults instantiates a new SystemUserImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *SystemUserImportRequest) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *SystemUserImportRequest) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *SystemUserImportRequest) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *SystemUserImportRequest) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetJwt

`func (o *SystemUserImportRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SystemUserImportRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SystemUserImportRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *SystemUserImportRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetSeed

`func (o *SystemUserImportRequest) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *SystemUserImportRequest) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *SystemUserImportRequest) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *SystemUserImportRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


