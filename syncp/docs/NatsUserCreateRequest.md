# NatsUserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtExpiresInSecs** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**SkGroupId** | **string** |  | 
**UserClaims** | Pointer to [**NatsUserClaimsRequest**](NatsUserClaimsRequest.md) |  | [optional] 

## Methods

### NewNatsUserCreateRequest

`func NewNatsUserCreateRequest(name string, skGroupId string, ) *NatsUserCreateRequest`

NewNatsUserCreateRequest instantiates a new NatsUserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatsUserCreateRequestWithDefaults

`func NewNatsUserCreateRequestWithDefaults() *NatsUserCreateRequest`

NewNatsUserCreateRequestWithDefaults instantiates a new NatsUserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtExpiresInSecs

`func (o *NatsUserCreateRequest) GetJwtExpiresInSecs() int64`

GetJwtExpiresInSecs returns the JwtExpiresInSecs field if non-nil, zero value otherwise.

### GetJwtExpiresInSecsOk

`func (o *NatsUserCreateRequest) GetJwtExpiresInSecsOk() (*int64, bool)`

GetJwtExpiresInSecsOk returns a tuple with the JwtExpiresInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtExpiresInSecs

`func (o *NatsUserCreateRequest) SetJwtExpiresInSecs(v int64)`

SetJwtExpiresInSecs sets JwtExpiresInSecs field to given value.

### HasJwtExpiresInSecs

`func (o *NatsUserCreateRequest) HasJwtExpiresInSecs() bool`

HasJwtExpiresInSecs returns a boolean if a field has been set.

### GetName

`func (o *NatsUserCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NatsUserCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NatsUserCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSkGroupId

`func (o *NatsUserCreateRequest) GetSkGroupId() string`

GetSkGroupId returns the SkGroupId field if non-nil, zero value otherwise.

### GetSkGroupIdOk

`func (o *NatsUserCreateRequest) GetSkGroupIdOk() (*string, bool)`

GetSkGroupIdOk returns a tuple with the SkGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkGroupId

`func (o *NatsUserCreateRequest) SetSkGroupId(v string)`

SetSkGroupId sets SkGroupId field to given value.


### GetUserClaims

`func (o *NatsUserCreateRequest) GetUserClaims() NatsUserClaimsRequest`

GetUserClaims returns the UserClaims field if non-nil, zero value otherwise.

### GetUserClaimsOk

`func (o *NatsUserCreateRequest) GetUserClaimsOk() (*NatsUserClaimsRequest, bool)`

GetUserClaimsOk returns a tuple with the UserClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClaims

`func (o *NatsUserCreateRequest) SetUserClaims(v NatsUserClaimsRequest)`

SetUserClaims sets UserClaims field to given value.

### HasUserClaims

`func (o *NatsUserCreateRequest) HasUserClaims() bool`

HasUserClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


