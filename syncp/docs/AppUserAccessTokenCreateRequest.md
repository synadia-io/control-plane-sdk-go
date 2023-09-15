# AppUserAccessTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | **NullableString** |  | 
**Name** | **string** |  | 

## Methods

### NewAppUserAccessTokenCreateRequest

`func NewAppUserAccessTokenCreateRequest(expires NullableString, name string, ) *AppUserAccessTokenCreateRequest`

NewAppUserAccessTokenCreateRequest instantiates a new AppUserAccessTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserAccessTokenCreateRequestWithDefaults

`func NewAppUserAccessTokenCreateRequestWithDefaults() *AppUserAccessTokenCreateRequest`

NewAppUserAccessTokenCreateRequestWithDefaults instantiates a new AppUserAccessTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *AppUserAccessTokenCreateRequest) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AppUserAccessTokenCreateRequest) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AppUserAccessTokenCreateRequest) SetExpires(v string)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *AppUserAccessTokenCreateRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *AppUserAccessTokenCreateRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *AppUserAccessTokenCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserAccessTokenCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserAccessTokenCreateRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


