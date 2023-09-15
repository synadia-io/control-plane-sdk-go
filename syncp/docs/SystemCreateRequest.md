# SystemCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JetstreamDomain** | Pointer to **NullableString** |  | [optional] 
**JetstreamEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Name** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSystemCreateRequest

`func NewSystemCreateRequest(name string, ) *SystemCreateRequest`

NewSystemCreateRequest instantiates a new SystemCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemCreateRequestWithDefaults

`func NewSystemCreateRequestWithDefaults() *SystemCreateRequest`

NewSystemCreateRequestWithDefaults instantiates a new SystemCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJetstreamDomain

`func (o *SystemCreateRequest) GetJetstreamDomain() string`

GetJetstreamDomain returns the JetstreamDomain field if non-nil, zero value otherwise.

### GetJetstreamDomainOk

`func (o *SystemCreateRequest) GetJetstreamDomainOk() (*string, bool)`

GetJetstreamDomainOk returns a tuple with the JetstreamDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamDomain

`func (o *SystemCreateRequest) SetJetstreamDomain(v string)`

SetJetstreamDomain sets JetstreamDomain field to given value.

### HasJetstreamDomain

`func (o *SystemCreateRequest) HasJetstreamDomain() bool`

HasJetstreamDomain returns a boolean if a field has been set.

### SetJetstreamDomainNil

`func (o *SystemCreateRequest) SetJetstreamDomainNil(b bool)`

 SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil

### UnsetJetstreamDomain
`func (o *SystemCreateRequest) UnsetJetstreamDomain()`

UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
### GetJetstreamEnabled

`func (o *SystemCreateRequest) GetJetstreamEnabled() bool`

GetJetstreamEnabled returns the JetstreamEnabled field if non-nil, zero value otherwise.

### GetJetstreamEnabledOk

`func (o *SystemCreateRequest) GetJetstreamEnabledOk() (*bool, bool)`

GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamEnabled

`func (o *SystemCreateRequest) SetJetstreamEnabled(v bool)`

SetJetstreamEnabled sets JetstreamEnabled field to given value.

### HasJetstreamEnabled

`func (o *SystemCreateRequest) HasJetstreamEnabled() bool`

HasJetstreamEnabled returns a boolean if a field has been set.

### GetName

`func (o *SystemCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *SystemCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SystemCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SystemCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SystemCreateRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SystemCreateRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SystemCreateRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


