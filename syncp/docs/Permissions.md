# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pub** | Pointer to [**Permission**](Permission.md) |  | [optional] 
**Resp** | Pointer to [**NullablePermissionsResp**](PermissionsResp.md) |  | [optional] 
**Sub** | Pointer to [**Permission**](Permission.md) |  | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPub

`func (o *Permissions) GetPub() Permission`

GetPub returns the Pub field if non-nil, zero value otherwise.

### GetPubOk

`func (o *Permissions) GetPubOk() (*Permission, bool)`

GetPubOk returns a tuple with the Pub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPub

`func (o *Permissions) SetPub(v Permission)`

SetPub sets Pub field to given value.

### HasPub

`func (o *Permissions) HasPub() bool`

HasPub returns a boolean if a field has been set.

### GetResp

`func (o *Permissions) GetResp() PermissionsResp`

GetResp returns the Resp field if non-nil, zero value otherwise.

### GetRespOk

`func (o *Permissions) GetRespOk() (*PermissionsResp, bool)`

GetRespOk returns a tuple with the Resp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResp

`func (o *Permissions) SetResp(v PermissionsResp)`

SetResp sets Resp field to given value.

### HasResp

`func (o *Permissions) HasResp() bool`

HasResp returns a boolean if a field has been set.

### SetRespNil

`func (o *Permissions) SetRespNil(b bool)`

 SetRespNil sets the value for Resp to be an explicit nil

### UnsetResp
`func (o *Permissions) UnsetResp()`

UnsetResp ensures that no value is present for Resp, not even an explicit nil
### GetSub

`func (o *Permissions) GetSub() Permission`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *Permissions) GetSubOk() (*Permission, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *Permissions) SetSub(v Permission)`

SetSub sets Sub field to given value.

### HasSub

`func (o *Permissions) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


