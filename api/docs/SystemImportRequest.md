# SystemImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JetstreamDomain** | Pointer to **NullableString** |  | [optional] 
**JetstreamEnabled** | Pointer to **bool** |  | [optional] [default to true]
**OperatorJwt** | **string** |  | 
**OperatorKey** | **string** |  | 
**SystemJwt** | **string** |  | 
**SystemKey** | **string** |  | 
**TeamId** | **string** |  | 

## Methods

### NewSystemImportRequest

`func NewSystemImportRequest(operatorJwt string, operatorKey string, systemJwt string, systemKey string, teamId string, ) *SystemImportRequest`

NewSystemImportRequest instantiates a new SystemImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemImportRequestWithDefaults

`func NewSystemImportRequestWithDefaults() *SystemImportRequest`

NewSystemImportRequestWithDefaults instantiates a new SystemImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJetstreamDomain

`func (o *SystemImportRequest) GetJetstreamDomain() string`

GetJetstreamDomain returns the JetstreamDomain field if non-nil, zero value otherwise.

### GetJetstreamDomainOk

`func (o *SystemImportRequest) GetJetstreamDomainOk() (*string, bool)`

GetJetstreamDomainOk returns a tuple with the JetstreamDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamDomain

`func (o *SystemImportRequest) SetJetstreamDomain(v string)`

SetJetstreamDomain sets JetstreamDomain field to given value.

### HasJetstreamDomain

`func (o *SystemImportRequest) HasJetstreamDomain() bool`

HasJetstreamDomain returns a boolean if a field has been set.

### SetJetstreamDomainNil

`func (o *SystemImportRequest) SetJetstreamDomainNil(b bool)`

 SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil

### UnsetJetstreamDomain
`func (o *SystemImportRequest) UnsetJetstreamDomain()`

UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
### GetJetstreamEnabled

`func (o *SystemImportRequest) GetJetstreamEnabled() bool`

GetJetstreamEnabled returns the JetstreamEnabled field if non-nil, zero value otherwise.

### GetJetstreamEnabledOk

`func (o *SystemImportRequest) GetJetstreamEnabledOk() (*bool, bool)`

GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamEnabled

`func (o *SystemImportRequest) SetJetstreamEnabled(v bool)`

SetJetstreamEnabled sets JetstreamEnabled field to given value.

### HasJetstreamEnabled

`func (o *SystemImportRequest) HasJetstreamEnabled() bool`

HasJetstreamEnabled returns a boolean if a field has been set.

### GetOperatorJwt

`func (o *SystemImportRequest) GetOperatorJwt() string`

GetOperatorJwt returns the OperatorJwt field if non-nil, zero value otherwise.

### GetOperatorJwtOk

`func (o *SystemImportRequest) GetOperatorJwtOk() (*string, bool)`

GetOperatorJwtOk returns a tuple with the OperatorJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorJwt

`func (o *SystemImportRequest) SetOperatorJwt(v string)`

SetOperatorJwt sets OperatorJwt field to given value.


### GetOperatorKey

`func (o *SystemImportRequest) GetOperatorKey() string`

GetOperatorKey returns the OperatorKey field if non-nil, zero value otherwise.

### GetOperatorKeyOk

`func (o *SystemImportRequest) GetOperatorKeyOk() (*string, bool)`

GetOperatorKeyOk returns a tuple with the OperatorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorKey

`func (o *SystemImportRequest) SetOperatorKey(v string)`

SetOperatorKey sets OperatorKey field to given value.


### GetSystemJwt

`func (o *SystemImportRequest) GetSystemJwt() string`

GetSystemJwt returns the SystemJwt field if non-nil, zero value otherwise.

### GetSystemJwtOk

`func (o *SystemImportRequest) GetSystemJwtOk() (*string, bool)`

GetSystemJwtOk returns a tuple with the SystemJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemJwt

`func (o *SystemImportRequest) SetSystemJwt(v string)`

SetSystemJwt sets SystemJwt field to given value.


### GetSystemKey

`func (o *SystemImportRequest) GetSystemKey() string`

GetSystemKey returns the SystemKey field if non-nil, zero value otherwise.

### GetSystemKeyOk

`func (o *SystemImportRequest) GetSystemKeyOk() (*string, bool)`

GetSystemKeyOk returns a tuple with the SystemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemKey

`func (o *SystemImportRequest) SetSystemKey(v string)`

SetSystemKey sets SystemKey field to given value.


### GetTeamId

`func (o *SystemImportRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SystemImportRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SystemImportRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


