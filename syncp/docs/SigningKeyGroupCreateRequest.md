# SigningKeyGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Scope** | Pointer to [**UserPermissionLimits**](UserPermissionLimits.md) |  | [optional] 
**Seed** | Pointer to **string** |  | [optional] 

## Methods

### NewSigningKeyGroupCreateRequest

`func NewSigningKeyGroupCreateRequest(name string, ) *SigningKeyGroupCreateRequest`

NewSigningKeyGroupCreateRequest instantiates a new SigningKeyGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyGroupCreateRequestWithDefaults

`func NewSigningKeyGroupCreateRequestWithDefaults() *SigningKeyGroupCreateRequest`

NewSigningKeyGroupCreateRequestWithDefaults instantiates a new SigningKeyGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SigningKeyGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SigningKeyGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SigningKeyGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *SigningKeyGroupCreateRequest) GetScope() UserPermissionLimits`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SigningKeyGroupCreateRequest) GetScopeOk() (*UserPermissionLimits, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SigningKeyGroupCreateRequest) SetScope(v UserPermissionLimits)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SigningKeyGroupCreateRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeed

`func (o *SigningKeyGroupCreateRequest) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *SigningKeyGroupCreateRequest) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *SigningKeyGroupCreateRequest) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *SigningKeyGroupCreateRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


