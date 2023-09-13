# AppUserAccessTokenUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewAppUserAccessTokenUpdateRequest

`func NewAppUserAccessTokenUpdateRequest() *AppUserAccessTokenUpdateRequest`

NewAppUserAccessTokenUpdateRequest instantiates a new AppUserAccessTokenUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserAccessTokenUpdateRequestWithDefaults

`func NewAppUserAccessTokenUpdateRequestWithDefaults() *AppUserAccessTokenUpdateRequest`

NewAppUserAccessTokenUpdateRequestWithDefaults instantiates a new AppUserAccessTokenUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *AppUserAccessTokenUpdateRequest) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AppUserAccessTokenUpdateRequest) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AppUserAccessTokenUpdateRequest) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AppUserAccessTokenUpdateRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *AppUserAccessTokenUpdateRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *AppUserAccessTokenUpdateRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *AppUserAccessTokenUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserAccessTokenUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserAccessTokenUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppUserAccessTokenUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


