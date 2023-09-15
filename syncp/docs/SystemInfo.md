# SystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JetstreamDomain** | Pointer to **NullableString** |  | [optional] 
**JetstreamEnabled** | **bool** |  | 
**Name** | **string** |  | 
**ServerUrls** | Pointer to **string** |  | [optional] 
**State** | [**SystemState**](SystemState.md) |  | 
**UserJwtExpiresInSecs** | **int64** |  | 

## Methods

### NewSystemInfo

`func NewSystemInfo(id string, jetstreamEnabled bool, name string, state SystemState, userJwtExpiresInSecs int64, ) *SystemInfo`

NewSystemInfo instantiates a new SystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInfoWithDefaults

`func NewSystemInfoWithDefaults() *SystemInfo`

NewSystemInfoWithDefaults instantiates a new SystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemInfo) SetId(v string)`

SetId sets Id field to given value.


### GetJetstreamDomain

`func (o *SystemInfo) GetJetstreamDomain() string`

GetJetstreamDomain returns the JetstreamDomain field if non-nil, zero value otherwise.

### GetJetstreamDomainOk

`func (o *SystemInfo) GetJetstreamDomainOk() (*string, bool)`

GetJetstreamDomainOk returns a tuple with the JetstreamDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamDomain

`func (o *SystemInfo) SetJetstreamDomain(v string)`

SetJetstreamDomain sets JetstreamDomain field to given value.

### HasJetstreamDomain

`func (o *SystemInfo) HasJetstreamDomain() bool`

HasJetstreamDomain returns a boolean if a field has been set.

### SetJetstreamDomainNil

`func (o *SystemInfo) SetJetstreamDomainNil(b bool)`

 SetJetstreamDomainNil sets the value for JetstreamDomain to be an explicit nil

### UnsetJetstreamDomain
`func (o *SystemInfo) UnsetJetstreamDomain()`

UnsetJetstreamDomain ensures that no value is present for JetstreamDomain, not even an explicit nil
### GetJetstreamEnabled

`func (o *SystemInfo) GetJetstreamEnabled() bool`

GetJetstreamEnabled returns the JetstreamEnabled field if non-nil, zero value otherwise.

### GetJetstreamEnabledOk

`func (o *SystemInfo) GetJetstreamEnabledOk() (*bool, bool)`

GetJetstreamEnabledOk returns a tuple with the JetstreamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstreamEnabled

`func (o *SystemInfo) SetJetstreamEnabled(v bool)`

SetJetstreamEnabled sets JetstreamEnabled field to given value.


### GetName

`func (o *SystemInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInfo) SetName(v string)`

SetName sets Name field to given value.


### GetServerUrls

`func (o *SystemInfo) GetServerUrls() string`

GetServerUrls returns the ServerUrls field if non-nil, zero value otherwise.

### GetServerUrlsOk

`func (o *SystemInfo) GetServerUrlsOk() (*string, bool)`

GetServerUrlsOk returns a tuple with the ServerUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrls

`func (o *SystemInfo) SetServerUrls(v string)`

SetServerUrls sets ServerUrls field to given value.

### HasServerUrls

`func (o *SystemInfo) HasServerUrls() bool`

HasServerUrls returns a boolean if a field has been set.

### GetState

`func (o *SystemInfo) GetState() SystemState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemInfo) GetStateOk() (*SystemState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemInfo) SetState(v SystemState)`

SetState sets State field to given value.


### GetUserJwtExpiresInSecs

`func (o *SystemInfo) GetUserJwtExpiresInSecs() int64`

GetUserJwtExpiresInSecs returns the UserJwtExpiresInSecs field if non-nil, zero value otherwise.

### GetUserJwtExpiresInSecsOk

`func (o *SystemInfo) GetUserJwtExpiresInSecsOk() (*int64, bool)`

GetUserJwtExpiresInSecsOk returns a tuple with the UserJwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJwtExpiresInSecs

`func (o *SystemInfo) SetUserJwtExpiresInSecs(v int64)`

SetUserJwtExpiresInSecs sets UserJwtExpiresInSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


