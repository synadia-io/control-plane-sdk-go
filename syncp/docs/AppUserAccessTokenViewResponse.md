# AppUserAccessTokenViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Expires** | **NullableString** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewAppUserAccessTokenViewResponse

`func NewAppUserAccessTokenViewResponse(created time.Time, expires NullableString, id string, name string, ) *AppUserAccessTokenViewResponse`

NewAppUserAccessTokenViewResponse instantiates a new AppUserAccessTokenViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserAccessTokenViewResponseWithDefaults

`func NewAppUserAccessTokenViewResponseWithDefaults() *AppUserAccessTokenViewResponse`

NewAppUserAccessTokenViewResponseWithDefaults instantiates a new AppUserAccessTokenViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppUserAccessTokenViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUserAccessTokenViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUserAccessTokenViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetExpires

`func (o *AppUserAccessTokenViewResponse) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AppUserAccessTokenViewResponse) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AppUserAccessTokenViewResponse) SetExpires(v string)`

SetExpires sets Expires field to given value.


### SetExpiresNil

`func (o *AppUserAccessTokenViewResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *AppUserAccessTokenViewResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetId

`func (o *AppUserAccessTokenViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUserAccessTokenViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUserAccessTokenViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppUserAccessTokenViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppUserAccessTokenViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppUserAccessTokenViewResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


