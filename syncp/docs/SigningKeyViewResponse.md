# SigningKeyViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Current** | **bool** |  | 
**Disabled** | **bool** |  | 
**DisabledAt** | Pointer to **time.Time** |  | [optional] 
**GroupId** | **string** |  | 
**Id** | **string** |  | 
**PublicKey** | **string** |  | 
**Rotated** | **bool** |  | 
**RotatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSigningKeyViewResponse

`func NewSigningKeyViewResponse(created time.Time, current bool, disabled bool, groupId string, id string, publicKey string, rotated bool, ) *SigningKeyViewResponse`

NewSigningKeyViewResponse instantiates a new SigningKeyViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSigningKeyViewResponseWithDefaults

`func NewSigningKeyViewResponseWithDefaults() *SigningKeyViewResponse`

NewSigningKeyViewResponseWithDefaults instantiates a new SigningKeyViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SigningKeyViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SigningKeyViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SigningKeyViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCurrent

`func (o *SigningKeyViewResponse) GetCurrent() bool`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *SigningKeyViewResponse) GetCurrentOk() (*bool, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *SigningKeyViewResponse) SetCurrent(v bool)`

SetCurrent sets Current field to given value.


### GetDisabled

`func (o *SigningKeyViewResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SigningKeyViewResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SigningKeyViewResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetDisabledAt

`func (o *SigningKeyViewResponse) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *SigningKeyViewResponse) GetDisabledAtOk() (*time.Time, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *SigningKeyViewResponse) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *SigningKeyViewResponse) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### GetGroupId

`func (o *SigningKeyViewResponse) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SigningKeyViewResponse) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SigningKeyViewResponse) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetId

`func (o *SigningKeyViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SigningKeyViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SigningKeyViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPublicKey

`func (o *SigningKeyViewResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *SigningKeyViewResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *SigningKeyViewResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetRotated

`func (o *SigningKeyViewResponse) GetRotated() bool`

GetRotated returns the Rotated field if non-nil, zero value otherwise.

### GetRotatedOk

`func (o *SigningKeyViewResponse) GetRotatedOk() (*bool, bool)`

GetRotatedOk returns a tuple with the Rotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotated

`func (o *SigningKeyViewResponse) SetRotated(v bool)`

SetRotated sets Rotated field to given value.


### GetRotatedAt

`func (o *SigningKeyViewResponse) GetRotatedAt() time.Time`

GetRotatedAt returns the RotatedAt field if non-nil, zero value otherwise.

### GetRotatedAtOk

`func (o *SigningKeyViewResponse) GetRotatedAtOk() (*time.Time, bool)`

GetRotatedAtOk returns a tuple with the RotatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedAt

`func (o *SigningKeyViewResponse) SetRotatedAt(v time.Time)`

SetRotatedAt sets RotatedAt field to given value.

### HasRotatedAt

`func (o *SigningKeyViewResponse) HasRotatedAt() bool`

HasRotatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


