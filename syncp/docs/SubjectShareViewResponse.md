# SubjectShareViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** |  | 
**Id** | **string** |  | 
**Jwt** | **string** |  | 
**JwtClaims** | [**ActivationClaims**](ActivationClaims.md) |  | 
**SubjectExportId** | **string** |  | 
**TargetAccountNkeyPublic** | **string** |  | 

## Methods

### NewSubjectShareViewResponse

`func NewSubjectShareViewResponse(created time.Time, id string, jwt string, jwtClaims ActivationClaims, subjectExportId string, targetAccountNkeyPublic string, ) *SubjectShareViewResponse`

NewSubjectShareViewResponse instantiates a new SubjectShareViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectShareViewResponseWithDefaults

`func NewSubjectShareViewResponseWithDefaults() *SubjectShareViewResponse`

NewSubjectShareViewResponseWithDefaults instantiates a new SubjectShareViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SubjectShareViewResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SubjectShareViewResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SubjectShareViewResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetId

`func (o *SubjectShareViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubjectShareViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubjectShareViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJwt

`func (o *SubjectShareViewResponse) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SubjectShareViewResponse) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SubjectShareViewResponse) SetJwt(v string)`

SetJwt sets Jwt field to given value.


### GetJwtClaims

`func (o *SubjectShareViewResponse) GetJwtClaims() ActivationClaims`

GetJwtClaims returns the JwtClaims field if non-nil, zero value otherwise.

### GetJwtClaimsOk

`func (o *SubjectShareViewResponse) GetJwtClaimsOk() (*ActivationClaims, bool)`

GetJwtClaimsOk returns a tuple with the JwtClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtClaims

`func (o *SubjectShareViewResponse) SetJwtClaims(v ActivationClaims)`

SetJwtClaims sets JwtClaims field to given value.


### GetSubjectExportId

`func (o *SubjectShareViewResponse) GetSubjectExportId() string`

GetSubjectExportId returns the SubjectExportId field if non-nil, zero value otherwise.

### GetSubjectExportIdOk

`func (o *SubjectShareViewResponse) GetSubjectExportIdOk() (*string, bool)`

GetSubjectExportIdOk returns a tuple with the SubjectExportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectExportId

`func (o *SubjectShareViewResponse) SetSubjectExportId(v string)`

SetSubjectExportId sets SubjectExportId field to given value.


### GetTargetAccountNkeyPublic

`func (o *SubjectShareViewResponse) GetTargetAccountNkeyPublic() string`

GetTargetAccountNkeyPublic returns the TargetAccountNkeyPublic field if non-nil, zero value otherwise.

### GetTargetAccountNkeyPublicOk

`func (o *SubjectShareViewResponse) GetTargetAccountNkeyPublicOk() (*string, bool)`

GetTargetAccountNkeyPublicOk returns a tuple with the TargetAccountNkeyPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountNkeyPublic

`func (o *SubjectShareViewResponse) SetTargetAccountNkeyPublic(v string)`

SetTargetAccountNkeyPublic sets TargetAccountNkeyPublic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


