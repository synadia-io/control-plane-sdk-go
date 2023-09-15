# SystemAccountImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | **string** |  | 
**Seed** | **string** |  | 
**SigningKeys** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSystemAccountImportRequest

`func NewSystemAccountImportRequest(jwt string, seed string, ) *SystemAccountImportRequest`

NewSystemAccountImportRequest instantiates a new SystemAccountImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemAccountImportRequestWithDefaults

`func NewSystemAccountImportRequestWithDefaults() *SystemAccountImportRequest`

NewSystemAccountImportRequestWithDefaults instantiates a new SystemAccountImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *SystemAccountImportRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SystemAccountImportRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SystemAccountImportRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetSeed

`func (o *SystemAccountImportRequest) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *SystemAccountImportRequest) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *SystemAccountImportRequest) SetSeed(v string)`

SetSeed sets Seed field to given value.


### GetSigningKeys

`func (o *SystemAccountImportRequest) GetSigningKeys() []string`

GetSigningKeys returns the SigningKeys field if non-nil, zero value otherwise.

### GetSigningKeysOk

`func (o *SystemAccountImportRequest) GetSigningKeysOk() (*[]string, bool)`

GetSigningKeysOk returns a tuple with the SigningKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeys

`func (o *SystemAccountImportRequest) SetSigningKeys(v []string)`

SetSigningKeys sets SigningKeys field to given value.

### HasSigningKeys

`func (o *SystemAccountImportRequest) HasSigningKeys() bool`

HasSigningKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


